package main

import (
	"crawler/project/internal/excel"
	"crawler/project/internal/service"
	"crawler/project/internal/utils"
	"fmt"

	"github.com/robfig/cron"
)

func main() {
	// webCrawler := service.NewWebCrawlerService()
	// webCrawler.GetDataWriteToTg(excel.GetCrawlerConfigFromExcel())

	c := cron.New()
	cronStr := utils.GetEnvData("RUN_TIMER")
	fmt.Println("排成設定值:", cronStr)
	err := c.AddFunc(cronStr, func() {
		webCrawler := service.NewWebCrawlerService()
		webCrawler.GetDataWriteToTg(excel.GetCrawlerConfigFromExcel())
		fmt.Println("爬蟲執行完成")
	})

	if err != nil {
		fmt.Errorf("爬蟲執行錯誤： %v", err)
		return
	}
	c.Start()

	defer c.Stop()
	select {}

}
