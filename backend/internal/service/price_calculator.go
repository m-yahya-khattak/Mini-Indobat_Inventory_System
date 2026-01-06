package service

// CalculateOrderPrice calculates the total price for an order.
// Formula: (Price × Quantity) × (1 - Discount Percent / 100)
func CalculateOrderPrice(price float64, quantity int, discountPercent float64) float64 {
	subtotal := price * float64(quantity)
	discountAmount := subtotal * (discountPercent / 100.0)
	totalPrice := subtotal - discountAmount
	
	if totalPrice < 0 {
		return 0
	}
	
	return totalPrice
}

