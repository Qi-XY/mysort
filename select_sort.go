package main

func selectSort(a []int) []int {
	for i := 0; i < len(a)-1; i++ {
		min := i
		for j := i + 1; j < len(a); j++ {
			if a[min] > a[j] {
				min = j
			}

		}
		a[i], a[min] = a[min], a[i]
	}
	return a

}
