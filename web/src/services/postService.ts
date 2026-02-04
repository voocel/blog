import apiClient from '@/services/apiClient';
import type { BlogPost } from '@/types';

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

    getPost: async (slug: string): Promise<BlogPost | undefined> => {
        try {
            // Public API: /posts/:slug (only published)
            const response = await apiClient.get(`/posts/${slug}`);
            return response.data;
        } catch (error) {
            console.error(`Failed to get post ${slug}`, error);
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

    getAdminPost: async (id: number): Promise<BlogPost | undefined> => {
        try {
            // Admin API: /admin/posts/:id (can view drafts)
            const response = await apiClient.get(`/admin/posts/${id}`);
            return response.data;
        } catch (error) {
            console.error(`Failed to get admin post ${id}`, error);
            return undefined;
        }
    },

    createPost: async (postData: Partial<BlogPost>): Promise<BlogPost> => {
        const response = await apiClient.post('/admin/posts', postData);
        return response.data;
    },

    updatePost: async (id: number, updatedFields: Partial<BlogPost>): Promise<BlogPost> => {
        const response = await apiClient.put(`/admin/posts/${id}`, updatedFields);
        return response.data;
    },

    deletePost: async (id: number): Promise<void> => {
        await apiClient.delete(`/admin/posts/${id}`);
    },

    // Like Endpoints
    getLikes: async (slug: string): Promise<number> => {
        try {
            const response = await apiClient.get('/likes', { params: { slug } });
            return response.data?.count ?? 0;
        } catch (error) {
            console.error(`Failed to get likes for ${slug}:`, error);
            return 0;
        }
    },

    like: async (slug: string): Promise<number> => {
        const response = await apiClient.post('/likes', null, { params: { slug } });
        return response.data?.count ?? 0;
    },
};
