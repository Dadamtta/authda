package dadamtta

import (
	"dadamtta/internal/product"
	"dadamtta/pkg/apis"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewProductCommand(router *gin.Engine) {
	service := product.NewService()

	// 상품 등록
	Register(router, service)
	// 상품 검색
	Search(router, service)
	// 상품 조회
	Get(router, service)
	// 상품 수정
	Update(router, service)
	// 상품 삭제
	Delete(router, service)

}

func Register(router *gin.Engine, service product.Service) {
	router.POST(`/v1/products`, func(c *gin.Context) {
		body := c.Request.Body
		bytes, err := io.ReadAll(body)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}

		var dto ProductRegisterFormRequest
		err = apis.BodyMapper[ProductRegisterFormRequest](bytes, &dto)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}

		productId, err := service.Register(dto.CategoryCode, dto.Label, dto.Price, dto.Description, dto.Content)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{
				"message": productId,
			})
			return
		}
	})
}

func Search(router *gin.Engine, service product.Service) {
	router.GET(`/v1/products`, func(c *gin.Context) {
		service.Search()
	})
}

func Get(router *gin.Engine, service product.Service) {
	router.GET(`/v1/products/:productId`, func(c *gin.Context) {
		service.Get("")
	})
}

func Update(router *gin.Engine, service product.Service) {
	router.PUT(`/v1/products/:productId`, func(c *gin.Context) {
		service.Update()
	})
}

func Delete(router *gin.Engine, service product.Service) {
	router.DELETE(`/v1/products/:productId`, func(c *gin.Context) {
		service.Delete("")
	})
}
