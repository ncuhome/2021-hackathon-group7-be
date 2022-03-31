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

	salt, err := util.RandHexStr(64) // desc 这里会随机 生成salt; salt 是2倍byteNum长度的乱码(只有数字和小写字母)
	if err != nil {
		model.ErrLog.Println(err)
		return ErrorServer
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
		return ErrorServer
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
		return ErrorServer
	}
	return SuccessCode
}

// 社团激活(注册)、修改资料
func OrgPostInfo(req *dto.OrgInfo, id uint) uint {
	ncuUser := &dao.User{ID: id}
	err := ncuUser.Retrieve()
	if err != nil {
		return ErrorCommitData
	}

	org := LeaderMap[ncuUser.Phone].Organization
	if org == "" {
		return ErrorToken
	}

	orgUser := &dao.User{Username: org}
	err = orgUser.Retrieve()
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return orgRegister(org, req)
		} else {
			model.ErrLog.Println(err)
			return ErrorServer
		}
	}

	orgPut(req, orgUser.ID)
	return SuccessCode
}

func Login(req *dto.Login) (*map[string]interface{}, uint) { //desc 这里 仅 区分 账号类型(云家园账号,社团账号)
	user := &dao.User{Username: req.Username}
	err := user.Retrieve()
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			// 这里没记录，调云家园接口，如果是云家园账号，就注册
			return NCUOSLogin(req)
		} else {
			model.ErrLog.Println(err)
			return nil, ErrorServer
		}
	}

	userInfo := &dao.UserInfo{
		UserID: user.ID,
	}
	err = userInfo.Retrieve()
	if err != nil {
		model.ErrLog.Println(err)
		return nil, ErrorServer
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
		return nil, ErrorLogin
	}

	token, err := model.Jwt.GenerateToken(user.LoginStatus, strconv.Itoa(int(user.ID)))
	if err != nil {
		model.ErrLog.Println(err)
		return nil, ErrorServer
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
		return nil, ErrorLogin
	}

	return NCUOSTokenLogin(&dto.Token{Token: NCUOS.Token})
}

func NCUOSRegister(NCUOSUser *model.NCUOSUserProfileBasic) (*map[string]interface{}, uint) {
	// 注册操作;密码没用，随机生成
	password, err := util.RandHexStr(8)
	if err != nil {
		model.ErrLog.Println(err)
		return nil, ErrorServer
	}

	salt, err := util.RandHexStr(64)
	if err != nil {
		model.ErrLog.Println(err)
		return nil, ErrorServer
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
		return nil, ErrorServer
	}

	token, err := model.Jwt.GenerateToken(user.LoginStatus, strconv.Itoa(int(user.ID)))
	if err != nil {
		model.ErrLog.Println(err)
		return nil, ErrorServer
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
		return nil, ErrorToken
	}

	user := &dao.User{
		Username: NCUOSUser.Username, // tip 这里的username是学号
	}
	err = user.Retrieve()
	if err != nil {
		// 首次登录进行注册操作,并返回相关信息
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return NCUOSRegister(NCUOSUser)
		}
		model.ErrLog.Println(err)
		return nil, ErrorServer
	}

	token, err := model.Jwt.GenerateToken(user.LoginStatus, strconv.Itoa(int(user.ID)))
	if err != nil {
		model.ErrLog.Println(err)
		return nil, ErrorServer
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
		return ErrorCommitData
	}
	for i := 0; i < usernameLen; i++ {
		if (username[i] < 'a' || 'z' < username[i]) && (username[i] < 'A' || 'Z' < username[i]) && (username[i] < '0' || '9' < username[i]) {
			return ErrorCommitData
		}
	}
	return SuccessCode
}

func CheckPassword(password string) uint {
	passwordLen := len(password)
	if passwordLen < 8 || passwordLen > 32 {
		return ErrorCommitData
	}
	//[33,126]覆盖了大小写字母、数字、普通可见符号
	for i := 0; i < passwordLen; i++ {
		if password[i] < 33 || password[i] > 126 {
			return ErrorCommitData
		}
	}
	return SuccessCode
}

//更新盐、个人登录状态、密码
func updatePassword(newPassword string, id uint) uint {
	saltStr, err := util.RandHexStr(64)
	if err != nil {
		model.ErrLog.Println(err)
		return ErrorServer
	}
	shaNewPassword := hex.EncodeToString(util.SHA512([]byte(newPassword + saltStr)))
	loginStatus, err := util.RandHexStr(8)
	if err != nil {
		model.ErrLog.Println(err)
		return ErrorServer
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
		return ErrorServer
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

// 获取用户角色 user, admin, team，并返回新token
func GetRole(id uint) (*map[string]interface{}, uint) {
	user := &dao.User{
		ID: id,
	}
	err := user.Retrieve()
	if err != nil {
		model.ErrLog.Println(err)
		return nil, ErrorServer
	}

	token, err := model.Jwt.GenerateToken(strconv.Itoa(int(id)), user.LoginStatus)
	if err != nil {
		model.ErrLog.Println(err)
		return nil, ErrorServer
	}

	data := &map[string]interface{}{
		"role":  "",
		"token": token,
	}

	if checkV(id) == SuccessCode {
		(*data)["role"] = "team"
		return data, SuccessCode
	}

	ncuUser := &dao.User{ID: id}
	err = ncuUser.Retrieve()
	if err != nil {
		return nil, ErrorCommitData
	}

	if _, ok := LeaderMap[user.Phone]; ok {
		(*data)["role"] = "admin"
		return data, SuccessCode
	}

	(*data)["role"] = "user"
	return data, SuccessCode
}
