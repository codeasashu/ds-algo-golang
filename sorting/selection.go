package sorting

import "fmt"

func SelectionSort(arr []int) []int {
	if len(arr) == 1 {
		return arr
	}
	for i := 0; i < len(arr)-2; i++ {
		mini := i
		for j := i; j < len(arr)-1; j++ {
			if arr[j] < arr[mini] {
				mini = j
			}
		}
		tmp := arr[i]
		arr[i] = arr[mini]
		arr[mini] = tmp
	}
	return arr
}

func RunSelectionSort() {
	samplearr := []int{4, 3, 2, 1, 5}
	newarr := SelectionSort(samplearr)
	fmt.Println(newarr)
}
