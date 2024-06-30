package sorting

import "fmt"

func findMinPos(arr []int) int {
	minimum := 0
	for i := range arr {
		if arr[i] < arr[minimum] {
			minimum = i
		}
	}
	return minimum
}

func swapArr(arr []int, minimum int) []int {
	tmp := arr[0]
	arr[0] = arr[minimum]
	arr[minimum] = tmp
	return arr
}

func SelectionSort(arr []int) []int {
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

func Run() {
	samplearr := []int{4, 3, 2, 1, 5}
	newarr := SelectionSort(samplearr)
	fmt.Println(newarr)
}
