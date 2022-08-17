package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var scanner = bufio.NewScanner(os.Stdin)
	arr := []int64{}

	for scanner.Scan() {
		num, err := strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			panic(err)
		}

		arr = append(arr, num)
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

	fmt.Println(arr)
}
