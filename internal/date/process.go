package date

import "time"

const Format = "2006-01-02"

//Now Сегодняшняя дата
func Now() string {
	return time.Now().Format(Format)
}