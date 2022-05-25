package dao

import (
	"errors"
	"github.com/FlyInThesky10/TikTok-Fly/internal/model"
	"github.com/jinzhu/gorm"
)

func (dao *Dao) UserAddUser(user *model.TiktokUser) error {
	u := model.TiktokUser{}
	err := dao.engine.Model(&model.TiktokUser{}).Where("username = ?", user.Username).Take(&u).Error
	if err == gorm.ErrRecordNotFound {
		return dao.engine.Model(&model.TiktokUser{}).Create(&*user).Error
	}
	return errors.New("username duplicated")
}
func (dao *Dao) FindUser(username string) (model.TiktokUser, error) {
	u := model.TiktokUser{}
	err := dao.engine.Model(&model.TiktokUser{}).Where("username = ?", username).Take(&u).Error
	if err != nil {
		return model.TiktokUser{}, err
	}
	return u, nil
}
