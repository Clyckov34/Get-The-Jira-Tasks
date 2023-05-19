package excel

import (
	"fmt"
	"rjt/internal/jira"
	"rjt/pkg/excel"

	"github.com/xuri/excelize/v2"
)

//Write создает файл excel
func Write(rsp *jira.Response, fp *jira.FlagParameters) (string, error) {
	var excelFile = &excel.Options{
		File:  excelize.NewFile(),
		Tasks: rsp.Tasks,
		Host:  fp.Host,
	}

	excelFile.ListAll()            // Создания общего листа с задачами
	excelFile.ListUser(&rsp.Users) // Создания листов по конкретным исполнителям

	fileName := fmt.Sprintf("%s %s - %s.xlsx", fp.Group, fp.DateStart, fp.DateEnd)

	// Создает файл excel
	if err := excelFile.Save(fileName); err != nil {
		return "", err
	}

	return fileName, nil
}
