package main

import "fmt"

// struct oluşturma
type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{"pie": 5.99, "cake": 3.99},
		tip:   0,
	}

	return b
}

// struct a tanımlı fonksiyon oluşturma
func (b bill) format() string {
	fs := "Bill breakedown \n"

	// toplayıcı değişken
	var total float64 = 0

	// structs tan türetilen b objesi içerisindeki itemsleri dönüp toplamak için for kullanıyoruz
	for k, v := range b.items {

		// 
		fs += fmt.Sprintf("%-25v ...$%v \n",k+":",v)
		total += v
	}

	// %0.2f değişkenin ondalıklı kısmının uzunluğunu belirtir 
	fs += fmt.Sprintf("%-25v ...$%0.2f","Total:",total)

	return fs
}