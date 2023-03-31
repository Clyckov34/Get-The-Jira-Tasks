package main

import (
	"flag"
	"log"
	"rjt/pkg/check"
	"rjt/internal/app"
)

var flagNew = &app.FlagParameters{}

func init() {
	var paramDefault = &app.FlagParameters{
		Host:     "https://brandquad.atlassian.net",
		UserName: "",
		Token:    "",
		Group:    "For Accounting Dep",
		Status:   "Done",
	}

	//Параметры запуска приложения
	flag.StringVar(&flagNew.Host, "host", paramDefault.Host, "Хост Jira Cloud")
	flag.StringVar(&flagNew.UserName, "user", paramDefault.UserName, "Логин пользователя")
	flag.StringVar(&flagNew.Token, "token", paramDefault.Token, "Токен")
	flag.StringVar(&flagNew.DateStart, "date_start", check.DateNow(), "Начальная дата запроса. Формат: YYYY-MM-DD")
	flag.StringVar(&flagNew.DateEnd, "date_end", check.DateNow(), "Конечная дата запроса. Формат: YYYY-MM-DD")
	flag.StringVar(&flagNew.Group, "group", paramDefault.Group, "Группа пользователей в Jira")
	flag.StringVar(&flagNew.Status, "status", paramDefault.Status, "Статус задач")
}

func main() {
	flag.Parse()
	if err := app.Run(flagNew); err != nil {
		log.Fatalln(err)
	}
}
