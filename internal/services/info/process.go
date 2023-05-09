package info

import "fmt"

func AppFlag() {
	fmt.Printf("Параметры запуска:\n" +
		"\t--host - Хост Jira Cloud. *\n" +
		"\t--user - Login. *\n" +
		"\t--token - Token. *\n" +
		"\t--group - Название группы в Jira. Список пользователей. *\n" +
		"\t--date_start - Начальная дата запроса. Формат: YYYY-MM-DD. Default: 'Сегоднешний день'.\n" +
		"\t--date_end - Конечная дата запроса Формат: YYYY-MM-DD. Default: 'Сегоднешний день'.\n" +
		"\t--status - Статус задачи. Default: To Do, In Progress, Done, Denied, Pause.\n" +
		"* - Обязательные параметры при запуске скрипта\n"+
		"Документация: https://github.com/Clyckov34/Get-The-Jira-Tasks/blob/main/README.md\n\n")
}
