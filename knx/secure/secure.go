package secure

import (
	"fmt"
	"sync"
	"time"
)

// DataSecure manages KNX Data Secure operations for a connection
type DataSecure struct {
	mu                         sync.RWMutex
	groupKeyTable              map[string][]byte // Group Address -> Key
	individualAddressTable     map[string]uint64 // Individual Address -> Last Sequence Number
	sequenceNumberSending      uint64            // Current sequence number for outgoing frames
}

// InitialSequenceNumber calculates the initial sequence number based on the time
// since a fixed reference timestamp (2018-01-05T00:00:00Z)
func InitialSequenceNumber() uint64 {
	// Reference timestamp: 2018-01-05T00:00:00Z
	referenceTime := time.Date(2018, 1, 5, 0, 0, 0, 0, time.UTC)
	return uint64(time.Since(referenceTime).Milliseconds())
}

// NewDataSecure creates a new DataSecure instance
func NewDataSecure(groupKeyTable map[string][]byte, individualAddressTable map[string]uint64) (*DataSecure, error) {
	sequenceNumber := InitialSequenceNumber()

	// Validate sequence number is within acceptable range (0 < seqNr < 0xFFFFFFFFFFFF)
	if sequenceNumber == 0 || sequenceNumber >= 0xFFFFFFFFFFFF {
		return nil, NewDataSecureError(
			fmt.Sprintf("initial sequence number out of range: %d (system time may not be set correctly)", sequenceNumber),
		)
	}

	return &DataSecure{
		groupKeyTable:          groupKeyTable,
		individualAddressTable: individualAddressTable,
		sequenceNumberSending:  sequenceNumber,
	}, nil
}

// NewDataSecureWithSequenceNumber creates a new DataSecure instance with a specific initial sequence number
func NewDataSecureWithSequenceNumber(
	groupKeyTable map[string][]byte,
	individualAddressTable map[string]uint64,
	lastSequenceNumber uint64,
) (*DataSecure, error) {
	// Validate sequence number is within acceptable range
	if lastSequenceNumber == 0 || lastSequenceNumber >= 0xFFFFFFFFFFFF {
		return nil, NewDataSecureError(
			fmt.Sprintf("sequence number out of range: %d", lastSequenceNumber),
		)
	}

	return &DataSecure{
		groupKeyTable:          groupKeyTable,
		individualAddressTable: individualAddressTable,
		sequenceNumberSending:  lastSequenceNumber,
	}, nil
}

// GetSequenceNumber returns the next sequence number for outgoing frames
func (ds *DataSecure) GetSequenceNumber() uint64 {
	ds.mu.Lock()
	defer ds.mu.Unlock()

	seq := ds.sequenceNumberSending
	ds.sequenceNumberSending++
	return seq
}

// CheckSequenceNumber validates an incoming sequence number from a source address
// Returns an error if the sequence number is invalid or sender is not known
func (ds *DataSecure) CheckSequenceNumber(sourceAddress string, receivedSequenceNumber uint64) error {
	ds.mu.Lock()
	defer ds.mu.Unlock()

	lastValidSequenceNumber, exists := ds.individualAddressTable[sourceAddress]
	if !exists {
		return NewDataSecureError(
			fmt.Sprintf("source address not found in Security Individual Address Table: %s", sourceAddress),
		)
	}

	if receivedSequenceNumber <= lastValidSequenceNumber {
		return NewDataSecureError(
			fmt.Sprintf("sequence number too low for %s: %d received, %d last valid",
				sourceAddress, receivedSequenceNumber, lastValidSequenceNumber),
		)
	}

	// Update the sequence number if validation passed
	ds.individualAddressTable[sourceAddress] = receivedSequenceNumber
	return nil
}

