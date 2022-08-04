package main

import "fmt"

func main() {
	var pos int
	var numbers = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("Arrayn elements are : ", numbers)

	fmt.Printf("Enter the position to remove :")
	fmt.Scan(&pos)

	numbers = append(numbers[:pos], numbers[pos+1:]...)
	fmt.Println("Array after removing  element : ", numbers)

	insertion()
}

func insertion() {
	var pos int
	num := 888
	var numbers = []int{100, 200, 300, 400, 500, 600, 700, 800, 900, 1000}
	fmt.Println("Arrayn elements are : ", numbers)

	fmt.Printf("Enter the position to insert :")
	fmt.Scan(&pos)

	numbers = append(numbers[:pos+1], numbers[pos:]...)
	numbers[pos] = num
	fmt.Println("Array after inserting  element : ", numbers)

}
