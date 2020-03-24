package services

import (
	"bytes"
	"os"
	"reflect"
	"testing"
)

var validCipher = [32]byte{120, 64, 164, 122, 131, 124, 47, 16, 195, 202, 234, 158, 141, 169, 80, 185, 1, 87, 187, 67,
	197, 18, 164, 102, 88, 96, 177, 113, 193, 184, 230, 24}

var validEncryptedString = []byte{242, 13, 178, 88, 15, 251, 179, 192, 23, 138, 168, 245, 189, 43, 31, 142, 32, 174, 176,
	90, 130, 131, 144, 107, 16, 129, 143, 80, 32, 197, 183, 37, 66, 74, 238, 190, 40, 253, 117, 68, 110, 127, 112, 7}

var validDecryptedString = []byte{101, 110, 99, 114, 121, 112, 116, 101, 100, 32, 115, 116, 114, 105, 110, 103}

func TestGetAESToken(t *testing.T) {
	err := os.Setenv("AES_TOKEN", "7840a47a837c2f10c3caea9e8da950b90157bb43c512a4665860b171c1b8e618")

	if err != nil {
		t.Fatalf("Unable to set env variable. Error: %v", err)
	} else {
		t.Logf("Env variable set successfuly")
	}

	decodedCipher, err := GetAESToken()

	if err != nil {
		t.Errorf("Unable to decode token to byte slice. Error: %v", err)
	} else {
		t.Logf("Cipher string decode successful")
	}

	if reflect.DeepEqual(decodedCipher, validCipher) {
		t.Logf("Cipher decode successful")
	} else {
		t.Error("Decoded cipher doesn't match")
	}
}

func TestEncrypt(t *testing.T) {
	_, err := Encrypt([]byte("encrypted string"), &validCipher)

	if err != nil {
		t.Errorf("Unable to encrypt string. Error: %v", err)
	} else {
		t.Logf("String encrypted successful")
	}
}

func TestDecrypt(t *testing.T) {
	decryptedString, err := Decrypt(validEncryptedString, &validCipher)

	if err != nil {
		t.Errorf("Unable to decrypt string. Error: %v", err)
	} else {
		t.Logf("String decrypted successful")
	}

	if bytes.Equal(decryptedString, validDecryptedString) {
		t.Log("Decrypted strings match")
	} else {
		t.Error("Decrypted strings don't match")
	}
}
