package main

import (
	"fmt"
)


func main(){
	names := []string{"Fatih","Can","Büşra","Berkay","Cesi","Mert"};

	// for 1
	x := 1;
	for x <= 5 {
		fmt.Println("Value:",x)
		x++
	}

	// for 2
	for i := 1; i <=5; i++{
		fmt.Println(i)
	}

	// for 3
	for i := 0; i < len(names); i++{
		fmt.Println(names[i])
	}

	// for 4
	for index,value:= range names{
		fmt.Printf("index: %v Value: %v \n", index+1, value)
	}

	// for 4
	for _,value:= range names{
		fmt.Printf("Value: %v \n",  value)
	}
}