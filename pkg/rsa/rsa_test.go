package rsa

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"testing"
)

func TestRSAPrivateKeyEncodeDecode(t *testing.T) {
	serverPrivateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		t.Error(err.Error())
		return
	}
	_, err = x509.MarshalPKCS8PrivateKey(serverPrivateKey)
	if err != nil {
		t.Error(err.Error())
		return
	}
	t.Log("정상")
}

func TestRSAEncodeDecode(t *testing.T) {
	// 서버에서 키 생성
	serverPrivateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		t.Error(err.Error())
		return
	}
	// 개인키는 저장한다.

	serverPublicKey := &serverPrivateKey.PublicKey

	data := "원피스는 사실 없다."

	println("원본 = ", serverPublicKey)
	// (서버 > 웹): 공개 키 전달
	// x509.MarshalPKCS1PublicKey(pub)
	publicX509, _ := x509.MarshalPKIXPublicKey(serverPublicKey)
	println("원본(마샬링) = ", publicX509)
	// publicX509String := string(publicX509)
	publicX509EncodedString := base64.StdEncoding.EncodeToString(publicX509)
	println("원본(인코딩) = ", publicX509)
	println("-----")

	publicX509DecodedBytes, _ := base64.StdEncoding.DecodeString(publicX509EncodedString)
	println("(디코딩) = ", publicX509DecodedBytes)
	publicKey, _ := x509.ParsePKIXPublicKey(publicX509DecodedBytes)
	println("(파싱) = ", publicX509DecodedBytes)
	back := publicKey.(*rsa.PublicKey)
	println("(복구) = ", back)
	ciphertext, err := rsa.EncryptPKCS1v15(
		rand.Reader,
		back,
		[]byte(data),
	)
	plaintext, err := rsa.DecryptPKCS1v15(
		rand.Reader,
		serverPrivateKey,
		ciphertext,
	)

	t.Log(string(plaintext))
}

func TestRSA(t *testing.T) {
	// (웹 > 서버): 클라이언트에서 공개키 요청

	// 서버에서 키 생성
	serverPrivateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		t.Error(err.Error())
		return
	}
	// 개인키는 저장한다.

	serverPublicKey := &serverPrivateKey.PublicKey
	// (서버 > 웹): 공개 키 전달
	// x509.MarshalPKCS1PublicKey(pub)
	publicX509, _ := x509.MarshalPKIXPublicKey(serverPublicKey)
	// publicX509String := string(publicX509)
	publicX509EncodedString := base64.StdEncoding.EncodeToString(publicX509)
	println(publicX509EncodedString)
	// 클라이언트에서는 데이터를 암호화한다.
	data := "원피스는 사실 없다."
	ciphertext, err := rsa.EncryptPKCS1v15(
		rand.Reader,
		serverPublicKey,
		[]byte(data),
	)

	// (웹 > 서버) 암호화된 데이터를 전송한다.
	plaintext, err := rsa.DecryptPKCS1v15(
		rand.Reader,
		serverPrivateKey,
		ciphertext,
	)

	t.Log(string(plaintext))
}
