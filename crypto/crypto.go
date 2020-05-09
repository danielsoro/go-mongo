package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"io"
	"io/ioutil"

	configuration "github.com/danielsoro/go-mongo/config"
)

func _makeNonce(gcm cipher.AEAD) ([]byte, error) {
	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}
	return nonce, nil
}

func _createCipher() (cipher.AEAD, error) {
	config := configuration.GetCryptoConfiguration()
	if len(config.Key) == 0 {
		return nil, errors.New("crypt key is not defined in your configuration")
	}

	key, err := ioutil.ReadFile(config.Key)
	if err != nil {
		return nil, err
	}

	hexDecode, err := hex.DecodeString(string(key))
	if err != nil {
		return nil, err
	}

	c, err := aes.NewCipher(hexDecode)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}

	return gcm, nil
}

// Encrypt string to AES formater with HEX return
func Encrypt(s string) (string, error) {
	aesgcm, err := _createCipher()
	if err != nil {
		return "", err
	}

	nonce, err := _makeNonce(aesgcm)
	if err != nil {
		return "", err
	}

	ciphertext := aesgcm.Seal(nonce, nonce, []byte(s), nil)
	encodedString := hex.EncodeToString(ciphertext)
	return encodedString, nil
}

// Decrypt string from AES formater with HEX
func Decrypt(s string) (string, error) {
	decodeString, err := hex.DecodeString(s)
	if err != nil {
		return "", err
	}

	aesgcm, err := _createCipher()
	if err != nil {
		return "", err
	}

	nonceSize := aesgcm.NonceSize()
	if len(decodeString) < nonceSize {
		return "", errors.New("nonce size > parameter size")
	}

	nonce, ciphertext := decodeString[:nonceSize], decodeString[nonceSize:]
	plaintext, err := aesgcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}
