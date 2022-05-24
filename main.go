package main

import (
	"fmt"
	"practice_go/dynamic_programming"
)

func main() {

	p := fmt.Println

	//fmt.Print(fact(4))

	//for i := 0; i < 10; i++ {
	//		fmt.Printf(",%d", fib(i))
	//}

	//p(checkPallindrome("ABCCDCDCDCDC"))
	//p(checkPallindrome("AABBAA"))
	//p(checkPallindrome("MALAYALAM"))

	arr := []int{1, 4, 1, 4, 5}

	//p(dynamic_programming.Longest_Increasing_SubSequence(arr))

	p(dynamic_programming.Sub_array_for_given_sum(arr, 5))

	//p(dynamic_programming.Longest_Common_SubSequence("Sumit", "Smnbt", len("Sumit"), len("Smnbt")))

	//stack.Print_stack()

}

func checkPallindrome(input_str string) bool {

	j := len(input_str) - 1

	for i := 0; i < len(input_str); i++ {

		if i >= j {
			break
		}
		if input_str[i] != input_str[j] {
			return false
		}
		j = j - 1
	}

	return true
}

func fact(n int) int {

	if n <= 2 {
		return n
	}

	return n * fact(n-1)
}

func fib(n int) int {

	if n == 0 || n == 1 {
		return n
	} else {
		return fib(n-1) + fib(n-2)
	}
}
