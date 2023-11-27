package dadamtta

import (
	"dadamtta/pkg/rsa"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func NewRSACommand(router *gin.Engine) {

	GetPublicKey(router)

}

func GetPublicKey(router *gin.Engine) {
	router.GET(`/v1/rsa`, func(c *gin.Context) {
		session := sessions.Default(c)
		base64EncodedPrivateKeyPem := session.Get("PrivateKey")
		var publicKey string = ""
		if base64EncodedPrivateKeyPem == nil {
			// 현재 세션으로 새로 만들기
			base64EncodedPrivateKeyPem, base64EncodedPublicKeyPem := rsa.GenerateRSA(2048)
			publicKey = base64EncodedPublicKeyPem
			session.Set("PrivateKey", base64EncodedPrivateKeyPem)
		} else {
			publicKeyString, err := rsa.GetBase64EncodedPublicKeyPem(base64EncodedPrivateKeyPem.(string))
			if err != nil {
				// todo 에러 처리
				c.JSON(http.StatusInternalServerError, gin.H{
					"message": err.Error(),
				})
				return
			}
			publicKey = publicKeyString
		}
		c.JSON(http.StatusOK, gin.H{
			"message": publicKey,
		})
		return
	})
}
