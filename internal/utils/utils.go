package utils

import (
	"crawler/project/internal/data"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"runtime"
	"strconv"
	"strings"
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

	os := strings.Trim(runtime.GOOS, " ")
	if os == "windows" {
		fileName = strings.ReplaceAll(fileName, "/", "\\")
		return fmt.Sprintf("%s\\%s", dir, fileName)
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
	// ShanghaiTimeZone, _ := time.LoadLocation("Asia/Shanghai")
	// localTime, err := time.ParseInLocation("2006-01-02 15:04:05", "2017-12-03 22:01:02", ShanghaiTimeZone)
	// newTime := time.Now().UTC().In(ShanghaiTimeZone)
	t, _ := time.ParseInLocation("2006-01-02 15:04:05", time.Now().Format("2006-01-02 15:04:05"), time.Local)
	return columnName + t.Format("2006-01-02 15:04") + "\n"
	// return columnName + localTime.Format("2006-01-02 15:04") + "\n"
}

func GetTimeSecond[T float64 | int64](second T) time.Duration {
	return time.Duration(second * T(time.Second))
}

func GetRetryWaittingTime() time.Duration {
	waittingTime := 1 * time.Minute
	OperateWaittingTime, _ := strconv.ParseFloat(GetEnvData("WEB_RETRY_WAITTING_MiNUTE"), 64)
	if OperateWaittingTime > 0 {
		waittingTime = GetTimeSecond[float64](OperateWaittingTime * 60)
	}
	return waittingTime
}

func GetRetryLimitTime() time.Duration {
	retryLimitTime := 1 * time.Minute
	OperateWaittingTime, _ := strconv.ParseFloat(GetEnvData("WEB_RETRY_LIMIT_MiNUTE"), 64)
	if OperateWaittingTime > 0 {
		retryLimitTime = GetTimeSecond[float64](OperateWaittingTime * 60)
	}
	return retryLimitTime
}

func GetIsDebug() bool {
	debug, err := strconv.ParseBool(strings.ToUpper(GetEnvData("ENV_DEBUG")))
	if err != nil {
		return false
	}

	return debug
}
func GetChromeDriveFilePath(filePath, fileName string) string {

	os := strings.Trim(runtime.GOOS, " ")
	if os == "windows" {
		filePath = fmt.Sprintf("%s\\%s", filePath, "windows")
		fileName = fileName + ".exe"
		return strings.ReplaceAll(fmt.Sprintf("%s\\%s", filePath, fileName), "\\\\", "\\")
	}

	switch os {
	case "darwin":
		filePath = filePath + "/mac"
	default:
		filePath = filePath + "/linux"

	}
	return fmt.Sprintf("%s/%s", filePath, fileName)
}
