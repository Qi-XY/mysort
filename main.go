package main

import "fmt"

func main() {
	nums := []int{4, 5, 6, 7, 8, 3, 2, 1}
	nums = selectSort(nums)
	fmt.Println(nums)

}
