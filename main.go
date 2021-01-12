package main

import (
	"fmt"
	"math/rand"
)

func main() {
	unsortedList := []int{1, 7, 4, 1, 5, 10, 3, 74, 2, 8, 6, 9}

	fmt.Println("Unsorted list: ", unsortedList)
	fmt.Println("Sorted list: ", sort(unsortedList))

	unsortedRngList := generateRandomArray(15)
	fmt.Println("\nUnsorted RNG list: ", unsortedRngList)
	fmt.Println("Sorted RNG list: ", sort(unsortedRngList))
}

func sort(arr []int) []int {
	for i := 1; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

func generateRandomArray(n int) []int {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = int(0 + rand.Float32()*101)
	}
	return arr
}
