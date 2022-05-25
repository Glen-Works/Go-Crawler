package main

import (
	"crawler/project/internal/excel"
	"crawler/project/internal/service"
	"crawler/project/internal/utils"
	"fmt"
	"log"
	"time"

	"github.com/robfig/cron"
)

func main() {
	c := cron.New()
	cronStr := utils.GetEnvData("RUN_TIMER")
	fmt.Println("排成設定值:", cronStr)
	StarCheck := false
	runLimitTime := time.Now().Add(utils.GetRetryLimitTime())
	runData(&StarCheck, runLimitTime)

	err := c.AddFunc(cronStr, func() {
		if !StarCheck {
			fmt.Println("wait last crawler")
			return
		}
		time.Sleep(time.Second * 30)
		StarCheck = false
		runLimitTime := time.Now().Add(utils.GetRetryLimitTime())
		runData(&StarCheck, runLimitTime)
		fmt.Println("爬蟲執行完成")
	})

	if err != nil {
		fmt.Errorf("爬蟲執行錯誤： %v", err)
		return
	}
	c.Start()

	defer c.Stop()
	for {
		fmt.Println("for loop")
		time.Sleep(2 * time.Second)
	}

}

func runData(starCheck *bool, runLimitTime time.Time) {
	defer func() {
		// 可以取得 err 的回傳值
		if r := recover(); r != nil {
			// 超過時間跳出
			if time.Now().After(runLimitTime) {
				log.Printf("\n抓取錯誤，錯誤:%s \n", r)
				return
			}
			fmt.Println("抓取錯誤，重新執行，錯誤:", r)
			runData(starCheck, runLimitTime)
		}
	}()
	webCrawler := service.NewWebCrawlerService()
	webCrawler.CrawlerSearch(excel.GetCrawlerConfigFromExcel())
	*starCheck = true
}
