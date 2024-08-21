# go-wc

CLI tool для подсчёта строк, слов или байтов в инпуте из stdin.

## Сборка

```bash
go build
```

## Как пользоваться

```bash
# Считать слова
cat main.go | ./go-wc

# Считать строки
cat main.go | ./go-wc -l

# Считать байты
cat main.go | ./go-wc -b
```

## Запуск тестов

```bash
make test

# или:
go build
go test -v
```