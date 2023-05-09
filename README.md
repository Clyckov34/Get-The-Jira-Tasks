<div>
    <center>
        <h1>Get The Jira Tasks</h1>
        <p>Выгрузка задач по сотрудникам из системы Jira Cloud формате Excel.xlsx</p>
    </center>
</div>
<div>
    <h2>Excel файл</h2>
    <p>Excel файл состоит:</p>
    <ul>
        <li>Общий лист задачами</li>
        <li>Отдельные листы с задачами по каждому сотруднику</li>
    </ul>
    <p>Таблица состоит:</p>
    <ul>
        <li>Ключ</li>
        <li>№ задачи</li>
        <li>URL задачи</li>
        <li>Название задачи</li>
        <li>Тип задачи</li>
        <li>Статус задачи</li>
        <li>Автор</li>
        <li>Исполнитель</li>
        <li>Задача Создана</li>
        <li>Задача Закрыта</li>
        <li>Задача Изменена</li>
    </ul>
</div>
<div>
    <h2>Настройка Jira Cloud в личном кабинете</h2>
    <ol>
        <li>Создать группу сотрудников по которым делать выгрузку задач</li>
        <ol>
            <li>Откройте панель администрации <a href="https://admin.atlassian.com/">https://admin.atlassian.com/</a></li>
            <li>Перейдите в раздел: <b>Groups</b> -> <b>Create group</b></li>
            <li>Создать группу например: <b>My Group</b></li>
            <li>Добавить сотрудников в группу</li>
        </ol>
        <li>Сгенерировать Token <a href="https://id.atlassian.com/manage/api-tokens">https://id.atlassian.com/manage/api-tokens</a></li>
    </ol>
</div>
<div>
<h2>Запуск утилиты с параметрами</h2>
<p>Выгрузка задач за сегодня</p>

```
go run cmd/app/main.go --host="https://myServer.atlassian.net" --user="Login" --token="Token" --group="My Group"
```
<p>Выгрузка задач за конкретный период</p>

```
go run cmd/app/main.go --host="https://myServer.atlassian.net" --user="Login" --token="Token" --group="My Group" --date_start="2022-10-01" --"date_end=2022-10-31"`
```
<p>Выгрузка задач со статусом "Done" за сегодня</p>

```
go run cmd/app/main.go --host="https://myServer.atlassian.net" --user="Login" --token="Token" --group="My Group" --date_start="2022-10-01" --"date_end=2022-10-31" --status="Done"
```
<p>Выгрузка задач с нескольками статусами: "Done" и "To Do" за сегодня:</p>

```
go run cmd/app/main.go --host="https://myServer.atlassian.net" --user="Login" --token="Token" --group="My Group"` --status='"Done", "To Do"'
```

<p>Важный момент: По умолчанию задачи выгружаются со статусами: </p>
<ul>
    <li>To Do</li>
    <li>In Progress</li>
    <li>Done</li>
    <li>Denied</li>
    <li>Pause</li>
<ul>
</div>
<div>
    <h2>Описание параметров</h2>
    <ul>
        <li><code>--host</code> - Хост Jira Cloud <b>*</b></li>
        <li><code>--user</code> - Логин пользователя <b>*</b></li>
        <li><code>--token</code> - Токен пользователя <b>*</b></li>
        <li><code>--group</code> - Название группы в Jira Cloud которую создавали. Список пользователей <b>*</b></li>
        <li><code>--date_start</code> - Начальная дата запроса. Формат: YYYY-MM-DD <code>Default</code>: <code>Сегоднешний день</code></li>
        <li><code>--date_end</code> - Конечная дата запроса Формат: YYYY-MM-DD. <code>Default</code>: <code>Сегоднешний день</code></li>
        <li><code>--status</code> - Статус задачи. <code>Default</code>: <code>To Do, In Progress, Done, Denied, Pause</code></li>
    </ul>
    <p><b>*</b> - <i>Обязательные параметры при запуске скрипта</i></p>
</div>
