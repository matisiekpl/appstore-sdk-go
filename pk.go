package appstore

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
)

//PrivateKey private key handler
type PrivateKey struct {
}

//LoadFromFile Load private key from file
func (pk *PrivateKey) LoadFromFile(path string) ([]byte, error) {
	return readFile(path)
}

//LoadFromContent Load private key from string
func (pk *PrivateKey) LoadFromContent(content string) ([]byte, error) {
	return []byte(content), nil
}

//LoadData Load private key from string or file path
func (pk *PrivateKey) LoadData(pathOrContent string) ([]byte, error) {
	if fileExists(pathOrContent) {
		return pk.LoadFromFile(pathOrContent)
	} else {
		return pk.LoadFromContent(pathOrContent)
	}
}

//Load and generate private key
func (pk *PrivateKey) Load(path string) (*ecdsa.PrivateKey, error) {
	data, err := pk.LoadData(path)
	if err != nil {
		return nil, err
	}
	return pk.ParseP8(data)
}

//ParseP8 Parse private key .p8
func (pk *PrivateKey) ParseP8(rawBytes []byte) (*ecdsa.PrivateKey, error) {
	block, err := pk.DecodePem(rawBytes)
	if err != nil {
		return nil, err
	}
	return pk.ParsePKCS8(block.Bytes)
}

//ParsePKCS8 Parse PKCS private key .p8
func (pk *PrivateKey) ParsePKCS8(rawBytes []byte) (*ecdsa.PrivateKey, error) {
	key, err := x509.ParsePKCS8PrivateKey(rawBytes)
	if err != nil {
		return nil, fmt.Errorf("PrivateKey.ParsePKCS8 wrong parse PKCS8: %v", err)
	}
	switch pk := key.(type) {
	case *ecdsa.PrivateKey:
		return pk, nil
	default:
		return nil, errors.New("PrivateKey.ParsePKCS8: AuthKey must be of type ecdsa.PrivateKey")
	}
}

//DecodePem Decode private key pem
func (pk *PrivateKey) DecodePem(rawBytes []byte) (*pem.Block, error) {
	block, _ := pem.Decode(rawBytes)
	if block == nil {
		return nil, errors.New("PrivateKey.DecodePem: AuthKey must be a valid .p8 PEM file")
	}
	return block, nil
}
