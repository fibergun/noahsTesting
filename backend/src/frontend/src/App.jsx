import {BrowserRouter, Routes, Route, Link, HashRouter} from "react-router-dom";
import Home from "./pages/home";
import Ping from "./pages/ping";
import Group from "./routers/Group.jsx";
import ProtectedRoute from "./routers/ProtectedRoute.jsx";
import MakeTask from "./pages/MakeTask.jsx";
import GetAllTasks from "./pages/GetAllTasks.jsx";
import GetRandomTask from "./pages/GetRandomTask.jsx";
import {Container, Nav, Navbar} from "react-bootstrap";

function App() {
  return (
    <HashRouter>
        <Navbar bg="light" expand="md">
            <Container>
                <Navbar.Brand as={Link} to="/">Noah's Spel</Navbar.Brand>
                <Navbar.Toggle aria-controls="main-nav" />
                <Navbar.Collapse id="main-nav">
                    <Nav>
                        <Nav.Link as={Link} to="/ping">Ping</Nav.Link>
                        <Nav.Link as={Link} to="/tasks/make">Maak Taak</Nav.Link>
                        <Nav.Link as={Link} to="/tasks/random">Krijg Taak</Nav.Link>
                        <Nav.Link as={Link} to="/tasks/list">Alle Taken</Nav.Link>
                    </Nav>
                </Navbar.Collapse>
            </Container>
        </Navbar>

        <div className="page-container">
      <Routes>
        <Route
          path="/"
          element={
            <ProtectedRoute>
              <Home />
            </ProtectedRoute>
          }
        />
        <Route
          path="/ping"
          element={
            <ProtectedRoute>
              <Ping />
            </ProtectedRoute>
          }
        />
        <Route
          path="/tasks/make"
          element={
            <ProtectedRoute>
              <MakeTask />
            </ProtectedRoute>
          }
        />
        <Route
          path="/tasks/random"
          element={
            <ProtectedRoute>
              <GetRandomTask />
            </ProtectedRoute>
          }
        />
        <Route
          path="/tasks/list"
          element={
            <ProtectedRoute>
              <GetAllTasks />
            </ProtectedRoute>
          }
        />
        <Route path="/:group/login" element={<Group />} />
      </Routes>
        </div>
    </HashRouter>
  );
}

export default App;
