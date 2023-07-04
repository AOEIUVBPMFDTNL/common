package github.com/AOEIUVBPMFDTNL/common

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"math"
	"regexp"
	"strconv"
	"strings"
	"text/template"
)

// Md5
func Md5(str []byte, b ...byte) string {
	h := md5.New()
	h.Write(str)
	return hex.EncodeToString(h.Sum(b))
}

// Xor
func Xor(input []byte, key []byte) []byte {
	output := make([]byte, len(input))
	for i := range input {
		output[i] = input[i] ^ key[i%len(key)]
	}
	return output
}

// Rot47
func Rot(input string) string {
	var result []string
	for i := range input[:] {
		j := int(input[i])
		if (j >= 33) && (j <= 126) {
			result = append(result, string(rune(33+((j+14)%94))))
		} else {
			result = append(result, string(input[i]))
		}
	}
	return strings.Join(result, "")
}

// 获取指定时间的零点
func GetZeroTime(date time.Time) time.Time {
	return time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())
}

// 获取指定时间的月份第一天零点
func GetMonthFirstDay(date time.Time) time.Time {
	return GetZeroTime(date.AddDate(0, 0, -date.Day()+1))
}

// 获取指定时间的月份最后一天零点
func GetMonthLastDay(date time.Time) time.Time {
	return GetMonthFirstDay(date).AddDate(0, 1, -1)
}

// 延迟指定工作日
func DelayWeekday(now time.Time, days, hour int) time.Time {
	for days > 0 {
		now = now.AddDate(0, 0, 1)
		if now.Weekday() != time.Sunday && now.Weekday() != time.Saturday {
			days--
		}
	}
	if hour < 0 {
		return now
	}
	return time.Date(now.Year(), now.Month(), now.Day(), hour, 0, 0, 0, time.Local)
}

// 四舍五入浮点数
func ToRound(val float64, precision int) float64 {
	p := math.Pow10(precision)
	return math.Floor(val*p+0.5) / p
}

// 转换为整数切片
func ToArray(array []string) (err error, sum int, res []int) {
	res = make([]int, len(array))
	var i int
	for k := range array {
		if i, err = strconv.Atoi(array[k]); err != nil {
			return
		} else {
			res[k] = i
			sum += i
		}
	}
	return
}

// 获取指定年份的生肖
func Zodiac(year int) string {
	return []string{"猴", "鸡", "狗", "猪", "鼠", "牛", "虎", "兔", "龙", "蛇", "马", "羊"}[year%12]
}

// 检查中国身份证
func CheckChinaIdCard(str string) bool {
	return regexp.MustCompile("(^\\d{15}$)|(^\\d{18}$)|(^\\d{17}(\\d|X|x)$)").MatchString(str)
}

// 检查台湾身份证
func CheckTaiwanIdCard(str string) bool {
	return regexp.MustCompile("[a-zA-Z][0-9]{9}").MatchString(str)
}

// 检查中国手机号
func CheckChinaPhoneNumber(str string) bool {
	return regexp.MustCompile("^1[3456789]{1}\\d{9}$").MatchString(str)
}

// 检查台湾手机号
func CheckTaiwanPhoneNumber(str string) bool {
	return regexp.MustCompile("^[0]?[9]\\d{8}$").MatchString(str)
}

// 解析转换文本消息模板
func ParseTemplate(name, character string, data any) (err error, result *bytes.Buffer) {
	tmpl, err := template.New(name).Parse(character)
	if err != nil {
		return
	}
	result = &bytes.Buffer{}
	err = tmpl.Execute(result, data)
	return
}
