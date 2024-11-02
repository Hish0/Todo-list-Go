import React from 'react';
import axios from 'axios'; // Import axios

const TaskItem = ({ task, onTaskToggle, onTaskDelete }) => {
  const handleToggle = async () => {
    try {
      await axios.put(`http://localhost:8080/tasks/${task.id}`, {
        completed: !task.completed
      });
      onTaskToggle(task.id);
    } catch (error) {
      console.error("Error toggling task:", error);
    }
  };

  const handleDelete = async () => {
    try {
      await axios.delete(`http://localhost:8080/tasks/${task.id}`);
      onTaskDelete(task.id);
    } catch (error) {
      console.error("Error deleting task:", error);
    }
  };

  return (
    <li className="task-item">
      <input
        type="checkbox"
        checked={task.completed}
        onChange={handleToggle}
      />
      <span className={task.completed ? 'completed' : ''}>{task.title}</span>
      <button onClick={handleDelete}>Delete</button>
    </li>
  );
};

export default TaskItem;

