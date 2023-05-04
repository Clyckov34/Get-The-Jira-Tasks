package excel

import (
	"github.com/andygrunwald/go-jira"
	"github.com/xuri/excelize/v2"
)

type Options struct {
	File  *excelize.File
	Tasks []jira.Issue
	Host  string
}

type FileName string