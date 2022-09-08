package bubble_sort

func BubbleSort(arNums []int) []int {
	var isTrue = true
	for isTrue {
		isTrue = false
		for i := 0; i < len(arNums)-1; i++ {
			if arNums[i] > arNums[i+1] {
				arNums[i], arNums[i+1] = arNums[i+1], arNums[i]
				isTrue = true
			}
		}
	}

	return arNums
}
