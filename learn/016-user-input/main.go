package main

import (
	"bufio"
	"fmt"
	"os"
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

	// kullanıcıdan  bir harf alıyoruz
	opt, _ := getInput("Chose option (e - add item, s - save bill, t - add tip): ", reader)

	fmt.Println(opt)
}

func main() {
	myBill := createBill()
	promptOptions(myBill)
}