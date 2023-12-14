package errorc

import "errors"

var (
	TokenGenerateError               = errors.New("TOKEN generate error.")
	DtoUnmarshalError                = errors.New("DTO unmarshaling error.")
	AuthorizedError                  = errors.New("AUTHORIZED error.")
	EntityNotFoundError              = errors.New("ENTITY not found error.")
	ProductNotOnSaleError            = errors.New("PRODUCT not on sale.")
	AppIsDeletedError                = errors.New("APP is deleted.")
	AppIsExpiredAtError              = errors.New("APP is expired.")
	POLICYProductOwnershipLimitError = errors.New("POLICY product ownership limit error.")
)
