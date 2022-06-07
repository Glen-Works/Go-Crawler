package main

import (
	"crawler/project/internal/excel"
	"crawler/project/internal/service"
	"crawler/project/internal/utils"
	"fmt"
	"os/exec"
	"runtime"
	"time"

	"net/http"
	_ "net/http/pprof"

	"github.com/robfig/cron"
)

func main() {
	c := cron.New()
	cronStr := utils.GetEnvData("RUN_TIMER")
	fmt.Println("排成設定值:", cronStr)
	StarCheck := false
	go func() {
		runData(&StarCheck)
	}()

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
	http.ListenAndServe("localhost:6060", nil)
	for {
		// fmt.Println("排成等待時間")
		// time.Sleep(1 * time.Second)
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
	CallClear()
	webCrawler := service.NewWebCrawlerService()
	*starCheck = webCrawler.CrawlerSearch(excel.GetCrawlerConfigFromExcel())
	fmt.Println("爬蟲執行完成")
}

var clear map[string]func() //create a map for storing clear funcs

func init() {
	clear = make(map[string]func()) //Initialize it
	clear["linux"] = func() {
		cmd := exec.Command("clear") //Linux example, its tested
		// cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["darwin"] = clear["linux"]
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
		// cmd := exec.Command("cmd", "cls") //Windows example, its tested
		// cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func CallClear() {
	value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
	if ok {                          //if we defined a clear func for that platform:
		value() //we execute it
	} else { //unsupported platform
		panic("clear cmd error")
	}
}
