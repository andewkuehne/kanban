import React, { useEffect, useState } from "react";
import axios from "axios";
import TaskList from "./components/TaskList";
import NewTaskForm from "./components/NewTaskForm";
import "./style.css"

function App() {
  const [tasks, setTasks] = useState([]);

  useEffect(() => {
    axios.get("/tasks").then((response) => {
      setTasks(response.data);
    });
  }, []);

  const handleAddTask = (task) => {
    axios.post("/tasks", task).then((response) => {
      setTasks([...tasks, response.data]);
    });
  };

  const handleUpdateTask = (id, updates) => {
    axios.put(`/tasks/${id}`, updates).then((response) => {
      setTasks(tasks.map((task) => (task.id === id ? response.data : task)));
    });
  };

  const handleDeleteTask = (id) => {
    axios.delete(`/tasks/${id}`).then(() => {
      setTasks(tasks.filter((task) => task.id !== id));
    });
  };

  return (
    <div className="app">
      <h1>Kanban Board</h1>
      <NewTaskForm onSubmit={handleAddTask} />
      <div className="task-list-container">
        <TaskList tasks={tasks.filter((task) => task.status === "To Do")} onUpdateTask={handleUpdateTask} onDeleteTask={handleDeleteTask} />
        <TaskList tasks={tasks.filter((task) => task.status === "In Progress")} onUpdateTask={handleUpdateTask} onDeleteTask={handleDeleteTask} />
        <TaskList tasks={tasks.filter((task) => task.status === "Done")} onUpdateTask={handleUpdateTask} onDeleteTask={handleDeleteTask} />
      </div>
    </div>
  );
}

export default App;
