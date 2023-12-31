import React from "react";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";

// Components
import Home from "./pages/Home/Home";
import NotFound from "./pages/NotFound/NotFound";
import Navbar from "./components/Navbar/Navbar";

function App() {
  return (
    <div className="min-h-screen bg-slate-50 ">
      <Router>
        <Navbar />
        <Routes>
          <Route path="/" element={<Home />} />
          <Route path="/*" element={<NotFound />} />
        </Routes>
      </Router>
    </div>
  );
}

export default App;
