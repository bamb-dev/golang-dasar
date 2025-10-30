package main

import (
	"fmt"
	"strconv"
)

func main() {
	// ! Dalam perubahan type data, harus bisa membedakan : parsing dan konfersi

	// parsing itu yang gk akan jadi masalah jika dikonfersi.
	// konfersi itu ada kemungkinan error,
	// yang ada kemungkinan error :
	// 1. string ke int, karena bisa jadi string bukan angka
	// 2. string ke boolean

	// * Int to float64   --> type data harus explisit
	a := 10
	var b float64 = float64(a)
	fmt.Println(b)

	// * Int to string
	var score int = 40 // ! harus int. tidak bisa int8, int16, ...
	var scoreString string = strconv.Itoa(score)
	fmt.Println(scoreString)

	// * string to int
	// * ini konfersi karena ada kemungkinan error
	var text string = "100"
	var number int
	number, error := strconv.Atoi(text)
	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Println(number)
	}

	// * Boolean to string
	truth := true
	var str string = strconv.FormatBool(truth)
	fmt.Println(truth)
	fmt.Println(str)

	// *  string to Boolean
	val, _ := strconv.ParseBool("true")

	fmt.Println(val)
}
