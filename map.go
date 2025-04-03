package main

import "fmt"

type student struct {
	name string
}

type BudiTempalte interface{
	About() string
}

func (s student) About() string{
	 
	return `"Hello , i'm name is ", ${s.name}`
	
}
func SumValue(val ...int)int {
	sum := 0
	for _, n := range val{
		sum += n
	}
	average := sum / len(val)
	return average
}
func main(){
	var Nm string = "Budi 123"
	fmt.Println(Nm)
	fmt.Println(&Nm)
	var budis = student{"Budi Raharjo"}
	fmt.Println(budis.About())
	
	fmt.Println(SumValue(7, 9, 5, 4))
}