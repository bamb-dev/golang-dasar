package main

import "fmt"

func add(a, b int) (hasil int) {
	hasil = a + b // bukan variabel baru
	return
}

func sum(angka ...int) int {
	total := 0
	for _, val := range angka {
		total += val
	}
	return total
}

func main() {
	fmt.Println(add(1, 2))
	fmt.Println(sum())
	fmt.Println(sum(1, 2, 3, 4, 5, 20000000))
}
