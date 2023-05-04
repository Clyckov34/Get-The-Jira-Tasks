package excel

import "time"

//parse обработка даты в формат
func parse(date []byte) interface{} {
	h, _ := time.Parse("\"2006-01-02T15:04:05.000-0700\"", string(date))
	timeString := h.Format("2006-01-02")

	if timeString != "0001-01-01" {
		return timeString
	}

	return nil
}
