package insertion_sort

func InsertSort(arNums []int) []int {
	var arr []int
	for _, val := range arNums {
		arr = append(arr, val)
		for i := len(arr) - 1; i >= 0; i-- {
			if i == 0 {
				break
			}
			if arr[i] < arr[i-1] {
				arr[i], arr[i-1] = arr[i-1], arr[i]
			} else {
				break
			}
		}
	}

	return arr
}
