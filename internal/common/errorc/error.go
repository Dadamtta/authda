package errorc

import "errors"

var (
	DtoUnmarshalError   = errors.New("DTO unmarshaling error.")
	AuthorizedError     = errors.New("AUTHORIZED error.")
	TokenGenerateError  = errors.New("TOKEN generate error.")
	EntityNotFoundError = errors.New("ENTITY not found error.")
)
