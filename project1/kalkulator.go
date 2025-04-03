package main

import "fmt"

func TemplateGui() {
	fmt.Println("====================")
	fmt.Println("-----KALKULATOR-----")
	fmt.Println("1. Persegi")
	fmt.Println("2. Persegi Panjang")
	fmt.Println("3. Lingkaran")
	fmt.Println("4. Exit")
	fmt.Println("--------------------")
	fmt.Println("Massukan Berupa Angka(1 - 3) : ")
}
type Shape struct{
	panjang int
	lebar int
}
func (s Shape) Persegi()(int, int){
	Luas := s.panjang * s.panjang
	Keliling := s.panjang * 4
	return Luas, Keliling
}
func main() {
	var  pilihan string
	
	TemplateGui()
	fmt.Scanln(&pilihan)
	switch(pilihan) {
		case "1" : {
			var shapes = Shape{4, 5}
			var lu, kel int = shapes.Persegi()
			fmt.Println(lu)
			fmt.Println(kel)
			break
			//Persegi()
		}
		default :
			fmt.Println("Tidak")
		
	}
	
}