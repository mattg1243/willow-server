package core

import (
	"fmt"
	"sort"

	"github.com/mattg1243/willow-server/db"
)

// Calculates amount user can pay themselves from retainer
// Returns payable amount and event ids to be marked as paid
func CalculatePayout(events []db.GetEventsRow) (int, []db.GetEventsRow) {
	retainer := 0
	charges := 0
	var markAsPaid []db.GetEventsRow
	// Calculate available retainer
	for i := 0; i < len(events); i++ {
		if !events[i].Charge {
			retainer += int(events[i].Amount)
		} else if events[i].Charge && events[i].Paid.Bool {
			retainer -= int(events[i].Amount)
		}
		// need to check here for paid charged events and subtract amount from retainer
	}
	if retainer == 0 {
		return 0, markAsPaid
	}
	// fmt.Printf("Total Retainer: %d\n", retainer)
	// Sort by amount to ensure deterministic order
	sort.Slice(events, func(i, j int) bool {
		return events[i].Amount < events[j].Amount
	})

	remainingRetainer := retainer
	for j := 0; j < len(events); j++ {
		if events[j].Charge && !events[j].Paid.Bool {
			eventAmount := int(events[j].Amount)

			// fmt.Printf("Processing event %s with amount %d\n", events[j].ID, eventAmount)

			// Check if this event can be fully paid with the remaining retainer.
			if eventAmount <= remainingRetainer {
				charges += eventAmount
				remainingRetainer -= eventAmount
				markAsPaid = append(markAsPaid, events[j])

				fmt.Printf("Added event %s to markAsPaid. Current charges: %d, Remaining Retainer: %d\n", events[j].ID, charges, remainingRetainer)
			} else {
				fmt.Printf("Skipping event %s because it would exceed remaining retainer.\n", events[j].ID)
			}
		}
	}

	// fmt.Printf("Final Charges: %d, Events to Mark as Paid: %d\n", charges, len(markAsPaid))
	return charges, markAsPaid
}
