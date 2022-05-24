package excel

import (
	"crawler/project/internal/utils"
	"log"
	"strings"

	"github.com/xuri/excelize/v2"
)

const (
	WebUrl int = iota
	GoogleAuthSecret
	Account
	Password
)

func GetCrawlerConfigFromExcel() []*utils.CrawlerConfig {

	var rowFist int = 0
	var columneName []string
	rowContent := make(map[string]string)

	var cc []*utils.CrawlerConfig
	rows := ExcelReader("", "config")

	for key, row := range rows {
		for colIndex, colCell := range row {
			if key == rowFist {
				columneName = append(columneName, colCell)
			} else {

				colCell = strings.ReplaceAll(colCell, " ", "")
				if columneName[colIndex] == "GoogleAuthSecret" {
					colCell = strings.ToUpper(colCell)
				}
				rowContent[columneName[colIndex]] = colCell
			}
		}

		if key == rowFist {
			continue
		}

		utils.Regex("tianshi.boinht.com:57859/admin.html#/", "[\\w.]+", 2)
		crawler := &utils.CrawlerConfig{
			WebUrl:           rowContent[columneName[WebUrl]],
			GoogleAuthSecret: rowContent[columneName[GoogleAuthSecret]],
			Account:          rowContent[columneName[Account]],
			Password:         rowContent[columneName[Password]],
		}
		cc = append(cc, crawler)
	}
	return cc
}

func ExcelReader(fileName, sheetName string) [][]string {
	fileName = utils.FilePath(fileName)
	// fmt.Println(fileName)

	f, err := excelize.OpenFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := f.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	rows, err := f.GetRows(sheetName)
	if err != nil {
		log.Fatal(err)
	}

	return rows

}
