package info

import "fmt"

func AppFlag() {
	fmt.Printf("Параметры запуска:\n" +
		"\t--host - Хост Jira Cloud.\n" +
		"\t--user - Login. ОБЯЗАТЕЛЬНЫЙ ПАРАМЕТР!\n" +
		"\t--token - Token. ОБЯЗАТЕЛЬНЫЙ ПАРАМЕТР!\n" +
		"\t--group - Название группы в Jira. Список пользователей.\n" +
		"\t--date_start - Начальная дата запроса. Формат: YYYY-MM-DD. Default: 'Сегоднешний день 00:00'.\n" +
		"\t--date_end - Конечная дата запроса Формат: YYYY-MM-DD. Default: 'Сегоднешний день 23:59'.\n" +
		"\t--status - Статус задачи. Default: 'Done'.\n" +
		"Документация: https://github.com/Clyckov34/JIra-Export-Task/blob/main/README.md\n\n")
}
