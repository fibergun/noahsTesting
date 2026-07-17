import { useState } from 'react'

function App() {
  const [status, setStatus] = useState(null)
  const [error, setError] = useState(null)

  const checkPing = async () => {
    setError(null)
    setStatus(null)
    try {
      const res = await fetch('/ping')
      if (!res.ok) throw new Error(`request failed: ${res.status}`)
      const data = await res.json()
      setStatus(data.status)
    } catch (err) {
      setError(err.message)
    }
  }

  return (
    <div>
      <h1>siteTesting</h1>
      <button onClick={checkPing}>Ping backend</button>
      {status && <p>Status: {status}</p>}
      {error && <p>Error: {error}</p>}
    </div>
  )
}

export default App
