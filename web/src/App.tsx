import React from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import './App.css';
import Dashboard from './components/Dashboard';
import Layout from './components/Layout';

function App() {
  return (
    <Router>
      <div className="App">
        <Layout>
          <Routes>
            <Route path="/" element={<Dashboard />} />
            <Route path="/expenses" element={<div>Expenses Module (Coming Soon)</div>} />
            <Route path="/tasks" element={<div>Tasks Module (Coming Soon)</div>} />
            <Route path="/notes" element={<div>Notes Module (Coming Soon)</div>} />
            <Route path="/pdf-parser" element={<div>PDF Parser Module (Coming Soon)</div>} />
          </Routes>
        </Layout>
      </div>
    </Router>
  );
}

export default App;
