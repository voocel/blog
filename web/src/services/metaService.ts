import apiClient from './apiClient';
import { Category, Tag, MediaFile, DashboardOverview } from '../types';

export const metaService = {
    // Categories
    getCategories: async (): Promise<Category[]> => {
        const response = await apiClient.get('/categories');
        return response.data;
    },
    addCategory: async (category: Category): Promise<Category> => {
        const response = await apiClient.post('/categories', category);
        return response.data;
    },
    deleteCategory: async (id: string): Promise<void> => {
        await apiClient.delete(`/categories/${id}`);
    },

    // Tags
    getTags: async (): Promise<Tag[]> => {
        const response = await apiClient.get('/tags');
        return response.data;
    },
    addTag: async (tag: Tag): Promise<Tag> => {
        const response = await apiClient.post('/tags', tag);
        return response.data;
    },
    deleteTag: async (id: string): Promise<void> => {
        await apiClient.delete(`/tags/${id}`);
    },

    // Files (Media)
    getFiles: async (): Promise<MediaFile[]> => {
        const response = await apiClient.get('/files');
        return response.data;
    },
    addFile: async (file: MediaFile): Promise<MediaFile> => {
        // The file is already uploaded via uploadService (POST /upload).
        // This method is called by BlogContext to update local state.
        // We just return the file as is, assuming the server already has it.
        return Promise.resolve(file);
    },
    deleteFile: async (id: string): Promise<void> => {
        await apiClient.delete(`/files/${id}`);
    },

    // Analytics
    logVisit: async (pagePath: string, postId?: string, postTitle?: string): Promise<void> => {
        await apiClient.post('/analytics/visit', { pagePath, postId, postTitle });
    },

    getVisitLogs: async (): Promise<any[]> => {
        const response = await apiClient.get('/analytics/logs');
        return response.data;
    },

    getDashboardOverview: async (): Promise<DashboardOverview> => {
        const response = await apiClient.get('/analytics/dashboard-overview');
        return response.data;
    }
};
