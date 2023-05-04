package jira

import "fmt"

// Filter структура (Запрос) фильтра
func Filter(m *JQL) string {
 // const filter = `status in (%v) AND created >= "%v 00:00" AND created <= "%v 23:59" AND assignee in (membersOf("%v")) ORDER BY created DESC` // Фильтр по созданным задачам
	const filter = `status in (%v) AND updated >= "%v 00:00" AND updated <= "%v 23:59" AND assignee in (membersOf("%v")) ORDER BY updated DESC` // Фильтр по обновленным задачам

	return fmt.Sprintf(filter, m.Status, m.DateStart, m.DateEnd, m.Group)
}
