
export interface BlogPost {
  id: string;
  slug: string;
  title: string;
  excerpt: string;
  content: string;
  author: string;
  // Scheduled publish time in RFC3339 (e.g. 2025-12-14T16:30:00+08:00)
  publishAt: string;
  categoryId: string;
  category: string;
  readTime: string;
  cover: string;
  tags: string[]; // Array of Tag IDs
  views: number;
  status: 'published' | 'draft';
}

export interface Category {
  id: string;
  name: string;
  slug: string;
  count: number;
}

export interface Tag {
  id: string;
  name: string;
}

export interface MediaFile {
  id: string;
  url: string;
  name: string;
  type: 'image' | 'video' | 'document';
  date: string;
}

export interface ChatMessage {
  role: 'user' | 'model';
  text: string;
  isError?: boolean;
}

export interface Comment {
  id: string;
  parentId: string | null;
  content: string;
  createdAt: string;
  user: {
    username: string;
    avatar?: string;
  };
  replies?: Comment[];
  replyToUser?: {
    username: string;
    avatar?: string;
  };
}

export interface User {
  id: string; // Ensure ID is present for management
  username: string;
  email: string;
  role: 'admin' | 'visitor';
  status?: 'active' | 'banned'; // Optional for backward compatibility if backend doesn't always send
  provider?: 'email' | 'google' | 'github'; // Auth provider
  avatar?: string;
}

export const Theme = {
  LIGHT: 'light',
  DARK: 'dark',
} as const;

export type Theme = typeof Theme[keyof typeof Theme];

export type AdminSection = 'overview' | 'posts' | 'categories' | 'tags' | 'files' | 'echoes' | 'users' | 'comments';

export interface VisitLog {
  id: string;
  pagePath: string;
  postId?: string;
  postTitle?: string;
  ip: string;
  location: string;
  timestamp: number;
  userAgent: string;
}

export interface DashboardOverview {
  counts: {
    posts: number;
    categories: number;
    tags: number;
    files: number;
  };
  recentPosts: BlogPost[];
  systemStatus: {
    storageUsage: number;
    aiQuota: number;
  };
}

export interface AuthResponse {
  access_token: string;
  refresh_token: string;
  expires_in: number;
  user: User;
}
