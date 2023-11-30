package dadamtta

import (
	"dadamtta/pkg/apis"
	"dadamtta/pkg/rsa"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func NewRSACommand(router *gin.Engine) {

	GetPublicKey(router)
	CheckDecrypt(router)

}

func CheckDecrypt(router *gin.Engine) {
	router.POST(`/v1/rsa/check`, func(c *gin.Context) {
		var test *string
		err := apis.BodyMapperWithDecrypt[string](c, test)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{
				"message": test,
			})
		}
	})
}

func GetPublicKey(router *gin.Engine) {
	router.GET(`/v1/rsa`, func(c *gin.Context) {
		session := sessions.Default(c)
		base64EncodedPrivateKeyPem := session.Get("PrivateKey")
		var publicKey string = ""
		if base64EncodedPrivateKeyPem == nil {
			base64EncodedPrivateKeyPem, base64EncodedPublicKeyPem := rsa.GenerateRSA(2048)
			publicKey = base64EncodedPublicKeyPem
			session.Set("PrivateKey", base64EncodedPrivateKeyPem)
			err := session.Save()

			if err != nil {
				println(err.Error())
			}
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
			"base64EncodedPublicKeyPem": publicKey,
		})
	})
}
