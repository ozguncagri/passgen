package helpers

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"io"

	"golang.org/x/crypto/pbkdf2"
)

// deriveKey generates a 32-byte AES key from password and salt
func deriveKey(password, salt []byte) []byte {
	return pbkdf2.Key(password, salt, 500_000, 32, sha256.New)
}

// Encrypt encrypts plaintext using a password and returns Base64 string
func Encrypt(plaintext []byte, password string) ([]byte, error) {
	// Generate random salt
	salt := make([]byte, 16)
	if _, err := io.ReadFull(rand.Reader, salt); err != nil {
		return []byte{}, err
	}

	key := deriveKey([]byte(password), salt)

	block, err := aes.NewCipher(key)
	if err != nil {
		return []byte{}, err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return []byte{}, err
	}

	nonce := make([]byte, aesGCM.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return []byte{}, err
	}

	ciphertext := aesGCM.Seal(nil, nonce, plaintext, nil)

	// Combine salt + nonce + ciphertext
	result := append(salt, nonce...)
	result = append(result, ciphertext...)

	return result, nil
}

// Decrypt decrypts a Base64-encoded string using a password
func Decrypt(ciphertext []byte, password string) ([]byte, error) {
	if len(ciphertext) < 16+12 { // minimum salt + nonce size
		return []byte{}, fmt.Errorf("invalid encrypted data")
	}

	salt := ciphertext[:16]
	key := deriveKey([]byte(password), salt)

	block, err := aes.NewCipher(key)
	if err != nil {
		return []byte{}, err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return []byte{}, err
	}

	nonce := ciphertext[16 : 16+aesGCM.NonceSize()]
	rawCiphertext := ciphertext[16+aesGCM.NonceSize():]

	plaintext, err := aesGCM.Open(nil, nonce, rawCiphertext, nil)
	if err != nil {
		return []byte{}, err
	}

	return plaintext, nil
}
