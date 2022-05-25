package service

import (
	"errors"
	"github.com/FlyInThesky10/TikTok-Fly/internal/model"
	"github.com/FlyInThesky10/TikTok-Fly/pkg/app"
	"github.com/FlyInThesky10/TikTok-Fly/pkg/util"
)

func (svc *Service) UserRegister(username, password string) error {
	encryptedPassword := util.EncodeMD5(password)
	err := svc.dao.UserAddUser(&model.TiktokUser{
		Username: username,
		Password: encryptedPassword,
	})
	return err
}
func (svc *Service) UserLogin(username, password string) (string, error) {
	encryptedPassword := util.EncodeMD5(password)
	user, err := svc.dao.FindUser(username)
	if err != nil {
		return "", err
	}
	if user.Password != encryptedPassword {
		return "", errors.New("password incorrect")
	}
	token, _, err := app.GenerateJWTToken(username, encryptedPassword)
	if err != nil {
		return "", errors.New("cannot generate token")
	}
	return token, nil
}
