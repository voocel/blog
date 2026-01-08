import React, { createContext, useContext, useState, useEffect, type ReactNode } from 'react';
import type { User } from '../types';
import { authService } from '../services/authService';

interface AuthContextType {
    user: User | null;
    isLoading: boolean;

    // Modal State
    isAuthModalOpen: boolean;
    setAuthModalOpen: (isOpen: boolean) => void;

    // Auth Actions
    login: (email: string, password?: string) => Promise<boolean>;
    register: (email: string, password: string) => Promise<boolean>;
    logout: () => void;
    updateUser: (user: User) => Promise<void>;
}

const AuthContext = createContext<AuthContextType | undefined>(undefined);

export const AuthProvider: React.FC<{ children: ReactNode }> = ({ children }) => {
    const [user, setUser] = useState<User | null>(null);
    const [isLoading, setIsLoading] = useState(true);
    const [isAuthModalOpen, setAuthModalOpen] = useState(false);

    // Fetch current user on mount
    useEffect(() => {
        const fetchUser = async () => {
            setIsLoading(true);
            try {
                const fetchedUser = await authService.getCurrentUser();
                setUser(fetchedUser);
            } catch (err) {
                console.error('Failed to get current user', err);
            } finally {
                setIsLoading(false);
            }
        };
        fetchUser();
    }, []);

    const login = async (email: string, password?: string) => {
        try {
            const user = await authService.login(email, password);
            if (user) {
                setUser(user);
                return true;
            }
            return false;
        } catch (err) {
            console.error("Login failed", err);
            throw err;
        }
    };

    const register = async (email: string, password: string) => {
        try {
            const user = await authService.register(email, password);
            if (user) {
                setUser(user);
                return true;
            }
            return false;
        } catch (err) {
            console.error("Registration failed", err);
            throw err;
        }
    };

    const logout = () => {
        authService.logout();
        setUser(null);
    };

    const updateUser = async (updatedUser: User) => {
        try {
            const user = await authService.updateProfile(updatedUser);
            setUser(user);
        } catch (err) {
            console.error("Update profile failed", err);
        }
    };

    return (
        <AuthContext.Provider value={{
            user,
            isLoading,
            isAuthModalOpen,
            setAuthModalOpen,
            login,
            register,
            logout,
            updateUser
        }}>
            {children}
        </AuthContext.Provider>
    );
};

export const useAuth = () => {
    const context = useContext(AuthContext);
    if (!context) {
        throw new Error('useAuth must be used within an AuthProvider');
    }
    return context;
};
