package response

import (
	"dadamtta/internal/common/errorc"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	ERROR_DTO_UNMARSHAL             = 100001 // 요청 에러
	ERROR_PRODUCT_NOTONSALE         = 200001 // 판매 중이 아닌 상품
	ERROR_POLICY_APP_OWNERSHIPLIMIT = 990001 //
	ERROR_AUTHORIZED                = 700001 // 인증 에러
	ERROR_ENTITY_NOTFOUND           = 900001 // DB 관련 에러
)

func HandleResponseErrorWithCustomMessage(context *gin.Context, err error, message string) bool {
	if err != nil {
		return false
	}
	if errors.Is(err, errorc.DtoUnmarshalError) {
		// 400
		context.AbortWithStatusJSON(http.StatusBadRequest, NewErrorResponse(ERROR_DTO_UNMARSHAL, message))
		return true
	} else if errors.Is(err, errorc.AuthorizedError) {
		// 401
		context.AbortWithStatusJSON(http.StatusUnauthorized, NewErrorResponse(ERROR_AUTHORIZED, message))
		return true
	} else if errors.Is(err, errorc.EntityNotFoundError) {
		// 401
		context.AbortWithStatusJSON(http.StatusUnauthorized, NewErrorResponse(ERROR_ENTITY_NOTFOUND, message))
		return true
	} else if errors.Is(err, errorc.ProductNotOnSaleError) {
		// 400
		context.AbortWithStatusJSON(http.StatusBadRequest, NewErrorResponse(ERROR_PRODUCT_NOTONSALE, message))
		return true
	} else if errors.Is(err, errorc.POLICYProductOwnershipLimitError) {
		// 422
		context.AbortWithStatusJSON(http.StatusUnprocessableEntity, NewErrorResponse(ERROR_PRODUCT_NOTONSALE, message))
		return true
	}
	return false
}

// func HandleResponseError(context *gin.Context, err error) bool {
// 	if err != nil {
// 		return false
// 	}
// 	if errors.Is(err, errorc.DtoUnmarshalError) {
// 		context.AbortWithStatusJSON(http.StatusBadRequest, NewErrorResponse(code, message))
// 	}
// 	return true
// }
