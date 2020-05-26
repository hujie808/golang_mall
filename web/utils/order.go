package utils

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateOrderNumber() string {

	year := time.Now().Year()     //年
	month := time.Now().Month()   //月
	day := time.Now().Day()       //日
	hour := time.Now().Hour()     //小时
	minute := time.Now().Minute() //分钟
	second := time.Now().Second() //秒

	stryear := Int2String(year)        //年
	strmonth := Int2String(int(month)) //月
	strday := Int2String(day)          //日
	strhour := Int2String(hour)        //小时
	strminute := Int2String(minute)    //分钟
	strsecond := Int2String(second)    //秒

	strmonth2 := fmt.Sprintf("%02s", strmonth)
	strday2 := fmt.Sprintf("%02s", strday)
	strhour2 := fmt.Sprintf("%02s", strhour)
	strminute2 := fmt.Sprintf("%02s", strminute)
	strsecond2 := fmt.Sprintf("%02s", strsecond)

	randnum := rand.Intn(999999-100000) + 100000
	strrandnum := Int2String(randnum)

	return stryear + strmonth2 + strday2 + strhour2 + strminute2 + strsecond2 + strrandnum
}
