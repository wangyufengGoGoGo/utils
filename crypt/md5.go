package crypt

import (
	"crypto/md5"
	"fmt"
)

func GetMD5(data string) string {
	hash := md5.New()
	hash.Write([]byte(data))
	return fmt.Sprintf("%x", hash.Sum([]byte(nil)))
}
