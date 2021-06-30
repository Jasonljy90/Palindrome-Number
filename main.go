package main

func isPalindrome(x int) bool {
	y := reverseNumber(x)
	if x == y {
		return true
	} else {
		return false
	}
}

func reverseNumber(num int) int {
	res := 0
	for num > 0 {
		remainder := num % 10
		res = (res * 10) + remainder
		num /= 10
	}
	return res
}
