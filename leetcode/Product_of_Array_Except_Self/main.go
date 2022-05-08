package main

func productExceptSelf(nums []int) []int {
	leftProduct := 1
	rightProduct := 1
	res := make([]int,len(nums))
	res[0] = 1
	res[len(nums)-1] = 1
	for r:= 1 ; r <= len(nums) -1 ; r++{
		rightProduct = nums[r-1]*rightProduct
		res[r] = rightProduct
	}
	for l:= len(nums)-2; l >= 0; l--{
		leftProduct = leftProduct*nums[l+1]
		res[l] = leftProduct*res[l]
	}
	return res
}
