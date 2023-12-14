package user

import (
	"dadamtta/internal/appl"
	"dadamtta/internal/common/errorc"
	"dadamtta/internal/payment_order"
	"dadamtta/internal/product"
	"dadamtta/pkg/utils/logger"
	"dadamtta/private/p_appl"
	"dadamtta/private/p_policy"
	"fmt"
	"time"
)

type AppType string

const (
	WEDDING AppType = "wedding"
)

type Service interface {
	SignUp(id, pwd, phone, email, name string) (string, error)
	SignIn(id, pwd string) error
	CreateApp(appType AppType, userId, productId string) (appId string, err error)
	UpdateAppData(appType AppType, userId, appId string, data *p_appl.WeddingInvitation) error
}

type service struct {
	repository             Repository
	productRepository      product.Repository
	applRepository         appl.Repository
	applDataRepository     appl.DataRepository[p_appl.WeddingInvitation]
	paymentOrderRepository payment_order.Repository
}

func NewService(userRepository Repository, applRepository appl.Repository, applDataRepository appl.DataRepository[p_appl.WeddingInvitation], productRepository product.Repository, paymentOrderRepository payment_order.Repository) Service {
	return &service{
		repository:             userRepository,
		applRepository:         applRepository,
		applDataRepository:     applDataRepository,
		productRepository:      productRepository,
		paymentOrderRepository: paymentOrderRepository,
	}
}

func (s *service) SignUp(id, pwd, phone, email, name string) (string, error) {
	println("회원가입 진행")
	newUser, err := GenerateUser(id, pwd, phone, email, name)
	if err != nil {
		return "", err
	}
	// todo - Save 에러
	err = s.repository.Save(*newUser)
	if err != nil {
		return "", err
	}
	return newUser.Id, err
}

func (s *service) SignIn(id, pwd string) error {
	println("로그인 진행")
	return nil
}

func (s *service) CreateApp(appType AppType, userId, productId string) (appId string, err error) {
	logger.Debug(fmt.Sprintf("Create App. User ID -> %s, Product ID -> %s", userId, productId))
	product := s.productRepository.FindById(productId)
	if product == nil {
		logger.Error(fmt.Sprintf("[Func] Not found product. Product ID -> %s", productId))
		err = errorc.EntityNotFoundError
		return
	}
	if !product.IsOnSale() {
		logger.Error(fmt.Sprintf("[Func] Invalid product. Product ID -> %s", productId))
		err = errorc.ProductNotOnSaleError
		return
	}
	var activeApplCount = 0
	appls := s.applRepository.FindAllByUserIdAndProductId(userId, productId)
	for _, appl := range appls {
		if appl.IsActive() {
			activeApplCount++
		}
	}
	if activeApplCount >= p_policy.PRODUCT_APP_OWNERSHIP_LIMIT {
		logger.Error(fmt.Sprintf("[Func] No more creation. User ID -> %s, Product ID -> %s", userId, productId))
		err = errorc.POLICYProductOwnershipLimitError
		return
	}
	appId, err = s.applRepository.Save(appl.GenerateAppl(userId, productId))
	return
}

func (s *service) UpdateAppData(appType AppType, userId, appId string, data *p_appl.WeddingInvitation) error {
	userApp := s.applRepository.FindById(appId)
	if userApp == nil {
		return errorc.EntityNotFoundError
	}
	if userApp.UserId != userId {
		return errorc.AuthorizedError
	}
	if !userApp.IsActive() {
		return errorc.AppIsDeletedError
	}
	today := time.Now()
	if today.After(userApp.ExpiredAt) {
		return errorc.AppIsExpiredAtError
	}
	// 앱데이터 조회
	s.applDataRepository.FindById("")
	// 결제된 앱인지 확인
	if s.paymentOrderRepository.ExistsByAppId(appId) {
		// // 결제되었으면 예식날짜 변경 불가 (정책)

	}

	// 앱데이터 업데이트 진행
	return nil
}
