package main

import "fmt"

func main() {
	arr := [4]int{1, 2, 3, 4}
	fmt.Println(arr)

	str := [5]string{"satu", "dua", "tiga", "empat", "lima"}
	fmt.Println(str)

	fmt.Println(str[3])
	fmt.Println(len(str))

	st := str[1:3]
	fmt.Println(st)

	str[1] = "Holla"
	fmt.Println(st)

	arr1 := [3]int{1, 2, 3}
	arr2 := [3]int{1, 2, 3}
	fmt.Println(arr1 == arr2)
	fmt.Println(arr1 != arr2)
}
