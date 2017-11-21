// Copyright 2016 Martijn Croonen. All rights reserved.
// Use of this source code is governed by the MIT license, a copy of which can
// be found in the LICENSE file.

package ece

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/binary"
	"errors"
)

// Encrypt encrypts |plaintext| using AEAD_AES_GCM_128 with the keys in |keys|
// adding |paddingLength| bytes of padding.
func Encrypt(plaintext []byte, keys *EncryptionKeys, paddingLength int) ([]byte, error) {
	if paddingLength < 0 || paddingLength > 65535 {
		return nil, errors.New("Padding should be between 0 and 65535.")
	}

	aes, err := aes.NewCipher(keys.cek)
	if err != nil {
		return nil, err
	}

	aesgcm, err := cipher.NewGCM(aes)
	if err != nil {
		return nil, err
	}

	record := make([]byte, 2+paddingLength+len(plaintext))
	binary.BigEndian.PutUint16(record, uint16(paddingLength))
	copy(record[2+paddingLength:], plaintext)

	var auth []byte
	return aesgcm.Seal(nil, keys.nonce, record, auth), nil
}
