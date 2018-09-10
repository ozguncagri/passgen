package helpers

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
)

// AES256Encrypt encrypts plaintext using password and returns it as ciphertext
func AES256Encrypt(password string, plaintext []byte) ([]byte, error) {
	// Convert password to sha-256 hash for guarantee algorithm to use aes-256
	passwordHash := sha256.New()
	passwordHash.Write([]byte(password))
	key := passwordHash.Sum(nil)

	// Use some part of hash as nonce for decryption
	nonce := key[0:12]

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	ciphertext := aesgcm.Seal(nil, nonce, plaintext, nil)

	return ciphertext, nil
}
