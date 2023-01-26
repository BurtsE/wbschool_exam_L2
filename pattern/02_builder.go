package main

import "fmt"

/*
Паттерн Builder относится к порождающим паттернам. Определяет процесс поэтапного построения сложного продукта
После того как будет построена последняя его часть, продукт можно использовать.
*/
type House struct {
	city, street string
	apNum        int

	sold  bool
	cost  int
	owner string
}

type HouseBuilder struct {
	house *House
}

type HouseAdressBuilder struct {
	HouseBuilder
}

func (a *HouseAdressBuilder) City(name string) *HouseAdressBuilder {
	a.house.city = name
	return a
}
func (a *HouseAdressBuilder) Street(name string) *HouseAdressBuilder {
	a.house.street = name
	return a
}
func (a *HouseAdressBuilder) ApNum(num int) *HouseAdressBuilder {
	a.house.apNum = num
	return a
}

type HouseSellingBuilder struct {
	HouseBuilder
}

func (s *HouseSellingBuilder) For(cost int) *HouseSellingBuilder {
	s.house.cost = cost
	return s
}
func (s *HouseSellingBuilder) To(name string) *HouseSellingBuilder {
	s.house.owner = name
	return s
}

func NewHouseBuilder() *HouseBuilder {
	return &HouseBuilder{
		house: &House{},
	}
}

func (b *HouseBuilder) Adress() *HouseAdressBuilder {
	return &HouseAdressBuilder{*b}
}
func (b *HouseBuilder) Sold() *HouseSellingBuilder {
	b.house.sold = true
	return &HouseSellingBuilder{*b}
}
func (b *HouseBuilder) Build() *House {
	return b.house
}

func main() {
	h := NewHouseBuilder()
	fmt.Println(h.house.cost)

	house := h.Adress().
		City("Moskow").Street("").ApNum(14).
		Sold().
		For(20000).To("Ivan").Build()

	fmt.Println(house.cost)
}
