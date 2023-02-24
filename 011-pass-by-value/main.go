package main

import "fmt"

// name değikenini değerini dewğiştirmek için bir fonksiyon olşuşturduk
func updateName(x string) string{
	// burada bize gelen parametrenin değerinin global değişmesini bekleriz fakat go 
	// fonksiyon içerisinde değişkenin bir kopyasını oluşturur ve biz onun değerini değiştiririz
	x = "updated name"

	// Değeri böyle alabilrisiniz
	return x
}

// lütfen klasör içerisindeki group-b resmine göz atın
// Burada map in değerini değiştirmek için fonksiyon oluşturduk
func updateMenu(y map[string]float64){
	// Burada mapın içinde var ise değeri kopyalıyoruz ve hafızadaki değerini değiştiriyoruz yok ise yeni oluşturuyoruz 
	y["coffee"] = 2.99
}

func main() {
	//- GROUP A Types -> string, int, bools, floats, arrays, structs
	
	// string değişken tanımladık
	name := "Fatih";
	
	// bu fonksiyon ile ismi güncelliyoruz
	updateName(name)
	
	// return ettiğimiz değeri alıyoruz
	fmt.Println(updateName(name))
	
	// burada name değişkeninin değerinin değişmesini bekleriz
	fmt.Println(name)


	//- GROUP B Types -> slices, maps, functions
	menu := map[string]float64{
		"pie": 5.95,
		"ice cream": 3.99,
	}

	// mapın içerisindeki bir itemin değerini değiştiyoruz
	updateMenu(menu)

	fmt.Println(menu)
}  