package vo

type LoginReq struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
	Captcha  string `form:"captcha" binding:"required"`
}

type LoginRes struct {
	AccessToken string `json:"accessToken"`
}
