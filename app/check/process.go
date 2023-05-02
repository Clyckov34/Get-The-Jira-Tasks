package check

import (
	"errors"
	"os"
)

//Args проверяет запуск скрипта с аргументов
func Args() error {
	if len(os.Args) < 3 {
		return errors.New("не указаны аргументы (Flags)")
	}

	return nil
}