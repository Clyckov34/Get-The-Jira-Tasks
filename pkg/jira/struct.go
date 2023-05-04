package jira

import (
	"github.com/andygrunwald/go-jira"
)

type JQL struct {
	Status    string
	DateStart string
	DateEnd   string
	Group     string
}

type Setting struct {
	Host *string
	Auth jira.BasicAuthTransport
}

type Response struct {
	Jira *jira.Client
}

