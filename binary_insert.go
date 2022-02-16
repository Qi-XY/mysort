package main

// BinaryInsertionSort 二分插入排序
func BinaryInsertionSort(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		left := 0      //左边界
		right := i - 1 //右边界：待排序元素的前一个位置
		//找到插入位置
		for left <= right {
			middle := (left + right) / 2
			if nums[middle] > nums[i] {
				right = middle - 1 //如果比中间位置小则缩小右边的范围为mid-1
			} else {
				left = middle + 1 //如果比中间位置大则扩大右边的范围为mid+1
			}
		}

		for j := i - 1; j >= left; j-- {
			nums[j], nums[j+1] = nums[j+1], nums[j]
		}
	}
	return nums
}
