package dto

type Register struct {
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
	Key   string `json:"key" form:"key" binding:"required"`
}

type OrgInfo struct {
	LogoUrl	 string `json:"logo_url" form:"logo_url"`
	// 长度大于8
	Password string `json:"password" form:"password"`
}

type Login struct {
	Username     string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

type Token struct {
	Token     string `json:"token" form:"token" binding:"required"`
}

type SetPassword struct {
	Password    string `json:"password" form:"password" binding:"required"`
	NewPassword string `json:"new_password" form:"new_password" binding:"required"`
}

type SetPasswordByEmailReq struct {
	Key         string `json:"key" form:"key" binding:"required"`
	NewPassword string `json:"new_password" form:"new_password" binding:"required"`
}
