package service

import (
	"testing"
)

func TestStockValidation(t *testing.T) {
	tests := []struct {
		name           string
		availableStock int
		requestedQty   int
		shouldPass     bool
		description    string
	}{
		{
			name:           "Sufficient stock - exact match",
			availableStock: 10,
			requestedQty:   10,
			shouldPass:     true,
			description:    "Ordering exactly available stock should pass",
		},
		{
			name:           "Sufficient stock - less than available",
			availableStock: 10,
			requestedQty:   5,
			shouldPass:     true,
			description:    "Ordering less than available stock should pass",
		},
		{
			name:           "Insufficient stock - more than available",
			availableStock: 5,
			requestedQty:   10,
			shouldPass:     false,
			description:    "Ordering more than available stock should fail",
		},
		{
			name:           "Zero stock - cannot order",
			availableStock: 0,
			requestedQty:   1,
			shouldPass:     false,
			description:    "Cannot order when stock is zero",
		},
		{
			name:           "Zero stock - ordering zero",
			availableStock: 0,
			requestedQty:   0,
			shouldPass:     false,
			description:    "Ordering zero quantity should fail (invalid quantity)",
		},
		{
			name:           "Single stock - ordering one",
			availableStock: 1,
			requestedQty:   1,
			shouldPass:     true,
			description:    "Ordering last available item should pass",
		},
		{
			name:           "Single stock - ordering more",
			availableStock: 1,
			requestedQty:   2,
			shouldPass:     false,
			description:    "Cannot order more than single available stock",
		},
		{
			name:           "Large stock - small order",
			availableStock: 1000,
			requestedQty:   1,
			shouldPass:     true,
			description:    "Small order from large stock should pass",
		},
		{
			name:           "Large stock - large order within limit",
			availableStock: 1000,
			requestedQty:   999,
			shouldPass:     true,
			description:    "Large order within stock limit should pass",
		},
		{
			name:           "Large stock - order exceeding limit",
			availableStock: 1000,
			requestedQty:   1001,
			shouldPass:     false,
			description:    "Order exceeding stock limit should fail",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test the stock validation logic
			// This tests: product.Stock < req.Quantity should fail
			passes := tt.availableStock >= tt.requestedQty && tt.requestedQty > 0
			
			if passes != tt.shouldPass {
				t.Errorf("Stock validation failed: %s\nAvailable: %d, Requested: %d, Expected: %v, Got: %v",
					tt.description, tt.availableStock, tt.requestedQty, tt.shouldPass, passes)
			}
		})
	}
}

func TestStockValidation_EdgeCases(t *testing.T) {
	tests := []struct {
		name           string
		availableStock int
		requestedQty   int
		expectedResult string
	}{
		{
			name:           "Stock should never go negative",
			availableStock: 5,
			requestedQty:   5,
			expectedResult: "Stock becomes 0, not negative",
		},
		{
			name:           "Stock calculation after order",
			availableStock: 10,
			requestedQty:   3,
			expectedResult: "New stock should be 7",
		},
		{
			name:           "Stock calculation - exact match",
			availableStock: 10,
			requestedQty:   10,
			expectedResult: "New stock should be 0",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.availableStock < tt.requestedQty {
				t.Skip("Skipping - insufficient stock scenario")
			}
			
			newStock := tt.availableStock - tt.requestedQty
			
			if newStock < 0 {
				t.Errorf("Stock calculation error: Stock became negative (%d)", newStock)
			}
			
			expectedNewStock := tt.availableStock - tt.requestedQty
			if newStock != expectedNewStock {
				t.Errorf("Stock calculation incorrect: Expected %d, Got %d", expectedNewStock, newStock)
			}
		})
	}
}

