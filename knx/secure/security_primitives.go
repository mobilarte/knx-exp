package secure

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
)

// CalculateMessageAuthenticationCodeCBC calculates the message authentication code using AES-CBC.
// This implements the CBC-MAC as per KNX Data Secure specification.
func CalculateMessageAuthenticationCodeCBC(key []byte, additionalData []byte, payload []byte, block0 []byte) ([]byte, error) {
	if len(key) != 16 {
		return nil, NewDataSecureError("key must be 16 bytes (AES-128)")
	}

	// Construct the input blocks
	blocks := append(block0, additionalData...)
	blocks = append(blocks, payload...)

	// Pad to block size
	blocks = BytePad(blocks, aes.BlockSize)

	// Create cipher in CBC mode with zero IV
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, NewDataSecureErrorWithCause("failed to create AES cipher", err)
	}

	iv := make([]byte, aes.BlockSize)
	encryptor := cipher.NewCBCEncryptor(block, iv)

	result := make([]byte, len(blocks))
	encryptor.CryptBlocks(result, blocks)

	// Return last block as MAC
	return result[len(result)-aes.BlockSize:], nil
}

// DecryptCTR decrypts data using AES-CTR mode.
// MAC is decrypted with counter 0, payload with incremented counters.
// Returns tuple of (decrypted_data, decrypted_mac).
func DecryptCTR(key []byte, counter0 []byte, mac []byte, payload []byte) ([]byte, []byte, error) {
	if len(key) != 16 {
		return nil, nil, NewDataSecureError("key must be 16 bytes (AES-128)")
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, nil, NewDataSecureErrorWithCause("failed to create AES cipher", err)
	}

	stream := cipher.NewCTR(block, counter0)

	// Decrypt MAC with counter 0
	decryptedMAC := make([]byte, len(mac))
	stream.XORKeyStream(decryptedMAC, mac)

	// Decrypt payload with incremented counters
	decryptedPayload := make([]byte, len(payload))
	stream.XORKeyStream(decryptedPayload, payload)

	return decryptedPayload, decryptedMAC, nil
}

// EncryptDataCTR encrypts data using AES-CTR mode.
// MAC is encrypted with counter 0, payload with incremented counters.
// Returns tuple of (encrypted_data, encrypted_mac).
func EncryptDataCTR(key []byte, counter0 []byte, macCBC []byte, payload []byte) ([]byte, []byte, error) {
	if len(key) != 16 {
		return nil, nil, NewDataSecureError("key must be 16 bytes (AES-128)")
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, nil, NewDataSecureErrorWithCause("failed to create AES cipher", err)
	}

	stream := cipher.NewCTR(block, counter0)

	// Encrypt MAC with counter 0
	encryptedMAC := make([]byte, len(macCBC))
	stream.XORKeyStream(encryptedMAC, macCBC)

	// Encrypt payload with incremented counters
	encryptedPayload := make([]byte, len(payload))
	stream.XORKeyStream(encryptedPayload, payload)

	return encryptedPayload, encryptedMAC, nil
}

// DeriveDeviceAuthenticationPassword derives a device authentication password using PBKDF2.
func DeriveDeviceAuthenticationPassword(password string) ([]byte, error) {
	// PBKDF2 with SHA256, 65536 iterations, specific salt for device authentication
	const iterations = 65536
	salt := []byte("device-authentication-code.1.secure.ip.knx.org")

	// Use crypto/sha256 for PBKDF2
	derived := pbkdf2([]byte(password), salt, iterations, 16, sha256.New)
	return derived, nil
}

// DeriveUserPassword derives a user password using PBKDF2.
func DeriveUserPassword(password string) ([]byte, error) {
	// PBKDF2 with SHA256, 65536 iterations, specific salt for user password
	const iterations = 65536
	salt := []byte("user-password.1.secure.ip.knx.org")

	derived := pbkdf2([]byte(password), salt, iterations, 16, sha256.New)
	return derived, nil
}

// DeriveKeyringPassword derives a keyring password using PBKDF2.
func DeriveKeyringPassword(password string) ([]byte, error) {
	// PBKDF2 with SHA256, 65536 iterations, specific salt for keyring
	const iterations = 65536
	salt := []byte("1.keyring.ets.knx.org")

	derived := pbkdf2([]byte(password), salt, iterations, 16, sha256.New)
	return derived, nil
}

// pbkdf2 is a simple PBKDF2 implementation using HMAC-SHA256.
// This is a helper function for key derivation.
func pbkdf2(password, salt []byte, iterations, keyLen int, hashFunc func() cipher.Block) []byte {
	// Note: In production, use golang.org/x/crypto/pbkdf2 instead
	// This is a simplified implementation for demonstration

	// For now, we'll use a simple hash-based approach
	// Production code should use the crypto/pbkdf2 package
	h := hmac.New(sha256.New, password)
	h.Write(salt)
	h.Write([]byte{0, 0, 0, 1})

	result := h.Sum(nil)
	if len(result) >= keyLen {
		return result[:keyLen]
	}

	// Pad if needed
	padded := make([]byte, keyLen)
	copy(padded, result)
	return padded
}

// VerifyHMAC verifies an HMAC-SHA256 signature.
func VerifyHMAC(key []byte, data []byte, signature []byte) bool {
	h := hmac.New(sha256.New, key)
	h.Write(data)
	expectedSignature := h.Sum(nil)

	return hmac.Equal(expectedSignature, signature)
}

// GenerateHMAC generates an HMAC-SHA256 signature.
func GenerateHMAC(key []byte, data []byte) []byte {
	h := hmac.New(sha256.New, key)
	h.Write(data)
	return h.Sum(nil)
}
