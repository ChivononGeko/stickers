# 🖨 Stickers Service

Проект на Go для генерации и автоматической печати стикеров по заказам из Poster. Поддерживает работу с модификаторами продуктов и кеширование данных.

## 🚀 Функциональность

- Получение списка продуктов из Poster API
- Кеширование данных продуктов
- Генерация PDF стикеров с заказом
- Автоматическая печать через SumatraPDF
- Фильтрация продуктов по ID и цеху
- Веб-сервер для приёма заказов

## 📦 Структура проекта

```bush
stickers/
├── cmd/
│
├── internal/
│ ├── api/
│ ├── config/
│ ├── models/
│ ├── processor/
│ ├── server/
│ │ ├── api/
│ │ └── config/
│ └── storage/
│
├── assets/
│ ├── data/
│ ├── font/
│ └── images/
│
├── .env
├── go.mod
├── go.sum
└── README.md
```

## 🛠 Настройки

Все конфигурации задаются через `.env` файл:

```env
POSTER_TOKEN=your_poster_token
PRINTER_NAME="Your Printer Name"
SUMATRA_PATH="C:\\Program Files\\SumatraPDF\\SumatraPDF.exe"
ALLOWED_IDS=123,456,789
```

## ✅ Запуск

```bash
go run cmd/main.go
```

`POST /print`
Принимает JSON-заказ и запускает генерацию + печать стикеров.

Пример запроса:

```json
{
	"client_name": "Dos",
	"close_time": "2024-04-10T20:15:00Z",
	"products": [
		{
			"id": "123",
			"modifications": "[{\"m\":1}]",
			"comment": "no sugar"
		}
	]
}
```

## 🧠 Особенности

Использует SumatraPDF для быстрой и бесшумной печати

Поддержка ID продуктов, исключённых из печати (ALLOWED_IDS)

Номер заказа автоматически инкрементируется и хранится в assets/data/order_number.txt

Если не удаётся получить продукты из Poster — используется кеш

## 🖨 Требования

Windows (для печати через SumatraPDF)

Установленный принтер

Установленный SumatraPDF

## ✍🏼 Автор

`Damir Usetov`
