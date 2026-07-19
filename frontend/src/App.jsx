import { BrowserRouter, Routes, Route, Link } from 'react-router-dom';
import Home from './pages/Home';
import Ping from './pages/Ping';

function App() {
  return (
      <BrowserRouter>
        <nav>
          <Link to="/">Home</Link>
          {' | '}
          <Link to="/ping">Ping</Link>
        </nav>

        <Routes>
          <Route path="/" element={<Home />} />
          <Route path="/ping" element={<Ping />} />
        </Routes>
      </BrowserRouter>
  );
}

export default App;