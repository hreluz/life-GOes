package invoiceitem

// Item contains of information of a invoiceitem
type Item struct {
	id      uint
	product string
	value   float64
}

// New returns a new Item
func New(id uint, product string, value float64) Item {
	return Item{id, product, value}
}

// Value getter of Item.Value
// func (i Item) Value() float64 {
// 	return i.value
// }

type Items []Item

func NewItems(items ...Item) Items {
	var is Items

	for _, item := range items {
		is = append(is, item)
	}

	return is
}

// SetTotal is the setter of Invoice.total
func (is Items) Total() float64 {
	var total float64

	for _, item := range is {
		total += item.value
	}

	return total
}
