package param

type ReqUserRegister struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}
type ReqUserLogin struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}
type ReqUserInfo struct {
	UserId int    `form:"user_id"`
	Token  string `form:"token" binding:"required"`
}

type ResUserRegister struct {
	Response
	Token  string `json:"token"`
	UserId int    `json:"user_id"`
}
type ResUserLogin struct {
	Response
	Token  string `json:"token"`
	UserId int    `json:"user_id"`
}
type ResUserInfo struct {
	Response
	User User `json:"user"`
}

type User struct {
	Id            int    `json:"id"`
	Name          string `json:"name"`
	FollowCount   int    `json:"follow_count"`
	FollowerCount int    `json:"follower_count"`
	IsFollow      bool   `json:"is_follow"`
}
