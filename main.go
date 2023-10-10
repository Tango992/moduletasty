package moduletasty

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"github.com/Tango992/moduletasty/config"
	"github.com/Tango992/moduletasty/handler"
)

func Start() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter connection with the following format:")
	fmt.Println("[USERNAME]:[PASSWORD]@tcp([HOST]:[PORT])/[DBNAME]")
	scanner.Scan()
	connection := scanner.Text()

	db, err := config.GetDB(connection)
	if err != nil {
		log.Fatal("Failed to connect db", err.Error())
	}
	defer db.Close()

	for {
		var option int
		displayMenu()

		scanner.Scan()
		if _, err := fmt.Sscanf(scanner.Text(), "%d", &option); err != nil {
			fmt.Println("Invalid input!")
			continue
		}

		switch option {
		case 1:
			values, err := handler.GetEmployeeInput()
			checkError(err)
			err = handler.AddEmployee(db, values)
			checkError(err)

		case 2:
			err = handler.ViewMenuItems(db)
			checkError(err)

		case 3:
			values, err := handler.GetOrderValues()
			checkError(err)

			err = handler.ProcessOrder(db, values)
			checkError(err)

		case 4:
			fmt.Println("Exiting program")
			os.Exit(0)
		}
	}
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func displayMenu() {
	fmt.Printf("\nOptions:\n")
	fmt.Println("1. Add Employee")
	fmt.Println("2. View Menu Items")
	fmt.Println("3. Process Order")
	fmt.Println("4. Exit")
	fmt.Print("Enter your option: ")
}