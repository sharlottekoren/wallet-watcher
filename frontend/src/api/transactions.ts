export type Transaction = {
  id: string;
  amount: number;
  category: string;
  description: string;
};

const BASE_URL = import.meta.env.VITE_API_URL;

export async function getTransactions(): Promise<Transaction[]> {
    const response = await fetch(`${BASE_URL}/transactions`);

    if (!response.ok) {
        throw new Error('Failed to fetch transactions');
    }

    return response.json();
}

export async function createTransaction(tx: {
    amount: number;
    category: string;
    description: string;
}): Promise<Transaction> {
    const response = await fetch(`${BASE_URL}/transactions`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(tx),
    });

    if (!response.ok) {
        throw new Error('Failed to create transaction');
    }

    return response.json();
}