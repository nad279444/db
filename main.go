package main

import (
	"encoding/json"
	"fmt"

	"github.com/nad279444/db/data"
	"github.com/nad279444/db/db"
	"github.com/nad279444/db/models"
)

func main() {
	dir := "./data"
	database, err := db.New(dir, nil)
	if err != nil {
		fmt.Println("Error creating db:", err)
		return
	}

	employees := data.Employees


	for _, emp := range employees {
		err := database.Write("users", emp.Name, emp)
		if err != nil {
			fmt.Println("Error writing record:", err)
		}
	}

	// ✅ Update example
	updatedUser := models.User{
		Name:    "Sammy",
		Age:     "23",
		Contact: "23344333",
		Company: "Hubtel",
		Address: models.Address{
			City:    "Accra",
			State:   "Airport",
			Country: "Ghana",
			PinCode: "410013",
		},
	}

	if err := database.Update("users", "Kwame", updatedUser); err != nil {
		fmt.Println("Error updating user:", err)
	} else {
		fmt.Println("User Kwame updated successfully!")
	}

	records, err := database.ReadAll("users")
	if err != nil {
		fmt.Println("Error reading records:", err)
		return
	}

	fmt.Println("All Records:")
	var allUsers []models.User
	for _, f := range records {
		var user models.User
		if err := json.Unmarshal([]byte(f), &user); err != nil {
			fmt.Println("Error unmarshalling user:", err)
			continue
		}
		allUsers = append(allUsers, user)
	}

	// if err := database.Delete("users", "John"); err != nil {
	// 	fmt.Println("Error", err)
	// }

	// if err := database.Delete("users", ""); err != nil {
	// 	fmt.Println("Error", err)
	// }

	fmt.Printf("%+v\n", allUsers)
}
