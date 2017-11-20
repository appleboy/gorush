// Copyright 2016 Martijn Croonen. All rights reserved.
// Use of this source code is governed by the MIT license, a copy of which can
// be found in the LICENSE file.

package webpush

import (
	"crypto/elliptic"
	"crypto/rand"
	"errors"
	"math/big"
)

var curve = elliptic.P256()

// KeyPair contains the public and private keys.
type KeyPair struct {
	privateKey []byte
	PublicKey  []byte
}

// GenerateKeys generates a fresh P-256 keypair.
func (k *KeyPair) GenerateKeys() error {
	var err error
	var x, y *big.Int

	k.privateKey, x, y, err = elliptic.GenerateKey(curve, rand.Reader)
	if err != nil {
		return err
	}

	k.PublicKey = elliptic.Marshal(curve, x, y)
	return nil
}

// SetPublicKey sets the public key part of the KeyPair. This should be an uncompressed
// NIST P-256 point.
func (k *KeyPair) SetPublicKey(publicKey []byte) error {
	if len(publicKey) != 65 {
		return errors.New("Incorrect key length. Use an uncompressed NIST P-256 point (65 bytes).")
	}
	k.PublicKey = publicKey
	return nil
}

// SetPrivateKey sets the private key part of the KeyPair.
func (k *KeyPair) SetPrivateKey(privateKey []byte) error {
	if len(privateKey) != 32 {
		return errors.New("Incorrect key length. The private key should be 32 bytes.")
	}
	k.privateKey = privateKey
	return nil
}

// CalculateSecret calculates the shared secret between the two KeyPairs using ECDH.
func CalculateSecret(sp *KeyPair, rp *KeyPair) []byte {
	x, y := elliptic.Unmarshal(curve, rp.PublicKey)
	result, _ := curve.ScalarMult(x, y, sp.privateKey)
	return result.Bytes()
}
