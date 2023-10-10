# Tasty Bites Module

## Background

Made as a personal project for FTGO Non-Graded Challenge

## Installation

Make sure you have initialized `go mod init` on the same folder of your project.\
Then, copy the following command:

```bash
go get -u github.com/Tango992/moduletasty
```

## Usage

Use the `moduletasty.Start()` method to use the module.

```go
import "github.com/Tango992/moduletasty"

func main() {
	moduletasty.Start()
}
```

You will be prompted to enter a database connection (currently for private usage)

```console
Enter connection with the following format:
[USERNAME]:[PASSWORD]@tcp([HOST]:[PORT])/[DBNAME]
```

Once the connection successfully establised, the CLI app is ready to use!

```console
Options:
1. Add Employee
2. View Menu Items
3. Process Order
4. Exit
Enter your option: 
```

## Happy Coding

![err nil](easteregg.jpg)