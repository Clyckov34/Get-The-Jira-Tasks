package check

import (
	"time"
)

const timeFormat = "2006-01-02"

//DateFormatParse формат даты. Проверка на правильный ввод данных
func DateFormatParse(date *string) (string, error) {
	t, err := time.Parse(timeFormat, *date)
	if err != nil {
		return "", err
	}

	return t.Format(timeFormat), nil
}

//DateNow Сегодняшняя дата
func DateNow() string {
	return time.Now().Format(timeFormat)
}
