import { BrowserRouter, Routes, Route, Link } from 'react-router-dom';
import Home from './pages/home';
import Ping from './pages/ping';
import Group from './routers/Group.jsx'
import ProtectedRoute from './routers/ProtectedRoute.jsx'
import MakeTask from "./pages/MakeTask.jsx";
import GetAllTasks from "./pages/GetAllTasks.jsx";

function App() {
  return (
      <BrowserRouter>
        <nav>
            <Link to="/">Home</Link>
            {' | '}
            <Link to="/ping">Ping</Link>
            {' | '}
            <Link to="/tasks/make">Make Task</Link>
            {' | '}
            <Link to="/tasks/list">All tasks</Link>
        </nav>

          <Routes>
              <Route path="/" element={
                  <ProtectedRoute><Home /></ProtectedRoute>
              } />
              <Route path="/ping" element={
                  <ProtectedRoute><Ping /></ProtectedRoute>
              } />
              <Route path="/tasks/make" element={
                  <ProtectedRoute><MakeTask /></ProtectedRoute>
              } />
              <Route path="/tasks/list" element={
                  <ProtectedRoute><GetAllTasks /></ProtectedRoute>
              } />
              <Route path="/:group/login" element={<Group />} />
          </Routes>
      </BrowserRouter>
  );
}

export default App;