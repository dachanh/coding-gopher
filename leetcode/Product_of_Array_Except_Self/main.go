package main

func productExceptSelf(nums []int) []int {
	res := make([]int,len(nums))
	isCheckZero := 0
	product := 1
	for it := 0; it < len(nums);it++{
		if nums[it] == 0{
			isCheckZero += 1
		}
	}
	if isCheckZero == 1{
		for it := 0 ;it < len(nums);it++{
			if nums[it] != 0{
				product = product*nums[it]
				res[it] = 0
			}
		}
		for it := 0; it < len(nums);it++{
			if nums[it] == 0{
				res[it] = product
			}
		}
		return res
	}else if isCheckZero > 1{
		for it := 0; it < len(nums);it++ {
			res[it] = 0
		}
		return  res
	}
	for it := 0 ;it < len(nums);it++ {
		product = product*nums[it]
	}
	for it := 0 ; it < len(nums);it++{
		res[it] = product/nums[it]
	}
	return res
}
