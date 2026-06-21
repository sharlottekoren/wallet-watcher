export interface Transaction {
  id: string;
  user_id: string;
  category_id: string;
  amount: number; // Stored as pence
  transaction_type: 'income' | 'expense';
  description: string;
  created_at: string; // ISO date string from Go's time.Time
}