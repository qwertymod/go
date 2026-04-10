package encrypter

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"os"
)


type Encrypter struct {
	Key string
}


func NewEncrypter() (*Encrypter) {
	key := os.Getenv("KEY")
	if key == "" {
		panic("Не передан параметр key в переменной окружения")
	}
	return &Encrypter{
		Key : key,
	}
}

func (enc *Encrypter) Encrypt(plainSt []byte) []byte {
	block, err := aes.NewCipher([]byte(enc.Key))
	if err != nil {
		panic(err.Error())
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	nonce := make([]byte, aesGCM.NonceSize())
	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		panic(err.Error())
	}

	return aesGCM.Seal(nonce, nonce, plainSt, nil)
}

func (enc *Encrypter) Decrypter(encryptedSt []byte) []byte {
	block, err := aes.NewCipher([]byte(enc.Key))
	if err != nil {
		panic(err.Error())
	}
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonceSize := aesGCM.NonceSize()
	nonce , cipherText:= encryptedSt[:nonceSize] , encryptedSt[nonceSize:]

	plainText , err := aesGCM.Open(nil, nonce , cipherText, nil)

	if err != nil {
		panic(err.Error())
	}

	return plainText
}