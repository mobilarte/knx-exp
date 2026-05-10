package secure

import (
	"crypto/sha256"
)

// BytesXOR performs XOR operation on two byte slices of equal length.
// Returns an error if the lengths are not equal.
func BytesXOR(a, b []byte) ([]byte, error) {
	if len(a) != len(b) {
		return nil, NewDataSecureError("XOR operands must have equal length")
	}

	result := make([]byte, len(a))
	for i := range a {
		result[i] = a[i] ^ b[i]
	}
	return result, nil
}

// BytePad pads data with zero bytes until its length is a multiple of blockSize.
func BytePad(data []byte, blockSize int) []byte {
	remainder := len(data) % blockSize
	if remainder == 0 {
		return data
	}
	padding := make([]byte, blockSize-remainder)
	return append(data, padding...)
}

// SHA256Hash calculates the SHA256 hash of data.
func SHA256Hash(data []byte) []byte {
	h := sha256.Sum256(data)
	return h[:]
}
