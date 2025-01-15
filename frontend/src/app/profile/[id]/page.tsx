"use client";

import React, { useState } from "react";
import Users from "@/app/lib/users/Users";
import { useAuth } from "@/app/context/AuthContext";
import { useRouter } from "next/navigation";


const UserProfilePage = () => {
    const [showOrders, setShowOrders] = useState(false);
    const {isAuthenticated, user} = useAuth();
    const router = useRouter()

    
    const toggleOrderHistory = () => {
        router.push(`${user?.userid}/orderhistory`);
    };

    return (
        <div className="min-h-screen flex items-center justify-center bg-gray-100">
            {isAuthenticated ? (
                <div className="w-full max-w-md p-8 space-y-6 bg-white rounded-lg shadow-md">
                <h1 className="text-2xl font-bold text-center mb-4 text-primary">User Profile</h1>
                
                
                <div className="space-y-4">
                    <p className="text-lg"><strong>Username:</strong> {user?.name}{user?.surname}</p>
                    <p className="text-lg"><strong>Email:</strong> {user?.email}</p>
                    <p className={`text-lg font-semibold ${user?.status ? "text-green-600" : "text-red-600"}`}>
                        <strong>Status:</strong> {user?.status}
                    </p>
                </div>

        
                <button
                    className="w-full bg-primary text-textColor py-2 px-4 rounded hover:bg-blue-600 transition-all"
                    onClick={toggleOrderHistory}
                >
                    {showOrders ? "Hide Order History" : "Show Order History"}
                </button>
            </div>
            ): (
                <h1>Not Found</h1>
            )}
        </div>
    );
}

export default UserProfilePage;