package main

import (
	"fmt"
)

func main(){
	age := 30;
	names := []string{"Fatih","Can","Büşra","Berkay","Cesi","Mert"};

	fmt.Println(age <= 50)
	fmt.Println(age >= 50)
	fmt.Println(age == 45)
	fmt.Println(age != 50)

	if age < 30{
		fmt.Println("30 Yaşından küçüksün")
	}else if age == 30{
		fmt.Println("30 Yaşındasın")
	}else{
		fmt.Println("30 Yaşından büyüksün")
	}

	for index, value := range names{
		if index == 1 {
			fmt.Println("Giriş Başarılı")
			continue
		}
		if index > 1 {
			fmt.Println("Giriş Başarısız")
			break
		}
		
		fmt.Printf("Merhaba %v \n",value)
	}
}