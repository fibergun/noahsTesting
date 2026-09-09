import React from "react";
import ReactDOM from "react-dom/client";
import App from "./App.jsx";
import { SessionProvider } from "./routers/Session.jsx";
import "./bootstrap.min.css";
import "./App.css";

ReactDOM.createRoot(document.getElementById("root")).render(
  <React.StrictMode>
    <SessionProvider>
      <App />
    </SessionProvider>
  </React.StrictMode>,
);
