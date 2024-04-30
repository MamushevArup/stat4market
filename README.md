# Installation and Usage

First clone repository on local
``` cd stat4market```
To run application use this command if you have docker installed.
Review .env to ensure that credentials true.
``` docker compose up --build```

### Requirements 1
Написать SQL-запросы для ClickHouse:

Выборки всех уникальных eventType у которых более 1000 событий.
Выборки событий которые произошли в первый день каждого месяца.
Выборки пользователей которые совершили более 3 различных eventType.

You can find the query to the clickhouse in the __event.sql__ file. 

### Requirements 2
Вывод событий по заданному eventType и временному диапазону.

Code impelemntation can be foun in __/internal/repository/clickhouse/select-event-type.go__ file as well.

### Endpoints
```plaintext
POST /api/event insert event to the clickhouse db
GET /swagger/* to see documentation server
```

### Features
 - Application builded using docker and docker compose
 - Migration applied at code level
 - Layered architecture approach applied
 - Code style check using linters
 - Swagger documentation implemented
