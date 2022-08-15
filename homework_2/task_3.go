package main

import (
	"fmt"
)

func main() {
	var a int
	var hundreds int
	var tens int
	var units int

	fmt.Println("Введите трехзначное число:")
	fmt.Scanln(&a)

	hundreds = a / 100
	tens = a / 10 % 10
	units = a % 10

	tHundreds := getPluralForm(hundreds, []string{"сотня", "сотни", "сотен"})
	tTens := getPluralForm(tens, []string{"десятка", "десятки", "десяток"})
	tUnits := getPluralForm(units, []string{"единица", "единицы", "единиц"})

	fmt.Printf("%d %s %d %s %d %s \n", hundreds, tHundreds, tens, tTens, units, tUnits)
}

/**
 * Склонение слов по числу
 */
func getPluralForm(n int, textForms []string) string {
	if n%10 == 1 && n%100 != 11 {
		return textForms[0]
	} else {
		if n%10 >= 2 && n%10 <= 4 && (n%100 < 10 || n%100 >= 20) {
			return textForms[1]
		} else {
			return textForms[2]
		}
	}
}
