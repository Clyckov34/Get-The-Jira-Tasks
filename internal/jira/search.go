package jira

import (
	"errors"

	jr "github.com/andygrunwald/go-jira"
)

// SearchTask поиск задач
func (m *Setting) SearchTask(client *jr.Client) ([]jr.Issue, error) {
	var (
		last   int
		issues []jr.Issue
	)

	for {
		opt := &jr.SearchOptions{
			MaxResults: 1000,
			StartAt:    last,
		}

		chunk, resp, err := client.Issue.Search(*m.JQL, opt)
		if err != nil {
			return nil, errors.New("ошибка: поиска")
		}

		total := resp.Total
		if issues == nil {
			issues = make([]jr.Issue, 0, total)
		}

		issues = append(issues, chunk...)
		last = resp.StartAt + len(chunk)

		if last >= total {
			return issues, nil
		}
	}
}

// GroupUser список сотрудников
func (m *Setting) GroupUser(client *jr.Client, nameGroup *string) ([]string, error) {
	group := make([]string, 0)

	list, _, err := client.Group.Get(*nameGroup)
	if err != nil {
		return nil, errors.New("ошибка: запрос в группу")
	}

	for _, value := range list {
		group = append(group, value.DisplayName)
	}
	return group, nil
}
