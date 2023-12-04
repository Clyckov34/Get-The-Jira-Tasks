package excel

import (
	"errors"
	jr "rjt/pkg/jira"

	"github.com/andygrunwald/go-jira"
)

// ListAll запись общего списока задач
func (m *Options) ListAll() {
	nameSheet := "Sheet1"

	m.setTableTitle(nameSheet)
	m.style(nameSheet)

	for key, value := range m.Tasks {
		m.setTableData(nameSheet, key, &value)
	}

	m.sheetRename(nameSheet, "Общий список задач")
}

// ListUser запись задач конкретных пользователей
func (m *Options) ListUser(userList *[]string) {
	var key2 int

	nameSheetRepair := "Repair"
	m.File.NewSheet(nameSheetRepair)
	m.style(nameSheetRepair)

	m.setTableTitleRepair(nameSheetRepair)

	for _, user := range *userList {
		data := make([]jira.Issue, 0)
		for _, value := range m.Tasks {
			if user == value.Fields.Assignee.DisplayName {
				data = append(data, value)
			} else {
				continue
			}
		}

		m.File.NewSheet(user)
		m.setTableTitle(user)
		m.style(user)

		keyTask := make([]string, 0)
		seen := map[string]bool{}

		for key, value := range data {
			m.setTableData(user, key, &value)

			kt := jr.KeyTask(value.Key)
			if !seen[kt] {
				keyTask = append(keyTask, kt)
			}
			seen[kt] = true
		}

		key2++
		m.setTableDataReport(nameSheetRepair, key2, user, keyTask)
		keyTask = nil
	}
}

// Save сохранения файла
func (m *Options) Save(fileName string) error {
	err := m.File.SaveAs(fileName)
	if err != nil {
		return errors.New("ошибка: сохранения файла")
	}
	return nil
}
