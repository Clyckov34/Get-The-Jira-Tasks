package app

import (
	"fmt"
	"log"
	"rjt/internal/services/check"
	"rjt/internal/services/excel"
	"rjt/internal/services/info"
	"rjt/internal/services/jira"
	"time"
)

// Run запуск приложения
func Run(fp *jira.FlagParameters) error {
	timerStart := time.Now()

	fmt.Println("Этап 1/4. Проверка параметров")

	// Проверка параметров на кол-во
	if err := check.Args(5); err != nil {
		info.AppFlag()
		log.Fatalln(err)
	}

	// Проверка формат даты
	if err := check.FormatDate(&fp.DateStart); err != nil {
		info.AppFlag()
		log.Fatalln(err)
	}

	if err := check.FormatDate(&fp.DateEnd); err != nil {
		info.AppFlag()
		log.Fatalln(err)
	}

	fmt.Println("Этап 2/4. Запрос API " + fp.Host)

	//Запрос в API Jira Cloud
	response, err := jira.Request(fp)
	if err != nil {
		return err
	}

	fmt.Println("Этап 3/4. Выгрузка задач в Excel файл")

	//Создания Excel файла
	fileName, err := excel.Write(response, fp)
	if err != nil {
		return err
	}

	fmt.Printf("Этап 4/4. Готово\n\nКол-во задач: %d\nКол-во сотрудников: %d\nФайл: %v\nВремя выполнения скрипта %.2fs\n", len(response.Tasks), len(response.Users), fileName, time.Since(timerStart).Seconds())
	return nil
}
