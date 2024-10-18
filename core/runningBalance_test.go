package core

import (
	"reflect"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/mattg1243/willow-server/db"
)

// TestRecalcRunningBalances tests the RecalcRunningBalances function with various scenarios, including retainer payments and charges.
func TestRecalcRunningBalances(t *testing.T) {
	// Define a set of test cases
	testCases := []struct {
		name     string            // Name of the test case
		events   []db.GetEventsRow // Input events for the test
		expected []int             // Expected running balances
	}{
		{
			name: "Single retainer payment",
			events: []db.GetEventsRow{
				{ID: uuid.New(), Amount: 100, Charge: false, Date: pgTimestamp("2024-10-08T09:00:00Z")}, // Retainer payment of 100
			},
			expected: []int{100}, // Running balance should be 100
		},
		{
			name: "Single charge event",
			events: []db.GetEventsRow{
				{ID: uuid.New(), Amount: 100, Charge: true, Date: pgTimestamp("2024-10-08T09:00:00Z")}, // Charge of 100
			},
			expected: []int{-100}, // Running balance should be -100
		},
		{
			name: "Multiple retainer payments and charges",
			events: []db.GetEventsRow{
				{ID: uuid.New(), Amount: 100, Charge: false, Date: pgTimestamp("2024-10-08T09:00:00Z")},   // Retainer payment of 100
				{ID: uuid.New(), Amount: 50, Charge: true, Date: pgTimestamp("2024-10-08T10:00:00Z")},    // Charge of 50
				{ID: uuid.New(), Amount: 30, Charge: true, Date: pgTimestamp("2024-10-08T11:00:00Z")},    // Charge of 30
				{ID: uuid.New(), Amount: 70, Charge: false, Date: pgTimestamp("2024-10-08T12:00:00Z")},   // Retainer payment of 70
			},
			expected: []int{100, 50, 20, 90}, // Cumulative running balances: 100 - 50 = 50; 50 - 30 = 20; 20 + 70 = 90
		},
		{
			name: "Multiple charges resulting in negative balance",
			events: []db.GetEventsRow{
				{ID: uuid.New(), Amount: 200, Charge: false, Date: pgTimestamp("2024-10-08T09:00:00Z")},  // Retainer payment of 200
				{ID: uuid.New(), Amount: 100, Charge: true, Date: pgTimestamp("2024-10-08T10:00:00Z")},   // Charge of 100
				{ID: uuid.New(), Amount: 150, Charge: true, Date: pgTimestamp("2024-10-08T11:00:00Z")},   // Charge of 150
				{ID: uuid.New(), Amount: 50, Charge: false, Date: pgTimestamp("2024-10-08T12:00:00Z")},   // Retainer payment of 50
			},
			expected: []int{200, 100, -50, 0}, // Cumulative running balances: 200 - 100 = 100; 100 - 150 = -50; -50 + 50 = 0
		},
		{
			name: "Events inserted out of order",
			events: []db.GetEventsRow{
				{ID: uuid.New(), Amount: 50, Charge: true, Date: pgTimestamp("2024-10-08T12:00:00Z")},    // Charge of 50 (inserted later)
				{ID: uuid.New(), Amount: 100, Charge: false, Date: pgTimestamp("2024-10-08T09:00:00Z")},  // Initial retainer payment of 100
			},
			expected: []int{100, 50}, // Running balance after sorting by date: 100 - 50 = 50
		},
		{
			name:     "No events",
			events:   []db.GetEventsRow{}, // No events
			expected: []int{},             // Running balance should be empty
		},
	}

	// Iterate through each test case
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Call the RecalcRunningBalances function with the test events
			result := RecalcRunningBalances(tc.events)

			// Compare the result with the expected running balances
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Test %s failed: got %v, expected %v", tc.name, result, tc.expected)
			}
		})
	}
}

// pgTimestamp is a helper function to convert a string timestamp into pgtype.Timestamptz for pgx v5.
func pgTimestamp(value string) pgtype.Timestamptz {
	t, err := time.Parse(time.RFC3339, value)
	if err != nil {
		panic("Failed to parse timestamp: " + err.Error())
	}

	// Create a pgtype.Timestamptz instance and set its fields directly
	return pgtype.Timestamptz{
		Time:  t,
		Valid: true,  // Set Valid to true since we're providing a valid timestamp
	}
}
