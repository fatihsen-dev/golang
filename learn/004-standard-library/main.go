package main

import (
	"fmt"
	"sort"
	"strings"
)

// go nun standart kütüphane için = https://pkg.go.dev/std

func main() {

	fmt.Println("============= Strings =============")

	greeting := "Hello there friends";
	// greeting içerisinde there stringi bulunursa true döndürür aksi takdirde false
	fmt.Println(strings.Contains(greeting,"there"))

	// Burada 1inci paramtredeki stringte 2 inci parametredeki değer var ise 2 inci parametreyi
	// 3 üncü parametre ile değiştiriyor aksi takdirde varsayılan stringi döndürüyor
	// var sayılan değerde değişiklik yapılmaz
	fmt.Println(strings.ReplaceAll(greeting,"Hello","hi"))

	// Metni küçük harflere çevirir
	fmt.Println(strings.ToLower(greeting))

	// Metni büyük harflere çevirir
	fmt.Println(strings.ToUpper(greeting))

	// Eşleşen değerin indexsini verir
	fmt.Println(strings.Index(greeting,"there"))

	// Split verdiğiniz değere göre stringi parçalar ve array haline getirir
	fmt.Println(strings.Split(greeting," "))


	fmt.Println("============= Sort ================")

	ages := []int{7,5,9,8,6,2,1,3,4}

	// sort.ints metodu ile sayıları küçükten büyüğe doğru sıralayabilirsiniz
	sort.Ints(ages)
	fmt.Println(ages)

	// sort.SearchInts metodu ile belirtiğiniz sayının indexsini bulabilirsiniz
	index := sort.SearchInts(ages,7)
	fmt.Println(index)

	// sort.Strings ile yazıları alfabetik şekilde sıralayabilirsiniz
	names := []string{"Fatih","Can","Büşra","Berkay","Cesi","Mert"}
	sort.Strings(names)
	fmt.Println(names)
}