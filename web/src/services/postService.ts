import apiClient from './apiClient';
import { BlogPost } from '../types';

export const postService = {
    getPosts: async (params?: { category?: string; tag?: string; status?: string; search?: string; page?: number; limit?: number }): Promise<BlogPost[]> => {
        const response = await apiClient.get('/posts', { params });
        // The API returns { data: BlogPost[], pagination: ... } or just BlogPost[] based on v1 spec.
        // Our frontend expects BlogPost[].
        // If the API returns the paginated structure, we need to extract .data.
        if (response.data.data && Array.isArray(response.data.data)) {
            return response.data.data;
        }
        return response.data;
    },

    getPost: async (id: string): Promise<BlogPost | undefined> => {
        try {
            const response = await apiClient.get(`/posts/${id}`);
            return response.data;
        } catch (error) {
            console.error(`Failed to get post ${id}`, error);
            return undefined;
        }
    },

    createPost: async (postData: any): Promise<BlogPost> => {
        const response = await apiClient.post('/posts', postData);
        return response.data;
    },

    updatePost: async (id: string, updatedFields: any): Promise<BlogPost> => {
        const response = await apiClient.put(`/posts/${id}`, updatedFields);
        return response.data;
    },

    deletePost: async (id: string): Promise<void> => {
        await apiClient.delete(`/posts/${id}`);
    }
};