// EncryptCEMI encrypts a CEMI frame with Data Secure if the destination group address has a key
// Returns the encrypted CEMI frame or the original frame if no key exists for the group address
func (ds *DataSecure) EncryptCEMI(
	sourceAddress string,
	destAddress string,
	payload []byte,
	frameFlags byte,
	tpciInt int,
) ([]byte, error) {
	ds.mu.RLock()
	key, exists := ds.groupKeyTable[destAddress]
	ds.mu.RUnlock()

	if !exists {
		// No key for this group address, return payload as-is
		return payload, nil
	}

	scf := &SecurityControlField{
		ToolAccess:      false,
		Algorithm:       CCMEncryption,
		SystemBroadcast: false,
		Service:         SAData,
	}

	sequenceNumber := ds.GetSequenceNumber()

	// Create address fields (4 bytes: 2 for source, 2 for dest)
	addressFieldsRaw := make([]byte, 4)
	copy(addressFieldsRaw[0:2], []byte(sourceAddress)[0:2]) // Simplified - in real code, parse addresses properly
	copy(addressFieldsRaw[2:4], []byte(destAddress)[0:2])

	secureData, err := NewSecureData(
		key,
		payload,
		scf,
		sequenceNumber,
		addressFieldsRaw,
		frameFlags,
		tpciInt,
	)
	if err != nil {
		return nil, err
	}

	// In real implementation, wrap in SecureAPDU here
	return secureData.ToKNX(), nil
}

// DecryptCEMI decrypts a secure CEMI frame with Data Secure
func (ds *DataSecure) DecryptCEMI(
	sourceAddress string,
	destAddress string,
	payload []byte,
	frameFlags byte,
	tpciInt int,
	scf *SecurityControlField,
) ([]byte, error) {
	// Check sequence number first
	secureData := &SecureData{}
	if err := secureData.FromKNX(payload); err != nil {
		return nil, err
	}

	sequenceNumber := binary.BigEndian.Uint64(append([]byte{0, 0}, secureData.SequenceNumberBytes...))
	if err := ds.CheckSequenceNumber(sourceAddress, sequenceNumber); err != nil {
		return nil, err
	}

	ds.mu.RLock()
	key, exists := ds.groupKeyTable[destAddress]
	ds.mu.RUnlock()

	if !exists {
		return nil, NewDataSecureError(
			fmt.Sprintf("no key found for group address %s from %s", destAddress, sourceAddress),
		)
	}

	// Create address fields
	addressFieldsRaw := make([]byte, 4)
	copy(addressFieldsRaw[0:2], []byte(sourceAddress)[0:2])
	copy(addressFieldsRaw[2:4], []byte(destAddress)[0:2])

	plainAPDU, err := secureData.GetPlainAPDU(
		key,
		scf,
		addressFieldsRaw,
		frameFlags,
		tpciInt,
	)
	if err != nil {
		return nil, err
	}

	return plainAPDU, nil
}

// GetGroupKey returns the key for a group address
func (ds *DataSecure) GetGroupKey(groupAddress string) ([]byte, bool) {
	ds.mu.RLock()
	defer ds.mu.RUnlock()

	key, exists := ds.groupKeyTable[groupAddress]
	return key, exists
}

// GetSenderSequenceNumber returns the last known sequence number for a sender
func (ds *DataSecure) GetSenderSequenceNumber(individualAddress string) (uint64, bool) {
	ds.mu.RLock()
	defer ds.mu.RUnlock()

	seq, exists := ds.individualAddressTable[individualAddress]
	return seq, exists
}

// SetSenderSequenceNumber sets the sequence number for a sender
func (ds *DataSecure) SetSenderSequenceNumber(individualAddress string, sequenceNumber uint64) {
	ds.mu.Lock()
	defer ds.mu.Unlock()

	ds.individualAddressTable[individualAddress] = sequenceNumber
}

// AddGroupKey adds or updates a group key
func (ds *DataSecure) AddGroupKey(groupAddress string, key []byte) error {
	if len(key) != 16 {
		return NewDataSecureError("group key must be 16 bytes")
	}

	ds.mu.Lock()
	defer ds.mu.Unlock()

	ds.groupKeyTable[groupAddress] = key
	return nil
}

// AddSender adds a sender to the individual address table
func (ds *DataSecure) AddSender(individualAddress string, initialSequenceNumber uint64) error {
	ds.mu.Lock()
	defer ds.mu.Unlock()

	ds.individualAddressTable[individualAddress] = initialSequenceNumber
	return nil
}

// Export imports internal binary encoding (needed for import statement)
import "encoding/binary"
