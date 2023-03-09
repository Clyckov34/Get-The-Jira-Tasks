package jira

import jr "github.com/andygrunwald/go-jira"

type Setting struct {
	Host *string
	Auth jr.BasicAuthTransport
	JQL  *string
}
