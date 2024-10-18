package core

import (
	"github.com/mattg1243/willow-server/db"
)

// Iterates over an array of events and returns the current balance
func CalculateBalance(events []db.GetEventsRow) int {
	balance := 0
	for i := 0; i < len(events); i++ {
		// Check if the event is a charge and adjust the balance accordingly
		if events[i].Charge {
			balance -= int(events[i].Amount)
		} else {
			balance += int(events[i].Amount)
		}
	}
	return balance
}