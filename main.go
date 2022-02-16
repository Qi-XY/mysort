package main

import "fmt"

func main() {
	nums := []int{2, 45, 34, 324, 3, 7, 555, 46}
	//	nums = selectSort(nums)
	//nums = insertSort(nums)
	nums = BinaryInsertionSort(nums)
	fmt.Println(nums)

}
