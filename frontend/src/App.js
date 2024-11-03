import React, { useState, useEffect } from 'react';
import './App.css';
import TaskList from './components/TaskList';
import LoginForm from './components/LoginForm';
import RegisterForm from './components/RegisterForm';
import TaskForm from './components/TaskForm';
import axios from 'axios'; // Import axios
import { jwtDecode } from 'jwt-decode';


function App() {
  const [isLoggedIn, setIsLoggedIn] = useState(false);
  const [user,setUser] = useState(null);
  const [tasks, setTasks] = useState([]); // State for tasks

  // // Check for token on initial load
  // useEffect(() => {
  //   const token = localStorage.getItem('token');
  //   if (token) {
  //     setIsLoggedIn(true);
  //     // Optionally fetch user data here if needed
  //   } else {
  //     // Redirect to login
  //     window.location.href = 'http://localhost:8080/auth';
  //   }
  // }, []);

  // Check for token on initial load
  useEffect(() => {
    const token = localStorage.getItem('token');
    if (token) {
      try {
        const decodedToken = jwtDecode(token);
        const isExpired = decodedToken.exp * 1000 < Date.now(); // Check if token is expired
        if (!isExpired) {
          setIsLoggedIn(true);
          // Optionally fetch user data here if needed
        } else {
          // Token is expired, clear it and redirect to login
          localStorage.removeItem('token');
          window.location.href = 'http://localhost:3000';
        }
      } catch (error) {
        console.error('Failed to decode token:', error);
        // Handle invalid token
        localStorage.removeItem('token');
        window.location.href = 'http://localhost:3000';
      }
    } else {
      // Redirect to login if no token
      window.location.href = 'http://localhost:3000';
    }
  }, []);

  // Handle login
  const handleLogin = (userData) => {
    // setUser(userData);
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
      const response = await axios.get('http://localhost:8080/tasks/', {
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
          <TaskForm onTaskSubmit={(newTask) => setTasks((prevTasks) => [...prevTasks, newTask])} />
          <TaskList tasks={tasks} setTasks={setTasks} /> 
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

