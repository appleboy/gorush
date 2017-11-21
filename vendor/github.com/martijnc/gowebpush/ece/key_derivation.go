// Copyright 2016 Martijn Croonen. All rights reserved.
// Use of this source code is governed by the MIT license, a copy of which can
// be found in the LICENSE file.

package ece

import (
	"bytes"
	"crypto/rand"
	"crypto/sha256"
	"encoding/binary"
	"io"

	"golang.org/x/crypto/hkdf"
)

// EncryptionKeys contains the salt, (optional) auth, encryption key, and nonce.
type EncryptionKeys struct {
	preSharedAuth []byte
	cek           []byte
	nonce         []byte
	salt          []byte
	isTest        bool
}

// SetPreSharedAuthSecret sets the (optional) pre-shared authentication secret that
// is to be used during key derivation.
func (ek *EncryptionKeys) SetPreSharedAuthSecret(preSharedAuthSecret []byte) {
	ek.preSharedAuth = preSharedAuthSecret
}

// GetSalt returns the generated salt.
func (ek *EncryptionKeys) GetSalt() []byte {
	return ek.salt
}

// CreateEncryptionKeys derives the encryption key and nonce from the input keying
// material.
func (ek *EncryptionKeys) CreateEncryptionKeys(secret []byte, context []byte) {
	cekInfo := buildInfoData("Content-Encoding: aesgcm", context)
	nonceInfo := buildInfoData("Content-Encoding: nonce", context)
	authInfo := []byte("Content-Encoding: auth\x00")

	// Generate new salt
	if !(ek.isTest && len(ek.salt) == 16) {
		ek.salt = generateSalt()
	}

	var prk []byte
	if len(ek.preSharedAuth) == 0 {
		prk = secret
	} else {
		prk = readHKDF(secret, ek.preSharedAuth, authInfo, 32)
	}

	ek.cek = readHKDF(prk, ek.salt, cekInfo, 16)
	ek.nonce = readHKDF(prk, ek.salt, nonceInfo, 12)
}

// BuildDHContext builds the context from the Diffie-Hellman keys that is needed
// to derive the encryption key and nonce.
// https://tools.ietf.org/html/draft-ietf-httpbis-encryption-encoding-00#section-4.2
func BuildDHContext(clientPublic []byte, serverPublic []byte) []byte {
	var buffer bytes.Buffer
	buffer.Write([]byte("\x00P-256\x00"))
	binary.Write(&buffer, binary.BigEndian, uint16(len(clientPublic)))
	buffer.Write(clientPublic)
	binary.Write(&buffer, binary.BigEndian, uint16(len(serverPublic)))
	buffer.Write(serverPublic)
	return buffer.Bytes()
}

// buildInfoData Merges label and context.
func buildInfoData(label string, context []byte) []byte {
	b := []byte(label)
	return append(b, context...)
}

// readHDKF runs hkdf (SHA-256) with the provided inputs and reads |length| bytes.
func readHKDF(secret, salt, info []byte, length int) []byte {
	hkdf := hkdf.New(sha256.New, secret, salt, info)
	output := make([]byte, length)
	io.ReadFull(hkdf, output)
	return output
}

// GenerateSalt generates 16 (cryptographically secure) random bytes to be used as salt.
func generateSalt() []byte {
	salt := make([]byte, 16)
	rand.Read(salt)
	return salt
}
