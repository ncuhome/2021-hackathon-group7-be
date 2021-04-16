package service

import (
	"fmt"
	"tudo/model"
	"tudo/model/dao"
	"tudo/model/dto"
	"tudo/util"
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
	SetPasswordByEmailPage = "https://nspyf.top/tudo/password"
)

func SendBindEmailKey(req *dto.Email, id uint) uint {
	user := &dao.User{
		ID: id,
	}
	_ = user.Retrieve()
	if user.Email != "" {
		return EmailRepeated
	}

	key, err := util.RandDecStr(6)
	if err != nil {
		model.ErrLog.Println(err)
		return ServerError
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
		return EmailSendingError
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
		return ServerError
	}
	err = cacheObj.SetData(10 * time.Minute)
	if err != nil {
		model.ErrLog.Println(err)
		return ServerError
	}

	return SuccessCode
}

func BindEmail(req *dto.EmailBind, id uint) uint {
	user := &dao.User{
		ID: id,
	}
	_ = user.Retrieve()
	if user.Email != "" {
		return EmailRepeated
	}

	data := &dao.EmailBindCache{}
	cacheObj := &dao.JsonCache{
		Data: data,
		ID:   strconv.Itoa(int(id)),
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

	if req.Key != data.Key {
		return CodeError
	}
	if req.Email != data.Email {
		return CodeError
	}

	//判断邮箱是否被其它用户绑定
	userByEmail := &dao.User{
		Email: req.Email,
	}
	_ = userByEmail.Retrieve()
	if userByEmail.ID != 0 {
		return EmailUsed
	}

	//删缓存
	err = cacheObj.DelData()
	if err != nil {
		model.ErrLog.Println(err)
		return ServerError
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
		return ServerError
	}

	return SuccessCode
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
		return ServerError
	}
	return SuccessCode
}

func SendPasswordEmailKey(req *dto.Email) uint {
	user := &dao.User{
		Email: req.Email,
	}
	_ = user.Retrieve()
	if user.ID == 0 {
		return EmailNotBinding
	}

	key, err := util.RandHexStr(16)
	if err != nil {
		model.ErrLog.Println(err)
		return ServerError
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
		return EmailSendingError
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
		return ServerError
	}
	err = cacheObj.SetData(10 * time.Minute)
	if err != nil {
		model.ErrLog.Println(err)
		return ServerError
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
			return CodeError
		}
		model.ErrLog.Println(err)
		return ServerError
	}

	//删缓存
	err = cacheObj.DelData()
	if err != nil {
		model.ErrLog.Println(err)
		return ServerError
	}

	return updatePassword(req.NewPassword, data.ID)
}
