package check

import (
	"errors"
	dt "rjt/internal/date"
	"time"
)

//FormatDate формат даты. Проверка на правильный ввод данных
func FormatDate(date *string) error {
	_, err := time.Parse(dt.Format, *date)
	if err != nil {
		return errors.New("ошибка: неверный формат даты. " + err.Error())
	}

	return nil
}
