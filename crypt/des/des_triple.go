package des

import (
	"crypto/cipher"
	"crypto/des"
	"github.com/wangyufengx/utils/crypt"
	"github.com/wangyufengx/utils/crypt/base"
)

func TripleEncrypt(origData, key []byte) ([]byte, error) {
	//key长度校验
	if len(key) != 24 {
		return nil, des.KeySizeError(len(key))
	}

	//创建3des密码块
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		return nil, err
	}

	//补码
	origData = crypt.PKCS5Padding(origData, block.BlockSize())
	//设置加密模式
	blockMode := cipher.NewCBCEncrypter(block, key[:8])

	encrypt := make([]byte, len(origData))

	//加密处理
	blockMode.CryptBlocks(encrypt, origData)
	return encrypt, nil
}

func TripleEncryptToBase64String(origData, key []byte) (string, error) {
	bytes, err := TripleEncrypt(origData, key)
	if err != nil {
		return "", err
	}
	return base.Base64Encode(bytes), nil
}

func TripleDecrypt(encrypt, key []byte) ([]byte, error) {
	//key长度校验
	if len(key) != 24 {
		return nil, des.KeySizeError(len(key))
	}

	//创建3des密码块
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		return nil, err
	}

	blockMode := cipher.NewCBCDecrypter(block, key[:8])
	origData := make([]byte, len(encrypt))

	blockMode.CryptBlocks(origData, encrypt)

	return crypt.PKCS5UnPadding(origData, 24)
}

func TripleDecryptString(encryptBase64Str string, key []byte) ([]byte, error) {
	encrypt, err := base.Base64Decode(encryptBase64Str)
	if err != nil {
		return nil, err
	}
	return TripleDecrypt(encrypt, key)
}
