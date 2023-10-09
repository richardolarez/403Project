package main

// Import the InitializeDatabase function
import (
	dbinitializer "github.com/SFWE403/UArizonaPharmacy/init"
)

func main() {
	err := dbinitializer.InitializeDatabase()
	if err != nil {
		panic(err)
	}
}
