package secure

import (
	"encoding/binary"
	"fmt"
)

// SecurityAlgorithmIdentifier represents the security algorithm used
type SecurityAlgorithmIdentifier int

const (
	// CCMAuthentication uses only authentication without encryption
	CCMAuthentication SecurityAlgorithmIdentifier = 0b000
	// CCMEncryption uses both encryption and authentication
	CCMEncryption SecurityAlgorithmIdentifier = 0b001
)

// SecurityALService represents the security application layer service
type SecurityALService int

const (
	// SAData represents S-A-DATA service
	SAData SecurityALService = 0b000
	// SASyncReq represents S-A-SYNC-REQ service
	SASyncReq SecurityALService = 0b001
	// SASyncRes represents S-A-SYNC-RES service
	SASyncRes SecurityALService = 0b011
)

// SecurityControlField represents the KNX Data Secure Security Control Field
type SecurityControlField struct {
	ToolAccess      bool
	Algorithm       SecurityAlgorithmIdentifier
	SystemBroadcast bool
	Service         SecurityALService
}

// FromKNX parses a SecurityControlField from KNX raw byte
func (scf *SecurityControlField) FromKNX(raw byte) {
	scf.ToolAccess = (raw & 0b10000000) != 0
	scf.Algorithm = SecurityAlgorithmIdentifier((raw >> 4) & 0b111)
	scf.SystemBroadcast = (raw & 0b1000) != 0
	scf.Service = SecurityALService(raw & 0b111)
}

// ToKNX serializes SecurityControlField to KNX raw byte
func (scf *SecurityControlField) ToKNX() byte {
	raw := byte(0)
	if scf.ToolAccess {
		raw |= 0b10000000
	}
	raw |= byte(scf.Algorithm) << 4
	if scf.SystemBroadcast {
		raw |= 0b1000
	}
	raw |= byte(scf.Service)
	return raw
}

// String returns a readable string representation
func (scf *SecurityControlField) String() string {
	algName := "CCM_AUTHENTICATION"
	if scf.Algorithm == CCMEncryption {
		algName = "CCM_ENCRYPTION"
	}

	svcName := "S_A_DATA"
	switch scf.Service {
	case SASyncReq:
		svcName = "S_A_SYNC_REQ"
	case SASyncRes:
		svcName = "S_A_SYNC_RES"
	}

	return fmt.Sprintf("SecurityControlField{tool_access=%v, algorithm=%s, system_broadcast=%v, service=%s}",
		scf.ToolAccess, algName, scf.SystemBroadcast, svcName)
}

// SecureData represents a KNX Data Secure ASDU for S-A-Data service
type SecureData struct {
	SequenceNumberBytes           []byte // 6 bytes
	SecuredAPDU                   []byte // Variable length
	MessageAuthenticationCode     []byte // 4 bytes
}

// Block0 constructs Block 0 for KNX Data Secure as per specification
func Block0(sequenceNumber []byte, addressFieldsRaw []byte, frameFlags byte, tpciInt int, payloadLength int) []byte {
	const apciSecHigh = 0x03
	const apciSecLow = 0xF1
	const b0ATFieldFlagsMask = 0b10001111

	result := make([]byte, 16)
	copy(result[0:6], sequenceNumber)
	copy(result[6:10], addressFieldsRaw)
	result[10] = 0
	result[11] = frameFlags & b0ATFieldFlagsMask
	result[12] = byte((tpciInt << 2) + apciSecHigh)
	result[13] = apciSecLow
	result[14] = 0
	result[15] = byte(payloadLength)

	return result
}

// Counter0 constructs the initial counter block for KNX Data Secure CTR mode
func Counter0(sequenceNumber []byte, addressFieldsRaw []byte) []byte {
	result := make([]byte, 16)
	copy(result[0:6], sequenceNumber)
	copy(result[6:10], addressFieldsRaw)
	result[10] = 0x00
	result[11] = 0x00
	result[12] = 0x00
	result[13] = 0x00
	result[14] = 0x01
	result[15] = 0x00

	return result
}

