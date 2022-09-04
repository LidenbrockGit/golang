package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

var maxRows = flag.Int64("max_rows", 0, "Количество строк")
var columns = flag.String("columns", "", "Введите выбираемы столбцы")

func main() {
	flag.Parse()
	filename := strings.TrimSpace(flag.Arg(0))

	// Проверяем существует ли файл
	if _, err := os.Stat(filename); err != nil {
		log.Fatalln("Не могу проверить существование файла:", err)
	}

	// Открываем файл
	file, err := os.OpenFile(filename, os.O_RDWR, 666)
	if err != nil {
		log.Fatalln("Не могу открыть файл:", err)
	}

	// Закрываем файл при выходе из ф-ции
	defer func() {
		err := file.Close()
		if err != nil {
			log.Println("Не могу закрыть файл:", err)
		}
	}()

	// Проверяем флаги
	if *maxRows < 0 || 100 < *maxRows {
		log.Fatalln("max_rows должен быть целым числом от 1 до 100")
	}

	// Проверяем столбцы
	if len(*columns) == 0 {
		log.Fatalln("columns не должен быть пустой строкой")
	}
	columns := strings.Split(*columns, ",")

	// Создаем мапу для переданных столбцов куда будем складывать номера
	// столбцов из файла
	colNumbers := make(map[string]int64, len(columns))
	for _, val := range columns {
		colNumbers[strings.TrimSpace(val)] = 0
	}

	// Чтение csv файла
	csvReader := csv.NewReader(file)
	fileCols, err := csvReader.Read()
	if err != nil {
		log.Fatalln("Не могу получить столбцы csv файла")
	}

	// Находим переданные столбцы среди столбцов файла, и
	// запоминаем индексы столбцов файла
	var countFoundColumns int
	for key, val := range fileCols {
		if _, ok := colNumbers[val]; ok {
			colNumbers[val] = int64(key)
			countFoundColumns++
		}
	}

	// Если найдены не все столбцы, выводим ошибку
	if countFoundColumns != len(columns) {
		var notFoundColumns []string
		for _, val := range columns {
			if colIndex, _ := colNumbers[val]; colIndex == 0 {
				notFoundColumns = append(notFoundColumns, val)
			}
		}
		log.Fatalln("Не все столбцы найдены:", notFoundColumns)
	}

	// Читаем строки и выводим в нужном нам формате
	for i := int64(1); i <= *maxRows; i++ {
		row, err := csvReader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatalln("Неизвестная ошибка при чтении файла", filename)
		}

		fmt.Printf("----------------- Строка %d\n", i)
		for key, col := range colNumbers {
			fmt.Println(key+":", row[col])
		}
		fmt.Printf("----------------- Конец строки\n\n")
	}
}
