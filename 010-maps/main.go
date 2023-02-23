package main

import "fmt"

func main() {

	menu := map[string]float64{
		"soup": 4.99,
		"pie": 7.99,
		"salad": 6.99,
		"toffee pudding": 3.55,
	}

	fmt.Println(menu)
	fmt.Println(menu["pie"])

	// loop maps
	for k,v :=range menu{
		println(k," - ",v)
	}

	// keyi sayı olan bir map
	phonebook := map[int]string{
		267584967: "mario",
		984759373: "luigi",
		845775485: "peach",
	}

	fmt.Println(phonebook)
	fmt.Println(phonebook[267584967])
	
	// map içerisinde bir itemin valuesini değiştirme
	phonebook[267584967] = "Fatih"
	fmt.Println(phonebook)

	// olmayan bir itemin değerini değiştirmeye çalışırsanız mape yeni olarak eklenir 
	phonebook[367584967] = "Test"
	fmt.Println(phonebook)
}