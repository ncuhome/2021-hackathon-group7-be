package service

import (
	"fmt"
	"io"
	"tudo/model"
	"tudo/util"
)

func PostPicture(file io.Reader, name string) (*map[string]string, uint) {
	randStr, err := util.RandHexStr(8)
	if err != nil {
		return nil, ErrorServer
	}

	filename := randStr + name

	err = model.OssObj.PutBytes(file, filename)
	if err != nil {
		fmt.Println(err)
		return nil, ErrorServer
	}

	data := &map[string]string{
		"file": model.OssBaseUrl + filename,
	}
	return data, SuccessCode
}
