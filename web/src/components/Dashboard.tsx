import React from 'react';
import './Dashboard.css';

const Dashboard: React.FC = () => {
  return (
    <div className="dashboard">
      <div className="dashboard-header">
        <h1>Welcome to NestMate</h1>
        <p>Your personal productivity and life management hub</p>
      </div>

      <div className="dashboard-grid">
        <div className="dashboard-card">
          <div className="card-header">
            <h3>💰 Expenses</h3>
          </div>
          <div className="card-content">
            <p>Track and categorize your income and expenses</p>
            <div className="card-stats">
              <span>Coming Soon</span>
            </div>
          </div>
        </div>

        <div className="dashboard-card">
          <div className="card-header">
            <h3>✅ Tasks</h3>
          </div>
          <div className="card-content">
            <p>Organize and track your tasks across different life areas</p>
            <div className="card-stats">
              <span>Coming Soon</span>
            </div>
          </div>
        </div>

        <div className="dashboard-card">
          <div className="card-header">
            <h3>📝 Notes</h3>
          </div>
          <div className="card-content">
            <p>Capture and organize your thoughts and learnings</p>
            <div className="card-stats">
              <span>Coming Soon</span>
            </div>
          </div>
        </div>

        <div className="dashboard-card">
          <div className="card-header">
            <h3>📄 PDF Parser</h3>
          </div>
          <div className="card-content">
            <p>Extract transaction data from bank statement PDFs</p>
            <div className="card-stats">
              <span>Coming Soon</span>
            </div>
          </div>
        </div>
      </div>

      <div className="dashboard-info">
        <div className="info-card">
          <h3>🚀 Getting Started</h3>
          <p>
            NestMate is your unified platform for managing expenses, tasks, and notes.
            Each module is designed to work seamlessly together while maintaining
            clean separation of concerns.
          </p>
          <ul>
            <li>Track expenses with automatic categorization</li>
            <li>Manage tasks with priorities and labels</li>
            <li>Organize notes with tags and search</li>
            <li>Import bank statements via PDF parsing</li>
          </ul>
        </div>
      </div>
    </div>
  );
};

export default Dashboard;