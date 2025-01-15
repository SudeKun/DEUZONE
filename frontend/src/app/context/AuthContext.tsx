"use client"
import React, { createContext, useState, useContext, ReactNode } from "react";
import { useRouter } from "next/navigation";
import Users from "../lib/users/Users";
import axios from "axios";


interface User {
    userid: string,
    email: string,
    password: string,
    status: string,
    name: string,
    surname: string
}

interface AuthContextType {
    isAuthenticated: boolean;
    user: User | null;
    login: (email: string, password: string) => void;
    logout: () => void;
}

const AuthContext = createContext<AuthContextType | undefined>(undefined);

export const AuthProvider = ({ children }: { children: ReactNode }) => {
    const [isAuthenticated, setIsAuthenticated] = useState(false);
    const [user, setUser] = useState<User | null>(null);
    const router = useRouter();

    const login = (email: string, password: string) => {

        const theuser = Users.find(u => u.email === email && u.password === password);

        if (theuser){
            setIsAuthenticated(true);
            setUser(theuser);
            localStorage.setItem("user", JSON.stringify(theuser));
            localStorage.setItem("isAuthenticated", "true");
            router.push("/");
        }
        else {
            alert("User Cannot Found!");
        }
            
    };

    const logout = () => {
        setUser(null);
        setIsAuthenticated(false);
        localStorage.removeItem("user");
        localStorage.removeItem("isAuthenticated");
        router.push("/");
    };

    return (
        <AuthContext.Provider value={{ isAuthenticated, user, login, logout }}>
            {children}
        </AuthContext.Provider>
    );
};

export const useAuth = () => {
    const context = useContext(AuthContext);
    if (!context) {
        throw new Error("useAuth must be used within an AuthProvider");
    }
    return context;
};
