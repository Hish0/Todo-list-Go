import React from 'react';
import axios from 'axios';

const TaskItem = ({ task, setTasks }) => {
  const handleToggle = async () => {
    try {
      // Send the PUT request to toggle the completed status
      await axios.put(`http://localhost:8080/tasks/${task.ID}`, {
        completed: !task.Completed
      }, {
        headers: {
          Authorization: `Bearer ${localStorage.getItem('token')}` 
        }
      });
      
      // Update the task in the state by toggling the completed status
      setTasks((prevTasks) =>
        prevTasks.map((t) =>
          t.ID === task.ID ? { ...t, Completed: !t.Completed } : t
        )
      );
    } catch (error) {
      console.error("Error toggling task:", error);
    }
  };

  const handleDelete = async () => {
    try {
      // Send the DELETE request to remove the task
      await axios.delete(`http://localhost:8080/tasks/${task.ID}`, {
        headers: {
          Authorization: `Bearer ${localStorage.getItem('token')}` 
        }
      });
      
      // Update the state to remove the task
      setTasks((prevTasks) => prevTasks.filter((t) => t.ID !== task.ID));
    } catch (error) {
      console.error("Error deleting task:", error);
    }
  };

  return (
    <li className="task-item">
      <input
        type="checkbox"
        checked={task.Completed}
        onChange={handleToggle}
      />
      <div>
        <span>{task.Completed ? '✔️' : '❌'}</span>
        <h3>{task.ID}</h3>
        <p>{task.Title}</p> {/* Display the title here */}
        <p>{task.Description}</p> {/* Display the description here */}
        <button onClick={handleDelete}>Delete</button>
      </div>
    </li>
  );
};

export default TaskItem;
