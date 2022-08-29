package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Чтение первого числа
func readFirstNumber(as []string) int {
	return stringToInt(as[0])
}

// Обработка строк
func processLine(as []string) {
	sum := 0
	for _, val := range as {
		number := stringToInt(val)
		sum = sum + number
	}
	fmt.Println(sum, " ")
}

// Чтение строки с разделением на части по пробелу
func readLineNumbs(in *bufio.Reader) []string {
	line, _ := in.ReadString('\n')
	line = strings.ReplaceAll(line, "\r", "")
	line = strings.ReplaceAll(line, "\n", "")
	numbs := strings.Split(line, " ")
	return numbs
}

// Функция конвертации строки в число
func stringToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
	}
	return i
}

func main() {
	// Стандартный поток ввода
	in := bufio.NewReader(os.Stdin)

	// Кол-во строк для обработки
	n := readFirstNumber(readLineNumbs(in))

	// Обработка строк
	for i := 0; i < n; i++ {
		processLine(readLineNumbs(in))
	}
}
