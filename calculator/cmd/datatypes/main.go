package main

import "fmt"

func main() {
	var title string
	var copies int
	var selling_price float64
	title = "For the Love of Go"
	copies = 99
	fmt.Println(title)
	fmt.Println(copies)
	var cost_price = 80.0
	selling_price = sale(cost_price,true,10.0)
	fmt.Println("Selling price after discount --->",selling_price)
}


func sale(cost_price float64,on_offer bool,discount float64) float64{
	var selling_price float64 = 0.0
	if on_offer{
		selling_price = cost_price * (1- discount/100)
	} else{
		selling_price = cost_price
	}
	



	return selling_price
}