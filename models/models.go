package models

type MenuItem struct {
	Name  string
	Price int
}

type Order struct {
	ItemName, Status string
	Quantity, Total  int
}
