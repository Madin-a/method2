package main

import "fmt"

type Purchases struct {
	purchase [3]float64
	percent float64
	cashback float64
}

func (summa *Purchases)Cashback() float64  {
	var a float64
	for _, value := range summa.purchase{
		a = a+value

	}
	summa.cashback = a*summa.percent/100
	return summa.cashback
}
func main() {
	p := Purchases{
		purchase: [3]float64{123, 321, 231},
		percent:  10,
		cashback: 0,
	}
	p.cashback = p.Cashback()
	fmt.Println(p)

}
