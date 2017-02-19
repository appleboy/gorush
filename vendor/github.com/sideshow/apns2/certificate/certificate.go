// Package certificate contains functions to load an Apple APNs PKCS#12
// or PEM certificate from either an in memory byte array or a local file.
package certificate

import (
	"crypto"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"io/ioutil"
	"strings"

	"golang.org/x/crypto/pkcs12"
)

// Possible errors when parsing a certificate.
var (
	ErrFailedToDecryptKey           = errors.New("failed to decrypt private key")
	ErrFailedToParsePKCS1PrivateKey = errors.New("failed to parse PKCS1 private key")
	ErrFailedToParseCertificate     = errors.New("failed to parse certificate PEM data")
	ErrNoPrivateKey                 = errors.New("no private key")
	ErrNoCertificate                = errors.New("no certificate")
)

// FromP12File loads a PKCS#12 certificate from a local file and returns a
// tls.Certificate.
//
// Use "" as the password argument if the pem certificate is not password
// protected.
func FromP12File(filename string, password string) (tls.Certificate, error) {
	p12bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return tls.Certificate{}, err
	}
	return FromP12Bytes(p12bytes, password)
}

// FromP12Bytes loads a PKCS#12 certificate from an in memory byte array and
// returns a tls.Certificate.
//
// Use "" as the password argument if the PKCS#12 certificate is not password
// protected.
func FromP12Bytes(bytes []byte, password string) (tls.Certificate, error) {
	key, cert, err := pkcs12.Decode(bytes, password)
	if err != nil {
		return tls.Certificate{}, err
	}
	return tls.Certificate{
		Certificate: [][]byte{cert.Raw},
		PrivateKey:  key,
		Leaf:        cert,
	}, nil
}

// FromPemFile loads a PEM certificate from a local file and returns a
// tls.Certificate. This function is similar to the crypto/tls LoadX509KeyPair
// function, however it supports PEM files with the cert and key combined
// in the same file, as well as password protected key files which are both
// common with APNs certificates.
//
// Use "" as the password argument if the PEM certificate is not password
// protected.
func FromPemFile(filename string, password string) (tls.Certificate, error) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return tls.Certificate{}, err
	}
	return FromPemBytes(bytes, password)
}

// FromPemBytes loads a PEM certificate from an in memory byte array and
// returns a tls.Certificate. This function is similar to the crypto/tls
// X509KeyPair function, however it supports PEM files with the cert and
// key combined, as well as password protected keys which are both common with
// APNs certificates.
//
// Use "" as the password argument if the PEM certificate is not password
// protected.
func FromPemBytes(bytes []byte, password string) (tls.Certificate, error) {
	var cert tls.Certificate
	var block *pem.Block
	for {
		block, bytes = pem.Decode(bytes)
		if block == nil {
			break
		}
		if block.Type == "CERTIFICATE" {
			cert.Certificate = append(cert.Certificate, block.Bytes)
		}
		if block.Type == "PRIVATE KEY" || strings.HasSuffix(block.Type, "PRIVATE KEY") {
			key, err := unencryptPrivateKey(block, password)
			if err != nil {
				return tls.Certificate{}, err
			}
			cert.PrivateKey = key
		}
	}
	if len(cert.Certificate) == 0 {
		return tls.Certificate{}, ErrNoCertificate
	}
	if cert.PrivateKey == nil {
		return tls.Certificate{}, ErrNoPrivateKey
	}
	if c, e := x509.ParseCertificate(cert.Certificate[0]); e == nil {
		cert.Leaf = c
	}
	return cert, nil
}

func unencryptPrivateKey(block *pem.Block, password string) (crypto.PrivateKey, error) {
	if x509.IsEncryptedPEMBlock(block) {
		bytes, err := x509.DecryptPEMBlock(block, []byte(password))
		if err != nil {
			return nil, ErrFailedToDecryptKey
		}
		return parsePrivateKey(bytes)
	}
	return parsePrivateKey(block.Bytes)
}

func parsePrivateKey(bytes []byte) (crypto.PrivateKey, error) {
	key, err := x509.ParsePKCS1PrivateKey(bytes)
	if err != nil {
		return nil, ErrFailedToParsePKCS1PrivateKey
	}
	return key, nil
}
