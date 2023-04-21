package app

import (
	"fmt"
	"rjt/pkg/excel"
	"rjt/pkg/jira"
	"rjt/app/date"
	"time"

	jr "github.com/andygrunwald/go-jira"
	"github.com/xuri/excelize/v2"
)

//const filter = `status in (%v) AND created >= "%v 00:00" AND created <= "%v 23:59" AND assignee in (membersOf("%v")) ORDER BY created DESC` // Фильтр по созданным задачам
const filter = `status in (%v) AND updated >= "%v 00:00" AND updated <= "%v 23:59" AND assignee in (membersOf("%v")) ORDER BY updated DESC` // Фильтр по обновленным задачам

// Run запуск приложения
func Run(fp *FlagParameters) error {
	timeStart := time.Now() //Начало таймера

	//Проверка введенного формата даты
	dateStart, err := date.Check(&fp.DateStart)
	if err != nil {
		return err
	}

	dateEnd, err := date.Check(&fp.DateEnd)
	if err != nil {
		return err
	}

	//Фильтр JQL в Jira
	jql := fmt.Sprintf(filter, fp.Status, dateStart, dateEnd, fp.Group)

	var server = &jira.Setting{
		Host: &fp.Host,
		Auth: jr.BasicAuthTransport{
			Username: fp.UserName,
			Password: fp.Token,
		},
		JQL: &jql,
	}

	//Подключения к Серверу
	fmt.Println("Загрузка: 10% - Подключение:", fp.Host)
	client, err := server.NewClient()
	if err != nil {
		return err
	}

	//Запрос в JQL
	fmt.Println("Загрузка: 30% - Запрос в JQL")
	issueList, err := server.GetTask(client)
	if err != nil {
		return err
	}

	//Запрос на список пользователей
	fmt.Println("Загрузка: 50% - Запрос на список пользователей")
	userList, err := server.GetGroupUser(client, &fp.Group)
	if err != nil {
		return err
	}

	var excelFile = &excel.Options{
		File:   excelize.NewFile(),
		Issues: &issueList,
		Host:   &fp.Host,
	}

	//Выгрузки задач в Excel
	fmt.Println("Загрузка: 60% - Выгрузка задач в excel таблицу")
	excelFile.ListAll()
	excelFile.ListUser(&userList)

	//Сохраняем файл Excel
	fmt.Printf("Загрузка: 90%% - Файл сохранен: %v - %v.xlsx\n", dateStart, dateEnd)
	err = excelFile.Save(&dateStart, &dateEnd)
	if err != nil {
		return err
	}

	fmt.Printf("\nКол-во задач: %d\nВремя выполнения скрипта %.2fs\nГотово\n", len(issueList), time.Since(timeStart).Seconds())
	return nil
}
