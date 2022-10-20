package main

import "fmt"

//Given an array nums. We define a running sum of an array as runningSum[i] = sum(nums[0]…nums[i]).
//Return the running sum of nums.

func runningSum(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	res := []int{nums[0]}
	for i := 1; i < len(nums); i++ {
		res = append(res, res[i-1]+nums[i])
	}
	return res
}

func main() {
	var n int
	var nums []int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		var tmp int
		fmt.Scan(&tmp)
		nums = append(nums, tmp)
	}
	fmt.Println(runningSum(nums))
}
