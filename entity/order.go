package entity

type Order struct {
	Order_id int
	Table_number int
	Employee_id int
	Order_date string
	Status string
}

type ProcessOrder struct {
	Employee_id int
	Table_number int
	Item_id int
	Quantity int
}