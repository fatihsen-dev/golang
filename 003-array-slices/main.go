package main

import "fmt"

func main() {

	fmt.Println("============== Array ==============")
	
	// değişken oluşturmanın farklı bir yolu
	test := "test"
	fmt.Print(test)

	// Burada 4 değerli tipi int olan bir array oluşturuyoruz 4 değerden fazlasında hata alırsınız
	// arraylere belirlediğiniz tip dışında farklı bir tip değer atayamazsınız
	var numberArray1 [4]int = [4]int{10,20,30,40};
	fmt.Println(numberArray1)

	// farklı bir kullanım türü
	var numberArray2 = [4]int{10,20,30,40};

	// array uzunluğunu len fonksiyonu ile öğrenebilirsiniz
	fmt.Printf("Array length (%v)", len(numberArray2))
	fmt.Println()
	
	stringArray1 := [6]string{"Fatih","Can","Büşra","Berkay","Cesi","Test"}

	// arrayın 5 inci indexindeki değeri değiştiriyoruz
	stringArray1[5] = "Mert";
	fmt.Println(stringArray1)


	fmt.Println("============= Slices ==============")

	// slices oluşturuyoruz
	var numberSlices = []int{10,15,30};

	// numberSlicesnin 1 inci indexteki değerini değiştiriyoruz
	numberSlices[1] = 20;

	// numberSlices içerisine yeni bir değer ekliyoruz
	numberSlices = append(numberSlices,40)

	fmt.Println(numberSlices, len(numberSlices))

	// stringArray1 arrayın içerisindeki 0 ıncı indexten 3 üncü indexteki değere kadar alıyoruz
	rangeOne := stringArray1[0:3];
	
	// stringArray1 arrayın içerisindeki 0 ıncı indexten sonuncu index kadar olan değerleri alıyoruz
	rangeTwo := stringArray1[3:];

	// stringArray1 arrayın içerisindeki 3 üncü indexten ilk indexse kadar olan değerleri alıyoruz
	rangeTheree := stringArray1[:3];

	
	fmt.Println((rangeOne))
	fmt.Println((rangeTwo))
	fmt.Println((rangeTheree))
}