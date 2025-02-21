"use client";

import { useEffect, useState } from "react";

type Task = {
    id: string;
    title: string;
    description: string;
    assigned_to: string;
};

export default function Dashboard() {
    const [tasks, setTasks] = useState<Task[]>([]);

    useEffect(() => {
        const fetchTasks = async () => {
            const token = localStorage.getItem("token");
            const response = await fetch("http://localhost:8080/api/tasks", {
                headers: {
                    Authorization: `Bearer ${token}`,
                },
            });
            const data = await response.json();
            if (response.ok) {
                setTasks(data.tasks);
            }
        };
        fetchTasks();
    }, []);

    return (
        <div className="min-h-screen bg-gray-100 p-8">
            <h1 className="text-2xl font-bold mb-6">Task Dashboard</h1>
            <div className="grid grid-cols-1 gap-4">
                {tasks.map((task) => (
                    <div key={task.id} className="bg-white p-4 rounded-lg shadow-md">
                        <h2 className="text-xl font-semibold">{task.title}</h2>
                        <p className="text-gray-600">{task.description}</p>
                        <p className="text-sm text-gray-500">Assigned to: {task.assigned_to}</p>
                    </div>
                ))}
            </div>
        </div>
    );
}