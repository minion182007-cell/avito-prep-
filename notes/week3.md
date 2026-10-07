# Неделя 3 — разбор

## Что получается

- Переписал `ping` и `csvstat` с нуля.
- Алгоритмы идут неплохо: 3Sum, Valid Parentheses, Queue via Stacks.
- Разделил `csvstat` на пакеты: `main.go` + `internal/stats`.

## Что даётся тяжело

**Ошибки.** `%w` и `%v`, `errors.Is` и `errors.As` (As нужен `&`),
sentinel-ошибки, обёртка через `fmt.Errorf("read csv: %w", err)`.

**Тесты.** Table-driven тесты, `t.Run`, анонимные структуры,
`strings.NewReader` вместо файла, пакет теста = пакет кода.

В `csvstat` трудности были именно в тестах и ошибках.

## Что повторить

- [ ] Ошибки: свой тип ошибки + `Unwrap`, `errors.Is` / `errors.As`.
- [ ] Тесты: написать table-driven тест с нуля без подсказок.

## Coverage

- `go test -cover ./projects/csvstat/...` — запускает тесты и считает долю выполненных строк.
- `go test -coverprofile=cover.out` + `go tool cover -html=cover.out` — красные строки не покрыты.
- Было меньше 100%: не покрыты ошибка `ReadAll` и `min = v64`.
- `min = v64` была красной, потому что в тестах числа шли по возрастанию.
  Добавил несортированный случай `3, 1, 2` и тест на битый CSV (`csv.ErrFieldCount`).
- Стало **100%**.
- 100% coverage ≠ нет багов: каждая строка выполнилась, но правильность проверяет только `want`.
