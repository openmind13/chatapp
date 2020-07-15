package utils

import (
	"crypto/md5"
	"encoding/hex"
)

// CreateHash hashes a string using md5
func CreateHash(strData string) string {
	hasher := md5.New()
	hasher.Write([]byte(strData))
	return hex.EncodeToString(hasher.Sum(nil))
}
