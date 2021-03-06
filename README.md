API цб ключевой ставки

GET | v1/keyrate | возвращает ключевую ставку на сегодня
Пример ответа:
```
{
    "current_page": 1,
    "data": [
    {
        "date": "2022-02-16T00:00:00+03:00",
        "rate": "8.50"
    },
    ],
    "from_date": "2022-02-16",
    "per_page": 3,
    "total": 0,
    "total_pages": 0
}
 ]
```

Если за текущий день нет ставки, ответ будет:
```
{
    "current_page": 1,
    "data": [],
    "from_date": "2022-02-16",
    "per_page": 3,
    "total": 0,
    "total_pages": 0
}
```

PARAMS | ?from_date=2021-12-01 | возвращает ключевую ставку от даты
Пример ответа:
```
{
    "current_page": 2,
    "data": [
        {
            "date": "2022-02-07T00:00:00+03:00",
            "rate": "8.50"
        },
        {
            "date": "2022-02-04T00:00:00+03:00",
            "rate": "8.50"
        },
        {
            "date": "2022-02-03T00:00:00+03:00",
            "rate": "8.50"
        }
    ],
    "from_date": "2022-02-16",
    "per_page": 3,
    "total": 31,
    "total_pages": 10
}
```
PARAMS | page=2 | Возвращает страницу 2

PARAMS | per_page=5 | Количество возвращаемых строк на одну страницу


Swagger http://localhost:8080/swagger/index.html

Документация * https://www.cbr.ru/DailyInfoWebServ/DailyInfo.asmx?op=KeyRateXML
