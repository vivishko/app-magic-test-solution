<p align="center">
  Тестовое в App Magic
</p>

## Task 1. Задание

Напишите программу на Go, которая в качестве входных данных получает набор текстов и в 3 потока выводит эти тексты в перевёрнутом виде. Каждая строка вывода должна содержать свой порядковый номер (1, 2, 3, ...) и номер потока (1, 2 или 3). Перед завершением работы программа должна вывести "done".

Тексты: "Hello", "qwerty", "Golang", "platypus", "тест", "level", "generics"

Пример вывода:

```bash
line 1, thread 1: "olleH"
line 2, thread 2: "gnaloG"
line 3, thread 1: "ytrewq"
...
done
```

Результат отправь, пожалуйста, ссылкой на gist.github.com или go.dev/play/ 

## Tech stack

Go

## Решение

запустить 3 горутины, через канал высылать данные в формате string, используя обработку рун обрабатывать строки внутри горутин, инкрементировать счетчик строки через мьютекс

## Инструкция по использованию

запустить существующую программу (в зависимости от того как называется файл, в этом случае main.go)

```bash
$ go run main.go
```

либо раскомментить строки 40-45 и закомментить 47-48 и запустить через 

```bash
$ go run main.go Hello qwerty Golang platypus тест level generics
```

**Результат:**

```bash
line 1, thread 3: "olleH"
line 2, thread 3: "supytalp"
line 3, thread 3: "тсет"
line 4, thread 3: "level"
line 5, thread 3: "scireneg"
line 6, thread 1: "ytrewq"
line 7, thread 2: "gnaloG"
done
```

## TODO
