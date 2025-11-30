
export const getAssetUrl = (path: string | undefined): string => {
    if (!path) return '';
    if (path.startsWith('http') || path.startsWith('https') || path.startsWith('data:')) {
        return path;
    }

    // Get the base URL from environment or default
    // We assume the API URL might contain /api/v1 but assets are served from root
    const apiUrl = import.meta.env.VITE_API_URL || 'http://localhost:8080/api/v1';

    // Extract origin (protocol + hostname + port)
    // This is a simple heuristic: remove /api/v1 suffix if present
    const origin = apiUrl.replace(/\/api\/v1\/?$/, '');

    // Ensure path starts with /
    const cleanPath = path.startsWith('/') ? path : `/${path}`;

    return `${origin}${cleanPath}`;
};
