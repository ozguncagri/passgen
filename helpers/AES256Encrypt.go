package helpers

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
)

// AES256Encrypt encrypts plaintext using password and returns it as ciphertext
func AES256Encrypt(password string, plaintext []byte) ([]byte, error) {
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

	// Encrypt plaintext using GCM with nonce
	ciphertext := aesgcm.Seal(nil, nonce, plaintext, nil)

	return ciphertext, nil
}
