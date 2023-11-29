package rsa

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"testing"
)

func TestA(t *testing.T) {
	privateKey, publicKey := GenerateRSA(2048)
	t.Log(privateKey)
	t.Log(publicKey)
}

func TestB(t *testing.T) {
	// base64EncryptedPrivateKeyPem := "LS0tLS1CRUdJTiBQUklWQVRFIEtFWS0tLS0tCk1JSUV2UUlCQURBTkJna3Foa2lHOXcwQkFRRUZBQVNDQktjd2dnU2pBZ0VBQW9JQkFRRFRDc0pMelBYMU9ZeksKcVdseGxmdThzL0JHb1RpcHAzK0wxcjVVeEt3YXUxS3RMdFlUb0VTeXYvUTBBSmFUVC9OVVpKdFh3NW1JWCs0awp4dkluYng5cFovSlBrckplNVQvMlYreUo1ai80a3lWVXVnbmlvQ0M1UXJvM2w2VWEwZXFBdFlTUDJrSDlTRnFMCnhUM3hHNlVlcWQwdXRmblNDTkp5OUUydStNbFVIcVBNdDdYa3lyaStWeWJkWkJSTDNUMzMxZVNlOEhEdjRjc1QKcENkaVJ5RC9OV1R3Q1hHU2pDNGtuTVJTN2JrV3Q1WVFrKyt2a2JHd0g2ejVnc1dRU0hSd3lXSGRabmY5eVJQawpSemNsVmZvU3IwOXRCaDBKZUUyd2NTMTNtaWd5TDQvbnRLK0ZxODBMb3RKZTZDdW9RSm1PMWZIV0crNTNybXRXCnphT3A0b0JmQWdNQkFBRUNnZ0VBVnliQ1JmdEdOeGFsQmF3Z1Z6L1F6WVNoWFFtSEZNaU82M2lxOE9hbkRBTC8KaVhVbUFzVFVtZHliUkQ1WjhFdTcrTVh6UUxNTjNEdUJaTDZqa1pWYUszVTVvNUs1Qm1jMW1zMkVhUTRrck1wWApTRDNyaHlNSllIZ2wzbjlKTjJJb3JTNUlwLzlCOFE3SUZhQkdqQ01XZFN0QW8rR1NZU0l6aGpWTlAweU9GNll2Ckc1MWR6ZlNNbThBVlF1Q0NRSTUyNXd3ZEVhMnFUWmJuVlNPODIrR0JBYjZFUm5tbW9zRmlyYnRQYklZV2NtSXQKWVlBL1h4enZzK0N1OGh1THBmeTNWVGdobERpTWdzREtOUTkyZ2pYYlI2cjlLK2NTbm5VR0RBbGt4blJsVkJUOQpDQmhudUIxVUhiSWpWSm14Tmh6Tm1HRXFkYzFkbVllRjg5aHNLdTVGQVFLQmdRRFlycmxJM0NvRklBbXY5T01qCndHOEM2Y0hsb2EzeHVySXdabG1vWFJZWVYxdWZzRE5vZHQwZVRkNGlnRmNJOEdqVmNEMHZxVEZDTDFKODNYYkMKTGkwMWora3pPVEdYZUU5bjIyVzFnSW8zSWljbW1HNnpGeXBSb1BlcWdBZFZvQWNqZjJ3UHpmWVdLOVpSSW8xSApIU0JSSHlWU3doR3JjU2JPc05WemtCUXRVUUtCZ1FENVZnWmJyRnhZV3NJa1VsQ2NzSURaNmtHcHpZOXJERSszCjdJWXV6NkdETnVCTURPMlkyQnV0R0RocWdmSC9KYkpEbEtIUEZ2dmd4dkk5b09kOEVQdHBzTlFwWVAxQU4wRm0KZ3hOOS91RnRjLzNJMEFMT25VQ01TRHY3Y0dsUmVCckt1UXpyaVZtN2RpMDBYVXo5WGFwdEEwTlN6Mmdqb2NMZQpUTmU1YWlXbXJ3S0JnUUNSMEJKemkyVGppZ202MU9xOU5sUDkrOXcxY1RXaXcrWUJWLzdobzdLb3NyQy9wMDBRCjhWd2ZkbnNUY0pZaE45YjBPdEdSN3lOc3pIZVArMkJkajhGd2wrVndTUDZzVlpQcDZlUnAzWlFsUWVCaXdwVk0KQXp6QkxiYzhqMmRnSVk1c1B1VXpSWUJpcnlGTC9CcGlUcUZldGJ6aUVpTGQ0Z2toY2pKUjBCVWo0UUtCZ0R1bgptVjNWbjBLRWRHNjhRdzA2TFpuMXNrQXVMbUl1VCtSeGNxbEFoQ1FndlFsMURtUERVTnZaTDZaTDhOcGE2b3ZBCncvWUlTYnM5M2pkZ2lBSG5Eai83VHNlWHJQL3dPSFdPWnR0K1gvRVFEMmc4WEcrU3NCN2dLWG0vSzNPU2lOWjMKNzk5cXZVMWN1UFFPWHAxQUM4YUpRVkpPelZ1NUdDYVh2QVZnb0NHQkFvR0FheDYzZGdLS3pHZ281ZGxGcXk3dgpGL2RPMGxmMGRMTC9HZjhrWU94dmxob3Z6THBCVG5OSHdUdmdTdUcyR1d2VERabFNqTUc0N2lmQnFRN3JMb3J4CmIzYXQ3NjQwMTdJVDNqbzQvVEtDTHdzanU0YTdpUjNVUG1WVnV6T1BCdStLTkQ5RmFici9XTXErRHNoSUN5NFoKcXRQQ0ZpUUxIZVVsZnBoalZuU3F0S0E9Ci0tLS0tRU5EIFBSSVZBVEUgS0VZLS0tLS0K"
	// base64EncodedEncryptedData := "WToSsWhE3n7GkBUs9tF0+yEtFtZt3qfWYfiQ8bbstm8zoRONMLJwNoB2iHqwn+QwgqebyaSCD7/o1RKsgslL5Szs2jOnUOGwBVY4uN+iQaA4WxkK3/vnCcGdGFA96oK7n2ABmJooDQ0vdyVmzfGItjDyU7flZbtczOPYf5jj4pWSBgBuK3751PYeTYEXAwDf4ghgh+HIPTR8OjtCRUwq+QbzLgc5NQ8wMgxxP5PwpCqM7cts+arfab17aXCkDzK1lkw9wIJPk7I/4W7Gp0BnljjM6Wq+cHbnnVy8+KacEFk3hiEzAGoAJw49Li2Y+FtIzqvYr4C/PR2wDZLAWkjIdQ=="
	base64EncryptedPrivateKeyPem := "LS0tLS1CRUdJTiBQUklWQVRFIEtFWS0tLS0tCk1JSUVvd0lCQUFLQ0FRRUFtd2lWUWxBSVdseThqSWx2ZlQyVzhPaXJlRjFsalZnUXdEWkxnS0RKTVZLdHluUnMKY2pUR21sVjRxOVRkOHAxbjZwY1ZhWVV1ZEZkSXlVaTdSc1dUcTQxdzdrOWJ0WU5YSFJTUzRYTUY3WWRrb0ZIcgowWkloV1p1R0hVdmZvYVZ1WWpKbVdvZWZzU2R4RDdzTXNOektnOU50S3Z5YjBZaURqV0ZQK2ZBOFROb2NJQnlUCjVQb0pHd2JGMTlmejdDUmswQjVmcVlkRVhWYk8zejA4NDRIaHlIZHQySHJuUXFOdnlxWGgwai85L1grc3FUd0EKVjVFQXRlQXZlS1k0bCszZlVRWU1ObTV4NUtDcHVRQnBiYW4wVG9tU2JUNHB1V2NSNVQ4dnB5QU9iK1Z3MkorbgpjSEtZUkp6cGM4VmJWQkhWYlJZNGtyQTE2WkRJZHd4TUVSNzVzUUlEQVFBQkFvSUJBSEZndllkWmNFSmxQNTZuCmc5bGJpVjkxU0xFbDNIYVp5T1pJenI2cTEzZ3l0dFl3Zm42bkNmY2tPck92WVFGbEMvSUx1NlVIdVFsc0s5YkkKVWUxWDRMNjlHYkd1WjZzKytoNVNSWGlLM2ZMdjZTODF1ZmZ4bm1JaDB0cnArZ29GS2N5MmZ4UUZ5MHVMUkt2MApaenVRdktuYm5TS3F5bWxhSkpyb1o4eXBXKy8zRHVhUjNLa21ZenBZcTdkbFl4L3JvTEx3TXFhOERQZjdJR0YrCmtZVmpNWkduK3NYb0hucTI0YWs4elptMU9STmRSbWNkL1h1aEhzNXllZ0ZndHRpb3BhdWk0UVhlbWU1ait0ODgKWmZDRXpJTCtWeWQ2VGUrODFlbGU0SHVlcC84UXU5bm1pa00xNXFTbFZ6aUEwMWNRT1gwbmFqQXFXVGwyR0VhWgpsZFRSYUlFQ2dZRUF4bklBbDBNYjhjZVU4RFBkb1B1S2J3MlBaV1htMmtPaXkyZVVUZHdCaHR1NEtWeVZUYWNGCkROcjQyT2hsQjhFU0RXV01jRlEyaGh6VEtQZTQ5M0I3OWsrRk5OdXB0NGY3TGtSbGxKOUtjRElDbytSUWk5QjQKanJ5akFaOSszdkpCTXpBUmNlcmhFQVpkclA4R1U0Z09ndnpEOVpoSWN0WDVLYVhiaG5JQTgvVUNnWUVBeC85aApFTE1kVHl2UCt4S0VNK09UUDNwc2w1ckM5citnVUE4YWRBZHUrQ1hwb2YxNzNGYVVjeFo3MExBUmVMQ2ZPUjFLClJCMVhXTGIwbjZ3OXdoREI5QzJ2R0FDUTJwalNqblFZajIySDA3aE81bmh0STZWeDdBZU51aGFCS1MzR0hGdHAKNkkzbWJhYW80NENLSnVmUHlTaW9NbHV0MVVKdHh4WlJwcVo0bFUwQ2dZQkRxMlJnTUZZN0lGKzllVEkrVHdocAp6aW56M2NmVzdmRjVneHlqWUQ1MUNqL0dldmUrdnhHTnpLM0c2WmhxQ0lCcmFSRjJ6SUM3UXVFT3UwcXJ0Z1BQCmZsWmlLL1czeS9ydlVBdURucjMzZkZaQ0pQMENjRmhyOUE5eDRqMlZNTzdpMGFWaFAwTDMvL081cGswMTl0TmIKLzIyak5nYWVnakN3N3dubzByOVljUUtCZ1FDRUorTFNFQ21yNlQyVG9OYXZHeEJ2Lzh2R0drUCtuUDJvS1hMSApmU244MjhDWWhGSFdkWGUrM1BxUUdlaFJvaFIwdjBBVjVuV1RiOElSeU5VK2FhaGdXOWU1dXBYOWZNS1YwMjF2ClNXZFpwZzZ0Q0tMRnpVdU1OaW1XNzh3RmsvTzNSVWlrblMvSkFUblVxMW1lLzhzMEY2T2RNeXVaSWo0OE1pbGUKUDk2cWVRS0JnR0srWW9zdEVTRGxxeGMrL1pzNk0wL1RqZ3N4NTVTU2NEaTFSRGtSMTVxb0pSK01ETUN3c28vZgoxTU5EZng3c080bS9iQXRFVldnb3IvSkF1aGxGU0tQNFpGeWdEZ25HSjcxdkk2anZEdmtQNHBldG9QYXpPei9ZCmJUNnQxRThYZTZJK2FDYm9Yc1dIdlh4dkNoNjlUbGxrSXZkTlU2aVJVVUxoMHYxVFNwaSsKLS0tLS1FTkQgUFJJVkFURSBLRVktLS0tLQo="
	base64EncodedEncryptedData2 := "WuOW/OTQryf40+nVJhKPta8kZAFidlCYTh5yRZcfnriwQwV5OsMSHWGuAm70sDF5aFmo3+PoVMyRxDM6eBAdrTUQd8MycNFEIaxARW+fmkYkyVANL2K9i1Jbsq9tFVX2ThEjftzyRexFwZBYWqzGsVT9D2cdyQ+mkSAf4PI8oTtOOrIATWNwFlAqBn8cTCzGpATbDdpc7gfGSBh4f24ETWEabEEOfV7yn/GR3SZqXRzs8Vu/sR7EkO9/zT6frAl7dffsq0qfJc+OTk0zka9nnWcz3o43Sa3iJIIeQFyhx1kRRwmd/nsBUsyXD/jFxqjcEkA23oC0NgNsCYehNE0pCA=="
	privateKey, err := DecodeBase64PrivateKeyPem(base64EncryptedPrivateKeyPem)
	if err != nil {
		t.Error(err.Error())
		return
	}

	publicKey := x509.MarshalPKCS1PublicKey(&privateKey.PublicKey)
	t.Log(string(publicKey))
	publicBlock := pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicKey,
	}
	t.Log(string(pem.EncodeToMemory(&publicBlock)))

	ciphertext, err := rsa.EncryptPKCS1v15(
		rand.Reader,
		&privateKey.PublicKey,
		[]byte("원피스는 사실 없다."),
	)
	if err != nil {
		t.Error(err.Error())
		return
	}

	plaintext2, err := rsa.DecryptPKCS1v15(
		rand.Reader,
		privateKey,
		ciphertext,
	)
	t.Log(string(plaintext2))

	encryptedData, err := base64.StdEncoding.DecodeString(base64EncodedEncryptedData2)
	if err != nil {
		t.Error(err.Error())
		return
	}

	plaintext, err := rsa.DecryptPKCS1v15(
		rand.Reader,
		privateKey,
		encryptedData,
	)

	if err != nil {
		t.Error(err.Error())
		return
	}
	t.Log(string(plaintext))
}

func TestC(t *testing.T) {
	serverPrivateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		t.Error(err.Error())
		return
	}
	x509PrivateKey := x509.MarshalPKCS1PrivateKey(serverPrivateKey)
	pemPrivateKey := pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: x509PrivateKey,
	}
	base64EncodedPrivateKeyPem := base64.StdEncoding.EncodeToString(pem.EncodeToMemory(&pemPrivateKey))
	println(base64EncodedPrivateKeyPem)

	serverPublicKey := &serverPrivateKey.PublicKey
	x509PublicKey := x509.MarshalPKCS1PublicKey(serverPublicKey)
	pemPublicKey := pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: x509PublicKey,
	}
	base64EncodedPublicKeyPem := base64.StdEncoding.EncodeToString(pem.EncodeToMemory(&pemPublicKey))
	println(base64EncodedPublicKeyPem)
}

func TestD(t *testing.T) {

}

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
	publicX509 := x509.MarshalPKCS1PublicKey(serverPublicKey)
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
