"use client";

import { useState } from "react";

type Message = {
    text: string;
    sender: "user" | "ai";
};

export default function Chat() {
    const [messages, setMessages] = useState<Message[]>([]);
    const [input, setInput] = useState("");

    const handleSend = async () => {
        if (input.trim()) {
            const userMessage: Message = { text: input, sender: "user" };
            setMessages([...messages, userMessage]);
            setInput("");

            // Call the backend API for AI suggestions
            const token = localStorage.getItem("token");
            const response = await fetch("http://localhost:8080/api/suggest", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: `Bearer ${token}`,
                },
                body: JSON.stringify({ prompt: input }),
            });
            const data = await response.json();
            if (response.ok) {
                const aiMessage: Message = { text: data.suggestion, sender: "ai" };
                setMessages((prev) => [...prev, aiMessage]);
            } else {
                alert("Failed to get AI suggestion");
            }
        }
    };

    return (
        <div className="min-h-screen bg-gray-100 p-8">
            <h1 className="text-2xl font-bold mb-6">AI Chat</h1>
            <div className="bg-white p-6 rounded-lg shadow-md">
                <div className="h-64 overflow-y-auto mb-4">
                    {messages.map((msg, index) => (
                        <div key={index} className={`mb-2 ${msg.sender === "user" ? "text-right" : "text-left"}`}>
                            <span className={`inline-block p-2 rounded-lg ${msg.sender === "user" ? "bg-blue-500 text-white" : "bg-gray-200"}`}>
                                {msg.text}
                            </span>
                        </div>
                    ))}
                </div>
                <div className="flex">
                    <input
                        type="text"
                        value={input}
                        onChange={(e) => setInput(e.target.value)}
                        className="flex-1 p-2 border border-gray-300 rounded-lg"
                        placeholder="Ask AI for suggestions"
                    />
                    <button
                        onClick={handleSend}
                        className="ml-2 bg-blue-500 text-white p-2 rounded-lg hover:bg-blue-600"
                    >
                        Send
                    </button>
                </div>
            </div>
        </div>
    );
}