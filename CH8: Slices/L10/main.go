package main

func sum(nums ...int) int {

	var endOfMonthBill int

	for i := 0; i < len(nums); i++ {
		endOfMonthBill += nums[i]
	}

	return endOfMonthBill
}
