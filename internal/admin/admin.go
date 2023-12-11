package admin

import "time"

type State int

const (
	UNACTIVE State = 0
	ACTIVE   State = 1
)

type Admin struct {
	Id        string
	HashedPwd string
	Email     string
	NickName  string
	Role      uint8
	Phone     string
	State     State
	CreatedAt time.Time
	UpdatedAt time.Time
}
