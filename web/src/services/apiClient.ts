import axios from 'axios';
import type { AuthResponse } from '../types';

// Token Management Helpers
export const getAccessToken = () => localStorage.getItem('access_token');
export const getRefreshToken = () => localStorage.getItem('refresh_token');
export const setTokens = (accessToken: string, refreshToken: string) => {
    localStorage.setItem('access_token', accessToken);
    localStorage.setItem('refresh_token', refreshToken);
};
export const clearTokens = () => {
    localStorage.removeItem('access_token');
    localStorage.removeItem('refresh_token');
    // Also clear legacy token if exists
    localStorage.removeItem('authToken');
};

// Create an axios instance
const apiClient = axios.create({
    // Use relative path by default so it works behind reverse proxy
    baseURL: import.meta.env.VITE_API_URL || '/api/v1',
    headers: {
        'Content-Type': 'application/json',
    },
});

// Add a request interceptor to attach the token
apiClient.interceptors.request.use(
    (config) => {
        const token = getAccessToken();
        if (token) {
            config.headers.Authorization = `Bearer ${token}`;
        }
        return config;
    },
    (error) => {
        return Promise.reject(error);
    }
);

// Flag to prevent multiple refresh requests
let isRefreshing = false;
// Queue for failed requests
let failedQueue: any[] = [];

const processQueue = (error: any, token: string | null = null) => {
    failedQueue.forEach(prom => {
        if (error) {
            prom.reject(error);
        } else {
            prom.resolve(token);
        }
    });
    failedQueue = [];
};

// Add a response interceptor for global error handling
apiClient.interceptors.response.use(
    (response) => response,
    async (error) => {
        const originalRequest = error.config;
        const status = error.response?.status as number | undefined;
        const data = error.response?.data as any;
        const msg: string = (() => {
            if (!data) return '';
            if (typeof data === 'string') return data;
            if (typeof data?.error === 'string') return data.error;
            if (typeof data?.message === 'string') return data.message;
            return '';
        })();
        const msgLower = msg.toLowerCase();

        // Token revoked / user banned: clear tokens and stop retrying.
        // Backend may return either standardized { error: "..." } or gin.H{"error": "..."}.
        if ((status === 401 && msgLower.includes('token revoked')) || (status === 403 && msgLower.includes('banned'))) {
            clearTokens();
            // window.location.href = '/login'; // Optional: redirect
            return Promise.reject(error);
        }

        // If error is 401 and we haven't retried yet
        if (status === 401 && !originalRequest._retry) {
            if (isRefreshing) {
                // If already refreshing, queue this request
                return new Promise(function (resolve, reject) {
                    failedQueue.push({ resolve, reject });
                }).then(token => {
                    originalRequest.headers['Authorization'] = 'Bearer ' + token;
                    return apiClient(originalRequest);
                }).catch(err => {
                    return Promise.reject(err);
                });
            }

            originalRequest._retry = true;
            isRefreshing = true;

            const refreshToken = getRefreshToken();

            if (!refreshToken) {
                clearTokens();
                // window.location.href = '/login'; // Optional: Redirect to login
                return Promise.reject(error);
            }

            try {
                // Call refresh endpoint
                // Note: We use axios.create() to avoid interceptors on this call
                const response = await axios.post(`${apiClient.defaults.baseURL}/auth/refresh`, {
                    refresh_token: refreshToken
                });

                const { access_token, refresh_token } = response.data as AuthResponse;

                setTokens(access_token, refresh_token);

                apiClient.defaults.headers.common['Authorization'] = 'Bearer ' + access_token;
                originalRequest.headers['Authorization'] = 'Bearer ' + access_token;

                processQueue(null, access_token);
                return apiClient(originalRequest);
            } catch (refreshError) {
                processQueue(refreshError, null);
                clearTokens();
                // window.location.href = '/login'; // Redirect to login on failed refresh
                return Promise.reject(refreshError);
            } finally {
                isRefreshing = false;
            }
        }

        return Promise.reject(error);
    }
);

export default apiClient;
