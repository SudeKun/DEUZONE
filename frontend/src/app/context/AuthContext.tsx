"use client"
import React, { createContext, useState, useContext, ReactNode } from "react";
import { useRouter } from "next/navigation";
import Users from "../lib/users/Users";


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

        const user = Users.find(u => u.email === email && u.password === password);

        if (user) {
            setUser(user);
            setIsAuthenticated(true);
            router.push("/");
        }
        else
            alert("User cannot found!");
    };

    const logout = () => {
        setUser(null);
        setIsAuthenticated(false);
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
