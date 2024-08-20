# go-wc

CLI tool для подсчёта слов или строк в инпуте из stdin.

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
```

## Запуск тестов

```bash
go build
go test -v
```