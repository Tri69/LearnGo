package main

import "fmt"

func main(){
	var Fruit [1]string
	var Person = [3]string{"Lisa", "Dewi", "Nisa"}
	
	//Instansiasi Jumlah Elemen Tak Terbatas
	//var Animal = [...]string{"Chicken", "Fish"}
	Fruit[0] = "Semangka"
	
	LengthData := len(Person)
	
	fmt.Println(Fruit[0])
	fmt.Println("Panjang Data = ", LengthData)
	fmt.Println(Person[0])
	fmt.Println(Person[1])
	fmt.Println(Person[2])
}