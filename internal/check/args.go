package check

import (
	"errors"
	"os"
)

//Args проверяет запуск скрипта с аргументов
func Args(number int) error {
	if len(os.Args) < 5 {
		return errors.New("не указаны все ОБЯЗАТЕЛЬНЫЙ ПАРАМЕТР! (Flags)")
	}

	return nil
}
