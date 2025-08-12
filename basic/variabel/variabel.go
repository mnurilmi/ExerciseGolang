package main

import "fmt"

type Product struct {
	Name      string
	Price     uint
	TotalSold uint
	Discount  uint8
}

var products []Product = []Product{
	{"Product A", 100000, 200, 0},
	{"Product B", 67000, 12, 20},
	{"Product C", 56000, 80, 0},
	{"Product D", 1000, 1350, 0},
	{"Product E", 20.000, 1, 0},
	{"Product F", 38455, 7, 15},
	{"Product G", 76000, 5644, 0},
	{"Product H", 530120, 30, 10},
	{"Product I", 143.000, 54, 0},
	{"Product J", 16000, 109, 0},
}

func main() {
	res := 0

	for _, v := range products {
		res += (int(v.Price) * int(v.TotalSold)) - (int(v.Price) * int(v.Discount) / 100 * int(v.TotalSold))
	}
	fmt.Println(res)
}

/*
You are tasked to calculate Shopee's annual gross profit from 10 different product.

Product	Price	Total Sold	Discount
Product A	Rp 100.000	200	0
Product B	Rp 67.000	12	20
Product C	Rp 56.000	80	0
Product D	Rp 1.000	1350	0
Product E	Rp 20.000	1	0
Product F	Rp 38.455	7	15
Product G	Rp 76.000	5644	0
Product H	Rp 530.120	30	10
Product I	Rp 143.000	54	0
Product J	Rp 16.000	109	0
The formula to calculate the gross profit is:

Gross Profit = Σ ((Price x total sold) - (price x discount x total sold))

You need to return the annual gross profit value.
*/
