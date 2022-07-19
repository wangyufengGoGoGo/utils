package crypt

import (
	"crypto/sha256"
	"fmt"
)

func GetSha256(data []byte) string {
	hash := sha256.New()
	hash.Write(data)
	return fmt.Sprintf("%x", hash.Sum(nil))
}

func GetSha256ByString(data string) string {
	hash := sha256.New()
	hash.Write([]byte(data))
	return fmt.Sprintf("%x", hash.Sum(nil))
}
