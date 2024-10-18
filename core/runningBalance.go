package core

import (
	"sort"

	"github.com/mattg1243/willow-server/db"
)

// Calculates the running balances of each event on insert / update
// of event that predates others. Must be passed ALL of a clients events to be accurate
func RecalcRunningBalances (events []db.GetEventsRow) []int {
	balance := 0
	runningBalances := []int{}

	// Make sure events are sorted by date
	sort.Slice(events, func (i, j int) bool {
		return events[i].Date.Time.Before(events[j].Date.Time)
	})

	// Iterate over the events and store the balances
	for i := 0; i < len(events); i++ {
		if !events[i].Charge {
			balance += int(events[i].Amount)
			runningBalances = append(runningBalances, balance)
		} else {
			balance -= int(events[i].Amount)
			runningBalances = append(runningBalances, balance)
		}
	}

	return runningBalances
}