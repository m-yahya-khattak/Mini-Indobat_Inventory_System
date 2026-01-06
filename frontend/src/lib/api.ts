import axios from 'axios';

const API_URL = process.env.NEXT_PUBLIC_API_URL || 'http://localhost:8080';

const apiClient = axios.create({
  baseURL: API_URL,
  headers: {
    'Content-Type': 'application/json',
  },
});

// Product types
export interface Product {
  id: number;
  name: string;
  stock: number;
  price: number;
  created_at?: string;
  updated_at?: string;
}

export interface CreateProductRequest {
  name: string;
  stock: number;
  price: number;
}

// Order types
export interface CreateOrderRequest {
  product_id: number;
  quantity: number;
  discount_percent: number;
}

export interface OrderResponse {
  order_id: number;
  product_id: number;
  quantity: number;
  total_price: number;
  message: string;
}

// API functions
export const productApi = {
  getAll: async (): Promise<Product[]> => {
    const response = await apiClient.get<Product[]>('/products');
    return response.data;
  },

  create: async (data: CreateProductRequest): Promise<Product> => {
    const response = await apiClient.post<Product>('/products', data);
    return response.data;
  },
};

export const orderApi = {
  create: async (data: CreateOrderRequest): Promise<OrderResponse> => {
    const response = await apiClient.post<OrderResponse>('/order', data);
    return response.data;
  },
};

export default apiClient;

