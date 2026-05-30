import { useEffect, useState } from "react";
import { getTransactions, createTransaction } from "./api/transactions";
import type { Transaction } from "./api/transactions";

function App() {
  const [transactions, setTransactions] = useState<Transaction[]>([]);

  const [amount, setAmount] = useState("");
  const [description, setDescription] = useState("");
  const [category, setCategory] = useState("");

  // Load transactions on page load
  useEffect(() => {
    async function load() {
      try {
        const data = await getTransactions();
        setTransactions(data);
      } catch (err) {
        console.error("Error loading transactions:", err);
      }
    }

    load();
  }, []);

  // Handle create transaction
  async function handleSubmit(e: React.FormEvent) {
    e.preventDefault();

    try {
      const newTx = await createTransaction({
        amount: parseFloat(amount),
        description,
        category,
      });

      // Add new transaction to top of list
      setTransactions((prev) => [newTx, ...prev]);

      // reset form
      setAmount("");
      setDescription("");
      setCategory("");
    } catch (err) {
      console.error("Error creating transaction:", err);
    }
  }

  return (
    <div style={{ padding: "20px", fontFamily: "Arial" }}>
      <h1>Wallet Watcher</h1>

      {/* CREATE FORM */}
      <form onSubmit={handleSubmit} style={{ marginBottom: "20px" }}>
        <input
          placeholder="amount"
          value={amount}
          onChange={(e) => setAmount(e.target.value)}
        />

        <input
          placeholder="description"
          value={description}
          onChange={(e) => setDescription(e.target.value)}
        />

        <input
          placeholder="category"
          value={category}
          onChange={(e) => setCategory(e.target.value)}
        />

        <button type="submit">Add Transaction</button>
      </form>

      {/* LIST */}
      {transactions.length === 0 ? (
        <p>No transactions yet</p>
      ) : (
        <ul>
          {transactions.map((t) => (
            <li key={t.id}>
              {t.category}: {t.description} - £{t.amount}
            </li>
          ))}
        </ul>
      )}
    </div>
  );
}

export default App;