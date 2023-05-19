package jira

import "github.com/andygrunwald/go-jira"

//FlagParameters параметры флагов программы
type FlagParameters struct {
	Host      string
	UserName  string
	Token     string
	DateStart string
	DateEnd   string
	Group     string
	Status    string
}

type Response struct {
	Tasks []jira.Issue
	Users []string
}
