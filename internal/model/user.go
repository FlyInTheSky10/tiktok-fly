package model

type TiktokUser struct {
	ID       uint `gorm:"primary_key"`
	Username string
	Password string
}
type TiktokRelation struct {
	ID         uint `gorm:"primary_key"`
	UserId     uint
	FollowerId uint
}
