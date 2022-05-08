package main

func productExceptSelf(nums []int) []int {
	leftProduct := make([]int,len(nums))
	rightProduct := make([]int,len(nums))
	res := make([]int,len(nums))
	rightProduct[0] = 1
	leftProduct[len(nums)-1] = 1
	for r:= 1 ; r <= len(nums) -1 ; r++{
		rightProduct[r] = nums[r-1]*rightProduct[r-1]
	}
	for l:= len(nums)-2; l >= 0; l--{
		leftProduct[l] = nums[l+1]*leftProduct[l+1]
	}
	for it:= 0; it < len(nums);it++{
		res[it] =rightProduct[it]*leftProduct[it]
	}
	return res
}
