package check

import (
	"testing"
	"time"
)

var dateNow = time.Now().Format("2006-01-02")

func TestDateFormatParse(t *testing.T) {
	_, err := DateFormatParse(&dateNow)
	if err != nil {
		t.Error("ошибка: не правильный формат даты")
	}
}

//TestDateNow проверка нынешней даты на формат
func TestDateNow(t *testing.T) {
	res := DateNow()
	if res != dateNow {
		t.Error("ошибка: не правильный формат даты")
	}
}
