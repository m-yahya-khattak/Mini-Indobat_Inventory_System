'use client';

import { useState, useEffect } from 'react';
import { productApi, orderApi, Product, CreateOrderRequest } from '@/lib/api';
import { calculateEstimatedPrice, formatCurrency } from '@/lib/utils';
import { toast } from 'react-toastify';

export default function OrderForm() {
  const [products, setProducts] = useState<Product[]>([]);
  const [selectedProductId, setSelectedProductId] = useState<number | ''>('');
  const [quantity, setQuantity] = useState<number>(1);
  const [discountPercent, setDiscountPercent] = useState<number>(0);
  // Separate display values for better UX (allow empty while typing)
  const [quantityDisplay, setQuantityDisplay] = useState<string>('1');
  const [discountDisplay, setDiscountDisplay] = useState<string>('0');
  const [loading, setLoading] = useState(false);
  const [fetchingProducts, setFetchingProducts] = useState(true);

  // Fetch products on mount
  useEffect(() => {
    const fetchProducts = async () => {
      try {
        setFetchingProducts(true);
        const data = await productApi.getAll();
        setProducts(data);
      } catch (err: any) {
        toast.error(
          <div className="flex items-start gap-3">
            <div className="flex-shrink-0">
              <svg className="w-5 h-5 text-red-600" fill="currentColor" viewBox="0 0 20 20">
                <path fillRule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zM8.707 7.293a1 1 0 00-1.414 1.414L8.586 10l-1.293 1.293a1 1 0 101.414 1.414L10 11.414l1.293 1.293a1 1 0 001.414-1.414L11.414 10l1.293-1.293a1 1 0 00-1.414-1.414L10 8.586 8.707 7.293z" clipRule="evenodd" />
              </svg>
            </div>
            <div className="flex-1">
              <p className="font-semibold text-slate-900">Failed to Load Products</p>
              <p className="text-sm text-slate-600 mt-0.5">{err.response?.data?.error || 'Failed to load products'}</p>
            </div>
          </div>,
          { icon: false }
        );
        console.error('Error fetching products:', err);
      } finally {
        setFetchingProducts(false);
      }
    };

    fetchProducts();
  }, []);

  // Get selected product
  const selectedProduct = products.find((p) => p.id === selectedProductId);

  // Check if order is valid (quantity doesn't exceed stock)
  const isOrderValid = selectedProduct ? quantity <= selectedProduct.stock && quantity > 0 : false;
  const exceedsStock = selectedProduct ? quantity > selectedProduct.stock : false;

  // Calculate estimated price in real-time
  const estimatedPrice = selectedProduct
    ? calculateEstimatedPrice(
        selectedProduct.price,
        quantity,
        discountPercent
      )
    : 0;

  // Handle form submission
  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();

    if (!selectedProductId) {
      toast.error(
        <div className="flex items-center gap-2">
          <svg className="w-5 h-5 text-red-600 flex-shrink-0" fill="currentColor" viewBox="0 0 20 20">
            <path fillRule="evenodd" d="M18 10a8 8 0 11-16 0 8 8 0 0116 0zm-7 4a1 1 0 11-2 0 1 1 0 012 0zm-1-9a1 1 0 00-1 1v4a1 1 0 102 0V6a1 1 0 00-1-1z" clipRule="evenodd" />
          </svg>
          <span className="font-medium text-slate-900">Please select a product</span>
        </div>,
        { icon: false }
      );
      return;
    }

    if (quantity <= 0) {
      toast.error(
        <div className="flex items-center gap-2">
          <svg className="w-5 h-5 text-red-600 flex-shrink-0" fill="currentColor" viewBox="0 0 20 20">
            <path fillRule="evenodd" d="M18 10a8 8 0 11-16 0 8 8 0 0116 0zm-7 4a1 1 0 11-2 0 1 1 0 012 0zm-1-9a1 1 0 00-1 1v4a1 1 0 102 0V6a1 1 0 00-1-1z" clipRule="evenodd" />
          </svg>
          <span className="font-medium text-slate-900">Quantity must be greater than 0</span>
        </div>,
        { icon: false }
      );
      return;
    }

    if (discountPercent < 0 || discountPercent > 100) {
      toast.error(
        <div className="flex items-center gap-2">
          <svg className="w-5 h-5 text-red-600 flex-shrink-0" fill="currentColor" viewBox="0 0 20 20">
            <path fillRule="evenodd" d="M18 10a8 8 0 11-16 0 8 8 0 0116 0zm-7 4a1 1 0 11-2 0 1 1 0 012 0zm-1-9a1 1 0 00-1 1v4a1 1 0 102 0V6a1 1 0 00-1-1z" clipRule="evenodd" />
          </svg>
          <span className="font-medium text-slate-900">Discount must be between 0 and 100</span>
        </div>,
        { icon: false }
      );
      return;
    }

    setLoading(true);

    try {
      const orderData: CreateOrderRequest = {
        product_id: Number(selectedProductId),
        quantity,
        discount_percent: discountPercent,
      };

      const response = await orderApi.create(orderData);
      
      // Show success toast
      toast.success(
        <div className="flex items-start gap-3">
          <div className="flex-shrink-0">
            <svg className="w-6 h-6 text-emerald-600" fill="currentColor" viewBox="0 0 20 20">
              <path fillRule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.707-9.293a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z" clipRule="evenodd" />
            </svg>
          </div>
          <div className="flex-1">
            <p className="font-semibold text-slate-900">Order Successful!</p>
            <p className="text-sm text-slate-600 mt-0.5">
              Total: <span className="font-bold text-emerald-600">{formatCurrency(response.total_price)}</span>
            </p>
          </div>
        </div>,
        {
          icon: false, // We're using custom icon
        }
      );

      // Reset form completely to prevent accidental duplicate orders
      setSelectedProductId('');
      setQuantity(1);
      setDiscountPercent(0);
      setQuantityDisplay('1');
      setDiscountDisplay('0');

      // Refresh products list (trigger parent refresh)
      window.dispatchEvent(new Event('productsUpdated'));
    } catch (err: any) {
      const errorMessage =
        err.response?.data?.error || 'Failed to create order';
      toast.error(
        <div className="flex items-start gap-3">
          <div className="flex-shrink-0">
            <svg className="w-5 h-5 text-red-600" fill="currentColor" viewBox="0 0 20 20">
              <path fillRule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zM8.707 7.293a1 1 0 00-1.414 1.414L8.586 10l-1.293 1.293a1 1 0 101.414 1.414L10 11.414l1.293 1.293a1 1 0 001.414-1.414L11.414 10l1.293-1.293a1 1 0 00-1.414-1.414L10 8.586 8.707 7.293z" clipRule="evenodd" />
            </svg>
          </div>
          <div className="flex-1">
            <p className="font-semibold text-slate-900">Order Failed</p>
            <p className="text-sm text-slate-600 mt-0.5">{errorMessage}</p>
          </div>
        </div>,
        { icon: false }
      );
      console.error('Error creating order:', err);
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="w-full">
      <div className="mb-6">
        <h2 className="text-2xl font-bold text-slate-900">Create Order</h2>
        <p className="text-sm text-slate-500 mt-1">Process a new purchase order</p>
      </div>
      
      <form onSubmit={handleSubmit} className="space-y-5">
        {/* Product Selection */}
        <div>
          <label
            htmlFor="product"
            className="block text-sm font-semibold text-slate-900 mb-2"
          >
            Select Drug
          </label>
          <select
            id="product"
            value={selectedProductId}
            onChange={(e) => setSelectedProductId(e.target.value ? Number(e.target.value) : '')}
            className="w-full px-4 py-3 border-2 border-slate-200 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500 bg-white text-slate-900 font-medium transition-colors disabled:bg-slate-50 disabled:text-slate-500"
            disabled={fetchingProducts || loading}
            required
          >
            <option value="">-- Select a product --</option>
            {products.map((product) => (
              <option key={product.id} value={product.id}>
                {product.name} (Stock: {product.stock}, Price: {formatCurrency(product.price)})
              </option>
            ))}
          </select>
        </div>

        {/* Quantity Input */}
        <div>
          <div className="flex justify-between items-center mb-2">
            <label
              htmlFor="quantity"
              className="block text-sm font-semibold text-slate-900"
            >
              Quantity
            </label>
            {selectedProduct && (
              <span className="text-xs text-slate-500 font-medium">
                Max: <span className="font-bold text-slate-700">{selectedProduct.stock}</span> available
              </span>
            )}
          </div>
          <input
            type="number"
            id="quantity"
            min="1"
            max={selectedProduct?.stock || undefined}
            value={quantityDisplay}
            onChange={(e) => {
              const value = e.target.value;
              setQuantityDisplay(value);
              // Update actual value if valid, but don't restrict typing
              const num = parseInt(value, 10);
              if (!isNaN(num) && num >= 1) {
                setQuantity(num);
              }
            }}
            onBlur={(e) => {
              // Validate and normalize on blur
              const num = parseInt(quantityDisplay, 10);
              if (isNaN(num) || num < 1) {
                setQuantity(1);
                setQuantityDisplay('1');
              } else if (selectedProduct && num > selectedProduct.stock) {
                setQuantity(selectedProduct.stock);
                setQuantityDisplay(selectedProduct.stock.toString());
              } else {
                setQuantity(num);
                setQuantityDisplay(num.toString());
              }
            }}
            onFocus={(e) => {
              // Select all text on focus for easy replacement
              e.target.select();
            }}
            className={`w-full px-4 py-3 border-2 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500 text-slate-900 font-medium transition-colors disabled:bg-slate-50 disabled:text-slate-500 ${
              exceedsStock
                ? 'border-red-300 bg-red-50 focus:border-red-500 focus:ring-red-500'
                : 'border-slate-200'
            }`}
            disabled={!selectedProductId || loading}
            required
          />
          {exceedsStock && (
            <p className="mt-1.5 text-sm text-red-600 font-medium flex items-center gap-1.5">
              <svg className="w-4 h-4" fill="currentColor" viewBox="0 0 20 20">
                <path fillRule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zM8.707 7.293a1 1 0 00-1.414 1.414L8.586 10l-1.293 1.293a1 1 0 101.414 1.414L10 11.414l1.293 1.293a1 1 0 001.414-1.414L11.414 10l1.293-1.293a1 1 0 00-1.414-1.414L10 8.586 8.707 7.293z" clipRule="evenodd" />
              </svg>
              Quantity exceeds available stock ({selectedProduct?.stock} units)
            </p>
          )}
        </div>

        {/* Discount Input */}
        <div>
          <label
            htmlFor="discount"
            className="block text-sm font-semibold text-slate-900 mb-2"
          >
            Discount (%)
          </label>
          <input
            type="number"
            id="discount"
            min="0"
            max="100"
            step="0.01"
            value={discountDisplay}
            onChange={(e) => {
              const value = e.target.value;
              setDiscountDisplay(value);
              // Update actual value if valid, but don't restrict typing
              const num = parseFloat(value);
              if (!isNaN(num)) {
                if (num >= 0 && num <= 100) {
                  setDiscountPercent(num);
                }
              }
            }}
            onBlur={(e) => {
              // Validate and normalize on blur
              const num = parseFloat(discountDisplay);
              if (isNaN(num) || num < 0) {
                setDiscountPercent(0);
                setDiscountDisplay('0');
              } else if (num > 100) {
                setDiscountPercent(100);
                setDiscountDisplay('100');
              } else {
                setDiscountPercent(num);
                // Keep decimal if user typed it, otherwise show as integer
                setDiscountDisplay(num % 1 === 0 ? num.toString() : num.toString());
              }
            }}
            onFocus={(e) => {
              // Select all text on focus for easy replacement
              e.target.select();
            }}
            className="w-full px-4 py-3 border-2 border-slate-200 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500 text-slate-900 font-medium transition-colors disabled:bg-slate-50 disabled:text-slate-500"
            disabled={!selectedProductId || loading}
            required
          />
        </div>

        {/* Estimated Price Display */}
        {selectedProduct && (
          <div className={`bg-gradient-to-r rounded-lg p-5 border-2 transition-all ${
            exceedsStock
              ? 'from-red-50 to-rose-50 border-red-200'
              : 'from-blue-50 to-indigo-50 border-blue-200'
          }`}>
            <div className="flex justify-between items-center mb-2">
              <span className="text-sm font-semibold text-slate-700">
                Estimated Price:
              </span>
              <span className={`text-2xl font-bold transition-colors ${
                exceedsStock ? 'text-red-700' : 'text-blue-700'
              }`}>
                {formatCurrency(estimatedPrice)}
              </span>
            </div>
            {selectedProduct.stock === 0 && (
              <div className="mt-3 pt-3 border-t border-red-200">
                <p className="text-sm text-red-700 font-medium flex items-center gap-2">
                  <svg className="w-4 h-4" fill="currentColor" viewBox="0 0 20 20">
                    <path fillRule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zM8.707 7.293a1 1 0 00-1.414 1.414L8.586 10l-1.293 1.293a1 1 0 101.414 1.414L10 11.414l1.293 1.293a1 1 0 001.414-1.414L11.414 10l1.293-1.293a1 1 0 00-1.414-1.414L10 8.586 8.707 7.293z" clipRule="evenodd" />
                  </svg>
                  Out of stock - Cannot place order
                </p>
              </div>
            )}
            {selectedProduct.stock > 0 && selectedProduct.stock < 10 && (
              <div className="mt-3 pt-3 border-t border-amber-200">
                <p className="text-sm text-amber-700 font-medium flex items-center gap-2">
                  <svg className="w-4 h-4" fill="currentColor" viewBox="0 0 20 20">
                    <path fillRule="evenodd" d="M8.257 3.099c.765-1.36 2.722-1.36 3.486 0l5.58 9.92c.75 1.334-.213 2.98-1.742 2.98H4.42c-1.53 0-2.493-1.646-1.743-2.98l5.58-9.92zM11 13a1 1 0 11-2 0 1 1 0 012 0zm-1-8a1 1 0 00-1 1v3a1 1 0 002 0V6a1 1 0 00-1-1z" clipRule="evenodd" />
                  </svg>
                  Low stock: Only {selectedProduct.stock} unit{selectedProduct.stock !== 1 ? 's' : ''} remaining
                </p>
              </div>
            )}
          </div>
        )}

        {/* Submit Button */}
        <button
          type="submit"
          disabled={!selectedProductId || loading || fetchingProducts || !isOrderValid || (selectedProduct?.stock === 0)}
          className="w-full px-4 py-3 bg-gradient-to-r from-blue-600 to-indigo-600 text-white rounded-lg font-semibold hover:from-blue-700 hover:to-indigo-700 transition-all shadow-lg hover:shadow-xl disabled:from-slate-400 disabled:to-slate-500 disabled:cursor-not-allowed disabled:shadow-none transform hover:scale-[1.02] disabled:transform-none"
        >
          {loading ? (
            <span className="flex items-center justify-center gap-2">
              <svg className="w-5 h-5 animate-spin" fill="none" viewBox="0 0 24 24">
                <circle className="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" strokeWidth="4"></circle>
                <path className="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
              </svg>
              Processing...
            </span>
          ) : selectedProduct?.stock === 0 ? (
            'Out of Stock'
          ) : !isOrderValid ? (
            'Invalid Quantity'
          ) : (
            'Submit Order'
          )}
        </button>
      </form>

    </div>
  );
}

