package handler

import (
	"database/sql"
	"fmt"
)

func ViewMenuItems(db *sql.DB) error {
	rows, err := db.Query("SELECT * FROM MenuItems")
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			item_id int
			name string
			description string
			price float32
			category string
		)
		if err = rows.Scan(&item_id, &name, &description, &price, &category); err != nil {
			return err
		}
		fmt.Printf("ID\t\t: %v\n", item_id)
		fmt.Printf("Name\t\t: %v\n", name)
		fmt.Printf("Description\t: %v\n", description)
		fmt.Printf("Price\t\t: %v\n", price)
		fmt.Printf("Category\t: %v\n\n", category)
	}
	return nil
}