package dadamtta

import (
	"dadamtta/internal/product"
	"dadamtta/internal/sql"
	"dadamtta/pkg/apis"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewProductCommand(router *gin.Engine, repository product.Repository) {
	service := product.NewService(repository)

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
		// Header에서 ID
		adminId := ""
		var dto ProductRegisterFormRequest
		err := apis.BodyMapper[ProductRegisterFormRequest](c, &dto)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}

		productId, err := service.Register(adminId, dto.CategoryCode, dto.Label, dto.Price, dto.Description, dto.Content)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		} else {
			c.JSON(http.StatusCreated, gin.H{
				"message": productId,
			})
			return
		}
	})
}

func Search(router *gin.Engine, service product.Service) {
	router.GET(`/v1/products`, func(c *gin.Context) {
		options, err := sql.NewSearchOptions(c.Param("page"), c.Param("listSize"), c.Param("sorter"), c.Param("component"), c.Param("q"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}

		service.Search(options)

		c.JSON(http.StatusOK, gin.H{
			"message": "",
		})
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
