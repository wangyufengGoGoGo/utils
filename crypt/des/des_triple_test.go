package des

import (
	"fmt"
	"github.com/wangyufengx/utils/crypt/base"
	"log"
	"testing"
)

func TestTripleDES(t *testing.T) {

	key := []byte("2fa6c1e92fa6c1e92fa6c1e9")
	str := []byte("Hello des!")

	t.Run("des triple encrypt decrypt", func(t *testing.T) {
		strEncrypted, err := TripleEncrypt(str, key)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Encrypted:", base.Base64Encode(strEncrypted))

		decrypt, err := TripleDecrypt(strEncrypted, key)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("decrypted:", string(decrypt))
	})

	t.Run("des triple str encrypt decrypt", func(t *testing.T) {
		encrypt, err := TripleEncryptToBase64String(str, key)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Encrypted:", encrypt)
		decrypt, err := TripleDecryptString(encrypt, key)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("decrypted:", string(decrypt))
	})

}
