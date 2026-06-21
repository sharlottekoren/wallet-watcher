import type { Transaction } from '../types';

const API_BASE_URL = 'http://localhost:8080';

export const api = {
  getTransactions: async (): Promise<Transaction[]> => {
    const response = await fetch(`${API_BASE_URL}/transactions`);
    if (!response.ok) {
      throw new Error('Network response was not ok');
    }
    return response.json();
  },

  createTransaction: async (transaction: Omit<Transaction, 'id' | 'user_id' | 'created_at'>): Promise<Transaction> => {
    const response = await fetch(`${API_BASE_URL}/transactions`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      // JSON.stringify turns our TypeScript object into a JSON string for the network
      body: JSON.stringify(transaction)
    });

    if (!response.ok) {
      throw new Error('Failed to create transaction');
    }

    // This satisfies TypeScript's requirement to return a value!
    return response.json();
  }
};