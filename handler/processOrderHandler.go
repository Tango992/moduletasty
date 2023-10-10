package handler

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"github.com/Tango992/moduletasty/entity"
)

func GetOrderValues() (entity.ProcessOrder, error) {
	scanner := bufio.NewScanner(os.Stdin)
	orderInput := entity.ProcessOrder{}

	fmt.Print("Enter employee ID: ")
	scanner.Scan()
	if _, err := fmt.Sscanf(scanner.Text(), "%d", &orderInput.Employee_id); err != nil {
		return entity.ProcessOrder{}, err
	}

	fmt.Print("Enter table number: ")
	scanner.Scan()
	if _, err := fmt.Sscanf(scanner.Text(), "%d", &orderInput.Table_number); err != nil {
		return entity.ProcessOrder{}, err
	}

	fmt.Print("Enter Item ID: ")
	scanner.Scan()
	if _, err := fmt.Sscanf(scanner.Text(), "%d", &orderInput.Item_id); err != nil {
		return entity.ProcessOrder{}, err
	}

	fmt.Print("Enter Quantity: ")
	scanner.Scan()
	if _, err := fmt.Sscanf(scanner.Text(), "%d", &orderInput.Quantity); err != nil {
		return entity.ProcessOrder{}, err
	}
	return orderInput, nil
}


func ProcessOrder(db *sql.DB, values entity.ProcessOrder) error {	
	_, err := db.Exec("INSERT INTO Orders (table_number, employee_id, order_date, status) VALUES (?, ?, CURRENT_DATE(), 'Placed')", values.Table_number, values.Employee_id)
	if err != nil {
		return err
	}

	currentOrderId, err1 := GetCurrentOrderId(db)
	if err1 != nil {
		return err1
	}

	subtotal, err2 := GetSubtotal(db, values)
	if err2 != nil {
		return err2
	}

	_, err3 := db.Exec("INSERT INTO OrderItems (order_id, item_id, quantity, subtotal) VALUES (?,?,?,?)", currentOrderId, values.Item_id, values.Quantity, subtotal)
	if err3 != nil {
		return err3
	}

	fmt.Printf("Order placed successfully! Subtotal: %v\n", subtotal)
	return nil
}


func GetCurrentOrderId(db *sql.DB) (int, error) {
	var currentOrderId int;

	row, err := db.Query("SELECT order_id FROM Orders ORDER BY order_id DESC LIMIT 1")
	if err != nil {
		return 0, err
	}
	defer row.Close()

	for row.Next() {
		if err = row.Scan(&currentOrderId); err != nil {
			return 0, err
		}
	}
	return currentOrderId, nil
}


func GetSubtotal(db *sql.DB, values entity.ProcessOrder) (float32, error) {
	var price float32

	row, err := db.Query("SELECT mi.price FROM OrderItems oi JOIN MenuItems mi ON oi.item_id = mi.item_id WHERE oi.item_id = ?", values.Item_id)
	if err != nil {
		return 0, err
	}
	defer row.Close()

	for row.Next() {
		if err = row.Scan(&price); err != nil {
			return 0, err
		}
	}
	return (price * float32(values.Quantity)), nil
}