package jira

import (
	"rjt/pkg/jira"

	jr "github.com/andygrunwald/go-jira"
)

//Request запрос в API Jira Cloud
func Request(fp *FlagParameters) (*Response, error) {
	var server = &jira.Setting{
		Host: &fp.Host,
		Auth: jr.BasicAuthTransport{
			Username: fp.UserName,
			Password: fp.Token,
		},
	}

	//Создания нового соединения
	client, err := jira.NewClient(server)
	if err != nil {
		return nil, err
	}

	//Параметры фильтра
	filter := &jira.JQL{
		Status:    fp.Status,
		DateStart: fp.DateStart,
		DateEnd:   fp.DateEnd,
		Group:     fp.Group,
	}

	//Поиск задач по фильту
	tasks, err := client.GetTask(filter.Update())
	if err != nil {
		return nil, err
	}

	//Выгрузка список пользователей
	users, err := client.GetGroupUser(&fp.Group)
	if err != nil {
		return nil, err
	}

	return &Response{Tasks: tasks, Users: users}, nil
}
