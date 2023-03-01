package main

import (
	"fmt"
)

func main() {
	myBill := newBill("Mario's bill")

	myBill.addItem("Tea",2.99)
	myBill.addItem("onion soup",4.50)
	myBill.addItem("toffe pudding",4.95)


	myBill.updateTip(20)

	fmt.Println(myBill.format())
}