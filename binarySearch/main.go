package main

import "fmt"

func binarySearch(arr []int, num int) int {
	var leftPointer = 0
	var rightPointer = len(arr) - 1
	for leftPointer <= rightPointer {

		var midPointer = int((leftPointer + rightPointer) / 2)
		var midValue = arr[midPointer]

		if midValue == num {
			fmt.Print("Value Fount at position ")
			return midPointer
		} else if midValue < num {
			leftPointer = midPointer + 1
		} else {
			rightPointer = midPointer - 1
		}
	}
	fmt.Print("Not Found")
	return -1
}

func main() {
	var array = []int{2, 9, 11, 21, 22, 32, 36, 48, 76}

	fmt.Println(binarySearch(array, 32))
}
