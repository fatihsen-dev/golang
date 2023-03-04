package main

import "fmt"

// fonksiyona name nin hafızadaki kopyalanmış adresini gönderip değerini değiştiriyoruz
func updateName(x *string){
	*x = "Test Name"
}

func main() {
	name := "Fatih"
	
	// pointer ile name nin hafızadaki adresini kopyalıyoruz
	m := &name
	fmt.Println("Memory address: ", m)
	fmt.Println("Value ot memory address: ", *m)

	// namenin hafızadaki yerini parametre olarak veriyoruz
	updateName(m)

	fmt.Println(name)
}