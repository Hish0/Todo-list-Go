import React, { useState, useEffect } from 'react';
import './App.css';
import TaskList from './components/TaskList'; 
import LoginForm from './components/LoginForm';
import RegisterForm from './components/RegisterForm';
import axios from 'axios'; // Import axios

function App() {
  const [isLoggedIn, setIsLoggedIn] = useState(false);
  const [user, setUser] = useState(null);
  const [tasks, setTasks] = useState([]); // State for tasks

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
      setTasks(response.data.tasks); // Update tasks state
    } catch (error) {
      console.error("Error fetching tasks:", error);
      // Handle error (e.g., show error message)
    }
  };

  // Handle task toggling
  const handleTaskToggle = (taskId) => {
    setTasks((prevTasks) => 
      prevTasks.map((task) => 
        task.id === taskId ? { ...task, completed: !task.completed } : task
      )
    );
  };

  // Handle task deletion
  const handleTaskDelete = (taskId) => {
    setTasks((prevTasks) => prevTasks.filter((task) => task.id !== taskId));
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
          <TaskList tasks={tasks} onTaskToggle={handleTaskToggle} onTaskDelete={handleTaskDelete} /> {/* Pass tasks and handlers */}
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

