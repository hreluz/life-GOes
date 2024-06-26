package course

import "fmt"

type course struct {
	name    string
	price   float64
	isFree  bool
	userIDs []uint
	classes map[uint]string
}

func (c *course) SetClasses(classes map[uint]string) {
	c.classes = classes
}

func (c *course) SetName(name string) {
	c.name = name
}

func (c *course) SetUserIDs(userIDs []uint) {
	c.userIDs = userIDs
}

func (c *course) SetPrice(price float64) {
	c.price = price
}

func (c *course) Name() string   { return c.name }
func (c *course) Price() float64 { return c.price }

func New(name string, price float64, isFree bool) *course {
	if price == 0 {
		price = 30
	}

	return &course{
		name:   name,
		price:  price,
		isFree: isFree,
	}
}

func (c *course) PrintClasses() {
	text := "Classes are: "
	for _, class := range c.classes {
		text += class + ", "
	}

	fmt.Println(text[:len(text)-2])
}
