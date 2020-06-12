package appstore

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"
)

type PrivateKey struct {
}

func (pk *PrivateKey) Load(path string) (string, error) {
	privateKey, err := ioutil.ReadFile(path)
	return string(privateKey), err
}

func (pk *PrivateKey) ParseP8(txt string) (*ecdsa.PrivateKey, error) {
	rawByte := []byte(txt)
	block, _ := pem.Decode(rawByte)
	if block == nil {
		return nil, ErrAuthKeyNotPem
	}
	key, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	switch pk := key.(type) {
	case *ecdsa.PrivateKey:
		return pk, nil
	default:
		return nil, ErrAuthKeyNotECDSA
	}
}
