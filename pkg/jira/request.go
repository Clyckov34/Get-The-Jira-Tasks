package jira

import (
	"errors"
	"github.com/andygrunwald/go-jira"
)

//NewClient авторизация клиента к сервису jira
func (m *Setting) NewClient() (*jira.Client, error){
	client, err := jira.NewClient(m.Auth.Client(), *m.Host)
	if err != nil {
		return nil, errors.New("ошибка: в подключении")
	}
	return client, nil
}
