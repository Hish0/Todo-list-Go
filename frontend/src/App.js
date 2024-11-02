import React, { useState, useEffect } from 'react';
import './App.css';
import TaskList from './components/TaskList';
import LoginForm from './components/LoginForm';
import RegisterForm from './components/RegisterForm';
import axios from 'axios'; // Import axios

function App() {
  const [isLoggedIn, setIsLoggedIn] = useState(false);
  const [user, setUser] = useState(null);

  // Handle login
  const handleLogin = (userData) => {
    setUser(userData);
    setIsLoggedIn(true);
  };

  // Handle logout
  const handleLogout = () => {
    setUser(null);
    setIsLoggedIn(false);
    // Clear the token from local storage
    localStorage.removeItem('token'); 
  };

  // Fetch tasks
  const fetchTasks = async () => {
    try {
      const response = await axios.get('/tasks', {
        headers: {
          Authorization: `Bearer ${localStorage.getItem('token')}` // Use token
        }
      });
      // ... (Handle the fetched tasks) ...
    } catch (error) {
      console.error("Error fetching tasks:", error);
      // Handle error (e.g., show error message)
    }
  };

  useEffect(() => {
    if (isLoggedIn) {
      fetchTasks();
    }
  }, [isLoggedIn]);

  return (
    <div className="App">
      <h1>Todo List App</h1>
      {isLoggedIn ? (
        <>
          <TaskList /> {/*  Will implement TaskList component */}
          <button onClick={handleLogout}>Logout</button>
        </>
      ) : (
        <>
          <LoginForm onLogin={handleLogin} />
          <RegisterForm />
        </>
      )}
    </div>
  );
}

export default App;

