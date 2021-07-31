package main

import "fmt"

/**
 * 1. 两数之和
 *
 * 给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那两个整数，并返回他们的数组下标。
 *
 * 你可以假设每种输入只会对应一个答案。但是，数组中同一个元素不能使用两遍。
 *
 * 示例:
 * 给定 nums = [2, 7, 11, 15], target = 9
 * 因为 nums[0] + nums[1] = 2 + 7 = 9
 * 所以返回 [0, 1]
 *
 * 题目链接：https://leetcode-cn.com/problems/two-sum
 */

func twoSum(nums []int, target int) []int {
	if nums == nil {
		return nil
	}
	for i := 0; i < len(nums); i++ {
		anotherNum := target - nums[i]
		for j := i + 1; j < len(nums); j++ {
			if nums[j] == anotherNum {
				return []int{i, j}
			}
		}
	}
	return nil
}

/*
方法2：一遍哈希表
*/
func twoSum2(nums []int, target int) []int {
	if nums == nil {
		return nil
	}
	m := make(map[int]int, len(nums)) // <元素，元素索引>
	for i, v := range nums {
		complement := target - v
		if j, ok := m[complement]; ok {
			return []int{i, j}
		}
		m[v] = i
	}
	return nil
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	res := twoSum(nums, target)
	fmt.Println(res)

	// 下面比较了一下，var nums1 []int 和 []int{}
	var nums1 []int           // 未初始化，所以为nil
	fmt.Println(nums1)        // []
	fmt.Println(len(nums1))   // 0
	fmt.Println(nums1 == nil) // true

	// []int{} 已经初始化
	fmt.Println([]int{})        // []
	fmt.Println([]int{} == nil) // false
	fmt.Println(len([]int{}))   // 0
}
