# API key rate of the central bank of the Russian Federation
# API ключевая ставка Российской Федерации

The service allows you to get key rates for a certain period in json format. The service takes information from the Central Bank of the Russian Federation. The service parses information in xml format and converts it to json format with pagination.

Сервис позволяет получить ключевые ставки за некоторый период в формате json. Сервис берёт информацию у центрального банка Российской Федерации. Сервис парсит информацию в формате xml и преобразует её в json формат с возможностью постраничного вывода.

## Install
1. Download packages:
```
make build
```

2. Change config:
* docker/app/.env

3. Run Swagger
```
make install-swag
```

```
make swag
```

4. Run:
```
make run
```

## API

GET | v1/keyrate | returns the key rate for today | возвращает ключевую ставку на сегодня 

Response:
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

If there is no set key rate for the current day: | Если за текущий день нет установленной ключевой ставки:

Response:
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

PARAMS | ?from_date=2021-12-01 | returns key rate from date | возвращает ключевую ставку от даты

Response:
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
PARAMS | page=2 | Returns page 2 | Возвращает страницу 2

PARAMS | per_page=5 | Number of returned rows per page | Количество возвращаемых строк на одну страницу


Swagger http://localhost:8080/swagger/index.html

Documentation of the Central Bank of the Russian Federation | Документация ЦБ РФ API * https://www.cbr.ru/DailyInfoWebServ/DailyInfo.asmx?op=KeyRateXML
