package helpers

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"io"
)

// AES256Encrypt encrypts plaintext using password and returns it as ciphertext
func AES256Encrypt(password string, plaintext []byte) ([]byte, error) {
	// Convert password to sha-256 hash for guarantee algorithm to use AES-256
	passwordHash := sha256.New()
	passwordHash.Write([]byte(password))
	key := passwordHash.Sum(nil)

	// Define AES cipher
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// Define GCM for AES cipher
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	//Create a nonce from secure random bytes with desired size from GCM
	nonce := make([]byte, aesGCM.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}

	// Encrypt plaintext using GCM with nonce
	ciphertext := aesGCM.Seal(nil, nonce, plaintext, nil)

	return ciphertext, nil
}

// AES256Decrypt decrypts ciphertext using password and returns it as plaintext
func AES256Decrypt(password string, encryptedData []byte) ([]byte, error) {
	// Convert password to sha-256 hash for guarantee algorithm to use AES-256
	passwordHash := sha256.New()
	passwordHash.Write([]byte(password))
	key := passwordHash.Sum(nil)

	// Define AES cipher
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// Define GCM for AES cipher
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	// Extract nonce from encrypted data
	nonce := encryptedData[:aesGCM.NonceSize()]

	// Extract ciphertext from encrypted data
	ciphertext := encryptedData[aesGCM.NonceSize():]

	// Decrypt ciphertext using GCM with nonce
	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}
