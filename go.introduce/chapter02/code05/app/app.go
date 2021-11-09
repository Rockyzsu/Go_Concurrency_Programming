package app

import (
	"go.introduce/chapter02/code05/app/config"
)

//GetVer 获取版本信息
func GetVer() string {
	return config.Cver
}
