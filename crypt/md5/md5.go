package md5

import (
	"crypto/md5"
	"fmt"
)

func GetMD5(data []byte) string {
	hash := md5.New()
	hash.Write(data)
	return fmt.Sprintf("%x", hash.Sum([]byte(nil)))
}

func GetMD5ByString(data string) string {
	hash := md5.New()
	hash.Write([]byte(data))
	return fmt.Sprintf("%x", hash.Sum([]byte(nil)))
}
