package main

import "fmt"

func main(){

	// %v: Değişkenin gerçek değerini yazdırır.
	// %T: Değişkenin türünü yazdırır.
	// %t: Boolean değişkenler için true veya false yazdırır.
	// %b: İkilik formatta yazdırır.
	// %c: Tek bir karakteri yazdırır.
	// %d: Onluk formatta yazdırır.
	// %o: Sekizlik formatta yazdırır.
	// %x: Onaltılık formatta yazdırır.
	// %f: Ondalık formatta yazdırır.
	// %e: E notasyonunda yazdırır.
	// %E: E notasyonunda yazdırır.
	// %s: Dizgi olarak yazdırır.
	// %q: Dizgi olarak yazdırır ve metin içinde gereken karakterleri düzgün şekilde yazdırır.
	// %x: Hexadecimal olarak yazdırır.
	// %p: Pointer adresini yazdırır.

	var name = "Fatih"
	var num1 = "25"
	var num2 = 25

	fmt.Println("Merhaba, Ben",name, " Yaşım",num1)
	fmt.Printf("Merhaba, Ben %q Yaşım %q \n",name,num1)
	fmt.Printf("Merhaba, Ben %v Yaşım %v \n",name,num1)
	fmt.Printf("Merhaba, Ben %v Yaşım %v \n",name,num1)
	fmt.Printf("Merhaba, Ben %T Yaşım Yaşım %d \n",name,num2)
	
	var str = fmt.Sprintln(23);  // printi değişkene atama
	fmt.Print(str)

	fmt.Printf("Değer: %q",25)
}