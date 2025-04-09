package main

import (
	"composition/store"
	"fmt"
)

func main() {
	boats := []*store.Boat{
		store.NewBoat("Kayak", 275, 1, false),
		store.NewBoat("Canoe", 400, 2, false),
		store.NewBoat("Tender", 650.25, 2, true),
	}
	crews := []*store.Crew{
		{Captain: "Captain Cook", FirstOfficer: "Alan"},
		{Captain: "Captain Henry", FirstOfficer: "Benjamin"},
		{Captain: "Captain Nico", FirstOfficer: "Charlie"},
	}

	for _, boat := range boats {
		fmt.Printf("Boat: %s, Price: %.2f\n", boat.Product.Name, boat.Price(0.2))
	}
	// generate rental boats from boats and crews
	rentalBoats := make([]*store.RentalBoat, len(boats))
	for i, boat := range boats {
		rentalBoats[i] = store.NewRentalBoat(boat.Name, boat.Price(0.2), boat.Capacity,
			boat.Motorized, true, crews[i].Captain, crews[i].FirstOfficer)
	}
	for _, rentalBoat := range rentalBoats {
		fmt.Printf("Rental Boat: %s, Price: %.2f, Crew: %s, %s\n",
			rentalBoat.Name, rentalBoat.Price(0.2), rentalBoat.Captain, rentalBoat.FirstOfficer)
	}

}
