package insertion_sort

import (
	"fmt"
	"math/rand"
	"testing"
)

type testOb struct {
	exm         []int
	exp         []int
	description string
}

var testArr = []testOb{
	{
		exm:         []int{9, 2, 3, 2, 1, 5, 6, 8},
		exp:         []int{1, 2, 2, 3, 5, 6, 8, 9},
		description: "#1 Успешно",
	},
	{
		exm:         []int{9, 2, 3, 2, 1, 5, 6, 8},
		exp:         []int{1, 2, 2, 3, 5, 6, 8},
		description: "#2 Ошибка длины массива",
	},
}

func TestInsertSort(t *testing.T) {
	for _, el := range testArr {
		result := InsertSort(el.exm)
		for k, j := range result {
			if k >= len(el.exp) {
				t.Errorf(
					"Индек выходит за пределы слайса; индекс - %d, длина массива %d, %s",
					k,
					len(el.exp),
					el.description,
				)
				continue
			}
			if el.exp[k] != j {
				t.Errorf("Итерация %d: %d != %d, %s", k, el.exp[k], j, el.description)
			}
		}
	}
}

func ExampleInsertSort() {
	fmt.Println(InsertSort(testArr[0].exm))
	// Output: [1 2 2 3 5 6 8 1]
}

// Чтобы запустить бенчмарк нужно закомментировать верхние тесты
func BenchmarkInsertSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var exmAr = make([]int, 0, 20)
		for j := 0; j < 20; j++ {
			exmAr = append(exmAr, rand.Intn(9))
		}
		InsertSort(exmAr)
	}
}
