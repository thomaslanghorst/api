package service

import (
	"errors"
	"github.com/HackIllinois/api/services/registration/config"
	"github.com/HackIllinois/api/services/registration/models"
	"github.com/HackIllinois/api/common/apirequest"
	"net/http"
)

/*
	Get basic user info from user serivce
*/
func GetUserInfo(id string) (*models.UserInfo, error) {
	var user_info models.UserInfo
	status, err := apirequest.Get(config.USER_SERVICE + "/user/" + id + "/", &user_info)

	if err != nil {
		return nil, err
	}

	if status != http.StatusOK {
		return nil, errors.New("User service failed to return information")
	}

	return &user_info, nil
}
