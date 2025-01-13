"use client";

import React, { useState } from "react";
import { useRouter } from "next/navigation";
import { useAuth } from "../context/AuthContext";
import Users from "../lib/users/Users";

export default function AuthPage() {
    const [isLogin, setIsLogin] = useState(true);
    const [mail, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const [name, setName] = useState('');
    const [surname, setSurname] = useState('');
    const router = useRouter();
    const { login } = useAuth();

    const handleFormSwitchLogin = () => setIsLogin(true);
    const handleFormSwitchRegister = () => setIsLogin(false);

    const handleSubmit = (e: React.FormEvent) => {
        e.preventDefault();

        router.push("/");
    };

    const handleLogin = () => {
        login(mail, password);
    };

    const handleRegister = () =>{
        const user = Users.find(u => u.email === mail);

        if (user)
            alert("Email has already exist!");
        else {
            let i = Users.length + 1;
            Users.push({
                userid: i.toString(),
                email: mail,
                password: password,
                status: "Active",
                name: name,
                surname: surname
            })
        }
    }


    return (
        <div className="min-h-screen flex items-center justify-center bg-gray-100">
            <div className="w-full max-w-md p-8 space-y-6 bg-white rounded-lg shadow-md">
                <div className="w-full flex flex-row justify-between items-center">
                    <button onClick={handleFormSwitchLogin} className={`text-2xl text-center border-b-2 w-1/2 ${isLogin ? "font-bold border-b-primary" : "font-normal"} `}>Login</button>
                    <button onClick={handleFormSwitchRegister} className={`text-2xl text-center border-b-2 w-1/2 ${!isLogin ? "font-bold border-b-primary" : "font-normal"} `}>Register</button>
                </div>

                <form onSubmit={handleSubmit} className="space-y-4">
                    {!isLogin && (
                        <div className="w-full flex flex-row justify-between">
                            <input
                            type="text"
                            placeholder="Name"
                            required
                            className="w-5/12 p-2 border rounded"
                            onChange={(e) => setName(e.target.value)}
                        />

                        <input
                            type="text"
                            placeholder="Surname"
                            required
                            className="w-5/12 p-2 border rounded"
                            onChange={(e) => setSurname(e.target.value)}
                        />
                        </div>
                    )}

                    <input
                        type="email"
                        placeholder="Email"
                        required
                        className="w-full p-2 border rounded"
                        onChange={(e) => setEmail(e.target.value)}
                    />
                    <input
                        type="password"
                        placeholder="Password"
                        required
                        className="w-full p-2 border rounded"
                        onChange={(e) => setPassword(e.target.value)}
                    />

                    <button
                        type="submit"
                        className="w-full p-2 bg-primary text-white rounded hover:bg-blue-600"
                        onClick={isLogin ? handleLogin : handleRegister}
                    >
                        {isLogin ? "Login" : "Register"}
                    </button>
                </form>
            </div>
        </div>
    );
}
