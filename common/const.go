package common

import "log"

const (
	DbTypeRestaurant = 1
	DbTypeUser       = 2
	DBTypeUpload     = 3
)
const (
	CurrentUser = "user"
)

type Requester interface {
	GetUserId() int
	GetEmail() string
	GetRole() string
}

func AppRecover() {
	if err := recover(); err != nil {
		log.Println("Recovery error:", err)
	}
}
