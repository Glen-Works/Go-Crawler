package service

import (
	"crawler/project/internal/data"
	"crawler/project/internal/utils"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/dgryski/dgoogauth"
	"github.com/tebeka/selenium"
)

type WebCrawler struct {
	timeBase             int64
	authWaitting         string
	crawlerDataPath      string
	varietySettingString string
}

func NewWebCrawlerService() *WebCrawler {
	return &WebCrawler{
		timeBase:             30,
		authWaitting:         "GOOGLE_AUTH_WAITTING_SECOND",
		crawlerDataPath:      "CRAWLER_DATA_PATH",
		varietySettingString: "VARIETY_SETTING_STRING",
	}
}

func (wc *WebCrawler) GetDataWriteToTg(cc []*utils.CrawlerConfig) {

	var wg sync.WaitGroup // 定义 WaitGroup
	for _, value := range cc {
		wg.Add(1) // 增加一个 wait 任务
		go func(s *utils.CrawlerConfig) {
			value.GoogleAuthCode = wc.search(value)
			crawlerData := wc.GetWebData(value)

			//read from json
			crawlerBefore := wc.ReadToJson(crawlerData, value.Account)
			fmt.Println("before data:", crawlerBefore)
			//write to json
			wc.WriteToJson(crawlerData, value.Account)
			fmt.Println("now data:", crawlerData)

			//variety
			varietyData := utils.VarietyTwo[map[string]string](crawlerData, crawlerBefore, utils.GetEnvData(wc.varietySettingString))
			fmt.Println("variety data:", varietyData)

			//TG
			crawlerStr := utils.StringFormatByList(crawlerData, "")
			varietyStr := utils.StringFormatByList(varietyData, utils.GetEnvData(wc.varietySettingString))
			fmt.Println(crawlerStr)
			fmt.Println(varietyStr)
			telegramRobotService := NewTelegramRobotService()
			telegramRobotService.sendMsg(crawlerStr)
			telegramRobotService.sendMsg(varietyStr)
			defer wg.Done()
		}(value)
	}
	wg.Wait()

}

func (wc *WebCrawler) getCrawlerFilePath(accountName string) string {
	return fmt.Sprintf("%s_%s%s", utils.FilePathByEnv(wc.crawlerDataPath), accountName, ".json")
}
func (wc *WebCrawler) ReadToJson(crawlerData map[string]string, accountName string) map[string]string {
	fielJsonName := wc.getCrawlerFilePath(accountName)
	file, _ := os.Open(fielJsonName)
	defer file.Close()
	decoder := json.NewDecoder(file)
	data := make(map[string]string)
	decoder.Decode(&data)
	utils.CheckKeyExist(data)
	return data
}

func (wc *WebCrawler) WriteToJson(crawlerData map[string]string, accountName string) {

	utils.CheckKeyExist(crawlerData)
	data := make(chan []byte, 1)
	jsonData, err := json.Marshal(crawlerData)
	if err != nil {
		log.Println(err, "\"格式化Json\"錯誤"+err.Error())
	}

	fielJsonName := wc.getCrawlerFilePath(accountName)
	data <- jsonData
	utils.WriteJsonFile(fielJsonName, data)
}

