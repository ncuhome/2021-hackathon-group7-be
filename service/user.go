package service

import (
	"encoding/hex"
	"errors"
	"gorm.io/gorm"
	"strconv"
	"tudo/model"
	"tudo/model/dao"
	"tudo/model/dto"
	"tudo/util"
)

func orgRegister(org string, req *dto.OrgInfo) uint {
	code := CheckPassword(req.Password)
	if code != SuccessCode {
		return code
	}

	salt, err := util.RandHexStr(64)
	if err != nil {
		model.ErrLog.Println(err)
		return ServerError
	}

	password := hex.EncodeToString(util.SHA512([]byte(req.Password + salt)))
	user := &dao.User{
		Username:    org,
		Password:    password,
		Salt:        salt,
		LoginStatus: "0",
	}

	userInfo := &dao.UserInfo{
		Nickname:     org,
		Avatar:       req.LogoUrl,
		Verification: "v",
	}
	err = user.CreateWith(userInfo)
	if err != nil {
		model.ErrLog.Println(err)
		return ServerError
	}
	return SuccessCode
}

func orgPut(req *dto.OrgInfo, id uint) uint {
	code := CheckPassword(req.Password)
	if code != SuccessCode {
		return code
	}

	updatePassword(req.Password, id)

	orgUserInfo := &dao.UserInfo{}
	orgUserInfo.ID = id

	change := &dao.UserInfo{
		Avatar: req.LogoUrl,
	}
	err := orgUserInfo.Update(change)
	if err != nil {
		model.ErrLog.Println(err)
		return ServerError
	}
	return SuccessCode
}

// 社团激活(注册)、修改资料
func OrgPostInfo(req *dto.OrgInfo, id uint) uint {
	ncuUser := &dao.User{ID: id}
	err := ncuUser.Retrieve()
	if err != nil {
		return CommitDataError
	}

	org := LeaderMap[ncuUser.Phone].Organization
	if org == "" {
		return TokenError
	}

	orgUser := &dao.User{Username: org}
	err = orgUser.Retrieve()
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return orgRegister(org, req)
		} else {
			model.ErrLog.Println(err)
			return ServerError
		}
	}

	orgPut(req, orgUser.ID)
	return SuccessCode
}

func Login(req *dto.Login) (*map[string]interface{}, uint) {
	user := &dao.User{Username: req.Username}
	err := user.Retrieve()
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			// 这里没记录，调云家园接口，如果是云家园账号，就注册
			return NCUOSLogin(req)
		} else {
			model.ErrLog.Println(err)
			return nil, ServerError
		}
	}

	userInfo := &dao.UserInfo{
		UserID: user.ID,
	}
	err = userInfo.Retrieve()
	if err != nil {
		model.ErrLog.Println(err)
		return nil, ServerError
	}

	// 社团账号
	if userInfo.Verification == "v" {
		return orgLogin(req, user)
	}

	// 非社团账号即云家园账号
	return NCUOSLogin(req)
}

func orgLogin(req *dto.Login, user *dao.User) (*map[string]interface{}, uint) {
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
		"id":           user.ID,
		"token":        token,
		"username":     user.Username,
		"verification": "v",
	}

	return data, SuccessCode
}

func NCUOSLogin(req *dto.Login) (*map[string]interface{}, uint) {
	NCUOS := &model.NCUOSOauth{}
	err := NCUOS.GetAccess(req.Username, req.Password)
	if err != nil {
		return nil, LoginError
	}

	return NCUOSTokenLogin(&dto.Token{Token: NCUOS.Token})
}

func NCUOSRegister(NCUOSUser *model.NCUOSUserProfileBasic) (*map[string]interface{}, uint) {
	// 注册操作;密码没用，随机生成
	password, err := util.RandHexStr(8)
	if err != nil {
		model.ErrLog.Println(err)
		return nil, ServerError
	}

	salt, err := util.RandHexStr(64)
	if err != nil {
		model.ErrLog.Println(err)
		return nil, ServerError
	}

	password = hex.EncodeToString(util.SHA512([]byte(password + salt)))
	user := &dao.User{
		Username:    NCUOSUser.Username,
		Password:    password,
		Salt:        salt,
		LoginStatus: "0",
		Phone:       NCUOSUser.Phone,
	}
	err = user.CreateWith(&dao.UserInfo{
		Nickname: NCUOSUser.Name,
		Sex:      NCUOSUser.Sex,
	})
	if err != nil {
		model.ErrLog.Println(err)
		return nil, ServerError
	}

	token, err := model.Jwt.GenerateToken(strconv.Itoa(int(user.ID)), user.LoginStatus)
	if err != nil {
		model.ErrLog.Println(err)
		return nil, ServerError
	}

	data := &map[string]interface{}{
		"id":       user.ID,
		"token":    token,
		"username": NCUOSUser.Username,
	}
	return data, SuccessCode
}

func NCUOSTokenLogin(req *dto.Token) (*map[string]interface{}, uint) {
	NCUOS := model.NCUOSOauth{
		Token: req.Token,
	}
	NCUOSUser, err := NCUOS.GetUserProfileBasic()
	if err != nil {
		return nil, TokenError
	}

	user := &dao.User{
		Username: NCUOSUser.Username,
	}
	err = user.Retrieve()
	if err != nil {
		// 首次登录进行注册操作,并返回相关信息
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return NCUOSRegister(NCUOSUser)
		}
		model.ErrLog.Println(err)
		return nil, ServerError
	}

	token, err := model.Jwt.GenerateToken(strconv.Itoa(int(user.ID)), user.LoginStatus)
	if err != nil {
		model.ErrLog.Println(err)
		return nil, ServerError
	}

	data := &map[string]interface{}{
		"id":       user.ID,
		"token":    token,
		"username": NCUOSUser.Username,
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

/*
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
*/

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

/*
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

*/
