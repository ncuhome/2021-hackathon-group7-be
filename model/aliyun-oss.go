package model

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"io"
)

type OssType struct {
	AccessKeyID     string
	AccessKeySecret string
	Endpoint        string
	Bucket          string
}

func (s *OssType) PutBytes(data io.Reader, fileName string) error {
	// 创建OSSClient实例。
	client, err := oss.New(s.Endpoint, s.AccessKeyID, s.AccessKeySecret)
	if err != nil {
		ErrLog.Println(err)
		return err
	}

	// 获取存储空间。
	bucket, err := client.Bucket(s.Bucket)
	if err != nil {
		ErrLog.Println(err)
		return err
	}

	// 上传Byte数组。
	err = bucket.PutObject(fileName, data)
	if err != nil {
		ErrLog.Println(err)
		return err
	}
	return nil
}
