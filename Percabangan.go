package main

import "fmt"

func main(){
	var value int = 19
	switchFruit := "Appel"
	
	if value < 5 {//default True / benar
		fmt.Println("value lebih kecil dari 5")
	}else if value > 5 {
		fmt.Println("Value lebih besar dari 5")
	}else {
		fmt.Println("Value bukan angkka")
	}
	
	switch(switchFruit) {
		case "Appel" : {
			fmt.Println("Ini adalah Buah Appel")
			break
		}
		case "Jeruk" : {
			fmt.Println("Ini adalah Buah Jeruk")
			break
		}
		default: {
			fmt.Println("Bukan Termasuk buah diatas")
		}
	}
	
	
}