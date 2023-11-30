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
	// base64EncryptedPrivateKeyPem := "LS0tLS1CRUdJTiBQUklWQVRFIEtFWS0tLS0tCk1JSUVvd0lCQUFLQ0FRRUFtd2lWUWxBSVdseThqSWx2ZlQyVzhPaXJlRjFsalZnUXdEWkxnS0RKTVZLdHluUnMKY2pUR21sVjRxOVRkOHAxbjZwY1ZhWVV1ZEZkSXlVaTdSc1dUcTQxdzdrOWJ0WU5YSFJTUzRYTUY3WWRrb0ZIcgowWkloV1p1R0hVdmZvYVZ1WWpKbVdvZWZzU2R4RDdzTXNOektnOU50S3Z5YjBZaURqV0ZQK2ZBOFROb2NJQnlUCjVQb0pHd2JGMTlmejdDUmswQjVmcVlkRVhWYk8zejA4NDRIaHlIZHQySHJuUXFOdnlxWGgwai85L1grc3FUd0EKVjVFQXRlQXZlS1k0bCszZlVRWU1ObTV4NUtDcHVRQnBiYW4wVG9tU2JUNHB1V2NSNVQ4dnB5QU9iK1Z3MkorbgpjSEtZUkp6cGM4VmJWQkhWYlJZNGtyQTE2WkRJZHd4TUVSNzVzUUlEQVFBQkFvSUJBSEZndllkWmNFSmxQNTZuCmc5bGJpVjkxU0xFbDNIYVp5T1pJenI2cTEzZ3l0dFl3Zm42bkNmY2tPck92WVFGbEMvSUx1NlVIdVFsc0s5YkkKVWUxWDRMNjlHYkd1WjZzKytoNVNSWGlLM2ZMdjZTODF1ZmZ4bm1JaDB0cnArZ29GS2N5MmZ4UUZ5MHVMUkt2MApaenVRdktuYm5TS3F5bWxhSkpyb1o4eXBXKy8zRHVhUjNLa21ZenBZcTdkbFl4L3JvTEx3TXFhOERQZjdJR0YrCmtZVmpNWkduK3NYb0hucTI0YWs4elptMU9STmRSbWNkL1h1aEhzNXllZ0ZndHRpb3BhdWk0UVhlbWU1ait0ODgKWmZDRXpJTCtWeWQ2VGUrODFlbGU0SHVlcC84UXU5bm1pa00xNXFTbFZ6aUEwMWNRT1gwbmFqQXFXVGwyR0VhWgpsZFRSYUlFQ2dZRUF4bklBbDBNYjhjZVU4RFBkb1B1S2J3MlBaV1htMmtPaXkyZVVUZHdCaHR1NEtWeVZUYWNGCkROcjQyT2hsQjhFU0RXV01jRlEyaGh6VEtQZTQ5M0I3OWsrRk5OdXB0NGY3TGtSbGxKOUtjRElDbytSUWk5QjQKanJ5akFaOSszdkpCTXpBUmNlcmhFQVpkclA4R1U0Z09ndnpEOVpoSWN0WDVLYVhiaG5JQTgvVUNnWUVBeC85aApFTE1kVHl2UCt4S0VNK09UUDNwc2w1ckM5citnVUE4YWRBZHUrQ1hwb2YxNzNGYVVjeFo3MExBUmVMQ2ZPUjFLClJCMVhXTGIwbjZ3OXdoREI5QzJ2R0FDUTJwalNqblFZajIySDA3aE81bmh0STZWeDdBZU51aGFCS1MzR0hGdHAKNkkzbWJhYW80NENLSnVmUHlTaW9NbHV0MVVKdHh4WlJwcVo0bFUwQ2dZQkRxMlJnTUZZN0lGKzllVEkrVHdocAp6aW56M2NmVzdmRjVneHlqWUQ1MUNqL0dldmUrdnhHTnpLM0c2WmhxQ0lCcmFSRjJ6SUM3UXVFT3UwcXJ0Z1BQCmZsWmlLL1czeS9ydlVBdURucjMzZkZaQ0pQMENjRmhyOUE5eDRqMlZNTzdpMGFWaFAwTDMvL081cGswMTl0TmIKLzIyak5nYWVnakN3N3dubzByOVljUUtCZ1FDRUorTFNFQ21yNlQyVG9OYXZHeEJ2Lzh2R0drUCtuUDJvS1hMSApmU244MjhDWWhGSFdkWGUrM1BxUUdlaFJvaFIwdjBBVjVuV1RiOElSeU5VK2FhaGdXOWU1dXBYOWZNS1YwMjF2ClNXZFpwZzZ0Q0tMRnpVdU1OaW1XNzh3RmsvTzNSVWlrblMvSkFUblVxMW1lLzhzMEY2T2RNeXVaSWo0OE1pbGUKUDk2cWVRS0JnR0srWW9zdEVTRGxxeGMrL1pzNk0wL1RqZ3N4NTVTU2NEaTFSRGtSMTVxb0pSK01ETUN3c28vZgoxTU5EZng3c080bS9iQXRFVldnb3IvSkF1aGxGU0tQNFpGeWdEZ25HSjcxdkk2anZEdmtQNHBldG9QYXpPei9ZCmJUNnQxRThYZTZJK2FDYm9Yc1dIdlh4dkNoNjlUbGxrSXZkTlU2aVJVVUxoMHYxVFNwaSsKLS0tLS1FTkQgUFJJVkFURSBLRVktLS0tLQo="
	// base64EncodedEncryptedData2 := "d4sFkmeYXAQDHGrwcgf6geyTdNMH8jcp6dlk32utY87qDvFbVYdUGvRHeX7fUWN4lcmzNtAKVBq/n6jnAtejwBzEva4N5DCxSN4IapjkIZztI/spBLC7ZOnK/MDLKhq3uW+ibqsEthIcxjBzbHm0hWUebFTWDvmjaou6UVgq5swrPK9K1FOl+KO7rXMOGgVPpPwh/j2HGePnOroOtbY1znYxuMeOVRrCykhHe2rGtO3+mRcM29G0qw6drDQEt8Bn9YtaVZYxu8w/NQwAIUk6kUI7ifdj4oIURnXezIC2dLm1HqMNwS6D9mWIjQGMdEvwxOtK4+Uu6xeIPaK5ON4vrg=="
	base64EncryptedPrivateKeyPem := "LS0tLS1CRUdJTiBQUklWQVRFIEtFWS0tLS0tCk1JSUVwQUlCQUFLQ0FRRUFvUzc4b2NsWWZxRXNLeVdVdWtaWnlUd3ZuTVB2eGFVVWF5cTVyaUQ5QmtsdjJWdTQKbVZxVHBzUlB5VnFNWDVKdlNIaWkvQldzWmdQdGJNaHAzc3M2TUxPcVFaVEZiNmFkWlQ3UGdwK2FNbXkxR2wxTQpPYk91WlAxREhvTlNjYUpqYUNqU3pwS05lQ2g2YUlvLytWQzdNNkdoVWt0ZFdUVVpkZjhXajRZalJzd0FCbm85CjJOMHdFWTU0bkhQWmZCZUsxWFh6dTVnL1BsN0NaUUdZbWxDNkRUTDBqdGpWN3I5Y1BXUDdJMElPM0NLQWh6Z0wKTXRZSmFwZHVSNUtraFFhMWJsUkZzc0pHZVAva2d0U3dwcFlocGVrbm1KRlJvdmsrQjkvVnBHU2hZby9FbnhXSwpoeFRNMlJJbmJkNksvMVVITUNSK0hvTERVQTdKcENjL0N3L3drUUlEQVFBQkFvSUJBUUNKcEZOT1lWOTF1M0s1ClliMVJHTlNCbVNqYmlTTFBOL3BvemJLRWRWMDlLaWNlaVVucGcwVEZTNjBLdUkxWTYvYWwvNGJIR0VjQlV1ZEUKVnh5NWlmaW1NRGI3QysxU1BBajZ0Wm1FNVlCTEFUUGlVTHZRSnBXTlhnMGNHRzZsNVZOWVZrdzI1VVZ6ejdWQQpZcGJnODhUNUFWUityeWNIRCtZdk1tZDBwc3FuSUNEaDdCZzNpTEs5Tm95aUQzTmhWSzMyZ2d3TmJJN3c1Ylc0CjFISjVVbHBNOHprNjBIcHJqclAxaFFudU9JeGV3QlZWS1UzTUd6TjlZZ090TStoRXhkdUZ6SGo1Qm05aDUxKzMKUFBwSEFYVTlmTGdHaS9mREd4YVBwMmVOMzUyMjB1SCtCbHNvTlBQRTBXbldnL1dzakg5dE05SC9neXdiRjc1cQpJUmlJdEJ3QkFvR0JBTkh5emk5VkVtTWhLOHREU1pGdGNOTHNtT01XUkhSU2pFVytPVHVKVlRmOHkrNTRBTGtECmE3TVBveFJYbEFyRk5vTVJLZkN5MUdORXJGZ2YxVXVPSHNwSU9yMXVCZS8yYStkc3lyYTlFZmdMVTZ0STR4N0YKWDcwSnM5c0c3czhTK0ZEc25WTGtBTDZHeVZUZExSS1llWEM2YnNNdVlZbXMwcVZVdmszTjhwL0JBb0dCQU1TSgo1UWlNREpZeHdDNU1tNzEyUHE4Y3h1NmJPOUFFK1NhSnZvdTA4dFpCekRxa252bDJ1a2N1eUdJTm5heUZ5RDNjCmhuNlBpVUZ3SEdiVGpVR0JJWlNtanptbDM1Y1F3S3dJRDdTWDcxK2dYYk1tc3c5THJqZkFJOWw1SExkZTlETjMKSmcyYkhBUHJxTlJKRkVGNTJwOW5abitCNFNqUDYvMGxUa3gwbklUUkFvR0FHdFgzWnZOdjB6MzFOQmtZdVQ4TQo4ZGtPUGRDbDFWeUwwNHBhRkhNY2NuaWN0SkdnUGdYaTZWTnhXL05KMkFxWFNpK1NkZExWOVE1ZEhUS01lQWpsCmhLT3c1bSsrMElwbUJvSFFjNFU5VzhCTzVKRC9mdnBwbjJjekZ5Y090V2RPV0VHMlpsY0FoVm9ET0JiQllTUWUKT2J1SzF1WUJ6N0JJaUp1SmE3YU1jY0VDZ1lBNEdiQnh2eUE4b3hqYlhteTZLUTE5aUxaVUY2VjhIRjlPRzlWegpKWVNIbWlLRXZzYk9LSkRGanRvTit4cjl5dlk4aWowdDFTVDFzOTIza0QxcTdFUThuVXFXeHYxS3JyS0FxSkJyCkdVOXZ1V1BscG05SU4yOFVaTmtXMDVaWjFWWFpkdkk2d2dLK0w3OXlVU1hpQnJsYjlQVHlBbzNWeTN2dU4zN2gKSGd3cmNRS0JnUUNDODhIRldKcWdkYzFmR0VHc1dYUWF3ekNDaUhpRk1xVUZ5blhMVEFESHJqeVpEQWpEaWs3bwpVUHlSYVFocGo2ZzQ1aUZQQVBYaHdrWkZ3em0wQ1BuT1J0VUx3R094YmIyNFpnMlBZMWtJbGJQU1pWTXV0elhvCk1vTDJDTUJhOXRyVHREVHozc0NNd2t5QjlZeHJpelVxMDc3OWRjcmFDWnRZWHd1Ym12UXZ0dz09Ci0tLS0tRU5EIFBSSVZBVEUgS0VZLS0tLS0K"
	base64EncodedEncryptedData2 := "HmcxFXjunYmJjCKkSnMRxCtmp5PJ3ycYuN+pc2aClFXu+XCRAXTuBw/hgzLb3enEitDK4LG4lpf6TSTpk2UrHKLrYg/GUywHCyx5WijpK3m80LLI5xzmUSlvladhTEOBfO7jSVxJxEbsvQ9gN/lCze6F1iXXDFtqJAJ+ONz9JQDj9FgZjHbf+3h/C930YBo1iQzieSP+x4R7zApUwqEwNTVKHO3K+fKrs/SK8m0IRwXkm62ebPKSewfDdwZnGIXewwgF/WEoowHdoWgHmyi9dYTX+QkCIuhgMs40GK4hLhspYLeRtPVSlO12CieLcA0p6A4SM6FEvjjhHADAGBi1TQ=="
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
