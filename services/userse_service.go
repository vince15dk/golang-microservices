package services

import (
	"mvc/project/domain"
	"mvc/project/utils"
)

func GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	return domain.GetUser(userId)
}
