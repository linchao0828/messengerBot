package utils

import (
	"fmt"
	"github.com/google/uuid"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// GetAge 根据身份证号获取年龄
func GetAge(idNo string) uint {
	if len(idNo) == 0 {
		return 0
	}

	birthDay := GetBirthDay(idNo)
	if birthDay == nil {
		return 0
	}
	now := time.Now()

	age := now.Year() - birthDay.Year()
	if now.Month() < birthDay.Month() {
		age = age - 1
	}
	if now.Month() == birthDay.Month() && now.Day() < birthDay.Day() {
		age = age - 1
	}

	if age <= 0 {
		return 0
	}

	return uint(age)
}

// GetBirthDay 根据身份证号获取生日（时间格式）
func GetBirthDay(idNo string) *time.Time {
	if len(idNo) == 0 {
		return nil
	}

	dayStr := idNo[6:14] + "000001"
	birthDay, err := time.Parse("20060102150405", dayStr)
	if err != nil {
		return nil
	}

	return &birthDay
}

func RandStrWithCapitalLetter(length int) string {
	str := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	var result []byte
	rand.Seed(time.Now().UnixNano() + int64(rand.Intn(100)))
	for i := 0; i < length; i++ {
		result = append(result, bytes[rand.Intn(len(bytes))])
	}
	return string(result)
}

func RandStrWithAllLetterAndNum(length int) string {
	str := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	bytes := []byte(str)
	var result []byte
	rand.Seed(time.Now().UnixNano() + int64(rand.Intn(100)))
	for i := 0; i < length; i++ {
		result = append(result, bytes[rand.Intn(len(bytes))])
	}
	return string(result)
}

// GenOid 目前定长24位
func GenOid() string {
	now := time.Now()
	date := now.Format("20060102")
	r := rand.Intn(1000)
	return fmt.Sprintf("%s%d%03d", date, now.UnixNano()/1e6, r)
}

// GenUuidWithoutBar 不带横杠的uuid
func GenUuidWithoutBar() string {
	return strings.Replace(uuid.NewString(), "-", "", -1)
}

func StrToi64(str string) int64 {
	ret, _ := strconv.ParseInt(str, 10, 64)
	return ret
}

func I64ToStr(num int64) string {
	return strconv.FormatInt(num, 10)
}

func StrTof64(str string) float64 {
	ret, _ := strconv.ParseFloat(str, 64)
	return ret
}
