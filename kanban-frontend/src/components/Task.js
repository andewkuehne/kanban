import React from "react";

function Task({ task }) {
  return (
    <div className="task">
      <h3>{task.title}</h3>
      <p>{task.description}</p>
      <span>{task.status}</span>
    </div>
  );
}

export default Task;
