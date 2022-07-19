package des

import (
	"fmt"
	"github.com/wangyufengx/utils/crypt/base"
	"log"
	"testing"
)

func TestCBCDES(t *testing.T) {
	key := []byte("2fa6c1e9")
	str := []byte("Hello des!")
	t.Run("des cbc encrypt decrypt", func(t *testing.T) {
		strEncrypted, err := CBCEncrypt(str, key)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Encrypted:", base.Base64Encode(strEncrypted))

		decrypt, err := CBCDecrypt(strEncrypted, key)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("decrypted:", string(decrypt))
	})

	t.Run("des cbc str encrypt decrypt", func(t *testing.T) {
		encrypt, err := CBCEncryptToBase64String(str, key)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Encrypted:", encrypt)
		decode, err := base.Base64Decode(encrypt)
		if err != nil {
			log.Fatal(err)
		}
		decrypt, err := CBCDecryptString(string(decode), key)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("decrypted:", string(decrypt))
	})

}
