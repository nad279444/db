package data

import "github.com/nad279444/db/models"

var Employees = []models.User{
	{
		Name:    "Kwame",
		Age:     "28",
		Contact: "0244123456",
		Company: "MEST Africa",
		Address: models.Address{
			City:    "Accra",
			State:   "Greater Accra",
			Country: "Ghana",
			PinCode: "410013",
		},
	},
	{
		Name:    "Abena",
		Age:     "25",
		Contact: "0209876543",
		Company: "MTN Ghana",
		Address: models.Address{
			City:    "Kumasi",
			State:   "Ashanti",
			Country: "Ghana",
			PinCode: "4100234",
		},
	},
	{
		Name:    "Kojo",
		Age:     "30",
		Contact: "0501234567",
		Company: "Ecobank Ghana",
		Address: models.Address{
			City:    "Takoradi",
			State:   "Western",
			Country: "Ghana",
			PinCode: "4100235",
		},
	},
	{
		Name:    "Efua",
		Age:     "27",
		Contact: "0547654321",
		Company: "Vodafone Ghana",
		Address: models.Address{
			City:    "Cape Coast",
			State:   "Central",
			Country: "Ghana",
			PinCode: "4100236",
		},
	},
	{
		Name:    "Yaw",
		Age:     "32",
		Contact: "0245678901",
		Company: "Ghana Revenue Authority",
		Address: models.Address{
			City:    "Tamale",
			State:   "Northern",
			Country: "Ghana",
			PinCode: "4100237",
		},
	},
}
