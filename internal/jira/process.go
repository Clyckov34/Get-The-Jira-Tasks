package jira

import jr "github.com/andygrunwald/go-jira"

type Setting struct {
	Host *string
	Auth jr.BasicAuthTransport
	JQL  *string
}

const Filter = `status in (%v) AND "[CHART] Date of First Response" >= "%v 00:00" 
AND "[CHART] Date of First Response" <= "%v 23:59" AND assignee in (membersOf("%v")) 
ORDER BY created DESC`
