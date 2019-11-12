package services

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"io"
	"os"
)

func GetAESToken() [32]byte {
	aesToken := os.Getenv("AES_TOKEN")

	keyStr, _ := hex.DecodeString(aesToken)
	var key [32]byte
	copy(key[:], keyStr)

	return key
}

func Encrypt(plaintext []byte, key *[32]byte) (cipherText []byte, err error) {
	block, err := aes.NewCipher(key[:])
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		return nil, err
	}

	return gcm.Seal(nonce, nonce, plaintext, nil), nil
}

func Decrypt(cipherText []byte, key *[32]byte) (plaintext []byte, err error) {
	block, err := aes.NewCipher(key[:])
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	if len(cipherText) < gcm.NonceSize() {
		return nil, errors.New("malformed cipher text")
	}

	return gcm.Open(nil,
		cipherText[:gcm.NonceSize()],
		cipherText[gcm.NonceSize():],
		nil,
	)
}