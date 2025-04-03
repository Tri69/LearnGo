package main

import "fmt"

func CetakPerson() {
	fmt.Println("Name : Budi")
	fmt.Println("Age : 18")
	fmt.Println("Address : Jl. Melati")
}
func Person(name string, address string) string{
	var cetak string = "Hello Name is", name, "Address is", address
	//ce  := "Age budi is", address
	return cetak
	
}
func main() {
	CetakPerson()
	fmt.Println(Person("Budi", "Jl. Melati"))
}