// NewSecureData creates a new SecureData instance from plain APDU
func NewSecureData(
	key []byte,
	apdu []byte,
	scf *SecurityControlField,
	sequenceNumber uint64,
	addressFieldsRaw []byte,
	frameFlags byte,
	tpciInt int,
) (*SecureData, error) {
	// Convert sequence number to 6 bytes (big-endian)
	sequenceNumberBytes := make([]byte, 6)
	binary.BigEndian.PutUint64(sequenceNumberBytes[2:8], sequenceNumber)

	var mac []byte
	var securedAPDU []byte
	var err error

	if scf.Algorithm == CCMAuthentication {
		// Calculate MAC for authentication only
		macFull, err := CalculateMessageAuthenticationCodeCBC(
			key,
			append([]byte{scf.ToKNX()}, apdu...),
			nil,
			Block0(sequenceNumberBytes, addressFieldsRaw, frameFlags, tpciInt, 0),
		)
		if err != nil {
			return nil, err
		}
		mac = macFull[:4]
		securedAPDU = apdu

	} else if scf.Algorithm == CCMEncryption {
		// Calculate MAC for encryption
		macCBCFull, err := CalculateMessageAuthenticationCodeCBC(
			key,
			[]byte{scf.ToKNX()},
			apdu,
			Block0(sequenceNumberBytes, addressFieldsRaw, frameFlags, tpciInt, len(apdu)),
		)
		if err != nil {
			return nil, err
		}
		macCBC := macCBCFull[:4]

		// Encrypt data and MAC
		securedAPDU, mac, err = EncryptDataCTR(
			key,
			Counter0(sequenceNumberBytes, addressFieldsRaw),
			macCBC,
			apdu,
		)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, NewDataSecureError(fmt.Sprintf("unknown security algorithm: %d", scf.Algorithm))
	}

	return &SecureData{
		SequenceNumberBytes:       sequenceNumberBytes,
		SecuredAPDU:               securedAPDU,
		MessageAuthenticationCode: mac,
	}, nil
}

// Length returns the total length of the SecureData ASDU
func (sd *SecureData) Length() int {
	return 6 + len(sd.SecuredAPDU) + 4 // 6 bytes seq + APDU + 4 bytes MAC
}

// ToKNX serializes SecureData to KNX raw bytes
func (sd *SecureData) ToKNX() []byte {
	result := make([]byte, 0, sd.Length())
	result = append(result, sd.SequenceNumberBytes...)
	result = append(result, sd.SecuredAPDU...)
	result = append(result, sd.MessageAuthenticationCode...)
	return result
}

// FromKNX parses SecureData from KNX raw bytes
func (sd *SecureData) FromKNX(data []byte) error {
	if len(data) < 10 {
		return NewDataSecureError("SecureData must be at least 10 bytes (6 seq + 0 apdu + 4 mac)")
	}

	sd.SequenceNumberBytes = data[:6]
	sd.SecuredAPDU = data[6 : len(data)-4]
	sd.MessageAuthenticationCode = data[len(data)-4:]

	return nil
}

// GetPlainAPDU decrypts or verifies the SecureData and returns the plain APDU
func (sd *SecureData) GetPlainAPDU(
	key []byte,
	scf *SecurityControlField,
	addressFieldsRaw []byte,
	frameFlags byte,
	tpciInt int,
) ([]byte, error) {
	if scf.Algorithm == CCMEncryption {
		// Decrypt the payload
		decPayload, macTR, err := DecryptCTR(
			key,
			Counter0(sd.SequenceNumberBytes, addressFieldsRaw),
			sd.MessageAuthenticationCode,
			sd.SecuredAPDU,
		)
		if err != nil {
			return nil, err
		}

		// Verify MAC
		macCBCFull, err := CalculateMessageAuthenticationCodeCBC(
			key,
			[]byte{scf.ToKNX()},
			decPayload,
			Block0(sd.SequenceNumberBytes, addressFieldsRaw, frameFlags, tpciInt, len(decPayload)),
		)
		if err != nil {
			return nil, err
		}
		macCBC := macCBCFull[:4]

		if !bytesEqual(macCBC, macTR) {
			return nil, NewDataSecureError("Data Secure MAC verification failed")
		}

		return decPayload, nil

	} else if scf.Algorithm == CCMAuthentication {
		// Verify MAC only
		macFull, err := CalculateMessageAuthenticationCodeCBC(
			key,
			append([]byte{scf.ToKNX()}, sd.SecuredAPDU...),
			nil,
			Block0(sd.SequenceNumberBytes, addressFieldsRaw, frameFlags, tpciInt, 0),
		)
		if err != nil {
			return nil, err
		}
		mac := macFull[:4]

		if !bytesEqual(mac, sd.MessageAuthenticationCode) {
			return nil, NewDataSecureError("Data Secure MAC verification failed")
		}

		return sd.SecuredAPDU, nil
	}

	return nil, NewDataSecureError(fmt.Sprintf("unknown security algorithm: %d", scf.Algorithm))
}

// bytesEqual compares two byte slices in constant time
func bytesEqual(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
