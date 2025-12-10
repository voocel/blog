
export interface BlogPost {
  id: string;
  title: string;
  excerpt: string;
  content: string;
  author: string;
  date: string;
  categoryId: string;
  category: string;
  readTime: string;
  imageUrl: string;
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
  username: string;
  email: string;
  role: 'admin' | 'visitor';
  avatar?: string;
}

export const Theme = {
  LIGHT: 'light',
  DARK: 'dark',
} as const;

export type Theme = typeof Theme[keyof typeof Theme];

export type AdminSection = 'overview' | 'posts' | 'categories' | 'tags' | 'files' | 'echoes';

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
