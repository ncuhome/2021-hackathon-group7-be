package model

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type NCUOSOauth struct {
	Token string `json:"token"`
}

type XbInfo struct {
	Dm string `json:"dm"`
}
type BaseInfo struct {
	Xm   string `json:"xm"`
	Xh   string `json:"xh"`
	Xb   XbInfo `json:"xb"`
	Yddh string `json:"yddh"`
}
type NCUOSUserProfileBasicInfo struct {
	BaseInfo BaseInfo `json:"base_info"`
}
type NCUOSUserProfileBasic struct {
	Username string `json:"username"`
	Name     string `json:"name"`
	Sex      string `json:"sex"`
	Phone    string `json:"phone"`
}

func (s *NCUOSOauth) GetAccess(username string, password string) error {
	url := "https://os.ncuos.com/api/user/token"
	method := "POST"

	payloadStr := fmt.Sprintf(`{"username":"%v","password":"%v"}`, username, password)
	payload := strings.NewReader(payloadStr)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		ErrLog.Println(err)
		return err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		ErrLog.Println(err)
		return err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		ErrLog.Println(err)
		return err
	}

	err = json.Unmarshal(body, s)
	if err != nil {
		ErrLog.Println(err)
		return err
	}

	if s.Token == "" {
		return errors.New("NCUOSOauth.GetAccess error")
	}

	return nil
}

// 先调用GetAccess
func (s *NCUOSOauth) GetUserProfileBasic() (*NCUOSUserProfileBasic, error) {
	url := "https://os.ncuos.com/api/user/profile/basic"
	method := "GET"

	payload := strings.NewReader(``)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		ErrLog.Println(err)
		return nil, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", "passport "+s.Token)

	res, err := client.Do(req)
	if err != nil {
		ErrLog.Println(err)
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		ErrLog.Println(err)
		return nil, err
	}

	userInfo := &NCUOSUserProfileBasicInfo{}
	err = json.Unmarshal(body, userInfo)
	if err != nil {
		ErrLog.Println(err)
		return nil, err
	}

	user := &NCUOSUserProfileBasic{
		Username: userInfo.BaseInfo.Xh,
		Name:     userInfo.BaseInfo.Xm,
		Sex:      userInfo.BaseInfo.Xb.Dm,
		Phone:    userInfo.BaseInfo.Yddh,
	}
	if user.Name == "" {
		return nil, errors.New("NCUOSOauth.GETGetUserProfileBasic error")
	}

	return user, nil
}

func (s *NCUOSOauth) GetUser(username string, password string) (*NCUOSUserProfileBasic, error) {
	err := s.GetAccess(username, password)
	if err != nil {
		ErrLog.Println(err)
		return nil, err
	}

	user, err := s.GetUserProfileBasic()
	if err != nil {
		ErrLog.Println(err)
		return nil, err
	}

	return user, nil
}
