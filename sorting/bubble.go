package sorting

import "fmt"

func BubbleSort(arr []int) []int {
	if len(arr) == 1 {
		return arr
	}
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if j+1 > len(arr)-1 {
				continue
			}
			if arr[j] > arr[j+1] {
				tmp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = tmp
			}
		}
	}
	return arr
}

func RunBubbleSort() {
	samplearr := []int{13, 46, 24, 52, 20, 9}
	newarr := BubbleSort(samplearr)
	fmt.Println(newarr)
}
