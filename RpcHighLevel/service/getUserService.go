package service

import (
	"errors"
	"highlevel/connection"
	"highlevel/structs"

	"gorm.io/gorm"
)

func GetUserService(userID string) (structs.UserInfoResponse, error) {
	var user structs.User
	result := connection.ReadDB.Where("user_id = ?", userID).First(&user)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return structs.UserInfoResponse{}, errors.New("user not found in database")
		}
		return structs.UserInfoResponse{}, result.Error
	}

	userInfoResponse := structs.UserInfoResponse{
		UserID:  user.UserID,
		EmailID: user.EmailID,
		Name:    user.Name,
		Role:    user.Role,
	}
	return userInfoResponse, nil
}
