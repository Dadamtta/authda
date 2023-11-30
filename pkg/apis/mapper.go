package apis

import (
	"bytes"
	"dadamtta/pkg/rsa"
	"encoding/base64"
	"encoding/json"
	"errors"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func BodyMapper[T any](bytes []byte, t *T) error {
	return json.Unmarshal([]byte(bytes), &t)
}

func BodyMapperWithDecrypt[T any](c *gin.Context, t *T) error {
	session := sessions.Default(c)
	base64EncodedPrivateKeyPem := session.Get("PrivateKey")
	if base64EncodedPrivateKeyPem == nil {
		return errors.New("Private Key 정보가 없음")
	}
	buf := new(bytes.Buffer)
	buf.ReadFrom(c.Request.Body)
	encodedRequestData := buf.String()
	data, err := base64.StdEncoding.DecodeString(encodedRequestData)
	if err != nil {
		return err
	}
	requestData, err := rsa.DecryptBase64EncodedPrivateKeyPem(data, base64EncodedPrivateKeyPem.(string))
	if err != nil {
		return errors.New(err.Error())
	}
	return BodyMapper[T]([]byte(requestData), t)
}
