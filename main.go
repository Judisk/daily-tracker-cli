package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

var reader = bufio.NewReader(os.Stdin)
var filename = "data.csv"

// сначала хотел сделать дату const но передумал

func main() {
	if len(os.Args) > 1 && os.Args[1] == "stats" {

		statsMode()
		return
	} else {
		defaultMode()
	}

}
func statsMode() {
	fmt.Println("Режим статистики")
	matrix := readCSV()
	sleepS := middleStat(matrix, 1)
	energyS := middleStat(matrix, 2)
	moodS := middleStat(matrix, 3)
	fmt.Printf("Среднее по сну =%2.f, энергии =%2.f, настроению =%2.f \n", sleepS, energyS, moodS)
}

func defaultMode() {

	fmt.Println("Введите дневные показатели:")
	sleep := ask(0, 5, "Оцените свой сон")
	energy := ask(0, 5, "Оцените свою энергию")
	mood := ask(0, 5, "Оцените свое настроение")
	storage := ask(0, 50, "Сколько осталось таблеток")
	if storage < 7 {
		fmt.Println("Возможно стоит пополнить запас")
	}
	saveInCSV(sleep, energy, mood, storage)

}

func middleStat(matrix [][]string, idx int) (result float64) {
	l := len(matrix)
	if l <= 1 {
		fmt.Println("Недостаточно данных")
		return 0
	}
	for r := 1; r < l; r++ {
		temp, err := strconv.ParseFloat(matrix[r][idx], 64)
		if err != nil {
			panic(err)
		}
		result += temp
	}
	return result / float64(l-1)
}

func input(min, max int) int {

	for {
		text, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		text = strings.TrimSpace(text)
		num, err := strconv.Atoi(text)
		if err != nil {
			fmt.Println("Ошибка: введите число")
			continue
		}
		if num < min || num > max {
			fmt.Printf("Введите число от %d до %d\n", min, max)
			continue
		}
		return num
	}

}

func ask(min, max int, text string) int {
	fmt.Printf("%s (%d-%d): ", text, min, max)
	return input(min, max)
}

// если это не стринг то в конце чатси ругается на то что это не string
func saveInCSV(sleep, energy, mood, storage int) {

	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()
	stat, err := file.Stat()
	if err != nil {
		panic(err)
	}
	if stat.Size() == 0 {
		if err := writer.Write([]string{"Дата", "Сон", "Энергия", "Настроение", "Таблеток осталось"}); err != nil {
			panic(err)
		}
	}
	date := time.Now().Format("2006-01-02")
	if err := writer.Write([]string{date,
		strconv.Itoa(sleep),
		strconv.Itoa(energy),
		strconv.Itoa(mood),
		strconv.Itoa(storage)}); err != nil {
		panic(err)
	}
}

func readCSV() [][]string {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}
	return records

}
