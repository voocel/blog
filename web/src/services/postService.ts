import apiClient from './apiClient';
import type { BlogPost } from '../types';

export const postService = {
    // Public Endpoints
    getPosts: async (params?: { category?: string; tag?: string; search?: string; page?: number; limit?: number }): Promise<{ data: BlogPost[]; pagination?: { total: number; page: number; limit: number; totalPages: number } }> => {
        // Public API: /posts
        const response = await apiClient.get('/posts', { params });
        if (response.data?.data && Array.isArray(response.data.data)) {
            return response.data;
        }
        if (Array.isArray(response.data)) {
            // Unpaginated response (legacy or no page param)
            return { data: response.data, pagination: undefined };
        }
        console.warn('getPosts: Expected array result, got:', response.data);
        return { data: [], pagination: undefined };
    },

    getPost: async (id: string): Promise<BlogPost | undefined> => {
        try {
            // Public API: /posts/:id (only published)
            const response = await apiClient.get(`/posts/${id}`);
            return response.data;
        } catch (error) {
            console.error(`Failed to get post ${id}`, error);
            return undefined;
        }
    },

    // Admin Endpoints
    getAdminPosts: async (params?: { category?: string; status?: string; search?: string; page?: number; limit?: number }): Promise<BlogPost[]> => {
        // Admin API: /admin/posts (all statuses)
        const response = await apiClient.get('/admin/posts', { params });
        if (response.data?.data && Array.isArray(response.data.data)) {
            return response.data.data;
        }
        if (Array.isArray(response.data)) {
            return response.data;
        }
        console.warn('getAdminPosts: Expected array result, got:', response.data);
        return [];
    },

    getAdminPost: async (id: string): Promise<BlogPost | undefined> => {
        try {
            // Admin API: /admin/posts/:id (can view drafts)
            const response = await apiClient.get(`/admin/posts/${id}`);
            return response.data;
        } catch (error) {
            console.error(`Failed to get admin post ${id}`, error);
            return undefined;
        }
    },

    createPost: async (postData: any): Promise<BlogPost> => {
        const response = await apiClient.post('/admin/posts', postData);
        return response.data;
    },

    updatePost: async (id: string, updatedFields: any): Promise<BlogPost> => {
        const response = await apiClient.put(`/admin/posts/${id}`, updatedFields);
        return response.data;
    },

    deletePost: async (id: string): Promise<void> => {
        await apiClient.delete(`/admin/posts/${id}`);
    }
};
