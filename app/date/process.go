package date

import (
	"time"
)

const format = "2006-01-02"

//FormatParseDate формат даты. Проверка на правильный ввод данных
func FormatParseDate(date *string) (string, error) {
	t, err := time.Parse(format, *date)
	if err != nil {
		return "", err
	}

	return t.Format(format), nil
}

//DateNow Сегодняшняя дата
func DateNow() string {
	return time.Now().Format(format)
}

//ParseDate обработка даты в формат
func ParseDate(date []byte) interface{} {
	h, _ := time.Parse("\"2006-01-02T15:04:05.000-0700\"", string(date))
	timeString := h.Format(format)

	if timeString != "0001-01-01" {
		return timeString
	}

	return nil
}