package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func reverse(x int) int {

	var isNegative bool

	if x < 0 {
		isNegative = true
	}

	xInStr := strconv.Itoa(x)
	arr := strings.Split(xInStr, "")

	// revert slice
    for i := len(arr)/2-1; i >= 0; i-- {
        opp := len(arr)-1-i
        arr[i], arr[opp] = arr[opp], arr[i]
    }

    if isNegative {
        arr = arr[:len(arr) - 1]
        fmt.Println(arr)
        arr = append([]string{"-"}, arr...)
    }

    str := strings.Join(arr, "")

    i, _ := strconv.Atoi(str)

    if i > math.MaxInt32 || i < math.MinInt32 {
		return 0
	}

	return i
}

func main(){
	var r int
	// r = reverse(1534236469)
	// fmt.Println(r) // since 9646324351 > math.MaxInt32(2147483647), it should return 0

	// r = reverse(123)
	// fmt.Println(r) // 321

	r = reverse(-123)
	fmt.Println(r) // -321

	// r = reverse(120)
	// fmt.Println(r) // 21
}