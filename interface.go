package main

import "fmt"

type HasName interface {
	GetName() string
}

type Person struct {
	Name string
}
type Animal struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}
func (animal Animal) GetName() string {
	return animal.Name
}

func SayHello(value HasName) {
	fmt.Println("Hallo ", value.GetName())
}

func main() {
	person := Person{Name: "Ahmad"}
	animal := Animal{Name: "Kucing"}
	SayHello(person)
	SayHello(animal)
}
