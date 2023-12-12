package response

import (
	"dadamtta/internal/common/errorc"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	ERROR_DTO_UNMARSHAL   = 100001 // 요청 에러
	ERROR_AUTHORIZED      = 700001 // 인증 에러
	ERROR_ENTITY_NOTFOUND = 900001 // DB 관련 에러
)

func HandleResponseErrorWithCustomMessage(context *gin.Context, err error, message string) bool {
	if err != nil {
		return false
	}
	if errors.Is(err, errorc.DtoUnmarshalError) {
		context.AbortWithStatusJSON(http.StatusBadRequest, NewErrorResponse(ERROR_DTO_UNMARSHAL, message))
		return true
	} else if errors.Is(err, errorc.AuthorizedError) {
		context.AbortWithStatusJSON(http.StatusUnauthorized, NewErrorResponse(ERROR_AUTHORIZED, message))
		return true
	} else if errors.Is(err, errorc.EntityNotFoundError) {
		context.AbortWithStatusJSON(http.StatusUnauthorized, NewErrorResponse(ERROR_ENTITY_NOTFOUND, message))
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
