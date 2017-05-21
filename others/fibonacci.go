package main

import "fmt"

func fibonacci_recursive(n int) int {
	if n < 2 {
		return n
	}

	return fibonacci_recursive(n-1) + fibonacci_recursive(n-2)
}

func fibonacci_loop(n int) []int {

	//var f_array = [n]int{1, 1}
	var f_array = make([]int, n)

	// will cause panic
	//f_array =[]int{1,1}

	f_array[0] = 1
	f_array[1] = 1

	fmt.Printf("len of f_array is %d \n",len(f_array))

	for i := 2; i < n; i++ {
		f_array[i] = f_array[i-1] + f_array[i-2]
	}

	return f_array[:]
}

func main() {

	for _, v := range fibonacci_loop(10) {
		fmt.Printf("%d, ", v)
	}
}
