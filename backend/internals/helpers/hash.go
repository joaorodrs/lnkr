package helpers

import (
	"crypto/sha256"
	"encoding/hex"
)

func GenerateHash(URL string) string {
	hash := sha256.Sum256([]byte(URL))
	return hex.EncodeToString(hash[:8])
}
