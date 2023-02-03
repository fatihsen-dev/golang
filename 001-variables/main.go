package main

import "fmt"

func main(){
	//! String variable
	fmt.Println("=========================");
	fmt.Println("STRING VARIABLE");

	// string tipinde bir değişkene string atama
	var string1 string =  "String 1";
	fmt.Println("String1:",string1);

	// değişkene string atama
	var string2 = "String 2";
	fmt.Println("String2:",string2);

	// tipi string olan bir değişken oluşturma
	var string3 string;
	fmt.Println("String3:",string3);

	// değişkene string atamanın farklı bir versionu
	string4 := "String 4";
	fmt.Println("String4:",string4);
	

	//! Int variable
	fmt.Println("=========================")
	fmt.Println("INT VARIABLE");

	// tipi int olan değişkene int değer atama
	var int1 int = 23;
	fmt.Println("int1:",int1);

	// değişkene int değer atama
	var int2 = 24;
	fmt.Println("int2:",int2);

	// tipi int olan değişken oluşturma
	var int3 int;
	fmt.Println("int3:",int3)

	// değişkene int atamanın farklı bir versionu
	int4 := 25;
	fmt.Println("int4:",int4);


	//! BIT Int variable
	fmt.Println("=========================")
	fmt.Println("BIT VARIABLE");

	//- INT
	var BitInt1 int8 = 100;
	fmt.Println("BitInt1:",BitInt1);

	// int16 -32768 ile 32767 arası değer alabilir
	var BitInt2 int16 = 1000;
	fmt.Println("BitInt2:",BitInt2);

	// int32 -2147483648 ile 2147483647 arası değer alabilir
	var BitInt3 int32 = 1000000000;
	fmt.Println("BitInt3:",BitInt3);

	// int64 -9223372036854775808 ile 9223372036854775807 arası değer alabilir
	var BitInt4 int64 = 100000000000000000;
	fmt.Println("BitInt4:",BitInt4);

	//- UINT
	// uint 0 ile 18446744073709551615 arası değer alabilir
	var BitInt5 uint = 20;
	fmt.Println("BitInt5:",BitInt5);

	// uint8 0 ile 255 arası değer alabilir
	var BitInt6 uint8 = 200;
	fmt.Println("BitInt6:",BitInt6);

	// uint16 0 ile 65535 arası değer alabilir
	var BitInt7 uint16 = 2000;
	fmt.Println("BitInt7:",BitInt7);

	// uint32 0 ile 4294967295 arası değer alabilir
	var BitInt8 uint32 = 2000000000;
	fmt.Println("BitInt8:",BitInt8);

	// uint64 0 ile 18446744073709551615 arası değer alabilir
	var BitInt9 uint64 = 200000000000000000;
	fmt.Println("BitInt9:",BitInt9);


	//! Float variable
	fmt.Println("=========================")
	fmt.Println("BIT VARIABLE");

	var float1 float32 = 99999999999999999999999999999999999999;  // ondalıklı sayı
	fmt.Println("Float1:",float1);

	var float2 float64 = 99999999999999999999999999999999999999;  // 
	fmt.Println("Float2:",float2);
}