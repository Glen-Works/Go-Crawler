##登入器步驟
需先開啟Boin登入器，啟動後才能訪問網頁
設定音波後台登入網址，注意port每次開啟會更換，e.g. http://tianshi.boinht.com:57626/admin.html#/
IP 請對方綁定
請對方給secret 設定到 Excel GoogleAuthSecret

##excel 相關設定(telegram、login)
設定excel相關資料，WebUrl(網址),GoogleAuthSecret,Account,Password,TelegramGroupId,TelegramToken，資料夾位於crawler_data
TelegramToken Telegram 機器人token 取得方式參考(https://teleme.io/articles/create_your_own_telegram_bot?utm_source=web_console&utm_medium=EmptyBoardBot)
TelegramGroupId Telegram 聊天群組 取得方式參考(https://stackoverflow.com/questions/32423837/telegram-bot-how-to-get-a-group-chat-id)
#TelegramGroupId=-536412709 //(測試)
#TelegramGroupId=-760794150 //(預設正式)

##env 相關設定
RUN_TIMER 計時器 請參考 golang cron 的套件設定方式 (https://pkg.go.dev/github.com/robfig/cron@v1.2.0)(預設每30分鐘執行一次)
GOOGLE_AUTH_WAITTING_SECOND GooglAuth 的刷新等待，開啟瀏覽器登入時使用，如果電腦效能不好請往上加
WEB_OPERATE_WAITTING_SECOND 網頁跳轉等待時間正常>2，等待頁面資料loading，如果電腦效能不好請往上加

##env 相關設定(目前不需調整)
ENV_DEBUG debug 狀態(true/false)
WEB_RETRY_WAITTING_MiNUTE 重新執行時間 
WEB_RETRY_LIMIT_MiNUTE 重新執行時間限制
SELENIUM_DRIVE_PATH GoogleDrive 檔案路徑 
SELENIUM_DRIVE_NAME GoogleDrive 檔名
SELENIUM_PORT GoogleDrive 使用的port 
CRAWLER_DATA_PATH 上一次資料的儲存路徑(json)
URL_DATA_PATH 爬蟲網址路徑
USER_AGENT 網頁相關瀏覽器設定值
VARIETY_SETTING_STRING (現在資料-過去資料) 輸出名稱 "變化" 

##google相關設定
查看chrome version: chrome://version
下載ChromeDriver (須依照 chrome version)，放入crawler_data/(mac/windows/linux)資料夾中，(檔案名稱為預設)：https://sites.google.com/chromium.org/driver/downloads?authuser=0
安裝 jre (linux) : brew install xvfb openjdk-11-jre