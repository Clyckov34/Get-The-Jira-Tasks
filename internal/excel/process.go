package excel

import (
	jr "github.com/andygrunwald/go-jira"
	"github.com/xuri/excelize/v2"
)

type Options struct {
	File   *excelize.File
	Issues *[]jr.Issue
	Host   *string
}
