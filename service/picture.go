package service

import (
	"io"
	"nspyf/model"
	"nspyf/util"
)

func PostPicture(file io.Reader,name string) (*map[string]string, uint) {
	randStr, err := util.RandHexStr(4)
	if err != nil {
		return nil,ServerError
	}

	filename := randStr + name

	//err = model.PictureObj.Post(file, filename)
	err = model.OssObj.PutBytes(file, filename)
	if err != nil {
		return nil,ServerError
	}

	data := &map[string]string{
		"filename": randStr + filename,
	}
	return data, SuccessCode
}