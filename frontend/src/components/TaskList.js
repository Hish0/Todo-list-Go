import React from 'react';
import TaskItem from './TaskItem';

const TaskList = ({ tasks, onTaskToggle, onTaskDelete }) => {
  return (
    <ul className="task-list">
      {tasks.map((task) => (
        <TaskItem
          key={task.id}
          task={task}
          onTaskToggle={onTaskToggle} // Pass handlers to TaskItem
          onTaskDelete={onTaskDelete}
        />
      ))}
    </ul>
  );
};

export default TaskList;

