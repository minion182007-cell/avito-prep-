# csvstat

Консольная утилита на Go: считает статистику по одной колонке CSV-файла.

## Запуск

Из корня репозитория:

```bash
go run ./projects/csvstat <файл> <колонка>
```

## Пример

```
$ go run ./projects/csvstat projects/csvstat/testdata/data.csv age
rows: 3
unique: 3
sum: 81
avg: 27
max: 31
min: 23
```

Для текстовой колонки выводятся только `rows` и `unique`.

## Ошибки

Сообщения пишутся в stderr, код выхода — 1.

| Ситуация | Сообщение |
|---|---|
| неверное число аргументов | `usage: csvstat <file> <column>` |
| файла нет | `error: file not found: <путь>` |
| колонки нет | `error: no such column: ...` |
| пустой файл | `error: file is empty: <путь>` |
| битый CSV | `error: read csv: ...` |

## Структура

```
csvstat/
├── main.go              — аргументы, открытие файла, вывод
├── internal/stats/      — чтение колонки и подсчёт статистики
│   ├── stats.go
│   └── stats_test.go
└── testdata/data.csv    — пример данных
```

## Тесты

```bash
go test -cover ./projects/csvstat/...
```

Покрытие пакета `stats` — 100%.