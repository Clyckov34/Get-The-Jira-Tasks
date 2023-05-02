package main

import (
	"flag"
	"log"
	"rjt/app"
	"rjt/app/check"
	"rjt/app/date"
	"rjt/app/info"
)

var flagNew = &app.FlagParameters{}

func init() {
	var paramDefault = &app.FlagParameters{
		Host:     "https://brandquad.atlassian.net",
		Group:    "For Accounting Dep",
		Status:   "Done",
	}

	//Параметры запуска приложения
	flag.StringVar(&flagNew.Host, "host", paramDefault.Host, "Хост Jira Cloud")
	flag.StringVar(&flagNew.UserName, "user", paramDefault.UserName, "Логин пользователя")
	flag.StringVar(&flagNew.Token, "token", paramDefault.Token, "Токен")
	flag.StringVar(&flagNew.DateStart, "date_start", date.Now(), "Начальная дата запроса. Формат: YYYY-MM-DD")
	flag.StringVar(&flagNew.DateEnd, "date_end", date.Now(), "Конечная дата запроса. Формат: YYYY-MM-DD")
	flag.StringVar(&flagNew.Group, "group", paramDefault.Group, "Группа пользователей в Jira")
	flag.StringVar(&flagNew.Status, "status", paramDefault.Status, "Статус задач")
}

func main() {
	if err := check.Args(); err != nil {
		info.AppFlag()
		log.Fatalln(err)
	}

	flag.Parse()

	if err := app.Run(flagNew); err != nil {
		log.Fatalln(err)
	}
}