package main

// Import the InitializeDatabase function
import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	dbinitializer "github.com/SFWE403/UArizonaPharmacy/init"
	"github.com/SFWE403/UArizonaPharmacy/internal/models"
)

func main() {
	err := dbinitializer.InitializeDatabase()
	if err != nil {
		panic(err)
	}

	// Create a new HTTP server
	server := http.Server{
		Addr: "localhost:8080",
	}

	// Define an endpoint to retrieve all pharmacies
	http.HandleFunc("/pharmacies", func(w http.ResponseWriter, r *http.Request) {
		// Create a list of pharmacies
		pharmacies := []*models.Pharmacy{
			models.NewPharmacy("Walgreens", "123 Main St", true, "https://www.walgreens.com", "Walgreens Inc.", 5551234, "9am-9pm"),
			models.NewPharmacy("CVS", "456 Elm St", true, "https://www.cvs.com", "CVS Health", 5555678, "8am-10pm"),
			models.NewPharmacy("Walmart", "789 Oak St", true, "https://www.walmart.com", "Walmart Inc.", 5559012, "24 hours"),
		}

		// Convert the list of pharmacies to JSON
		pharmaciesJSON, err := json.Marshal(pharmacies)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Set the Content-Type header to application/json
		w.Header().Set("Content-Type", "application/json")

		// Write the JSON response to the client
		w.Write(pharmaciesJSON)
	})

	// Define an endpoint to retrieve all inventory items
	http.HandleFunc("/inventory", func(w http.ResponseWriter, r *http.Request) {
		// Create a list of inventory items
		inventory := []*models.InventoryItem{
			models.NewInventoryItem(1, "Aspirin", "Pain reliever", 4.99, 100),
			models.NewInventoryItem(2, "Ibuprofen", "Pain reliever", 6.99, 50),
			models.NewInventoryItem(3, "Acetaminophen", "Pain reliever", 5.99, 75),
			models.NewInventoryItem(4, "Benadryl", "Antihistamine", 8.99, 25),
			models.NewInventoryItem(5, "Claritin", "Antihistamine", 12.99, 20),
		}

		// Convert the list of inventory items to JSON
		inventoryJSON, err := json.Marshal(inventory)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Set the Content-Type header to application/json
		w.Header().Set("Content-Type", "application/json")

		// Write the JSON response to the client
		w.Write(inventoryJSON)
	})

	// Define an endpoint to checkout a sale
	http.HandleFunc("/checkout", func(w http.ResponseWriter, r *http.Request) {
		// Read the request body
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Parse the request body into a Sale struct
		var sale models.Sale
		err = json.Unmarshal(body, &sale)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Call the checkout function with the Sale struct
		total, err := models.Checkout(&sale)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Convert the total to JSON
		totalJSON, err := json.Marshal(total)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Set the Content-Type header to application/json
		w.Header().Set("Content-Type", "application/json")

		// Write the JSON response to the client
		w.Write(totalJSON)
	})

	// Start the server
	server.ListenAndServe()
}
