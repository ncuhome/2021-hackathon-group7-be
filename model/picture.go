package model

import (
	"io"
	"os"
)

type Picture struct {
	Path string
}

func (s *Picture) Post(data io.Reader, fileName string) error {
	file, err := os.Create(s.Path + fileName)
	if err != nil {
		return err
	}
	_, err = io.Copy(file, data)
	if err != nil {
		return err
	}
	return nil
}
