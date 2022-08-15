package main

import "fmt"
import "math"

func main() {
	var a int

	fmt.Println("Введите площадь круга")
	fmt.Scanln(&a)
    
	fmt.Println("Диаметр круга:", 2*math.Sqrt(float64(a)/math.Pi))
	fmt.Println("Длина окружности:", math.Sqrt(float64(a)*4*math.Pi))
}
