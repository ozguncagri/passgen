package storage

import (
	"bytes"
	"encoding/gob"
	"io/ioutil"
	"passgen/helpers"
)

// Save function encodes and encrypts storage collection and saves it to the file
func Save(storagePassword string, memoryStorage *PassgenStorage, storagePath string) error {
	// Define encode buffer
	var buf bytes.Buffer

	// Define encoder
	encoder := gob.NewEncoder(&buf)

	// Serialize store object
	err := encoder.Encode(memoryStorage)
	if err != nil {
		return err
	}

	// Encrypt encoded storage
	encryptedStorage, err := helpers.AES256Encrypt(storagePassword, buf.Bytes())
	if err != nil {
		return err
	}

	// Write encrypted storage to file
	err = ioutil.WriteFile(storagePath, encryptedStorage, 0777)
	if err != nil {
		return err
	}

	return nil
}
