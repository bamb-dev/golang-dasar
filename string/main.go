package main

import (
	"fmt"
	"strings"
)

func main()  {
	// string punya banyak buil-in functionnya

	name := "Ibrohim"
	fmt.Println(len(name))
	fmt.Println(name[0])
	fmt.Println(strings.ToUpper("haji suka makan bubur"))
	fruit := "mangga,pisang,maggis"
	fmt.Println(strings.Split(fruit, ","))
	fmt.Println(strings.HasPrefix("Hallo Dunia", "Hallo"))
	
}