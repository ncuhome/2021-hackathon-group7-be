package service

import (
	"fmt"
	"nspyf/model"
	"nspyf/model/dao"
	"nspyf/model/dto"
	"nspyf/util"
	"strconv"
	"time"
)

const (
	BindEmailFormat = "尊敬的：%v<br>" +
		"您收到这封邮件，是因为有用户进行邮箱绑定操作时填写了您的邮箱。<br>" +
		"若非本人操作，请忽略这封邮件。<br>" +
		"尊敬的用户：%v<br>" +
		"以下是您邮箱绑定的验证码，验证码有效期为10分钟<br>" +
		"验证码：%v<br>" +
		"若验证码过期请到绑定邮箱相关页面重新发送<br>" +
		"该邮件由系统发出，请勿回复"
	PasswordEmailFormat = "尊敬的：%v<br>" +
		"您收到这封邮件，是因为有用户通过邮箱更改密码时填写了您的邮箱。<br>" +
		"若非本人操作，请忽略这封邮件。<br>" +
		"尊敬的用户：%v<br>" +
		"以下是您重设密码的链接，链接有效期为10分钟<br>" +
		`重设密码链接：<a href="%v">%v</a><br>` +
		"若非链接形式，请复制该链接粘贴到浏览器网址栏访问<br>" +
		"若链接过期请到忘记密码相关页面重新发送<br>" +
		"该邮件由系统发出，请勿回复"
	SetPasswordByEmailPage = "https://nspyf.top/password"
)

func SendBindEmailKey(req *dto.Email, id uint) uint {
	user := &dao.User{
		ID: id,
	}
	_ = user.Retrieve()
	if user.Email != "" {
		return 8
	}

	key, err := util.RandDecStr(6)
	if err != nil {
		model.ErrLog.Println(err)
		return 1
	}

	content := fmt.Sprintf(BindEmailFormat, req.Email, user.Username, key)

	err = model.Email.Send(
		req.Email,
		user.Username,
		"邮箱绑定",
		"text/html",

		//TODO frontend url
		//TODO html template

		content,
	)
	if err != nil {
		model.ErrLog.Println(err)
		return 9
	}

	//确保邮箱有效再设缓存
	//删除缓存并设置新缓存
	cacheObj := &dao.JsonCache{
		Data: &dao.EmailBindCache{
			ID:    id,
			Email: req.Email,
			Key:   key,
		},
		ID: strconv.Itoa(int(id)),
	}
	err = cacheObj.DelData()
	if err != nil {
		model.ErrLog.Println(err)
		return 1
	}
	err = cacheObj.SetData(10 * time.Minute)
	if err != nil {
		model.ErrLog.Println(err)
		return 1
	}

	return 0
}

func BindEmail(req *dto.EmailBind, id uint) uint {
	user := &dao.User{
		ID: id,
	}
	_ = user.Retrieve()
	if user.Email != "" {
		return 8
	}

	data := &dao.EmailBindCache{}
	cacheObj := &dao.JsonCache{
		Data: data,
		ID:   strconv.Itoa(int(id)),
	}
	err := cacheObj.GetData()

	if err != nil {
		if err == dao.CacheNil {
			return 10
		} else {
			model.ErrLog.Println(err)
			return 2
		}
	}

	if req.Key != data.Key {
		return 11
	}
	if req.Email != data.Email {
		return 12
	}

	//判断邮箱是否被其它用户绑定
	userByEmail := &dao.User{
		Email: req.Email,
	}
	_ = userByEmail.Retrieve()
	if userByEmail.ID != 0 {
		return 13
	}

	//删缓存
	err = cacheObj.DelData()
	if err != nil {
		model.ErrLog.Println(err)
		return 1
	}

	//将邮箱写入数据库
	user = &dao.User{
		ID: id,
	}
	newUser := &dao.User{
		Email: req.Email,
	}
	err = user.Update(newUser)
	if err != nil {
		model.ErrLog.Println(err)
		return 1
	}

	return 0
}

func RemoveEmail(id uint) uint {
	user := &dao.User{
		ID: id,
	}
	err := user.Update(map[string]interface{}{
		"Email": "",
	})
	if err != nil {
		model.ErrLog.Println(err)
		return 1
	}
	return 0
}

func SendPasswordEmailKey(req *dto.Email) uint {
	user := &dao.User{
		Email: req.Email,
	}
	_ = user.Retrieve()
	if user.ID == 0 {
		return 15
	}

	key, err := util.RandHexStr(16)
	if err != nil {
		model.ErrLog.Println(err)
		return 1
	}
	url := SetPasswordByEmailPage + "?key=" + key
	content := fmt.Sprintf(PasswordEmailFormat, req.Email, user.Username, url, url)

	err = model.Email.Send(
		req.Email,
		user.Username,
		"更改密码",
		"text/html",

		//TODO frontend url
		//TODO html template

		content,
	)
	if err != nil {
		model.ErrLog.Println(err)
		return 9
	}

	//确保邮箱有效再设缓存
	//删除缓存并设置新缓存
	cacheObj := &dao.JsonCache{
		Data: &dao.EmailPasswordCache{
			ID:  user.ID,
			Key: key,
		},
		ID: key,
	}
	err = cacheObj.DelData()
	if err != nil {
		model.ErrLog.Println(err)
		return 1
	}
	err = cacheObj.SetData(10 * time.Minute)
	if err != nil {
		model.ErrLog.Println(err)
		return 1
	}

	return 0
}

func SetPasswordByEmail(req *dto.SetPasswordByEmailReq) uint {
	code := CheckPassword(req.NewPassword)
	if code != 0 {
		return code
	}

	data := &dao.EmailPasswordCache{}
	cacheObj := &dao.JsonCache{
		Data: data,
		ID:   req.Key,
	}
	err := cacheObj.GetData()
	if err != nil {
		if err == dao.CacheNil {
			return 16
		}
		model.ErrLog.Println(err)
		return 1
	}

	//删缓存
	err = cacheObj.DelData()
	if err != nil {
		model.ErrLog.Println(err)
		return 1
	}

	return updatePassword(req.NewPassword, data.ID)
}
