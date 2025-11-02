package main

import "fmt"

func main() {
	// leteral
	m := map[string]string{
		"nama":  "sairony",
		"kelas": "A",
	}
	fmt.Println(m)
	m["jurusan"] = "Teknik Informatika"
	fmt.Println(m)
	fmt.Println(m)

	// ! Fucntion
	/*
		m := make(map[string]int)
		m :=  make(map[string]int{"a" : 1})
		m["b"] = 2
		val := m["b"]
		val, ok := m["c"]
		delete(m, "a")
		len(m)
	*/

	_, ok := m["umur"]
	fmt.Println(ok)

}
