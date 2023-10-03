package main

import (
	initializer "github.com/SFWE403/UArizonaPharmacy/init"
)

// Import the InitializeDatabase function

func main() {
	err := initializer.InitializeDatabase()
	if err != nil {
		panic(err)
	}
}
