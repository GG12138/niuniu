package card

import (

	"strconv"
)

type Card struct {
	data  int
	value int
	color int
}
func NewCard(data , color , value int) *Card{
	card := new(Card)
	card.data = data
	card.color = color
	card.value = value
	return card
}

func (c Card) GetData() int{
	return c.data
}
func (c Card) String() string{
	color := c.getColor()
	value := c.getValue()
	data := c.GetData()
	return "[" + strconv.Itoa(data)+ " " + color + value + "]"
}

func (c Card) getValue() string{
	if c.value == 1 {
		return "A"
	}
	if c.value == 11 {
		return "J"
	}
	if c.value == 12 {
		return "Q"
	}
	if c.value == 13 {
		return "K"
	}
	return strconv.Itoa(c.value)
}

func (c Card) getColor() string{
	if c.color == 1 {
		return "红桃"
	}
	if c.color == 2 {
		return "黑桃"
	}
	if c.color == 3 {
		return "方块"
	}
	if c.color == 4 {
		return "梅花"
	}
	return ""
}


