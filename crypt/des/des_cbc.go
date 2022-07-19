package des

import (
	"crypto/cipher"
	"crypto/des"
	"github.com/wangyufengx/utils/crypt"
	"github.com/wangyufengx/utils/crypt/base"
)

//CBCEncrypt DES对称加密CBC模式加密
func CBCEncrypt(origData, key []byte) ([]byte, error) {
	//校验秘钥长度
	if len(key) != des.BlockSize {
		return nil, des.KeySizeError(des.BlockSize)
	}

	//创建一个密码块
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}

	//补码
	origData = crypt.PKCS5Padding(origData, block.BlockSize())

	//设计加密模式
	blockMode := cipher.NewCBCEncrypter(block, key)
	//存放加密后的密文
	encrypt := make([]byte, len(origData))
	//块加密
	blockMode.CryptBlocks(encrypt, origData)
	return encrypt, nil
}

//CBCEncryptToBase64String DES对称加密CBC模式加密返回base64编码
func CBCEncryptToBase64String(origData, key []byte) (string, error) {
	encrypt, err := CBCEncrypt(origData, key)
	if err != nil {
		return "", err
	}
	return base.Base64Encode(encrypt), nil
}

//CBCDecrypt DES解密
func CBCDecrypt(encrypt, key []byte) ([]byte, error) {
	//校验秘钥长度
	if len(key) != des.BlockSize {
		return nil, des.KeySizeError(des.BlockSize)
	}

	//创建一个密码块
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockMode := cipher.NewCBCDecrypter(block, key)

	//创建缓冲，存放解密后的数据
	orgData := make([]byte, len(encrypt))
	//开始解密
	blockMode.CryptBlocks(orgData, encrypt)
	//去掉编码
	orgData, err = crypt.PKCS5UnPadding(orgData, block.BlockSize())
	if err != nil {
		return nil, err
	}
	return orgData, nil
}

//CBCDecryptString DES解密
func CBCDecryptString(encryptBase64Str string, key []byte) ([]byte, error) {
	encrypt, err := base.Base64Decode(encryptBase64Str)
	if err != nil {
		return nil, err
	}
	return CBCDecrypt(encrypt, key)
}
