
export interface BlogPost {
  id: number;
  slug: string;
  title: string;
  excerpt: string;
  content: string;
  author: string;
  // Scheduled publish time in RFC3339 (e.g. 2025-12-14T16:30:00+08:00)
  publishAt: string;
  categoryId: number;
  category: string;
  readTime: string;
  cover: string;
  tags: number[]; // Tag IDs for editing; API response returns string[] names (JS handles automatically)
  views: number;
  status: 'published' | 'draft';
}

export interface Category {
  id: number;
  name: string;
  slug: string;
  count: number;
}

export interface Tag {
  id: number;
  name: string;
}

export interface MediaFile {
  id: number;
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
  id: number;
  parentId: number | null;
  content: string;
  createdAt: string;
  postTitle?: string;  // Available in admin list
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
  id: number;
  username: string;
  email: string;
  role: 'admin' | 'visitor';
  status?: 'active' | 'banned';
  provider?: 'email' | 'google' | 'github';
  avatar?: string;
}

export const Theme = {
  LIGHT: 'light',
  DARK: 'dark',
} as const;

export type Theme = typeof Theme[keyof typeof Theme];

export type AdminSection = 'overview' | 'posts' | 'categories' | 'tags' | 'files' | 'echoes' | 'users' | 'comments';

export interface VisitLog {
  id: number;
  pagePath: string;
  postId?: number;
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
