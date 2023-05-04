package excel

import (
	"rjt/internal/services/jira"
	"rjt/pkg/excel"

	"github.com/xuri/excelize/v2"
)

//Write создает файл excel
func Write(rsp *jira.Response, fp *jira.FlagParameters) (excel.FileName, error) {
	var excelFile = &excel.Options{
		File:  excelize.NewFile(),
		Tasks: rsp.Tasks,
		Host:  fp.Host,
	}

	excelFile.ListAll()            // Создания общего листа с задачами
	excelFile.ListUser(&rsp.Users) // Создания листов по конкретным исполнителям

	// Создает файл excel
	fileName, err := excelFile.Save(&fp.DateStart, &fp.DateEnd)
	if err != nil {
		return "", err
	}

	return fileName, nil
}
