import React, { useState, useEffect } from 'react';
import axios from 'axios';

const App = () => {
    const [tasks, setTasks] = useState([]); 
    const [taskTitle, setTaskTitle] = useState(''); 
    const [taskDescription, setTaskDescription] = useState(''); 
    const [isEditing, setIsEditing] = useState(false); 
    const [currentTaskId, setCurrentTaskId] = useState(null); 
    
    useEffect(() => {
        fetchTasks(); 
    }, []); 

    // Fetch all tasks from the backend 
    const fetchTasks = async () => {
        try {
            const response = await axios.get('http://localhost:8080/tasks'); 
            setTasks(response.data); 
        } catch (error) {
            console.error('There was an error fetching the tasks:', error); 
        }
    }; 

    // Create a new task 
    const createTask = async () => {
        try {
            const newTask = { title: taskTitle, description: taskDescription, completed: false}; 
            const response = await axios.post('http://localhost:8080/tasks', newTask); 
            setTasks([...tasks, response.data]); 
            setTaskTitle(''); 
            setTaskDescription(''); 
        } catch (error) {
            console.error('Error creating task:', error); 
        }
    }; 

    // Update an existing task 
    const updateTask = async () => {
        try {
            const updatedTask = { title: taskTitle, description: taskDescription, completed: false}; 
            await axios.post(`http://localhost:8080/tasks/${currentTaskId}`, updatedTask); 
            const updatedTasks = tasks.map(task => 
                task.id === currentTaskId ? { ...task, ...updatedTask } : task 
            ); 
            setTasks(updatedTasks); 
            resetForm();
        } catch (error) {
            console.error('Error updating task:', error); 
        }
    }; 

    // Delete a task 
    const deleteTask = async (id) => {
        try {
            await axios.delete(`http://localhost:8080/tasks/${id}`); 
            setTasks(tasks.filter(task => task.id !== id)); 
        } catch (error) {
            console.error('Error deleting task:', error); 
        }
    }; 

    // Set up form for editing a task 
    const editTask = (task) => {
        setTaskTitle(task.title); 
        setTaskDescription(task.description); 
        setIsEditing(true); 
        setCurrentTaskId(task.id); 
    }; 

    // Reset form fields 
    const resetForm = () => {
        setTaskTitle(''); 
        setTaskDescription(''); 
        setIsEditing(false); 
        setCurrentTaskId(null); 
    }; 

    return (
        <div className="App">
          <h1>Task Manager</h1>
    
          {/* Form for adding or updating a task */}
          <div>
            <input
              type="text"
              value={taskTitle}
              onChange={(e) => setTaskTitle(e.target.value)}
              placeholder="Task Title"
            />
            <input
              type="text"
              value={taskDescription}
              onChange={(e) => setTaskDescription(e.target.value)}
              placeholder="Task Description"
            />
            {isEditing ? (
              <button onClick={updateTask}>Update Task</button>
            ) : (
              <button onClick={createTask}>Create Task</button>
            )}
            <button onClick={resetForm}>Cancel</button>
          </div>
    
          {/* Task List */}
          <ul>
            {tasks.map(task => (
              <li key={task.id}>
                <div>
                  <span>{task.title}</span> - <span>{task.description}</span>
                  <button onClick={() => editTask(task)}>Edit</button>
                  <button onClick={() => deleteTask(task.id)}>Delete</button>
                </div>
              </li>
            ))}
          </ul>
        </div>
      );
}; 

export default App;
