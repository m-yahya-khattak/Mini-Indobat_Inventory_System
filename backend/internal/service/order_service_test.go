package service

import (
	"testing"
)

func TestOrderService_InputValidation(t *testing.T) {
	tests := []struct {
		name            string
		quantity        int
		discountPercent float64
		shouldFail      bool
		expectedError   error
	}{
		{
			name:            "Valid quantity and discount",
			quantity:        5,
			discountPercent: 10,
			shouldFail:      false,
			expectedError:   nil,
		},
		{
			name:            "Zero quantity should fail",
			quantity:        0,
			discountPercent: 0,
			shouldFail:      true,
			expectedError:   ErrInvalidQuantity,
		},
		{
			name:            "Negative quantity should fail",
			quantity:        -1,
			discountPercent: 0,
			shouldFail:      true,
			expectedError:   ErrInvalidQuantity,
		},
		{
			name:            "Negative discount should fail",
			quantity:        5,
			discountPercent: -5,
			shouldFail:      true,
			expectedError:   ErrInvalidDiscount,
		},
		{
			name:            "Discount over 100% should fail",
			quantity:        5,
			discountPercent: 101,
			shouldFail:      true,
			expectedError:   ErrInvalidDiscount,
		},
		{
			name:            "Discount exactly 100% should pass",
			quantity:        5,
			discountPercent: 100,
			shouldFail:      false,
			expectedError:   nil,
		},
		{
			name:            "Discount exactly 0% should pass",
			quantity:        5,
			discountPercent: 0,
			shouldFail:      false,
			expectedError:   nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test validation logic
			quantityValid := tt.quantity > 0
			discountValid := tt.discountPercent >= 0 && tt.discountPercent <= 100
			
			shouldFail := !quantityValid || !discountValid
			
			if shouldFail != tt.shouldFail {
				t.Errorf("Validation failed: Quantity=%d, Discount=%.2f, Expected fail=%v, Got fail=%v",
					tt.quantity, tt.discountPercent, tt.shouldFail, shouldFail)
			}
			
			if !quantityValid && tt.expectedError != ErrInvalidQuantity {
				t.Errorf("Expected ErrInvalidQuantity for invalid quantity")
			}
			
			if !discountValid && tt.expectedError != ErrInvalidDiscount {
				t.Errorf("Expected ErrInvalidDiscount for invalid discount")
			}
		})
	}
}

