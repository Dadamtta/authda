package appl

import "time"

type Repository interface {
	Save(userId, productId string) error
	DeleteById(applId string) error
	UpdateExpiredAt(expiredAt time.Time) error
	FindById(applId string) *Appl
	FindAllByUserId(userId string) *[]Appl
}

type ImageRepository interface {
	Save(applId, originFilePath, resizingFilePath string, fileSize uint) error
	FindAllByApplId(applId string) *[]ApplImage
	UpdateState(imageId string, state ImageState) error
}
