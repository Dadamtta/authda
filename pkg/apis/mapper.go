package apis

import (
	"bytes"
	"dadamtta/pkg/rsa"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func BodyMapper[T any](c *gin.Context, t *T) error {
	body := c.Request.Body
	bytes, err := io.ReadAll(body)
	if err != nil {
		return err
	}
	return BytesMapper[T]([]byte(bytes), t)
}

func BytesMapper[T any](bytes []byte, t *T) error {
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
	return BytesMapper[T]([]byte(requestData), t)
}
