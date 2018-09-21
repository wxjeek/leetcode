package main

import (
	`fmt`
)

/**

Given a sorted array and a target value, return the index if the target is found. If not, return the index where it would be if it were inserted in order.

You may assume no duplicates in the array.

Example 1:

Input: [1,3,5,6], 5
Output: 2
Example 2:

Input: [1,3,5,6], 2
Output: 1
Example 3:

Input: [1,3,5,6], 7
Output: 4
Example 4:

Input: [1,3,5,6], 0
Output: 0

 */

func searchInsert(nums []int, target int) int {
	if len(nums)==0 ||target<nums[0]{
		return 0
	}
	for i:=0;i<len(nums)-1;i++{
		v := nums[i]
		if v<target&&target<nums[i+1] {
			return i+1
		}
		if nums[i]==target {
			return i
		}
		if nums[i+1]==target {
			return i+1
		}
	}
	return len(nums)
}

func main()  {
	r := searchInsert([]int{2,3,4,5,6,7},8)
	fmt.Println(r)
}