package rsa

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
)

func GenerateRSA(bits int) (base64EncodedPrivateKeyPem, base64EncodedPublicKeyPem string) {
	serverPrivateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		// todo 에러
		return
	}
	base64EncodedPrivateKeyPem, err = GetBase64EncodedPrivateKeyPem(*serverPrivateKey)
	if err != nil {
		// todo 에러
		return
	}
	base64EncodedPublicKeyPem, err = GetBase64EncodedPublicKeyPem(base64EncodedPrivateKeyPem)
	return
}

func DecryptBase64EncodedPrivateKeyPem(encryptedData []byte, base64EncodedPrivateKeyPem string) (data string, err error) {
	privateKeyPem, err := base64.StdEncoding.DecodeString(base64EncodedPrivateKeyPem)
	if err != nil {
		return
	}
	privateKeyBlock, _ := pem.Decode(privateKeyPem)
	privateKeyAny, err := x509.ParsePKCS8PrivateKey(privateKeyBlock.Bytes)
	if err != nil {
		return
	}
	decryptedBytes, err := rsa.DecryptPKCS1v15(
		rand.Reader,
		privateKeyAny.(*rsa.PrivateKey),
		encryptedData,
	)
	data = string(decryptedBytes)
	return
}

func GetBase64EncodedPrivateKeyPem(privateKey rsa.PrivateKey) (base64EncodedPrivateKeyPem string, err error) {
	privateX509, err := x509.MarshalPKCS8PrivateKey(privateKey)
	if err != nil {
		// todo 에러
		return
	}

	privateBlock := pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: privateX509,
	}
	base64EncodedPrivateKeyPem = base64.StdEncoding.EncodeToString(pem.EncodeToMemory(&privateBlock))
	return
}

func GetBase64EncodedPublicKeyPem(base64EncodedPrivateKeyPem string) (base64EncodedPublicKeyPem string, err error) {
	privateKeyPem, err := base64.StdEncoding.DecodeString(base64EncodedPrivateKeyPem)
	if err != nil {
		return
	}
	privateKeyBlock, _ := pem.Decode(privateKeyPem)
	privateKeyAny, err := x509.ParsePKCS8PrivateKey(privateKeyBlock.Bytes)
	if err != nil {
		return
	}
	privateKey := privateKeyAny.(*rsa.PrivateKey)
	publicX509, err := x509.MarshalPKIXPublicKey(privateKey.PublicKey)
	if err != nil {
		// todo 에러
		return
	}
	publicBlock := pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicX509,
	}
	base64EncodedPublicKeyPem = base64.StdEncoding.EncodeToString(pem.EncodeToMemory(&publicBlock))
	return
}
