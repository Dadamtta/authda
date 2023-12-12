package payment_order

type Repository interface {
	ExistsByAppId(appId string) bool
}
