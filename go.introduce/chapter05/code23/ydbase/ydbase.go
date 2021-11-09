package ydbase

import (
	"strconv"
	"strings"
)

func Insert(stu Student) bool {

	len1 := len(studentTable)
	cap1 := cap(studentTable)
	if len1 < cap1 {
		studentTable = append(studentTable, stu)
	}
	studentTable[len1] = stu
	return true
}
func Count() int {

	return len(studentTable)
}
func Modify(stu Student) bool {

	for i, v := range studentTable {
		if v.XH == stu.XH {
			//v = stu //修改的不是同一个对象
			studentTable[i] = stu
			break
		}
	}
	return true
}
func Delete(xh string) bool {

	for i, v := range studentTable {
		if v.XH == xh {
			studentTable = append(studentTable[:i], studentTable[i+1:]...)
			break
		}
	}
	return true
}

//支持 >= == <= like
func Query(exp string) []Student {
	stu := make([]Student, 0)

	where := make([]string, 3)
	if strings.Contains(exp, ">=") {
		where = strings.Split(exp, ">=")
		if len(where) == 2 {
			for _, v := range studentTable {
				if strings.ToUpper(strings.Trim(where[0], " ")) == "AGE" {
					age, _ := strconv.Atoi(where[1])
					if v.Age >= age {
						stu = append(stu, v)
					}
				}

			}
		}

	}
	if strings.Contains(exp, "<=") {
		where = strings.Split(exp, "<=")
		if len(where) == 2 {
			for _, v := range studentTable {
				if strings.ToUpper(strings.Trim(where[0], " ")) == "AGE" {
					age, _ := strconv.Atoi(where[1])
					if v.Age <= age {
						stu = append(stu, v)
					}
				}

			}
		}
	}
	if strings.Contains(exp, "==") {
		where = strings.Split(exp, "==")
		if len(where) == 2 {
			for _, v := range studentTable {
				if strings.ToUpper(strings.Trim(where[0], " ")) == "XH" {
					if v.XH == strings.Trim(where[1], " ") {
						stu = append(stu, v)
					}
				}
				if strings.ToUpper(strings.Trim(where[0], " ")) == "NAME" {
					if v.Name == strings.Trim(where[1], " ") {
						stu = append(stu, v)
					}
				}
			}
		}
	}
	if strings.Contains(exp, "like") {
		where = strings.Split(exp, "like")
		if len(where) == 2 {
			for _, v := range studentTable {
				if strings.ToUpper(strings.Trim(where[0], " ")) == "XH" {
					if strings.Contains(v.XH, strings.Trim(where[1], " ")) {
						stu = append(stu, v)
					}
				}
				if strings.ToUpper(strings.Trim(where[0], " ")) == "NAME" {
					if strings.Contains(v.Name, strings.Trim(where[1], " ")) {
						stu = append(stu, v)
					}
				}
			}
		}
	}
	return stu
}
