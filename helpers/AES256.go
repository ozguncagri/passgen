package helpers

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"io"

	"golang.org/x/crypto/argon2"
)

// deriveKeyArgon2id derives a 32-byte AES key using Argon2id with given parameters.
func deriveKeyArgon2id(password, salt []byte, time uint32, memoryKiB uint32, parallelism uint8) []byte {
	return argon2.IDKey(password, salt, time, memoryKiB, parallelism, 32)
}

// Encrypt encrypts plaintext using a password and returns Base64 string
func Encrypt(plaintext []byte, password string) ([]byte, error) {
	// Generate 32 byte salt
	salt := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, salt); err != nil {
		return []byte{}, err
	}

	// Argon2 parameters
	timeCost := uint32(3)          // 3 Iterations
	memoryKiB := uint32(64 * 1024) // 64MB Memory Usage
	parallelism := uint8(4)        // 4 Parallelism

	key := deriveKeyArgon2id([]byte(password), salt, timeCost, memoryKiB, parallelism)

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

	// Construct blob: salt | time | memory | parallelism | nonce | ciphertext
	buf := new(bytes.Buffer)
	if _, err := buf.Write(salt); err != nil {
		return []byte{}, err
	}
	if err := binary.Write(buf, binary.BigEndian, timeCost); err != nil {
		return []byte{}, err
	}
	if err := binary.Write(buf, binary.BigEndian, memoryKiB); err != nil {
		return []byte{}, err
	}
	if err := buf.WriteByte(parallelism); err != nil {
		return []byte{}, err
	}
	if _, err := buf.Write(nonce); err != nil {
		return []byte{}, err
	}
	if _, err := buf.Write(ciphertext); err != nil {
		return []byte{}, err
	}

	return buf.Bytes(), nil
}

// Decrypt decrypts a Base64-encoded string using a password
func Decrypt(data []byte, password string) ([]byte, error) {
	// salt (32 bytes) + time (4 bytes uint32) + memory (4 bytes uint32) + parallelism (1 bytes uint8) + nonce (12 bytes)
	minSize := 32 + 4 + 4 + 1 + 12
	if len(data) < minSize {
		return []byte{}, fmt.Errorf("invalid encrypted data")
	}

	offset := 0

	salt := data[offset : offset+32]
	offset += 32

	timeCost := binary.BigEndian.Uint32(data[offset : offset+4])
	offset += 4

	memoryKiB := binary.BigEndian.Uint32(data[offset : offset+4])
	offset += 4

	parallelism := data[offset]
	offset += 1

	nonce := data[offset : offset+12]
	offset += 12

	ciphertext := data[offset:]

	key := deriveKeyArgon2id([]byte(password), salt, timeCost, memoryKiB, parallelism)

	block, err := aes.NewCipher(key)
	if err != nil {
		return []byte{}, err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return []byte{}, err
	}

	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return []byte{}, err
	}

	return plaintext, nil
}
