package main

func sumDigitalNumber(N int) int {
	numTemp := N
	sum := 0
	for numTemp != 0 {
		sum += numTemp % 10
		numTemp /= 10
	}
	return sum
}

func Solution(N int) int {
	sumDigital := sumDigitalNumber(N)
	execpectNumber := sumDigital * 2
	itNumber := N + 1
	result := 0
	for true {
		tempSumDigitalNumber := sumDigitalNumber(itNumber)
		if tempSumDigitalNumber == execpectNumber {
			result = itNumber
			break
		}
		itNumber++
	}
	return result
}
