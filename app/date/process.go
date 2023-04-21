package date

import (
	"time"
)

const format = "2006-01-02"

//Check формат даты. Проверка на правильный ввод данных
func Check(date *string) (string, error) {
	t, err := time.Parse(format, *date)
	if err != nil {
		return "", err
	}

	return t.Format(format), nil
}

//Now Сегодняшняя дата
func Now() string {
	return time.Now().Format(format)
}

//Parse обработка даты в формат
func Parse(date []byte) interface{} {
	h, _ := time.Parse("\"2006-01-02T15:04:05.000-0700\"", string(date))
	timeString := h.Format(format)

	if timeString != "0001-01-01" {
		return timeString
	}

	return nil
}