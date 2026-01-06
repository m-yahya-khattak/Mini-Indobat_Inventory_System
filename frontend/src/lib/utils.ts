/**
 * Calculate estimated price for an order
 * Formula: (Price × Quantity) - (Discount %)
 */
export function calculateEstimatedPrice(
  price: number,
  quantity: number,
  discountPercent: number
): number {
  if (price <= 0 || quantity <= 0) return 0;
  
  const subtotal = price * quantity;
  const discountAmount = subtotal * (discountPercent / 100);
  const total = subtotal - discountAmount;
  
  return Math.max(0, total); // Ensure non-negative
}

/**
 * Format number as currency (Indonesian Rupiah)
 */
export function formatCurrency(amount: number): string {
  return new Intl.NumberFormat('id-ID', {
    style: 'currency',
    currency: 'IDR',
    minimumFractionDigits: 0,
  }).format(amount);
}

/**
 * Format number with thousand separators
 */
export function formatNumber(num: number): string {
  return new Intl.NumberFormat('id-ID').format(num);
}

