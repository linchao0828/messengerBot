package utils

import (
	"crypto"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"path/filepath"
)

func NewPrivateKeyFromPemFile(pemFile string) (*rsa.PrivateKey, error) {
	keyPath, err := filepath.Abs(pemFile)
	if err != nil {
		return nil, err
	}
	b, err := ioutil.ReadFile(keyPath)
	if err != nil {
		return nil, err
	}
	block, _ := pem.Decode(b)
	if block == nil {
		return nil, errors.New("no PEM data is found")
	}

	var pk interface{}
	switch block.Type {
	case "RSA PRIVATE KEY":
		pk, err = x509.ParsePKCS1PrivateKey(block.Bytes)
	case "PRIVATE KEY":
		pk, err = x509.ParsePKCS8PrivateKey(block.Bytes)
	}

	if err != nil {
		return nil, err
	}

	return pk.(*rsa.PrivateKey), nil
}

func NewPublicKeyFromDerFile(pemFile string) (*rsa.PublicKey, error) {
	keyPath, err := filepath.Abs(pemFile)

	if err != nil {
		return nil, err
	}

	b, err := ioutil.ReadFile(keyPath)

	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(b)

	if block == nil {
		return nil, errors.New("no PEM data is found")
	}

	cert, err := x509.ParseCertificate(block.Bytes)

	if err != nil {
		return nil, err
	}

	return cert.PublicKey.(*rsa.PublicKey), nil
}

func ParsePKCS8PrivateKeyFromString(privateKeyStr string) (*rsa.PrivateKey, error) {
	decodedKey, err := base64.StdEncoding.DecodeString(privateKeyStr)
	if err != nil {
		return nil, err
	}

	parsedKey, err := x509.ParsePKCS8PrivateKey(decodedKey)
	if err != nil {
		return nil, err
	}

	key, ok := parsedKey.(*rsa.PrivateKey)
	if !ok {
		return nil, fmt.Errorf("key is not a RSA private key")
	}

	return key, nil
}

func ParseCertificateFromString(publicKeyStr string) (*rsa.PublicKey, error) {
	decodedKey, err := base64.StdEncoding.DecodeString(publicKeyStr)
	if err != nil {
		return nil, err
	}

	parsedKey, err := x509.ParsePKIXPublicKey(decodedKey)
	if err != nil {
		return nil, err
	}

	key, ok := parsedKey.(*rsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("key is not a RSA public key")
	}

	return key, nil
}

func RsaEncrypt(pubKey *rsa.PublicKey, origData []byte) ([]byte, error) {
	return rsa.EncryptPKCS1v15(rand.Reader, pubKey, origData)
}

func RsaDecrypt(prvKey *rsa.PrivateKey, ciphertext []byte) ([]byte, error) {
	return rsa.DecryptPKCS1v15(rand.Reader, prvKey, ciphertext)
}

func RsaSign(prvKey *rsa.PrivateKey, origData []byte) ([]byte, error) {
	h := crypto.SHA1.New()
	h.Write(origData)
	return rsa.SignPKCS1v15(rand.Reader, prvKey, crypto.SHA1, h.Sum(nil))
}

func RsaVerify(pubKey *rsa.PublicKey, origData, sign []byte) error {
	h := crypto.SHA1.New()
	h.Write(origData)
	return rsa.VerifyPKCS1v15(pubKey, crypto.SHA1, h.Sum(nil), sign)
}

func RsaSignWithMD5(prvKey *rsa.PrivateKey, origData []byte) ([]byte, error) {
	hash := md5.New()
	hash.Write(origData)
	return rsa.SignPKCS1v15(rand.Reader, prvKey, crypto.MD5, hash.Sum(nil))
}
