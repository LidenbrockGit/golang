package main

import "fmt"

func main() {
	var a, b int

	fmt.Println("Введите сторону прямоугольника:")
	fmt.Scanln(&a)

	fmt.Println("Введите вторую сторону прямоугольника:")
	fmt.Scanln(&b)

	fmt.Println("Площадь прямоугольника:", a*b)
}
