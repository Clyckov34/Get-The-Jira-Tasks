package main

import (
	"flag"
	"log"
	"rjt/internal/app"
	"rjt/internal/date"
	"rjt/internal/jira"
)

var flagNew = &jira.FlagParameters{}

func init() {
	//Параметры запуска приложения
	flag.StringVar(&flagNew.Host, "host", "", "Хост Jira Cloud")
	flag.StringVar(&flagNew.UserName, "user", "", "Логин пользователя")
	flag.StringVar(&flagNew.Token, "token", "", "Токен")
	flag.StringVar(&flagNew.DateStart, "date_start", date.Now(), "Начальная дата запроса. Формат: YYYY-MM-DD")
	flag.StringVar(&flagNew.DateEnd, "date_end", date.Now(), "Конечная дата запроса. Формат: YYYY-MM-DD")
	flag.StringVar(&flagNew.Group, "group", "", "Группа пользователей в Jira")
	flag.StringVar(&flagNew.Status, "status", `"To Do", "In Progress", "Done", "Denied", "Pause"`, "Статус задач")
}

func main() {
	flag.Parse()

	//Запуск приложения
	if err := app.Run(flagNew); err != nil {
		log.Fatalln(err)
	}
}
