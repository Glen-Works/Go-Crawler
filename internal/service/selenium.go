package service

import (
	"crawler/project/internal/utils"
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"
	"time"

	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

type Selenium struct {
	seleniumPath           string
	seleniumName           string
	seleniumPort           string
	browserAgent           string
	urlDataPath            string
	urlReportPath          string
	webOperateWaittingTime string
}

func NewWebSeleniumService() *Selenium {
	return &Selenium{
		seleniumPath:           "SELENIUM_DRIVE_PATH",
		seleniumName:           "SELENIUM_DRIVE_NAME",
		seleniumPort:           "SELENIUM_PORT",
		browserAgent:           "USER_AGENT",
		urlDataPath:            "URL_DATA_PATH",
		urlReportPath:          "URL_REPORT_PATH",
		webOperateWaittingTime: "WEB_OPERATE_WAITTING_SECOND",
	}
}

type delegFunc func() (interface{}, error)

func (se *Selenium) ExecuteFunc(fun delegFunc) (interface{}, error) {
	waittingTime := 1 * time.Second
	OperateWaittingTime, _ := strconv.ParseFloat(utils.GetEnvData(se.webOperateWaittingTime), 64)
	if OperateWaittingTime > 0 {
		// waittingTime = time.Duration(OperateWaittingTime * float64(time.Second))
		waittingTime = utils.GetTimeSecond[float64](OperateWaittingTime)
	}

	time.Sleep(waittingTime)
	Data, err := fun()
	time.Sleep(waittingTime)
	return Data, err
}

func (se *Selenium) SeleniumWebDriverSetting(sPort int) (selenium.WebDriver, error) {
	//呼叫瀏覽器
	//設定瀏覽器相容性，我們設定瀏覽器名稱為chrome
	caps := selenium.Capabilities{"browserName": "chrome"}

	arg := []string{
		"--headless", // 设置Chrome无头模式
		fmt.Sprintf("--user-agent=%s", utils.GetEnvData(se.browserAgent)), // 模拟user-agent，防反爬
		"--ignore-certificate-errors",
	}

	if runtime.GOOS == "windows" {
		arg = append(arg, "--disable-gpu")
	}

	chromeCaps := chrome.Capabilities{
		Path: "",
		Args: arg,
	}

	//以上是设置浏览器参数
	caps.AddChrome(chromeCaps)

	//呼叫瀏覽器urlPrefix: 測試參考：DefaultURLPrefix = "http://127.0.0.1:4444/wd/hub"
	return selenium.NewRemote(caps, fmt.Sprintf("http://127.0.0.1:%d/wd/hub", sPort))

}

func (se *Selenium) SeleniumServiceSetting(path string, port int) (*selenium.Service, error) {

	opts := []selenium.ServiceOption{
		// selenium.Output(os.Stderr), // Output debug information to STDERR.
	}

	if utils.GetIsDebug() {
		opts = append(opts, selenium.Output(os.Stderr)) // Output debug information to STDERR.
	}

	if runtime.GOOS == "linux" {
		opts = append(opts, selenium.StartFrameBuffer()) // Start an X frame buffer for the browser to run in.
	}

	// selenium.SetDebug(true)

	service, err := selenium.NewChromeDriverService(path, port, opts...)

	return service, err
}

func (se *Selenium) GetSelPathAndPort() (string, int) {

	// path := utils.GetEnvData(se.seleniumPath)
	path := utils.GetChromeDriveFilePath(utils.GetEnvData(se.seleniumPath), utils.GetEnvData(se.seleniumName))
	port := utils.GetEnvData(se.seleniumPort)

	sPath := utils.FilePath(path)
	sPort, err := strconv.Atoi(port)
	if err != nil {
		log.Fatal(err, ", port need to be integer")
	}
	return sPath, sPort
}
