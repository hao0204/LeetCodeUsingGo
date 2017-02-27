package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(convertToBase7(-7))
}

func convertToBase7(num int) string {
	if num == 0 {
		return "0"
	}
	flag := false
    	if num < 0 {
		flag = true
    		num = -num
	}
	res := ""
    	for num > 0 {
		a := num % 7;
		res = strconv.Itoa(a) + res
		num /= 7
	}
	if flag {
		res = "-" + res
	}
	return res
}