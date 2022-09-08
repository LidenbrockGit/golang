package fibonacci

import "testing"
import "github.com/stretchr/testify/assert"

func TestFibonacci(t *testing.T) {
	assert.Equal(t, 55, Fib(10, map[int]int{}), "Они должны быть равны")

	assert.Equal(t, 0, Fib(0, map[int]int{}), "Они должны быть равны")

	assert.NotEqual(t, 34, Fib(10, map[int]int{}), "Они не должны быть равны")

	assert.Equal(t, 0, Fib(-3, map[int]int{}), "Они должны быть равны")
}
