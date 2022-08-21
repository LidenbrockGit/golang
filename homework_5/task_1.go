package main

import (
	"fmt"
)

var iterationsCount int64 = 0

func main() {
	fmt.Println("Введите n-е число Фибоначчи:")

	var num int64
	_, err := fmt.Scanln(&num)
	if err != nil {
		panic(err)
	}

	fmt.Println(fib(num))
	fmt.Println("Количество итераций:", iterationsCount)
}

func fib(num int64) int64 {
	iterationsCount++
	if num <= 1 {
		return num
	} else {
		return fib(num-1) + fib(num-2)
	}
}
