package controller

import (
	"nspyf/util"
)

type GinConfig struct {
	Port string `json:"port"`
}

var GinConfigObj GinConfig

func GinInit(path string) {
	if err := util.ReadJSON(path, &GinConfigObj); err != nil {
		panic(err)
	}
	return
}