func (wc *WebCrawler) GetWebData(cc *utils.CrawlerConfig) map[string]string {

	seleniumService := NewWebSeleniumService()
	//取得網址與埠號
	sPath, sPort := seleniumService.GetSelPathAndPort()
	//取得chrome服務端
	service, err := seleniumService.SeleniumServiceSetting(sPath, sPort)

	if err != nil {
		log.Println(err) // panic is used only as an example and is not otherwise recommended.
	}
	//延遲關閉服務
	defer service.Stop()

	wd, err := seleniumService.SeleniumWebDriverSetting(sPort)
	if err != nil {
		log.Println(err)
	}
	//延遲退出chrome
	defer wd.Quit()

	//對頁面元素進行操作
	if err := wd.Get(cc.WebUrl); err != nil {
		log.Println(err, "\"登入頁面\"錯誤")
	}

	wdInput, err := wd.FindElements(selenium.ByTagName, "input")
	if err != nil {
		log.Println(err, "取登入\"input\"錯誤")
	}

	if len(wdInput) != 3 {
		log.Println(err, fmt.Sprintf("登入 input 數量錯誤，len(%d),%s，(e.g. account,password,googleToken)", wdInput, wdInput))
	}

	//寫入帳密跟google驗證碼
	for _, value := range wdInput {
		attrValue, _ := value.GetAttribute("placeholder")
		switch attrValue {
		case "请输入账号":
			// fmt.Println("请输入账号")
			value.SendKeys(cc.Account)
		case "请输入密码":
			// fmt.Println("请输入密码")
			value.SendKeys(cc.Password)
		case "请输入谷歌验证码":
			// fmt.Println("请输入谷歌验证码")
			value.SendKeys(cc.GoogleAuthCode)
		}
		// fmt.Println()
	}

	//登入按鈕
	btn, err := wd.FindElement(selenium.ByClassName, "ivu-btn-long")
	if err != nil {
		log.Println(err, "抓取\"login button\" 錯誤")
	}

	//送出
	btn.Click()

	//取得相關cookie值
	cookies, err := wd.GetCookies()
	if err != nil {
		log.Println(err, "抓取\"cookies\"錯誤")
	}

	//設定相關cookie值
	for _, cookie := range cookies {
		wd.AddCookie(&cookie)
	}

	//處理url 字尾/(組合起來，會多一個)
	url := cc.WebUrl
	if url[len(url)-1:] == "/" {
		url = url[:len(url)-1]
	}
	newUrl := url + utils.GetEnvData(seleniumService.urlDataPath)

	//所有會跳轉或送出，需等待後再執行
	wdInterface, err := seleniumService.ExecuteFunc(func() (interface{}, error) {
		return wd, wd.Get(newUrl)
	})
	wd = wdInterface.(selenium.WebDriver)
	if err != nil {
		log.Println(err, "\"/LiveTotalData\"錯誤")
	}

	tagDataItem, err := wd.FindElements(selenium.ByClassName, "dataItem")
	if err != nil {
		log.Println(err, "取登入\"input\"錯誤")
	}

	return wc.getDataInfoByDom(tagDataItem)

}

func (wc *WebCrawler) getDataInfoByDom(tagDataItem []selenium.WebElement) map[string]string {

	var dataInfo = make(map[string]string)
	//取出特定的span標題
	for _, tagValue := range tagDataItem {

		tagSpan, err := tagValue.FindElement(selenium.ByTagName, "span")
		if err != nil {
			log.Println(err, "取資料\"span tag\"錯誤")
			continue
		}

		span, err := tagSpan.Text()
		if err != nil {
			log.Println(err, "取資料\"span 文字\"錯誤")
			continue
		}
		spanText := strings.Join(utils.Regex(span, "[\u4E00-\u9FA5]+", -1), "")
		DataName := data.GetSearchTitleKey(data.TITLE_NAME, spanText)
		if len(DataName) == 0 {
			continue
		}

		tr, err := tagValue.FindElements(selenium.ByTagName, "tr")
		if err != nil {
			log.Println(err, "取資料\"tr tag\"錯誤")
			continue
		}

		for _, trValue := range tr {

			trText, err := trValue.Text()
			if err != nil {
				log.Println(err, "取資料\"tr text\"錯誤")
				continue
			}

			columnName := strings.Join(utils.Regex(trText, "[\u4E00-\u9FA5]+[(]*[\u4E00-\u9FA5]*[)]*[\u4E00-\u9FA5]*", -1), "")
			for _, DataNameValue := range DataName {
				checkColumnName := data.GetSearchInfo(data.COLUMN_NAME, DataNameValue)
				if columnName == checkColumnName {
					td, err := trValue.FindElements(selenium.ByTagName, "td")
					if err != nil {
						log.Println(err, "取資料\"td tag\"錯誤")
						continue
					}
					tdText, err := td[1].Text()
					if err != nil {
						log.Println(err, "取資料\"td text\"錯誤")
						continue
					}
					dataInfo[DataNameValue] = tdText
					DataName = utils.DeleteArrayByValue[string](DataName, DataNameValue)
					break
				}
			}

		}

	}

	//睡眠1秒後退出
	time.Sleep(1 * time.Second)
	return dataInfo
}

func (wc *WebCrawler) search(cc *utils.CrawlerConfig) string {

	count, remainTime := wc.getUnixCountAndDuration(wc.timeBase)
	authWaitting, _ := strconv.ParseFloat(utils.GetEnvData(wc.authWaitting), 64)
	if delayTime := float64(wc.timeBase - remainTime); delayTime <= authWaitting {
		delayTime += 2
		time.Sleep(time.Duration(delayTime * float64(time.Second)))
		count, _ = wc.getUnixCountAndDuration(wc.timeBase)
		// fmt.Println("delayTime " + strconv.FormatInt(delayTime, 10))
	}

	authCode := dgoogauth.ComputeCode(cc.GoogleAuthSecret, count)
	return strconv.Itoa(authCode)

}

func (wc *WebCrawler) getUnixCountAndDuration(timeBase int64) (int64, int64) {
	timeNow := time.Now().Unix()
	return timeNow / timeBase, timeNow % timeBase
}
