package main

import "fmt"


// Паттерн применяется для создания упрощенного интерфейса взаимодействия со сложной системой, состоящей из множества классов
// Позволяет уменьшить связность системы
type Phone interface {
	ModelN() string
	Price() int
}

type Iphone struct{}

func (p Iphone) ModelN() string {
	return "iPhone 11"
}
func (p Iphone) Price() int {
	return 30000
}

type Samsung struct{}

func (s Samsung) ModelN() string {
	return "Samsung galaxy note 2"
}
func (s Samsung) Price() int {
	return 25000
}

type Blackberry struct{}

func (b Blackberry) ModelN() string {
	return "Blackberry Z10"
}
func (b Blackberry) Price() int {
	return 15000
}

type ShopKeeper struct {
	iphone     Phone
	samsung    Phone
	blackberry Phone
}

func (sk ShopKeeper) iphoneSale() {
	fmt.Println(sk.iphone.ModelN())
	fmt.Println(sk.iphone.Price())
}
func (sk ShopKeeper) samsungSale() {
	fmt.Println(sk.samsung.ModelN())
	fmt.Println(sk.samsung.Price())
}
func (sk ShopKeeper) blackberrySale() {
	fmt.Println(sk.blackberry.ModelN())
	fmt.Println(sk.blackberry.Price())
}
func newShopKeeper() ShopKeeper {
	return ShopKeeper{
		iphone:     Iphone{},
		samsung:    Samsung{},
		blackberry: Blackberry{},
	}
}

type FacadePatternClient struct {
	choice int
	sk     ShopKeeper
}

func (c *FacadePatternClient) sell() {
	fmt.Println("========= Mobile Shop ============")
	fmt.Println("         1. IPHONE.               ")
	fmt.Println("         2. SAMSUNG.              ")
	fmt.Println("         3. BLACKBERRY.           ")
	fmt.Println("         4. Exit.                 ")
	fmt.Println("Enter your choice:")
	for c.choice != 4 {
		fmt.Scan(&c.choice)
		switch c.choice {
		case 1:
			c.sk.iphoneSale()
		case 2:
			c.sk.samsungSale()
		case 3:
			c.sk.blackberrySale()
		default:
			fmt.Println("Nothing You purchased")
		}
	}
}

func NewClient() FacadePatternClient {
	return FacadePatternClient{
		choice: 0,
		sk:     newShopKeeper(),
	}
}
func main() {
	cl := NewClient()
	cl.sell()
}
