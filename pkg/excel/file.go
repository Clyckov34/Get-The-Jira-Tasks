package excel

import (
	"errors"

	"github.com/andygrunwald/go-jira"
)

//ListAll запись общего списока задач
func (m *Options) ListAll() {
	nameSheet := "Sheet1"

	m.setTableTitle(nameSheet)
	m.style(nameSheet)

	for key, value := range m.Tasks {
		m.setTableData(nameSheet, key, &value)
	}

	m.sheetRename(nameSheet, "Общий список задач")
}

//ListUser запись задач конкретных пользователей
func (m *Options) ListUser(userList *[]string) {
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

		for key, value := range data {
			m.setTableData(user, key, &value)
		}
	}
}

//Save сохранения файла
func (m *Options) Save(fileName string) error {
	err := m.File.SaveAs(fileName)
	if err != nil {
		return errors.New("ошибка: сохранения файла")
	}
	return nil
}
