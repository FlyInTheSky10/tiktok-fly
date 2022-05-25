package param

type ReqUserRegister struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}
type ReqUserLogin struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

type ResUserRegister struct {
	Token string `json:"token"`
}
type ResUserLogin struct {
	Token string `json:"token"`
}
