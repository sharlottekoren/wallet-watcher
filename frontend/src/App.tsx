import { useEffect, useState } from 'react';
import { api } from './services/api';
import type { Transaction } from './types';

function App() {
  const [transactions, setTransactions] = useState<Transaction[]>([]);
  const [loading, setLoading] = useState<boolean>(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    api.getTransactions()
      .then((data) => {
        setTransactions(data);
        setLoading(false);
      })
      .catch((err) => {
        setError(err.message);
        setLoading(false);
      });
  }, []);

  if (loading) return <div>Loading your wallet data...</div>;
  if (error) return <div>Error fetching financial data: {error}</div>;

  return (
    <div style={{ padding: '2rem', fontFamily: 'sans-serif' }}>
      <h1>Wallet Watcher Dashboard</h1>
      
      {/* Transaction List Container */}
      <div style={{ marginTop: '2rem' }}>
        <h2>Recent Transactions</h2>
        {transactions.length === 0 ? (
          <p>No transactions recorded yet. Add your first expense!</p>
        ) : (
          <ul>
            {transactions.map((tx) => (
              <li key={tx.id} style={{ margin: '0.5rem 0' }}>
                <strong>{tx.description}</strong> - 
                <span style={{ color: tx.transaction_type === 'income' ? 'green' : 'red', marginLeft: '0.5rem' }}>
                  {tx.transaction_type === 'income' ? '+' : '-'}£{(tx.amount / 100).toFixed(2)}
                </span>
              </li>
            ))}
          </ul>
        )}
      </div>
    </div>
  );
}

export default App;