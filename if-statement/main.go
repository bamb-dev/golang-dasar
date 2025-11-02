package main

import "fmt"

func main() {
	angka := 2
	// ! Di if tidak perlum memberi kurung di sekitar pengkondisian

	// * contoh if short, if dengan membuat variabel baru dan hanya bisa digunakan pada block if saja
	if v := angka * 2; v > 2 {
		fmt.Println("v lebih besar dari 2")
	} else {
		fmt.Println("v lebih kecil dari 2")
	}
}
