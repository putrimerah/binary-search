package main

import "fmt"

func binarySarch(nums []int, target int) int {
	var low int = 0
	var high int = len(nums) - 1

	midtmp := (high - low) / 2
	fmt.Println(midtmp)
	for low <= high {
		mid := low + (high-low)/2

		if nums[mid] == target {
			return mid
		} else if target < nums[mid] {
			high = mid - 1
		} else if target > nums[mid] {
			low = mid + 1
		}
	}
	return -1
}

func main() {
	binary := []int{1, 2, 3, 4, 5}
	target := 5
	search := binarySarch(binary, target)

	fmt.Println(search)
}
