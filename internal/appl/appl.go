package appl

import (
	"dadamtta/private/p_policy"
	"time"

	"github.com/google/uuid"
)

type ApplState int

const (
	APPL_DELETED ApplState = 0
	APPL_ACTIVE  ApplState = 1
)

type Appl struct {
	ID        string
	UserId    string
	ProductId string
	ExpiredAt time.Time
	State     ApplState
	CreatedAt time.Time
	UpdatedAt time.Time
}

func GenerateAppl(userId, productId string) Appl {
	expiredAt := time.Now()
	expiredAt.Add(time.Hour * p_policy.APP_EXPIRED_AT_HOURS)
	return Appl{
		ID:        uuid.New().String(),
		UserId:    userId,
		ProductId: productId,
		State:     APPL_ACTIVE,
		CreatedAt: time.Now(),
		ExpiredAt: expiredAt,
	}
}

func (a *Appl) IsActive() bool {
	return a.State == APPL_ACTIVE
}

type ImageState int

const (
	IMAGE_DELETED  ImageState = 0
	IMAGE_ACTIVE   ImageState = 1
	IMAGE_UNACTIVE ImageState = 2
	IMAGE_ERROR    ImageState = 3
)

type ApplImage struct {
	Id               string
	ApplId           string
	OriginFilePath   string
	ResizingFilePath string
	fileSize         uint
	State            ImageState
	CreatedAt        time.Time
	UpdatedAt        time.Time
}
