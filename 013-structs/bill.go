package main

// struct oluşturma 
type bill struct {
	name string
	item map[string]float64
	tip  float64
}

func newBill(name string) bill {


	// struct türetme
	b := bill{
		name: name,
		item: map[string]float64{},
		tip:  0,
	}

	return b
}