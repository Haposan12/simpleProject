package services

import (
	"errors"
	"project/freelance/danspro/simpleProject/app/models"
	"project/freelance/danspro/simpleProject/app/utils/oauth"
)

func LoginService(req models.LoginRequest) (models.LoginResponseModel, error) {
	var resp models.LoginResponseModel
	userId := checkUsernameAndPassword(req.Username, req.Password)
	if userId == 0 {
		err := errors.New("data not found")
		return resp, err
	}

	//set oauth
	tokenString, err := oauth.GenerateToken(models.User{
		Id:       userId,
		Name:     req.Username,
		Password: req.Password,
	})

	if err != nil {
		return resp, err
	}

	resp.Id = userId
	resp.Username = req.Username
	resp.AccessToken = tokenString

	return resp, nil
}

// checkUsernameAndPassword: check username and password by select from database. And return the id
func checkUsernameAndPassword(username, password string) int {
	user := models.GetUser(username, password)

	return user.Id
}
