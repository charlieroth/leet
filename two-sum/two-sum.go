package twosum

import (
	"fmt"
)

// Given an array of integers nums and an integer target, return indices of the
// two numbers such that they add up to target.
//
// You may assume that each input would have exactly one solution, and you may
// not use the same element twice.
//
// You can return the answer in any order.

func twoSum(nums []int, target int) []int {
	var answer []int
	curr := 0

	for i := 1; i < len(nums); i++ {
		if i == curr {
			curr += 1
			continue
		}

		sum := nums[i] + nums[curr]
		fmt.Printf("i=%d, curr=%d, sum=%d\n", i, curr, sum)
		if sum == target {
			answer = []int{i, curr}
			break
		}

		curr += 1
	}

	return answer
}

func test1() {
	fmt.Println("test 01")
	nums := []int{2, 7, 11, 15}
	target := 9
	result := twoSum(nums, target)
	expected := []int{0, 1}
	fmt.Printf("%v =?= %v\n", expected, result)
}

func test2() {
	fmt.Println("test 02")
	nums := []int{3, 2, 4}
	target := 6
	result := twoSum(nums, target)
	expected := []int{1, 2}
	fmt.Printf("%v =?= %v\n", expected, result)
}

func test3() {
	fmt.Println("test 03")
	nums := []int{3, 3}
	target := 6
	result := twoSum(nums, target)
	expected := []int{0, 1}
	fmt.Printf("%v =?= %v\n", expected, result)
}

func test4() {
	fmt.Println("test 04")
	nums := []int{3, 2, 3}
	target := 6
	result := twoSum(nums, target)
	expected := []int{0, 2}
	fmt.Printf("%v =?= %v\n", expected, result)
}

func Proof() {
	// test1()
	// test2()
	// test3()
	test4()
}
