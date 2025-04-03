package main 

import "fmt"

func main(){
	var dataPerson = []string{"Budi", "Andy", "Edo"}
	var DataInt = [...]int{4, 3, 5, 6, 9, 4}
	for i:= 0; i < len(dataPerson) ; i++ {
		fmt.Println( " Hello ", dataPerson[i])
	}
	
	sum := 0
	for i, nilai := range DataInt{
		sum += nilai
		fmt.Println(i)
	}
	
	v := 6
	
	
	
	fmt.Println(sum * 4)
}