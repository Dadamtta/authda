package auth

import "testing"

func TestJwTokenGenerateToken(t *testing.T) {
	payload := map[string]any{
		"name": "홍길동",
		"age":  "240",
	}
	token, err := New("secret-key").GenerateToken(payload)

	if err != nil {
		t.Log("에러 발생", err.Error())
	} else {
		t.Log(token)
	}
}

func TestJwTokenIsValid(t *testing.T) {
	tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZ2UiOiIyNDAiLCJuYW1lIjoi7ZmN6ri464-ZIn0.uX_R6457XUjCKHA7l5eiw1K3hy48SdV717_C3ofs-Zg"
	isValid := New("secret-key").IsValid(tokenString)
	if isValid {
		t.Log("유효")
	} else {
		t.Log("유효하지 않음")
	}
}

func TestJwTokenParse(t *testing.T) {
	tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZ2UiOiIyNDAiLCJuYW1lIjoi7ZmN6ri464-ZIn0.uX_R6457XUjCKHA7l5eiw1K3hy48SdV717_C3ofs-Zg"
	claims, err := New("secret-key").Parse(tokenString)
	if err != nil {
		t.Log("에러 발생", err.Error())
	} else {
		for k, v := range claims {
			t.Log("키 = ", k, ", 값 = ", v)
		}
	}
}
