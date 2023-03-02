package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Fazla kod yazmak yerine yeniden kullanılabilir fonksiyon oluşturmak daha mantıklı
// hem kodunuzun daha temiz görünmesini hemde daha az kod yazmanızı sağlar
func getInput(prompt string,r *bufio.Reader) (string, error){
	fmt.Print((prompt))

	input, err := r.ReadString('\n')

	return strings.TrimSpace(input),err
}

func createBill() bill{

	// kullanıcıdan girdi almak için bir reader oluşturuyoruz
	// bufio paketi bu işi bizim için yapıyor girdinin terminaldenmi başka 
	// şekildemi yapılacağını os paketi ile belirliyoruz
	reader := bufio.NewReader(os.Stdin)

	name, _ := getInput("Create a new bill name: ",reader)

	// Veri almadan önce mesaj mesaj gösteriyoruz
	//- fmt.Print("Create a new bill name: ")

	// Burada işimize şuan yaramicak şeyleri almıyoruz
	// verinin hangi tuştan sonra alacağını belirtiyoruz
	// alt satıra \n ile geçildği için bizde enter
	// yapılınca alacağımızı belirtiyoruz
	//- name, _ := reader.ReadString('\n')

	// fazla olan boşlukları strings paketinde bulunan TrimSpace fonksiyonu ile  siliyoruz
	//- name = strings.TrimSpace(name)

	// oluşturduğumuz veriyi bir değişkene atıyoruz
	b := newBill(name)
	fmt.Println("create the bill - ",b.name)

	// veriyi return ediyoruz
	return b
}

func promptOptions(b bill){
	reader := bufio.NewReader(os.Stdin)

	opt, _ := getInput("Chose option (a - add item, s - save bill, t - add tip): ", reader)

	switch opt{
	case "a":
		name, _ := getInput("Item name: ",reader)
		price, _ := getInput("İtem price: ",reader)

		p, err := strconv.ParseFloat(price,64)

		if err != nil{
			fmt.Println("the price must be a number")
			promptOptions(b)
		}
		b.addItem(name,p)
		fmt.Println("item added - ",name,price)
		promptOptions(b)
	case "t":
		tip, _ := getInput("Enter tip amount ($): ",reader)
		
		t, err := strconv.ParseFloat(tip,64)

		if err != nil{
			fmt.Println("the tip must be a number")
			promptOptions(b)
		}
		b.updateTip(t)

		fmt.Println("tip added - ", tip)
		promptOptions(b)
	case "s":
		b.save()
		fmt.Println("you saved the file - ",b.name)
	default:
		fmt.Println("that not a invalid option")
		promptOptions(b)
	}
}

func main() {
	myBill := createBill()
	promptOptions(myBill)
}