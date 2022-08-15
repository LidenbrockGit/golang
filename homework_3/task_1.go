package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	var a float64
	var b float64
	var operator string

	fmt.Println("Введите операцию (+, -, *, ^, /, f)")
	fmt.Scanln(&operator)

	fmt.Println("Введите цифру")
	fmt.Scanln(&a)

	if operator != "f" {
		fmt.Println("Введите вторую цифру")
		fmt.Scanln(&b)
	}

	switch operator {
	case "+":
		fmt.Printf("Сумма: %.2f \n", a+b)
	case "-":
		fmt.Printf("Разница: %.2f \n", a-b)
	case "*":
		fmt.Printf("Произведение: %.2f \n", a*b)
	case "/":
		if b == 0 {
			fmt.Println("Нельзя делить на ноль")
			os.Exit(1)
		}
		fmt.Printf("Частное: %.2f \n", a/b)
	case "^":
		fmt.Printf("Результат возведения в степень: %.2f \n", math.Pow(a, b))
	case "f":
		fmt.Printf("Факториал числа: %d \n", factorial(uint(a)))
	default:
		fmt.Println("Операция выбрана неверно")
		os.Exit(1)
	}
}

func factorial(a uint) uint {
	if a == 0 {
		return 1
	}

	return a * factorial(a-1)
}
