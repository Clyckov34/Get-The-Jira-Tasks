package jira

import (
	"errors"

	"github.com/andygrunwald/go-jira"
)

// GetTask получить список задач
func (m *Response) GetTask(jql string) ([]jira.Issue, error) {
	var (
		last   int
		issues []jira.Issue
	)

	for {
		opt := &jira.SearchOptions{
			MaxResults: 1000,
			StartAt:    last,
		}

		chunk, resp, err := m.Jira.Issue.Search(jql, opt)
		if err != nil {
			return nil, errors.New("ошибка: поиска")
		}

		total := resp.Total
		if issues == nil {
			issues = make([]jira.Issue, 0, total)
		}

		issues = append(issues, chunk...)
		last = resp.StartAt + len(chunk)

		if last >= total {
			return issues, nil
		}
	}
}

// GetGroupUser получить список сотрудников
func (m *Response) GetGroupUser(nameGroup *string) ([]string, error) {
	group := make([]string, 0)

	list, _, err := m.Jira.Group.Get(*nameGroup)
	if err != nil {
		return nil, errors.New("ошибка: запрос в группу")
	}

	for _, value := range list {
		group = append(group, value.DisplayName)
	}
	return group, nil
}
