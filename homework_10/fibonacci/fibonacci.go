package fibonacci

func Fib(num int, innerRes map[int]int) int {
	innerNum, exist := innerRes[num]
	if exist {
		return innerNum
	}

	if num <= 1 {
		return num
	} else {
		result := Fib(num-1, innerRes) + Fib(num-2, innerRes)
		innerRes[num] = result
		return result
	}
}
