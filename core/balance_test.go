package core

import (
	"testing"

	"github.com/google/uuid"
	"github.com/mattg1243/willow-server/db" // Replace with your actual db package path
)

// Helper function to create a GetEventRow for testing
func createTestEvent(id string, amount int32, charge bool) db.GetEventsRow {
	return db.GetEventsRow{
		ID:          uuid.MustParse(id),
		ClientID:    uuid.New(),
		EventTypeID: uuid.New(),
		Rate:        0,
		Amount:      amount,
		RunningBalance:  0, // Replace with actual value or leave as 0 if not needed
		Charge:      charge,
	}
}

// TestCalculateBalance tests the CalculateBalance function with various inputs.
func TestCalculateBalance(t *testing.T) {
	tests := []struct {
		name     string
		events   []db.GetEventsRow
		expected int
	}{
		{
			name: "Single positive amount without charge",
			events: []db.GetEventsRow{
				createTestEvent("11111111-1111-1111-1111-111111111111", 100, false),
			},
			expected: 100,
		},
		{
			name: "Single negative amount with charge",
			events: []db.GetEventsRow{
				createTestEvent("22222222-2222-2222-2222-222222222222", 50, true),
			},
			expected: -50,
		},
		{
			name: "Multiple events with mixed charges",
			events: []db.GetEventsRow{
				createTestEvent("33333333-3333-3333-3333-333333333333", 200, false),
				createTestEvent("44444444-4444-4444-4444-444444444444", 100, true),
				createTestEvent("55555555-5555-5555-5555-555555555555", 50, false),
			},
			expected: 150, // 200 - 100 + 50
		},
		{
			name: "No events",
			events: []db.GetEventsRow{},
			expected: 0,
		},
		{
			name: "Multiple charges resulting in negative balance",
			events: []db.GetEventsRow{
				createTestEvent("66666666-6666-6666-6666-666666666666", 50, true),
				createTestEvent("77777777-7777-7777-7777-777777777777", 75, true),
				createTestEvent("88888888-8888-8888-8888-888888888888", 25, true),
			},
			expected: -150, // -50 - 75 - 25
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CalculateBalance(tt.events)
			if result != tt.expected {
				t.Errorf("Test '%s' failed: got %d, want %d", tt.name, result, tt.expected)
			} else {
				t.Logf("Test '%s' passed: got %d", tt.name, result)
			}
		})
	}
}
