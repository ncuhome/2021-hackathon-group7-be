package service

import (
	"encoding/hex"
	"strconv"
	"strings"
	"tudo/model"
	"tudo/model/dao"
	"tudo/model/dto"
	"tudo/util"
)

// 邮箱注册
func Register(req *dto.Register) uint {
	// 账号检查交给邮箱和验证码
	code := CheckPassword(req.Password)
	if code != SuccessCode {
		return code
	}

	user := &dao.User{
		Email: req.Username,
	}
	_ = user.Retrieve()
	if user.ID != 0 {
		return EmailUsed
	}

	// 获取邮箱验证码数据
	data := &dao.EmailBindCache{}
	cacheObj := &dao.JsonCache{
		Data: data,
		ID:   req.Username,
	}
	err := cacheObj.GetData()

	if err != nil {
		if err == dao.CacheNil {
			return CodeError
		} else {
			model.ErrLog.Println(err)
			return CommitDataError
		}
	}

	// 检验验证码
	if req.Key != data.Key {
		return CodeError
	}
	if req.Username != data.Email {
		return CodeError
	}

	//删缓存
	err = cacheObj.DelData()
	if err != nil {
		model.ErrLog.Println(err)
		return ServerError
	}

	// 予以注册
	salt, err := util.RandHexStr(64)
	if err != nil {
		model.ErrLog.Println(err)
		return ServerError
	}

	password := hex.EncodeToString(util.SHA512([]byte(req.Password + salt)))

	user = &dao.User{
		Username:    req.Username,
		Password:    password,
		Email:       req.Username,
		Salt:        salt,
		LoginStatus: "0",
	}
	err = user.Create()
	if err != nil {
		model.ErrLog.Println(err)
		return ServerError
	}
	return SuccessCode
}

func Login(req *dto.Login) (*map[string]interface{}, uint) {
	user := &dao.User{}
	if strings.Contains(req.User, "@") {
		user.Email = req.User
	} else {
		user.Username = req.User
	}
	_ = user.Retrieve()
	if user.ID == 0 {
		return nil, LoginError
	}

	password := hex.EncodeToString(util.SHA512([]byte(req.Password + user.Salt)))
	if user.Password != password {
		return nil, LoginError
	}

	token, err := model.Jwt.GenerateToken(strconv.Itoa(int(user.ID)), user.LoginStatus)
	if err != nil {
		model.ErrLog.Println(err)
		return nil, ServerError
	}

	data := &map[string]interface{}{
		"id":       user.ID,
		"token":    token,
		"username": user.Username,
	}

	return data, SuccessCode
}

func CheckUsername(username string) uint {
	usernameLen := len(username)
	if usernameLen < 2 || usernameLen > 16 {
		return CommitDataError
	}
	for i := 0; i < usernameLen; i++ {
		if (username[i] < 'a' || 'z' < username[i]) && (username[i] < 'A' || 'Z' < username[i]) && (username[i] < '0' || '9' < username[i]) {
			return CommitDataError
		}
	}
	return SuccessCode
}

func CheckPassword(password string) uint {
	passwordLen := len(password)
	if passwordLen < 8 || passwordLen > 32 {
		return CommitDataError
	}
	//[33,126]覆盖了大小写字母、数字、普通可见符号
	for i := 0; i < passwordLen; i++ {
		if password[i] < 33 || password[i] > 126 {
			return CommitDataError
		}
	}
	return SuccessCode
}

func SetPassword(req *dto.SetPassword, id uint) uint {
	code := CheckPassword(req.NewPassword)
	if code != SuccessCode {
		return code
	}

	user := &dao.User{
		ID: id,
	}
	err := user.Retrieve()
	if err != nil {
		return ServerError
	}

	shaPassword := hex.EncodeToString(util.SHA512([]byte(req.Password + user.Salt)))
	if user.Password != shaPassword {
		return OldPasswordError
	}

	return updatePassword(req.NewPassword, id)
}

//更新盐、个人登录状态、密码
func updatePassword(newPassword string, id uint) uint {
	saltStr, err := util.RandHexStr(64)
	if err != nil {
		model.ErrLog.Println(err)
		return ServerError
	}
	shaNewPassword := hex.EncodeToString(util.SHA512([]byte(newPassword + saltStr)))
	loginStatus, err := util.RandHexStr(8)
	if err != nil {
		model.ErrLog.Println(err)
		return ServerError
	}

	user := &dao.User{
		ID: id,
	}
	err = user.Update(map[string]interface{}{
		"Password":    shaNewPassword,
		"Salt":        saltStr,
		"LoginStatus": loginStatus,
	})
	if err != nil {
		model.ErrLog.Println(err)
		return ServerError
	}

	//删用户缓存
	Data := &dao.UserData{
		ID: id,
	}
	err = Data.DelCache()
	if err != nil {
		model.ErrLog.Println(err)
	}

	return SuccessCode
}

func GetEmail(id uint) (*map[string]interface{}, uint) {
	user := &dao.User{}
	user.ID = id
	err := user.Retrieve()
	if err != nil {
		return nil, ServerError
	}

	data := &map[string]interface{}{
		"email": user.Email,
	}
	return data, SuccessCode
}
