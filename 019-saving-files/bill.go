package main

import (
	"fmt"
	"os"
)

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

func newBill(name string) bill {

	// struct türetme
	b := bill{
		name:  name,
		items: map[string]float64{"pie": 5.99, "cake": 3.99},
		tip:   0,
	}

	return b
}

// structa tanımlı fonksiyon oluşturma (itemslerin fiyatlarını toplama)
func (b *bill) format() string {
	fs := "Bill breakedown \n"
	var total float64 = 0

	for k, v := range b.items {
		fs += fmt.Sprintf("%-25v ...$%v \n",k+":",v)
		total += v
	}

	// add tip
	fs += fmt.Sprintf("%-25v ...$%v \n","Tip:",b.tip)
	
	// total
	fs += fmt.Sprintf("%-25v ...$%0.2f","Total:",total)

	return fs
}

// struct içerisindeki tip değişkeninin değerini değiştirme (pointer ile)
func (b *bill) updateTip(tip float64){
	b.tip = tip
}

// items içerisine yeri bir item ekleme
func (b *bill) addItem(name string,price float64){
	b.items[name] = price
}

func (b *bill) save(){
	data := []byte(b.format())

	err := os.WriteFile(b.name+".txt",data,0644)

	if err != nil {
		panic(err)
	}
	fmt.Println("bill was saved to file")
}