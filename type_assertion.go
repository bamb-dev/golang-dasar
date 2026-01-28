package main

import "fmt"

func random() any {
	return true
}

func main() {
	var result any = random()
<<<<<<< HEAD

=======
>>>>>>> 57a66b4015696082a1efa820149800c3c4794088
	//var resultString string = result.(string)
	//fmt.Println(resultString)
	//var resultInt int = result.(int)
	//fmt.Println(resultInt)

	switch value := result.(type) {
	case string:
		fmt.Println("String", value)
	case int:
		fmt.Println("Int", value)
	default:
		fmt.Println("Unknown", value)
	}
}
