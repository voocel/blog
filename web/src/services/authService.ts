import apiClient from './apiClient';
import { User } from '../types';

export const authService = {
    login: async (email: string, password?: string): Promise<User> => {
        try {
            const response = await apiClient.post('/auth/login', { email, password });
            const { token, user } = response.data;
            if (token) {
                localStorage.setItem('authToken', token);
            }
            return user;
        } catch (error) {
            console.error('Login failed:', error);
            throw error;
        }
    },

    register: async (email: string, password: string): Promise<User | null> => {
        try {
            const response = await apiClient.post('/auth/register', { email, password });
            const { token, user } = response.data;
            if (token) {
                localStorage.setItem('authToken', token);
            }
            return user;
        } catch (error) {
            console.error('Registration failed:', error);
            throw error;
        }
    },

    getCurrentUser: async (): Promise<User | null> => {
        try {
            const token = localStorage.getItem('authToken');
            if (!token) return null;

            const response = await apiClient.get('/auth/me');
            return response.data;
        } catch (error) {
            console.error('Failed to get current user:', error);
            localStorage.removeItem('authToken');
            return null;
        }
    },

    updateProfile: async (user: Partial<User>): Promise<User> => {
        const response = await apiClient.put('/users/profile', user);
        return response.data;
    },

    logout: () => {
        localStorage.removeItem('authToken');
    }
};
