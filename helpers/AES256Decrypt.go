package helpers

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
)

// AES256Decrypt decrypts ciphertext using password and returns it as plaintext
func AES256Decrypt(password string, ciphertext []byte) ([]byte, error) {
	// Convert password to sha-256 hash for guarantee algorithm to use AES-256
	passwordHash := sha256.New()
	passwordHash.Write([]byte(password))
	key := passwordHash.Sum(nil)

	// Use static part of hash as nonce for successful decryption
	nonce := key[0:12]

	// Define AES cipher
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// Define GCM for AES cipher
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	// Decrypt ciphertext using GCM with nonce
	plaintext, err := aesgcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}
