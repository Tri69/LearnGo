package main

import "fmt"

func main(){
	dataAnimal := [4]string{
		"Harimau",
		"Elang",
		"Ayam",
		"Bebek",
	}
	
	dataSlice := dataAnimal[0:2]
	lengthSlice := len(dataSlice)
	CapasitySlice := cap(dataSlice)
	
	fmt.Println(dataSlice)
	fmt.Println("Panjang Data", lengthSlice)
	fmt.Println("Kapasitas Data", CapasitySlice)
}