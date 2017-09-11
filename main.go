package main

import (
	"./card"
	"fmt"
)

func main(){
	mCard := make([]*card.Card,0)
	value := []int{1 ,2 ,3 ,4 ,5 ,6 ,7 ,8 ,9 ,10 ,11 ,12 ,13}
	data := 0
	for i := 1; i <= 4; i++ {
		for _, v := range value {
			data++
			Card := card.NewCard(data, i, v)
			mCard = append(mCard, Card)
		}
	}
	fmt.Println(mCard)
}
