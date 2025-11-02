package main

import "fmt"

/*
! Struktur Data Mirip Array Tapi Lebih Flexsibel karena Ukurannya Bisa berubah (Dinamis)
	Ada tiga Istilah dalam slice
		1. Pointer
		2. Len
		3. Cap
*/

func main() {
	arr := []int{10, 20, 30, 40, 50, 60}

	// slice := arr[:3]
	// slice := arr[:]
	// slice := arr[2:]
	// slice := arr[1:4]

	slice := arr[1:4] // dari index 1 sampai sebelum 4
	fmt.Println(slice)
	fmt.Println(slice[0])
	fmt.Println("panjang : ", len(slice))
	fmt.Println("kapasitas : ", cap(slice))

	// ! Function
	/*
		- cap()
		- len()
		- append()
		- make()
		- copy()
		- arr[i,j]
	*/
	s := make([]int, 3, 5)
	fmt.Println(s)

	s2 := append(s, 60, 70)
	fmt.Println(s2)

	s3 := make([]int, 4)
	s4 := copy(s3, s2)
	fmt.Println("hasil copy : ", s3)
	fmt.Println("jumlah element yang tersalin : ", s4)

}
