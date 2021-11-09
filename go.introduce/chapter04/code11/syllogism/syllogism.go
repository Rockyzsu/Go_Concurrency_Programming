package syllogism

import (
	"errors"
	"fmt"
	"strconv"
)

func isJudgeChar(n string) bool {
	switch n {
	case "A", "E", "I", "O":
		{
			return true
		}
	default:
		return false
	}
}

//IsValid 判断三段式的模式是否有效
//
//gz是第几格 1,2,3,4
//s,m,p取值范围为A,E,I,O
func IsValid(gz int, s, m, p string) (bool, error) {
	if gz > 4 || gz < 1 {
		return false, errors.New("格子数只能是1-4")
	}
	if !isJudgeChar(s) {
		return false, errors.New("判断只能是A,E,I,O")
	}
	if !isJudgeChar(m) {
		return false, errors.New("判断只能是A,E,I,O")
	}
	if !isJudgeChar(p) {
		return false, errors.New("判断只能是A,E,I,O")
	}
	ret := s + m + p + strconv.Itoa(gz)
	fmt.Println(ret)
	switch ret {
	case "AAA1", "EAE1", "AII1", "EIO1":
		{
			return true, nil
		}
	case "AEE2", "EAE2", "AOO2", "EIO2":
		{
			return true, nil
		}
	case "AAI3", "EAO3", "AII3", "EIO3", "IAI3", "OAO3":
		{
			return true, nil
		}
	case "AAI4", "EAO4", "AEE4", "EIO4", "IAI4":
		{
			return true, nil
		}
	default:
		return false, nil
	}

}
