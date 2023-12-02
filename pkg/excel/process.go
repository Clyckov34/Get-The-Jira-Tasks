package excel

import (
	"fmt"

	jr "rjt/pkg/jira"

	"github.com/andygrunwald/go-jira"
)

// setTableData загаловки столбцов
func (m *Options) setTableTitle(nameSheet string) {
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
	_ = m.File.SetCellValue(nameSheet, "K1", "Задача Изменена")
}

// setTableData запись данные в ячейки
func (m *Options) setTableData(sheet string, key int, value *jira.Issue) {
	dateCreateTask, _ := value.Fields.Created.MarshalJSON()
	dateResolutionTask, _ := value.Fields.Resolutiondate.MarshalJSON()
	dateUpdateTask, _ := value.Fields.Updated.MarshalJSON()

	_ = m.File.SetCellValue(sheet, fmt.Sprintf("A%d", key+2), jr.KeyTask(value.Key))
	_ = m.File.SetCellValue(sheet, fmt.Sprintf("B%d", key+2), value.Key)
	_ = m.File.SetCellFormula(sheet, fmt.Sprintf("C%d", key+2), `HYPERLINK("`+m.Host+`/browse/`+value.Key+`")`)
	_ = m.File.SetCellValue(sheet, fmt.Sprintf("D%d", key+2), value.Fields.Summary)
	_ = m.File.SetCellValue(sheet, fmt.Sprintf("E%d", key+2), value.Fields.Type.Name)
	_ = m.File.SetCellValue(sheet, fmt.Sprintf("F%d", key+2), value.Fields.Status.Name)
	_ = m.File.SetCellValue(sheet, fmt.Sprintf("G%d", key+2), value.Fields.Creator.DisplayName)
	_ = m.File.SetCellValue(sheet, fmt.Sprintf("H%d", key+2), value.Fields.Assignee.DisplayName)
	_ = m.File.SetCellValue(sheet, fmt.Sprintf("I%d", key+2), parse(dateCreateTask))
	_ = m.File.SetCellValue(sheet, fmt.Sprintf("J%d", key+2), parse(dateResolutionTask))
	_ = m.File.SetCellValue(sheet, fmt.Sprintf("K%d", key+2), parse(dateUpdateTask))
}

// setTableDataRepor запись данные в ячейки отчет
func (m *Options) setTableDataReport(sheet string, key int, assignee string, value []string) {
	_ = m.File.SetCellValue(sheet, fmt.Sprintf("A%d", key), assignee)
	_ = m.File.SetCellValue(sheet, fmt.Sprintf("B%d", key), value)

	var	pim, eci string

	if len(value) == 0 {
		pim = "-"
	} else {
		for _, v := range value {
			if v == "BFBV2" || v == "SPR" {
				eci = "ECI"
			} else if v == "BIT" || v == "ETP" || v == "IP" || v == "PIM" {
				pim = "PIM"
			} else {
				pim = "-"
			}
		}
	}

	res := pim + " " + eci
	_ = m.File.SetCellValue(sheet, fmt.Sprintf("C%d", key), res)
}

// sheetRename переменовывает лист в таблице
func (m *Options) sheetRename(oldName, newName string) {
	m.File.SetSheetName(oldName, newName)
}

// style стиль таблицы
func (m *Options) style(nameSheet string) {
	_ = m.File.SetColWidth(nameSheet, "A", "B", 20)
	_ = m.File.SetColWidth(nameSheet, "C", "D", 60)
	_ = m.File.SetColWidth(nameSheet, "E", "F", 15)
	_ = m.File.SetColWidth(nameSheet, "G", "K", 25)
}
