package core

import (
	"fmt"
	"testing"

	"github.com/mattg1243/willow-server/db" // Replace with your actual db package path
	"github.com/stretchr/testify/assert"
)

// TestCalculatePayout tests the CalculatePayout function with various scenarios.

// TestCalculatePayout tests the CalculatePayout function with various scenarios.
// TestCalculatePayout tests the CalculatePayout function with various scenarios.
func TestCalculatePayout(t *testing.T) {
	tests := []struct {
		name        string
		events      []db.GetEventsRow
		expectedAmt int
	}{
		{
			name: "Multiple retainers and charges within limits",
			events: []db.GetEventsRow{
				createTestEvent("11111111-1111-1111-1111-111111111111", 500, false), // Retainer payment
				createTestEvent("22222222-2222-2222-2222-222222222222", 300, false), // Another retainer payment
				createTestEvent("33333333-3333-3333-3333-333333333333", 200, true),  // Chargeable event
				createTestEvent("44444444-4444-4444-4444-444444444444", 250, true),  // Chargeable event
				createTestEvent("55555555-5555-5555-5555-555555555555", 100, true),  // Chargeable event
				createTestEvent("66666666-6666-6666-6666-666666666666", 50, true),   // Chargeable event
				createTestEvent("77777777-7777-7777-7777-777777777777", 150, true),  // Chargeable event
			},
			expectedAmt: 750, // Up to the combined retainers of 800 (500 + 300)
		},
		{
			name: "Multiple retainers with overcharges and partial fulfillment",
			events: []db.GetEventsRow{
				createTestEvent("33333333-3333-3333-3333-333333333333", 500, false), // Retainer payment
				createTestEvent("44444444-4444-4444-4444-444444444444", 500, false), // Retainer payment
				createTestEvent("55555555-5555-5555-5555-555555555555", 100, true),  // Chargeable event
				createTestEvent("66666666-6666-6666-6666-666666666666", 250, true),  // Chargeable event
				createTestEvent("77777777-7777-7777-7777-777777777777", 300, true),  // Chargeable event
				createTestEvent("88888888-8888-8888-8888-888888888888", 600, true),  // Chargeable event
				createTestEvent("99999999-9999-9999-9999-999999999999", 200, true),  // Chargeable event
				createTestEvent("aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa", 100, true),  // Chargeable event
			},
			expectedAmt: 950, // Charges up to the retainer limit of 1000, can't pay the 600 event fully
		},
		{
			name: "Overcharges with no retainer payments",
			events: []db.GetEventsRow{
				createTestEvent("11111111-1111-1111-1111-111111111111", 100, true), // Chargeable event
				createTestEvent("22222222-2222-2222-2222-222222222222", 200, true), // Chargeable event
				createTestEvent("33333333-3333-3333-3333-333333333333", 300, true), // Chargeable event
			},
			expectedAmt: 0, // No retainer payments, so no payout
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Call the function under test
			actualAmt, actualEvents := CalculatePayout(tt.events)

			// Calculate the total of the events in actualIDs
			totalCharged := 0
			for _, actualEvent := range actualEvents {
				for _, event := range tt.events {
					if event.ID == actualEvent.ID {
						totalCharged += int(event.Amount)
					}
				}
			}

			// Assert that the total charged is equal to the calculated payout
			assert.Equal(t, tt.expectedAmt, totalCharged, fmt.Sprintf("Test '%s' failed: Total charged %d does not match expected amount %d", tt.name, totalCharged, tt.expectedAmt))

			// Ensure that no chargeable events exceed the total retainer amount
			assert.True(t, totalCharged <= tt.expectedAmt, "Total charged exceeds the expected payout amount")

			// Ensure that all returned events are chargeable events
			for _, actualEvent := range actualEvents {
				found := false
				for _, event := range tt.events {
					if event.ID == actualEvent.ID && event.Charge {
						found = true
						break
					}
				}
				assert.True(t, found, fmt.Sprintf("Non-chargeable event %s found in result set", actualEvent.ID))
			}

			// Print detailed information for the test case
			t.Logf("Test '%s' passed: Calculated Payout: %d, Total Charged: %d, Event IDs Used: %v", tt.name, actualAmt, totalCharged, actualEvents)
		})
	}
}
