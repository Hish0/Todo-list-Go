import React from 'react';
import TaskItem from './TaskItem';

const TaskList = ({ tasks, setTasks }) => {
  return (
    <ul className="task-list">
      {tasks.map((task) => (
        <TaskItem
          key={task.ID}
          task={task}
          setTasks={setTasks} // Pass `setTasks` down to each TaskItem
        />
      ))}
    </ul>
  );
};

export default TaskList;

