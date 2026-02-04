import apiClient from '@/services/apiClient';

export interface UploadResult {
    url: string;
    filename: string;
    type: string;
}

/**
 * Uploads an image to the backend.
 */
export const uploadImage = async (file: File): Promise<UploadResult> => {
    const formData = new FormData();
    formData.append('file', file);

    const response = await apiClient.post('/admin/upload', formData, {
        headers: {
            'Content-Type': 'multipart/form-data',
        },
    });

    // API returns { id, url, name, type, date }
    // We map it to UploadResult
    return {
        url: response.data.url,
        filename: response.data.name,
        type: response.data.type
    };
};
