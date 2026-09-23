# csvstat

Консольная утилита: считает статистику по одной колонке CSV-файла.

## Запуск

Из корня репозитория:

```bash
go run ./projects/csvstat <файл> <колонка>
```

## Пример

$ go run ./projects/csvstat projects/csvstat/testdata/data.csv age
rows: 3
unique: 3
sum: 81
avg: 27
min: 23
max: 31


Для текстовой колонки выводятся только `rows` и `unique`.

## Тесты

```bash
go test ./projects/csvstat -v
```