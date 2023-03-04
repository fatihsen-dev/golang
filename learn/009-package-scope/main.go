package main

import "fmt"

// Buradaki değişkene diğer fonksiyonlardan erişim sağlanabilir (global scope)
var score = 99.5;

func main() {

	// Fonksiyon içerisinde tanımlanmış değişkene farklı fonksiyonlardan ulaşılamaz
	var score = 99.5;

	sayHello("Fatih")

	for _, v := range points {
		fmt.Println(v)
	}

	showScore()
}