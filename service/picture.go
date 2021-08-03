package service

import (
	"io"
	"tudo/model"
	"tudo/util"
)

const pictureBaseUrl = "https://nspyf.oss-cn-shanghai.aliyuncs.com/"

func PostPicture(file io.Reader, name string) (*map[string]string, uint) {
	randStr, err := util.RandHexStr(4)
	if err != nil {
		return nil, ServerError
	}

	filename := randStr + name

	err = model.OssObj.PutBytes(file, filename)
	if err != nil {
		return nil, ServerError
	}

	data := &map[string]string{
		"file": pictureBaseUrl + filename,
	}
	return data, SuccessCode
}
