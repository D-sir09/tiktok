package utils

import (
	"strconv"
	"time"
)

func GetMonthDay() string {

	//获取月份和日，MM-DD
	month := time.Now().Format("01")
	timeDay := time.Now().Day()
	day := strconv.Itoa(timeDay)
	if timeDay < 10 {
		day = "0" + day
	}
	now := month + "-" + day
	return now
}
