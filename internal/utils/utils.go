package utils

import (
	"crawler/project/internal/data"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strconv"
	"time"
)

func Regex(RegexStr, filterRule string, GetCount int) []string {
	reg, _ := regexp.Compile(filterRule)
	return reg.FindAllString(RegexStr, GetCount)
}

func FilePath(fileName string) string {
	if fileName == "" {
		fileName = "crawler.xlsx"
	}

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	return fmt.Sprintf("%s/%s", dir, fileName)
}

func FilePathByEnv(fileName string) string {
	return FilePath(GetEnvData(fileName))
}

func WriteJsonFile(fielJsonName string, jsonData <-chan []byte) {
	ioutil.WriteFile(fielJsonName, <-jsonData, os.ModePerm)
}

func TypeToString[T interface{}](a []T) string {
	var steamStr string
	for _, v := range a {
		steamStr = fmt.Sprintf("%v %s", steamStr, v)
	}
	return steamStr
}

func DeleteArrayByValue[T int | string](arr []T, value T) []T {

	for i := 0; i < len(arr); i++ {
		if arr[i] == value {
			arr = append(arr[:i], arr[i+1:]...)
		}
	}
	return arr
}

func CheckKeyExist[T map[string]string](crawlerData T) {

	for _, val := range data.GetSearchList() {
		if _, ok := crawlerData[val]; !ok {
			crawlerData[val] = "0"
		}
	}
}

func VarietyTwo[T map[string]string](crawler, crawlerBefore T, settingStr string) T {
	var variety = make(T)
	for _, val := range data.GetSearchList() {
		nowValue := getArrayFirst(Regex(crawler[val], "[0-9]+", 1), 0)
		beforeValue := getArrayFirst(Regex(crawlerBefore[val], "[0-9]+", 1), 0)

		variety[val+settingStr] = strconv.Itoa(nowValue - beforeValue)

	}
	return variety
}

func getArrayFirst(data []string, returnData int) int {

	if len(data) == 0 {
		return returnData
	}

	dataValue, err := strconv.Atoi(data[0])
	if err != nil {
		return returnData
	}

	return dataValue
}

func StringFormatByList(crawlerData map[string]string, settingStr string) string {
	var crawlerStr string
	for _, val := range data.GetSearchList() {
		crawlerStr += fmt.Sprintf("%s:%s \n", val+settingStr, crawlerData[val+settingStr])
	}
	return crawlerStr
}

func GetTimeNow(columnName string) string {
	ShanghaiTimeZone, _ := time.LoadLocation("Asia/Shanghai")
	newTime := time.Now().UTC().In(ShanghaiTimeZone)
	return columnName + newTime.Format("2006-01-02 15:04") + "\n"
}
