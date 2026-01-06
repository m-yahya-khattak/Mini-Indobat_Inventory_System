'use client';

import { useEffect, useState } from 'react';
import ProductTable from '@/components/ProductTable';
import OrderForm from '@/components/OrderForm';

export default function Home() {
  const [refreshKey, setRefreshKey] = useState(0);

  useEffect(() => {
    const handleProductsUpdated = () => {
      setRefreshKey((prev) => prev + 1);
    };

    window.addEventListener('productsUpdated', handleProductsUpdated);
    return () => {
      window.removeEventListener('productsUpdated', handleProductsUpdated);
    };
  }, []);

  return (
    <div className="min-h-screen bg-gradient-to-br from-slate-50 to-slate-100">
      <div className="container mx-auto px-4 py-8 max-w-7xl">
        <header className="mb-10">
          <h1 className="text-4xl font-bold text-slate-900 mb-2 tracking-tight">
            Mini-Indobat Inventory System
          </h1>
          <p className="text-slate-600 text-lg">
            Manage pharmaceutical inventory and process orders
          </p>
        </header>

        <div className="grid grid-cols-1 lg:grid-cols-2 gap-6">
          {/* Left Column - Product Table (Dashboard) */}
          <div className="bg-white rounded-xl shadow-lg border border-slate-200 p-6">
            <ProductTable key={refreshKey} />
          </div>

          {/* Right Column - Order Form */}
          <div className="bg-white rounded-xl shadow-lg border border-slate-200 p-6">
            <OrderForm />
          </div>
        </div>
      </div>
    </div>
  );
}
