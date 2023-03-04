package main

import (
	"fmt"
)

func main() {
	myBill := newBill("Mario's bill")

	fmt.Println(myBill.format())
}