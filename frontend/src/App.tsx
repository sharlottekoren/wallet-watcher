import { useEffect, useState } from 'react'
import reactLogo from './assets/react.svg'
import viteLogo from './assets/vite.svg'
import heroImg from './assets/hero.png'
import { getTransactions } from './api/transactions'
import type { Transaction } from './api/transactions'
import './App.css'

function App() {
  const [count, setCount] = useState(0)
  const [transactions, setTransactions] = useState<Transaction[]>([])
  const [loading, setLoading] = useState(true)

  useEffect(() => {
    async function load() {
      try {
        const data = await getTransactions()
        setTransactions(data)
      } catch (error) {
        console.error('Error fetching transactions:', error)
      } finally {
        setLoading(false)
      }
    }

    load()
  }, [])

  if (loading) return <p>Loading transactions...</p>

  return (
    <div style={{ padding: '20px' }}>
      <h1>Wallet Watcher</h1>

      <ul>
        {transactions.map(tx => (
          <li key={tx.id}>
            <strong>{tx.category}</strong>: {tx.description} - ${tx.amount}
          </li>
        ))}
      </ul>
    </div>
  )
}

export default App
