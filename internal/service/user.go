package service

import (
	"errors"
	"github.com/FlyInThesky10/TikTok-Fly/internal/model"
	"github.com/FlyInThesky10/TikTok-Fly/internal/service/param"
	"github.com/FlyInThesky10/TikTok-Fly/pkg/app"
	"github.com/FlyInThesky10/TikTok-Fly/pkg/util"
	"github.com/jinzhu/gorm"
)

func (svc *Service) UserRegister(username, password string) (int, error) {
	encryptedPassword := util.EncodeMD5(password)
	u := model.TiktokUser{
		Username: username,
		Password: encryptedPassword,
	}
	err := svc.dao.UserAddUser(&u)
	return int(u.ID), err
}
func (svc *Service) UserLogin(username, password string) (string, int, error) {
	encryptedPassword := util.EncodeMD5(password)
	user, err := svc.dao.FindUserByName(username)
	if err != nil {
		return "", 0, err
	}
	if user.Password != encryptedPassword {
		return "", 0, errors.New("password incorrect")
	}
	token, _, err := app.GenerateJWTToken(username, encryptedPassword, int(user.ID))
	if err != nil {
		return "", 0, errors.New("cannot generate token")
	}
	return token, int(user.ID), nil
}
func (svc *Service) UserInfo(userId, myId int) (param.ResUserInfo, error) {
	followCount, followerCount, err := svc.dao.UserFollowAndFollower(userId)
	if err != nil {
		return param.ResUserInfo{}, err
	}
	user, err := svc.dao.FindUserById(userId)
	if err != nil {
		return param.ResUserInfo{}, err
	}
	relation, err := svc.dao.UserFindRelation(userId, myId)
	if err != nil && err != gorm.ErrRecordNotFound {
		return param.ResUserInfo{}, err
	}
	return param.ResUserInfo{
		Response: param.Response{StatusCode: 0, StatusMsg: "success"},
		User: param.User{
			Id:            userId,
			Name:          user.Username,
			FollowCount:   followCount,
			FollowerCount: followerCount,
			IsFollow:      relation,
		},
	}, nil
}
