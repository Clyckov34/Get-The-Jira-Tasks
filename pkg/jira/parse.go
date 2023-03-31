package jira

import "strings"

// keyTask преобразовывает номер ключа ABCD-1234 -> ABCD
func KeyTask(nameKey string) string {
	key := strings.Split(nameKey, "-")
	return key[0]
}
