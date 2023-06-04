package global

import (
	"crypto/sha256"
	"encoding/hex"
)

func Sha256Encode(str string) string {
	hash := sha256.New()
	hash.Write([]byte(str))
	bytes := hash.Sum(nil)
	return hex.EncodeToString(bytes)
}
