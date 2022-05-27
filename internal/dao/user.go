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
func (dao *Dao) FindUserByName(username string) (model.TiktokUser, error) {
	u := model.TiktokUser{}
	err := dao.engine.Model(&model.TiktokUser{}).Where("username = ?", username).Take(&u).Error
	if err != nil {
		return model.TiktokUser{}, err
	}
	return u, nil
}
func (dao *Dao) FindUserById(userId int) (model.TiktokUser, error) {
	u := model.TiktokUser{}
	err := dao.engine.Model(&model.TiktokUser{}).Where("id = ?", userId).Take(&u).Error
	if err != nil {
		return model.TiktokUser{}, err
	}
	return u, nil
}
func (dao *Dao) UserFollowAndFollower(userId int) (int, int, error) {
	var followCount, followerCount int
	err := dao.engine.Model(&model.TiktokRelation{}).Where("follower_id = ?", userId).Count(&followCount).Error
	if err != nil {
		return 0, 0, err
	}
	err = dao.engine.Model(&model.TiktokRelation{}).Where("user_id = ?", userId).Count(&followerCount).Error
	if err != nil {
		return 0, 0, err
	}
	return followCount, followerCount, nil
}
func (dao *Dao) UserFindRelation(userId int, followerId int) (bool, error) {
	r := model.TiktokRelation{}
	err := dao.engine.Model(&model.TiktokRelation{}).Where("user_id = ? AND follower_id = ?", userId, followerId).Take(&r).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	if err == gorm.ErrRecordNotFound {
		return false, nil
	}
	return true, nil
}
