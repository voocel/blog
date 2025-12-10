import apiClient, { setTokens, clearTokens, getRefreshToken, getAccessToken } from './apiClient';
import type { User, AuthResponse } from '../types';

export const authService = {
    login: async (email: string, password?: string): Promise<User> => {
        try {
            const response = await apiClient.post('/auth/login', { email, password });
            const { access_token, refresh_token, user } = response.data as AuthResponse;
            setTokens(access_token, refresh_token);
            return user;
        } catch (error) {
            console.error('Login failed:', error);
            throw error;
        }
    },

    register: async (email: string, password: string): Promise<User | null> => {
        try {
            const response = await apiClient.post('/auth/register', { email, password });
            const { access_token, refresh_token, user } = response.data as AuthResponse;
            setTokens(access_token, refresh_token);
            return user;
        } catch (error) {
            console.error('Registration failed:', error);
            throw error;
        }
    },

    getCurrentUser: async (): Promise<User | null> => {
        const token = getAccessToken();
        if (!token) {
            return null;
        }

        try {
            const response = await apiClient.get('/auth/me');
            return response.data;
        } catch (error) {
            console.error('Failed to get current user:', error);
            return null;
        }
    },

    updateProfile: async (user: Partial<User>): Promise<User> => {
        const response = await apiClient.put('/users/profile', user);
        return response.data;
    },

    logout: () => {
        clearTokens();
    },

    refreshToken: async (): Promise<void> => {
        const refreshToken = getRefreshToken();
        if (!refreshToken) throw new Error("No refresh token");

        const response = await apiClient.post('/auth/refresh', { refresh_token: refreshToken });
        const { access_token, refresh_token: new_refresh_token } = response.data as AuthResponse;
        setTokens(access_token, new_refresh_token);
    }
};
