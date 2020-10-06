package appstore_sdk

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
)

type PrivateKey struct {
}

func (pk *PrivateKey) LoadFile(path string) ([]byte, error) {
	privateKey, err := ioutil.ReadFile(path)
	return privateKey, err
}

func (pk *PrivateKey) Load(path string) (*ecdsa.PrivateKey, error) {
	data, err := pk.LoadFile(path)
	if err != nil {
		return nil, err
	}
	return pk.ParseP8(data)
}

func (pk *PrivateKey) ParseP8(rawBytes []byte) (*ecdsa.PrivateKey, error) {
	block, err := pk.DecodePem(rawBytes)
	if err != nil {
		return nil, err
	}
	return pk.ParsePKCS8(block.Bytes)
}

func (pk *PrivateKey) ParsePKCS8(rawBytes []byte) (*ecdsa.PrivateKey, error) {
	key, err := x509.ParsePKCS8PrivateKey(rawBytes)
	if err != nil {
		return nil, fmt.Errorf("PrivateKey@ParsePKCS8 wrong parse PKCS8: %v", err)
	}
	switch pk := key.(type) {
	case *ecdsa.PrivateKey:
		return pk, nil
	default:
		return nil, errors.New("PrivateKey@ParsePKCS8: AuthKey must be of type ecdsa.PrivateKey")
	}
}

func (pk *PrivateKey) DecodePem(rawBytes []byte) (*pem.Block, error) {
	block, _ := pem.Decode(rawBytes)
	if block == nil {
		return nil, errors.New("PrivateKey@DecodePem: AuthKey must be a valid .p8 PEM file")
	}
	return block, nil
}
