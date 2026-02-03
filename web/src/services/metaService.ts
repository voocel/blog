import apiClient from './apiClient';
import type { Category, Tag, MediaFile, DashboardOverview } from '../types';

export const metaService = {
    // Categories
    getCategories: async (): Promise<Category[]> => {
        const response = await apiClient.get('/categories');
        return response.data;
    },
    addCategory: async (category: Category): Promise<Category> => {
        const response = await apiClient.post('/admin/categories', category);
        return response.data;
    },
    deleteCategory: async (id: number): Promise<void> => {
        await apiClient.delete(`/admin/categories/${id}`);
    },

    // Tags
    getTags: async (): Promise<Tag[]> => {
        const response = await apiClient.get('/tags');
        return response.data;
    },
    addTag: async (tag: Tag): Promise<Tag> => {
        const response = await apiClient.post('/admin/tags', tag);
        return response.data;
    },
    deleteTag: async (id: number): Promise<void> => {
        await apiClient.delete(`/admin/tags/${id}`);
    },

    // Files (Media)
    getFiles: async (): Promise<MediaFile[]> => {
        const response = await apiClient.get('/admin/files');
        return response.data;
    },
    addFile: async (file: MediaFile): Promise<MediaFile> => {
        // The file is already uploaded via uploadService (POST /upload).
        // This method is called by BlogContext to update local state.
        // We just return the file as is, assuming the server already has it.
        return Promise.resolve(file);
    },
    deleteFile: async (id: number): Promise<void> => {
        await apiClient.delete(`/admin/files/${id}`);
    },

    // Analytics
    logVisit: async (pagePath: string, postId?: number, postTitle?: string): Promise<void> => {
        await apiClient.post('/analytics/visit', { pagePath, postId, postTitle });
    },

    getVisitLogs: async (): Promise<any[]> => {
        const response = await apiClient.get('/admin/analytics/logs');
        return response.data;
    },

    getDashboardOverview: async (): Promise<DashboardOverview> => {
        const response = await apiClient.get('/admin/analytics/dashboard-overview');
        return response.data;
    }
};
