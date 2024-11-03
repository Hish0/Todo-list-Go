import React, { useState } from 'react';
import axios from 'axios'; // Import axios

const TaskForm = ({ onTaskSubmit }) => {
  const [title, setTitle] = useState('');
  const [description, setDescription] = useState('');
  const [error, setError] = useState(null);

  const handleSubmit = async (event) => {
    event.preventDefault();
    setError(null);

    try {
    console.log(localStorage.getItem('token'));
      const response = await axios.post('http://localhost:8080/tasks/', {
        title,
        description
      }, {
        headers: {
          Authorization: `Bearer ${localStorage.getItem('token')}` // Use token
        },
        timeout: 10000 // timeout set to 10 seconds
      });

      // // Call the onTaskSubmit callback with the new task to update the tasks in App
      // onTaskSubmit(response.data.task);

      // Reset the form
      setTitle('');
      setDescription('');
      
      // Call the onTaskSubmit callback to update tasks state
      onTaskSubmit(response.data.task);
    } catch (error) {
        console.error("Error:", error); // Log the entire error object
        if (error.response) {
            setError(error.response.data.error || 'An unexpected error occurred');
        } else if (error.request) {
            console.log(error.request); // Log the request object for more details
            setError('No response from the server.');
        } else {
            setError('Error setting up the request: ' + error.message);
        }
    }
    
  };

  return (
    <form onSubmit={handleSubmit}>
      <h2>Add New Task</h2>
      {error && <div className="error">{error}</div>}
      <div>
        <label htmlFor="title">Title:</label>
        <input
          type="text"
          id="title"
          value={title}
          onChange={(e) => setTitle(e.target.value)}
          required
        />
      </div>
      <div>
        <label htmlFor="description">Description:</label>
        <textarea
          id="description"
          value={description}
          onChange={(e) => setDescription(e.target.value)}
        />
      </div>
      <button type="submit">Add Task</button>
    </form>
  );
};

export default TaskForm;

