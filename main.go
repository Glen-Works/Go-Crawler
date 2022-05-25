package main

import (
	"crawler/project/internal/excel"
	"crawler/project/internal/service"
	"crawler/project/internal/utils"
	"fmt"
	"time"

	"github.com/robfig/cron"
)

func main() {
	c := cron.New()
	cronStr := utils.GetEnvData("RUN_TIMER")
	fmt.Println("排成設定值:", cronStr)
	StarCheck := false
	runData(&StarCheck)

	err := c.AddFunc(cronStr, func() {
		if !StarCheck {
			fmt.Printf("\n等待上一個爬蟲執行，此次時間:%v、\n", time.Now())
			return
		}

		StarCheck = false
		runData(&StarCheck)

	})

	if err != nil {
		fmt.Errorf("爬蟲執行錯誤： %v", err)
		return
	}
	c.Start()

	defer c.Stop()
	for {
		// fmt.Println("排成等待時間")
		time.Sleep(1 * time.Minute)
	}

}

func runData(starCheck *bool) {
	// func runData(starCheck *bool, runLimitTime time.Time) {
	// defer func() {
	// 	// 可以取得 err 的回傳值
	// 	if r := recover(); r != nil {
	// 		// 超過時間跳出
	// 		if time.Now().After(runLimitTime) {
	// 			log.Printf("\n爬蟲抓取錯誤，錯誤:%s \n", r)
	// 			return
	// 		}
	// 		fmt.Println("爬蟲抓取錯誤，重新執行，錯誤:", r)
	// 		runData(starCheck, runLimitTime)
	// 	}
	// }()
	webCrawler := service.NewWebCrawlerService()
	*starCheck = webCrawler.CrawlerSearch(excel.GetCrawlerConfigFromExcel())
	fmt.Println("爬蟲執行完成")
}
