package hash

import (
	"crypto/sha256"
	"encoding/hex"
	"golang.org/x/crypto/sha3"
)

const (
	Size = 32
)

// SHA256 returns the SHA2-256 checksum of the data.
func SHA256(data []byte) []byte {
	b := sha256.Sum256(data)
	return b[:]
}

// SHA256Hex returns hex encode a SHA2-256 hash on data
func SHA256Hex(data []byte) string {
	return hex.EncodeToString(SHA256(data))
}

// Keccak256 returns the SHA3-256 digest of the data.
func Keccak256Hash(data []byte) []byte {
	b := sha3.Sum256(data)
	return b[:]
}

// Keccak256Hex returns hex encode a SHA3-256 hash on data
func Keccak256Hex(data []byte) string {
	return hex.EncodeToString(Keccak256Hash(data))
}
