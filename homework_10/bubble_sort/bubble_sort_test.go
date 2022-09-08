package bubble_sort

import (
	"fmt"
	"math/rand"
	"testing"
)

func ExampleBubbleSort() {
	fmt.Println(BubbleSort([]int{6, 4, 2, 8, 1, 9}))
	// Output: [1 2 4 6 8 1]
}

// Чтобы запустить бенчмарк нужно закомментировать верхние тесты
func BenchmarkBubbleSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var exmAr = make([]int, 0, 20)
		for j := 0; j < 20; j++ {
			exmAr = append(exmAr, rand.Intn(9))
		}
		BubbleSort(exmAr)
	}
}
