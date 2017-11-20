// Copyright 2016 Martijn Croonen. All rights reserved.
// Use of this source code is governed by the MIT license, a copy of which can
// be found in the LICENSE file.

package ece

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

// CryptoKeyHeader is a small struct that can be used to define the values in
// the "Crypto-Key" header. https://tools.ietf.org/html/draft-ietf-httpbis-encryption-encoding-00#section-4
type CryptoKeyHeader struct {
	keyid  string
	aesgcm string
	dh     string
}

// EncryptionHeader is a small struct that can be used to define the values in
// the "Encryption" header. https://tools.ietf.org/html/draft-ietf-httpbis-encryption-encoding-00#section-3
type EncryptionHeader struct {
	keyid string
	salt  string
	rs    int
}

// CreateRequest creates a new http.Request with the ECE headers set correctly based on the
// |cryptoKey| and |encryption| parameters.
func CreateRequest(client http.Client, url string, data []byte, cryptoKey *CryptoKeyHeader, encryption *EncryptionHeader, ttl int) *http.Request {
	r, _ := http.NewRequest("POST", url, bytes.NewBuffer(data))
	// Required by Firefox' push server but breaks Chrome's push server.
	r.Header.Add("Content-Encoding", "aesgcm")
	r.Header.Add("Crypto-Key", cryptoKey.toString())
	r.Header.Add("Encryption", encryption.toString())
	r.Header.Add("TTL", strconv.Itoa(ttl))
	return r
}

// SetDHKey sets the Diffie-Hellman (public) key.
// https://tools.ietf.org/html/draft-ietf-httpbis-encryption-encoding-00#section-4.2
func (ckh *CryptoKeyHeader) SetDHKey(publicKey []byte) {
	ckh.dh = strings.TrimRight(base64.URLEncoding.EncodeToString(publicKey), "=")
}

// SetExplicitKey sets the explicit encryption key.
// https://tools.ietf.org/html/draft-ietf-httpbis-encryption-encoding-00#section-4.1
func (ckh *CryptoKeyHeader) SetExplicitKey(key []byte) {
	ckh.aesgcm = base64.URLEncoding.EncodeToString(key)
}

// SetKeyID sets the keyid.
// https://tools.ietf.org/html/draft-ietf-httpbis-encryption-encoding-00#section-4.3
func (ckh *CryptoKeyHeader) SetKeyID(keyID string) {
	ckh.keyid = keyID
}

// toString formats and returns the output string for the crypto-key header.
func (ckh *CryptoKeyHeader) toString() string {
	var items []string
	if len(ckh.keyid) != 0 {
		items = append(items, fmt.Sprintf("keyid=%s", ckh.keyid))
	}
	if len(ckh.dh) != 0 {
		items = append(items, fmt.Sprintf("dh=%s", ckh.dh))
	}
	if len(ckh.aesgcm) != 0 {
		items = append(items, fmt.Sprintf("aesgcm=%s", ckh.aesgcm))
	}
	return strings.Join(items, ";")
}

// SetSalt sets the encryption salt that should be used for key derivation.
// https://tools.ietf.org/html/draft-ietf-httpbis-encryption-encoding-00#section-3.1
func (eh *EncryptionHeader) SetSalt(salt []byte) {
	eh.salt = strings.TrimRight(base64.URLEncoding.EncodeToString(salt), "=")
}

// SetRecordSize sets the record size for the message.
// https://tools.ietf.org/html/draft-ietf-httpbis-encryption-encoding-00#section-3.1
func (eh *EncryptionHeader) SetRecordSize(size int) {
	eh.rs = size
}

// SetKeyID set the keyid.
func (eh *EncryptionHeader) SetKeyID(keyID string) {
	eh.keyid = keyID
}

// toString formats and returns the output string for the encryption header.
func (eh *EncryptionHeader) toString() string {
	var items []string
	if len(eh.keyid) != 0 {
		items = append(items, fmt.Sprintf("keyid=%s", eh.keyid))
	}
	if eh.rs != 0 {
		items = append(items, fmt.Sprintf("rs=%d", eh.rs))
	}
	if len(eh.salt) != 0 {
		items = append(items, fmt.Sprintf("salt=%s", eh.salt))
	}
	return strings.Join(items, ";")
}
