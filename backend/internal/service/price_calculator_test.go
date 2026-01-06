package service

import (
	"math"
	"testing"
)

func TestCalculateOrderPrice(t *testing.T) {
	tests := []struct {
		name            string
		price           float64
		quantity        int
		discountPercent float64
		expected        float64
	}{
		{
			name:            "Normal order without discount",
			price:           5000,
			quantity:        5,
			discountPercent: 0,
			expected:        25000,
		},
		{
			name:            "Order with 10% discount",
			price:           5000,
			quantity:        5,
			discountPercent: 10,
			expected:        22500, // 25000 - (25000 * 0.1) = 22500
		},
		{
			name:            "Order with 50% discount",
			price:           10000,
			quantity:        2,
			discountPercent: 50,
			expected:        10000, // 20000 - (20000 * 0.5) = 10000
		},
		{
			name:            "Order with 100% discount (free)",
			price:           5000,
			quantity:        3,
			discountPercent: 100,
			expected:        0, // 15000 - (15000 * 1.0) = 0
		},
		{
			name:            "Order with decimal discount",
			price:           10000,
			quantity:        1,
			discountPercent: 12.5,
			expected:        8750, // 10000 - (10000 * 0.125) = 8750
		},
		{
			name:            "Single item with small discount",
			price:           7500,
			quantity:        1,
			discountPercent: 5,
			expected:        7125, // 7500 - (7500 * 0.05) = 7125
		},
		{
			name:            "Large quantity order",
			price:           1000,
			quantity:        100,
			discountPercent: 15,
			expected:        85000, // 100000 - (100000 * 0.15) = 85000
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CalculateOrderPrice(tt.price, tt.quantity, tt.discountPercent)
			
			// Use tolerance for floating point comparison
			if math.Abs(result-tt.expected) > 0.01 {
				t.Errorf("CalculateOrderPrice(%v, %d, %v) = %v, want %v",
					tt.price, tt.quantity, tt.discountPercent, result, tt.expected)
			}
		})
	}
}

func TestCalculateOrderPrice_EdgeCases(t *testing.T) {
	tests := []struct {
		name            string
		price           float64
		quantity        int
		discountPercent float64
		expected        float64
	}{
		{
			name:            "Zero price",
			price:           0,
			quantity:        5,
			discountPercent: 10,
			expected:        0,
		},
		{
			name:            "Zero quantity",
			price:           5000,
			quantity:        0,
			discountPercent: 10,
			expected:        0,
		},
		{
			name:            "Negative price should result in zero",
			price:           -1000,
			quantity:        5,
			discountPercent: 10,
			expected:        0, // Should not go negative
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CalculateOrderPrice(tt.price, tt.quantity, tt.discountPercent)
			
			if result != tt.expected {
				t.Errorf("CalculateOrderPrice(%v, %d, %v) = %v, want %v",
					tt.price, tt.quantity, tt.discountPercent, result, tt.expected)
			}
		})
	}
}

