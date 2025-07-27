package main

import "fmt"

func circularRob(nums []int) int {
	n := len(nums)

	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}

	return max(robLinear(nums[:n-1]), robLinear(nums[1:]))
}

func robLinear(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}

	prev2 := nums[0]
	prev1 := max(nums[0], nums[1])

	for i := 2; i < n; i++ {
		current := max(prev1, prev2+nums[i])
		prev2 = prev1
		prev1 = current
	}

	return prev1
}

func main() {
	nums := []int{2, 3, 4, 7, 3, 5}
	res := circularRob(nums)
	fmt.Println(res)
}
