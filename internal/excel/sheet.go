package excel

import (
	"fmt"
	"github.com/andygrunwald/go-jira"
	"strings"
	"time"
)

//setTitle загаловки столбцов
func (m *Options) setTitle(nameSheet string) {
	_ = m.File.SetCellValue(nameSheet, "A1", "Ключ")
	_ = m.File.SetCellValue(nameSheet, "B1", "№ Задача")
	_ = m.File.SetCellValue(nameSheet, "C1", "URL задачи")
	_ = m.File.SetCellValue(nameSheet, "D1", "Название задачи")
	_ = m.File.SetCellValue(nameSheet, "E1", "Тип задачи")
	_ = m.File.SetCellValue(nameSheet, "F1", "Статус задачи")
	_ = m.File.SetCellValue(nameSheet, "G1", "Автор")
	_ = m.File.SetCellValue(nameSheet, "H1", "Исполнитель")
	_ = m.File.SetCellValue(nameSheet, "I1", "Задача Создана")
	_ = m.File.SetCellValue(nameSheet, "J1", "Задача Закрыта")
}

//setData запись данные в ячейки
func (m *Options) setData(sheet string, key int, value *jira.Issue) {
	createTask, _ := value.Fields.Created.MarshalJSON()
	resolutionTask, _ := value.Fields.Resolutiondate.MarshalJSON()

	_ = m.File.SetCellValue(sheet, fmt.Sprintf("A%d", key+2), keyTask(value.Key))
	_ = m.File.SetCellValue(sheet, fmt.Sprintf("B%d", key+2), value.Key)
	_ = m.File.SetCellFormula(sheet, fmt.Sprintf("C%d", key+2), `HYPERLINK("`+*m.Host+`/browse/`+value.Key+`")`)
	_ = m.File.SetCellValue(sheet, fmt.Sprintf("D%d", key+2), value.Fields.Summary)
	_ = m.File.SetCellValue(sheet, fmt.Sprintf("E%d", key+2), value.Fields.Type.Name)
	_ = m.File.SetCellValue(sheet, fmt.Sprintf("F%d", key+2), value.Fields.Status.Name)
	_ = m.File.SetCellValue(sheet, fmt.Sprintf("G%d", key+2), value.Fields.Creator.DisplayName)
	_ = m.File.SetCellValue(sheet, fmt.Sprintf("H%d", key+2), value.Fields.Assignee.DisplayName)
	_ = m.File.SetCellValue(sheet, fmt.Sprintf("I%d", key+2), parseDate(createTask))
	_ = m.File.SetCellValue(sheet, fmt.Sprintf("J%d", key+2), parseDate(resolutionTask))
}

//keyTask преобразовывает номер ключа ABCD-1234 -> ABCD
func keyTask(nameKey string) string {
	key := strings.Split(nameKey, "-")
	return key[0]
}

//parseDate обработка даты в формат
func parseDate(date []byte) interface{} {
	h, _ := time.Parse("\"2006-01-02T15:04:05.000-0700\"", string(date))
	timeString := h.Format("2006-01-02")

	if timeString != "0001-01-01" {
		return timeString
	}

	return nil
}

//sheetRename переменовывает лист в таблице
func (m *Options) sheetRename(oldName, newName string) {
	m.File.SetSheetName(oldName, newName)
}

//style стиль таблицы
func (m *Options) style(nameSheet string) {
	_ = m.File.SetColWidth(nameSheet, "A", "B", 20)
	_ = m.File.SetColWidth(nameSheet, "C", "D", 60)
	_ = m.File.SetColWidth(nameSheet, "E", "F", 15)
	_ = m.File.SetColWidth(nameSheet, "G", "K", 25)
}
