import React from "react";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import Login from "./Page/login.jsx";
import InfoPage from "./Page/InfoPage.jsx";

function App() {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<Login />} />
        <Route path="/login" element={<Login />} />
        <Route path="/InfoPage" element={<InfoPage />} />
      </Routes>
    </Router>
  );
}

export default App;