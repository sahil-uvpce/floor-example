package main

import (
	"fmt"
)

func checkPrime(x int64) []int64 {
	var i, j int64
	var numbers []int64
	for i = 2; i < x; i++ {
		for j = 2; j <= (i / j); j++ {
			if i%j == 0 {
				numbers = append(numbers, i)
				break // if factor found, not prime
			}
		}
		if j > (i / j) {
			//fmt.Printf("%d is prime\n", i)
		}
	}
	return numbers
}

func checkNoPrime(x int64) bool {
	var i, flag int64
	flag = 0
	for i = 2; i < x; i++ {
		if x%i == 0 {
			flag = 1
			break
		}
	}
	if flag == 0 {
		return true
	}
	return false
}

func checkDigitPrime(x []int64) []int64 {
	var numbers []int64
	var result bool
	for _, value := range x {
		//fmt.Println("slice item", i, "is", value)
		a := value
		for a > 0 {
			result = false
			no := a % 10
			//fmt.Println("x:", no)
			if no == 0 || no == 1 {
			} else {
				result = checkNoPrime(no)
			}
			if result == true {
				break
			}
			a = a / 10
		}
		if result != true {
			numbers = append(numbers, value)
		}
	}
	return numbers
}

func main() {
	var totalFloor int64
	//totalFloor = 69
	fmt.Println("Enter Total Floor NUmber:")
	fmt.Scan(&totalFloor)
	number := checkPrime(totalFloor + 1)
	result := checkDigitPrime(number)
	fmt.Println("from this floor you can buy any 7 floor which is satisfied our conditions:", result)
}
