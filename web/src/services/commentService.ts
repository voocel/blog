import type { Comment } from '@/types';

interface CommentResponse {
    data: Comment[];
    pagination: {
        total: number;
        page: number;
        limit: number;
        totalPages: number;
    };
}

import apiClient from '@/services/apiClient';

// ... interface CommentResponse remains the same ...

export const commentService = {
    getComments: async (postSlug: string, page = 1, limit = 20): Promise<CommentResponse> => {
        try {
            const response = await apiClient.get(`/posts/${postSlug}/comments`, {
                params: { page, limit, withReplies: true }
            });
            // Handle both structure: { data: [...], pagination: {...} } or direct array [...]
            if (response.data.data && Array.isArray(response.data.data)) {
                return response.data;
            }
            if (Array.isArray(response.data)) {
                return {
                    data: response.data,
                    pagination: { total: response.data.length, page, limit, totalPages: 1 }
                };
            }
            return { data: [], pagination: { total: 0, page, limit, totalPages: 0 } };
        } catch (error) {
            console.error('Failed to fetch comments:', error);
            return { data: [], pagination: { total: 0, page, limit, totalPages: 0 } };
        }
    },

    createComment: async (postSlug: string, content: string, parentId?: number): Promise<Comment> => {
        const payload: { content: string; parentId?: number } = { content };
        if (parentId) {
            payload.parentId = parentId;
        }
        const response = await apiClient.post(`/posts/${postSlug}/comments`, payload);
        return response.data;
    },

    // Admin Methods
    getAllComments: async (): Promise<Comment[]> => {
        const response = await apiClient.get('/admin/comments');
        return response.data;
    },

    deleteComment: async (id: number): Promise<void> => {
        await apiClient.delete(`/admin/comments/${id}`);
    }
};
