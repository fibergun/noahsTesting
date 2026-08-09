import { BrowserRouter, Routes, Route, Link } from 'react-router-dom';
import Home from './pages/home';
import Ping from './pages/ping';
import Login from './pages/login.jsx'
import ProtectedRoute from './routers/ProtectedRoute.jsx'

function App() {
  return (
      <BrowserRouter>
        <nav>
          <Link to="/">Home</Link>
          {' | '}
          <Link to="/ping">Ping</Link>
        </nav>

          <Routes>
              <Route path="/" element={
                  <ProtectedRoute><Home /></ProtectedRoute>
              } />
              <Route path="/ping" element={
                  <ProtectedRoute><Ping /></ProtectedRoute>
              } />
              <Route path="/:group/login" element={<Login />} />
          </Routes>
      </BrowserRouter>
  );
}

export default App;