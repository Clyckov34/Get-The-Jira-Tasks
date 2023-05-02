package info

import "fmt"

func AppFlag() {
	fmt.Printf("Параметры запуска:\n"+ 
	"\t--date_start - Начальная дата запроса. Формат: YYYY-MM-DD. Default: 'Сегоднешний день 00:00'.\n"+
	"\t--date_end - Конечная дата запроса Формат: YYYY-MM-DD. Default: 'Сегоднешний день 23:59'.\n"+
	"\t--group - Название группы в Jira. Список пользователей. Default: 'My Group'.\n"+
	"\t--status - Статус задачи. Default: 'Done'.\n"+
	"\t--host - Хост Jira Cloud. Default: 'https://myserver.atlassian.net'.\n"+
	"\t--token - Token. ОБЯЗАТЕЛЬНЫЙ ПАРАМЕТР!\n"+
	"\t--user - Login. ОБЯЗАТЕЛЬНЫЙ ПАРАМЕТР!\n"+
	"Документация: https://github.com/Clyckov34/JIra-Export-Task/blob/main/README.md\n\n")
}