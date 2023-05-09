package jira

import "fmt"

// UpdateString Фильтр по обновленным задачам
func (m *JQL) UpdateString() string {
	const filter = `status in (%v) AND updated >= "%v 00:00" AND updated <= "%v 23:59" AND assignee in (membersOf("%v")) ORDER BY updated DESC`
	return fmt.Sprintf(filter, m.Status, m.DateStart, m.DateEnd, m.Group)
}

// CreateString Фильтр по созданным задачам
func (m *JQL) CreateString() string {
	const filter = `status in (%v) AND created >= "%v 00:00" AND created <= "%v 23:59" AND assignee in (membersOf("%v")) ORDER BY created DESC` // Фильтр по созданным задачам
	return fmt.Sprintf(filter, m.Status, m.DateStart, m.DateEnd, m.Group)
}
