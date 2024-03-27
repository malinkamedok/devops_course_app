# Микросервис парсинга курса валют ЦБ РФ

Приложение, отдающее курс валюты по ЦБ РФ за определенную дату. Для получения курсов валют используется официальное API ЦБ РФ.

#### Получение курса валюты за определенную дату

<details>
 <summary><code>GET</code> <code><b>/</b></code> <code>info</code> <code><b>/</b></code> <code>currency</code></summary>

##### Parameters

> | name     | type     | data type | example    | description                 |
> |----------|----------|-----------|------------|-----------------------------|
> | currency | required | string    | `USD`        | Валюта в стандарте ISO 4217 |
> | date     | optional | string    | `2016-01-06` | Дата в формате YYYY-MM-DD   |

##### Example output

```json 
{
    "data": {
      "USD": 33.4013
    },
    "service": "currency"
}
```

</details>

#### Получение информации о погоде в определенном городе

<details>
 <summary><code>GET</code> <code><b>/</b></code> <code>info</code> <code><b>/</b></code> <code>weather</code></summary>

##### Parameters

> | name | type     | data type | example         | description                         |
> |------|----------|-----------|-----------------|-------------------------------------|
> | city | required | string    | `SaintPetersburg` | Страна, город, адрес или координаты |
> | from | optional | string    | `2024-03-20`      | Дата в формате YYYY-MM-DD           |
> | to   | optional | string    | `2024-03-26`      | Дата в формате YYYY-MM-DD           |

##### Example output

```json 
{
  "data": {
    "city": "SaintPetersburg",
    "from": "2024-03-25",
    "to": "2024-03-26",
    "temperature_c": {
      "average": 2.35,
      "median": 2.35,
      "min": -1,
      "max": 8
    },
    "humidity": {
      "average": 86.5,
      "median": 86.5,
      "min": 85.6,
      "max": 87.4
    },
    "pressure_mb": {
      "average": 1004.8,
      "median": 1004.8,
      "min": 1000.7,
      "max": 1008.9
    }
  },
  "service": "weather"
}
```

</details>


## Структура проекта

```bash
.
├── .github         
│   └── workflows             # CI
├── cmd
│   └── main                  # Точка входа в приложение
├── docs                      # Проектная документация OpenApi
├── internal
│   ├── app                   # Настройки приложения
│   ├── config                # Парсинг переменных окружения (стандартный порт)
│   ├── controller
│   │   └── http
│   │       └── v1            # Endpoints 
│   ├── entity                # Сущности
│   └── usecase               # Бизнес-логика приложения
│       ├── cbrf              # Обработка данных с ЦБ РФ
│       └── visualcrossing    # Обработка данных с VS
└── pkg
    ├── httpserver            # Конфигурации для работы с HTTP сервером
    └── web                   # Конфигурации для обработки JSON-ответов
```

## Документация и запуск

Для запуска выполнить сборку приложения

### docker compose

```bash
docker-compose up --build
```
### linux machine

```bash
go mod tidy
go build -o app cmd/main/main.go 
./app
```

### Документация
[OpenApi](https://malinkamedok.github.io/devops_course_app/)

<details>
 <summary>Переменные окружения</summary>

##### Parameters

> | name    | type     | example                     | description                             |
> |---------|----------|-----------------------------|-----------------------------------------|
> | PORT    | optional | `8000`                      | Порт приложения. default = 8000         |
> | API_KEY | required | `AAAAAAAAAAAAAAA123BBBBBBB` | API ключ для сервиса visualcrossing     |

</details>



