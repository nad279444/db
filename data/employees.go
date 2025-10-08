package data


import "github.com/nad279444/db/models"

// Employees is sample data you can import from main.go
var Employees = []models.User{
	{
		Name:    "John",
		Age:     "23",
		Contact: "23344333",
		Company: "Myrl Tech",
		Address: models.Address{
			City:    "Bangalore",
			State:   "Karnataka",
			Country: "India",
			PinCode: "410013",
		},
	},
	{
		Name:    "Paul",
		Age:     "25",
		Contact: "23344333",
		Company: "Google",
		Address: models.Address{
			City:    "San Francisco",
			State:   "California",
			Country: "USA",
			PinCode: "410013",
		},
	},
	{
		Name:    "Robert",
		Age:     "27",
		Contact: "23344333",
		Company: "Microsoft",
		Address: models.Address{
			City:    "Bangalore",
			State:   "Karnataka",
			Country: "India",
			PinCode: "410013",
		},
	},
}
