package dto

// Username即是Emmail
type Register struct {
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
	Key   string `json:"key" form:"key" binding:"required"`
}

type Login struct {
	//user may be username or email
	User     string `json:"user" form:"user" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

type SetPassword struct {
	Password    string `json:"password" form:"password" binding:"required"`
	NewPassword string `json:"new_password" form:"new_password" binding:"required"`
}

type SetPasswordByEmailReq struct {
	Key         string `json:"key" form:"key" binding:"required"`
	NewPassword string `json:"new_password" form:"new_password" binding:"required"`
}
