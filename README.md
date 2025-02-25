# go-tiny-url

[![Linting](https://github.com/xgmsx/go-tiny-url/actions/workflows/golangci-lint.yml/badge.svg?branch=main)](https://github.com/xgmsx/go-tiny-url/actions/workflows/golangci-lint.yml)
[![Tests](https://github.com/xgmsx/go-tiny-url/actions/workflows/coverage.yml/badge.svg?branch=main)](https://github.com/xgmsx/go-tiny-url/actions/workflows/coverage.yml)
[![CodeQL](https://github.com/xgmsx/go-tiny-url/actions/workflows/codeql.yml/badge.svg?branch=main)](https://github.com/xgmsx/go-tiny-url/actions/workflows/codeql.yml)
[![Coverage_Report](https://img.shields.io/badge/Coverage_Report-57.9%25-yellow)](https://xgmsx.github.io/go-tiny-url)

Проект для демонстрации сервиса на Golang с использованием технологий gRPC, REST HTTP, Kafka, Postgres, Observability, Unit-testing.

Структура:
* [internal](internal) содержит пакеты относящиеся к проекту - app, config, shortener.
* [pkg](pkg) содержит универсальные пакеты, которые можно переиспользовать - logger, http-сервер, клиенты к postgres, redis, kafka и т.д. 

Table of Contents:
- [Features](#features)
- [Installation](#installation)
- [Quick start](#quick-start)
- [Usage](#usage)
- [Metrics](#metrics)
- [Traces](#traces)
- [Environment variables](#environment-variables)

## Features

- **Поддержка протоколов**: HTTP и gRPC интерфейсы для взаимодействия с сервисом.
- **Интеграция с Kafka**: Отправка и получение сообщений в kafka для взаимодействия с сервисом.
- **Хранение данных**: Данные о созданных ссылках хранятся в Postgres SQL.
- **Кэширование**: Данные о созданных и запрашиваемых ссылках кешируются в Redis для снижения нагрузки на БД.  
- **Трассировки (observability)**: Данные трейсов входящих запросов отправляются в Jaeger для анализа производительности распределенных систем.
- **Метрики (observability)**: Данные метрик входящих запросов отправляются в Prometheus для анализа производительности сервиса в Grafana.
- **Логи (observability)**: Информация об ошибках передается в Sentry. Работа сервиса логируется в формате JSON.

## Installation

<div class="termy">

```console
git clone https://github.com/xgmsx/go-tiny-url
cd go-tiny-url
cp ./configs/.env_example ./configs/.env
cd ./configs/.env_localhost_example ./configs/.env_localhost
```

</div>

## Quick start

#### Установка инструментов для разработки

```shell
make install
```

Проверка работы установленных инструментов: 

```shell
make generate
make fmt
make lint
```

#### Запуск сервиса с зависимостями (postgres, redis, kafka, jaeger, prometheus, grafana) в Docker

<div class="termy">

```console
$ docker compose up -d --build

 ✔ Network go-tiny-url_default         Created                                                                                                                                                                       0.1s 
 ✔ Container go-tiny-url-app-1         Started                                                                                                                                                                       0.4s 
 ✔ Container go-tiny-url-postgres-1    Started                                                                                                                                                                       0.2s 
 ✔ Container go-tiny-url-redis-1       Started                                                                                                                                                                       0.2s 
...                                                                                                                                                                     0.3s 
```

</div>

#### Локальный запуск сервиса в Linux и MacOS

<div class="termy">

```console
# Linux, MacOS
$ export $(grep -v '^#' ./configs/.env_localhost | xargs) && go run ./cmd/app

# Windows Git Bash
$ env $(grep -v '^#' ./configs/.env_localhost | xargs) go run ./cmd/app

2025/01/01 12:00:00 maxprocs: Leaving GOMAXPROCS=8: CPU quota undefined
12:00:00 INF Logger initialized
12:00:00 INF App starting...
...
12:00:00 INF App started
```

</div>

## Usage

#### Migrations

**Warning**: При первом запуске сервиса нужно выполнить команды  `make migrate-install` и `make migrate-up`:

<div class="termy">

```console
$ make migrate-up

migrate -database "postgres://login:pass@localhost:5432/app-db?sslmode=disable" -path "./migrations" up
20241212085257/u init (19.9317ms)
```

</div>

#### HTTP-запросы

**Note**: Выполнять запросы можно в веб-интерфейсе http://localhost:8000/swagger

Установка утилиты `curl`:
```shell
# Linux
$ sudo apt-get install curl

# MacOS
$ brew install curl

# Windows
$ scoop install curl
```

Создание новой короткой ссылки:
```shell
curl -X 'POST' \
  'http://localhost:8000/api/shortener/v1/link' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "url": "https://google.com"
}'

# {"url":"https://google.com","alias":"IFIYr0OGRKeqF9jPUIbwww","expired_at":"2025-01-02T12:00:00.000000000Z"}
```

Получение полной ссылки:
```shell
curl -X 'GET' \
  'http://localhost:8000/api/shortener/v1/link/IFIYr0OGRKeqF9jPUIbwww' \
  -H 'accept: application/json'

# {"url": "https://google.com", "alias": "IFIYr0OGRKeqF9jPUIbwww", "expired_at": "2025-01-02T12:00:00.000000000Z"}
```

Переход по короткой ссылке:
```shell
# Linux and MacOS
open http://localhost:8000/api/shortener/v1/link/IFIYr0OGRKeqF9jPUIbwww/redirect

# Windows:
explorer http://localhost:8000/api/shortener/v1/link/IFIYr0OGRKeqF9jPUIbwww/redirect
```

#### gRPC-запросы

Установка утилиты `grpcurl`:
```shell
# Linux
$ sudo apt-get install grpcurl

# MacOS
$ brew install grpcurl

# Windows
$ scoop install grpcurl
```

Создание новой короткой ссылки:
```shell
$ grpcurl -d '{"url": "https://google.com"}' -plaintext localhost:50051 shortener_v1.Shortener/CreateLink

# {"url": "https://google.com", "alias": "IFIYr0OGRKeqF9jPUIbwww", "expired_at": "2025-01-02T12:00:00.000000000Z"}
```

Получение полной ссылки:
```shell
$ grpcurl -d '{"alias": "IFIYr0OGRKeqF9jPUIbwww"}' -plaintext localhost:50051 shortener_v1.Shortener/FetchLink

# {"url": "https://google.com", "alias": "IFIYr0OGRKeqF9jPUIbwww", "expired_at": "2025-01-02T12:00:00.000000000Z"}
```

#### Redis UI

Текущее содержимое cache в Redis можно посмотреть через веб-интерфейс http://localhost:8081
![screen-redis-ui.png](assets/screen-redis-ui.png)

#### Kafka UI

Посмотреть сообщения в Kafka можно через веб-интерфейс http://localhost:8383/ui/clusters/local/all-topics
![screen-kafka-ui.png](assets/screen-kafka-ui.png)

В топике [links-created](http://localhost:8383/ui/clusters/local/all-topics/links-created) содержатся информация о всех созданных коротких ссылках.

Сервис получает сообщения из [links-requested](http://localhost:8383/ui/clusters/local/all-topics/links-requested) и автоматически создает короткую ссылку при получении нового сообщения.

## Metrics

Посмотреть метрики сервиса можно в Grafana: http://localhost:3000/d/golang-metrics-dashboard/golang-metrics

![screen-grafana-dashboard.png](assets/screen-grafana-dashboard.png)

## Traces

Посмотреть трейсы сервиса можно в Jaeger: http://localhost:16686

![screen-jaeger-1.png](assets/screen-jaeger-1.png)

![screen-jaeger-2.png](assets/screen-jaeger-2.png)

## Environment variables

| Name                  | Type   | Expected | Default       | Description                              |
|-----------------------|--------|----------|---------------|------------------------------------------|
| APP_NAME              | string |          | url-shortener | service name                             |
| APP_VERSION           | string |          | 0.0.0         | service version                          |
| APP_ENV               | string |          | DEV           | service environment (DEV, PROD, etc)     |
| LOGGER_LEVEL          | string |          | error         | logging level (debug, info, warn, error) |
| LOGGER_PRETTY_CONSOLE | bool   |          | false         | logging format (text/json)               |
| SENTRY_DSN            | string |          |               | sentry DSN (disabled if empty)           |
