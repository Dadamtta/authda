package appl

import "time"

type Repository interface {
	Save(appl Appl) (appId string, err error)
	DeleteById(applId string) error
	UpdateExpiredAt(expiredAt time.Time) error
	FindById(applId string) *Appl
	FindAllByUserId(userId string) *[]Appl
	FindAllByUserIdAndProductId(userId, productId string) []Appl
}

type ImageRepository interface {
	Save(applId, originFilePath, resizingFilePath string, fileSize uint) error
	FindAllByApplId(applId string) *[]ApplImage
	UpdateState(imageId string, state ImageState) error
}

type DataRepository[T any] interface {
	FindById(applId string)
	SaveAndUpdate(t *T) error
}
