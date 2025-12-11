import React, { createContext, useContext, useState, useEffect, type ReactNode } from 'react';
import type { BlogPost, User, Category, Tag, MediaFile, VisitLog, DashboardOverview } from '../types';
import { authService } from '../services/authService';
import { postService } from '../services/postService';
import { metaService } from '../services/metaService';



interface BlogContextType {
  posts: BlogPost[];
  categories: Category[];
  tags: Tag[];
  files: MediaFile[];
  user: User | null;
  visitLogs: VisitLog[];
  dashboardStats: DashboardOverview | null;

  // Loading State
  isLoading: boolean;
  error: string | null;

  // Modal State
  isAuthModalOpen: boolean;
  setAuthModalOpen: (isOpen: boolean) => void;

  // CRUD Operations
  addPost: (post: BlogPost) => Promise<void>;
  updatePost: (id: string, post: Partial<BlogPost>) => Promise<void>;
  deletePost: (id: string) => Promise<void>;

  addCategory: (category: Category) => Promise<void>;
  deleteCategory: (id: string) => Promise<void>;

  addTag: (tag: Tag) => Promise<void>;
  deleteTag: (id: string) => Promise<void>;

  addFile: (file: MediaFile) => Promise<void>;
  deleteFile: (id: string) => Promise<void>;

  // Auth
  login: (email: string, password?: string) => Promise<boolean>;
  register: (email: string, password: string) => Promise<boolean>;
  logout: () => void;
  updateUser: (user: User) => Promise<void>;

  // Navigation
  // activePostId: string | null; // Deprecated in favor of URL routing
  // setActivePostId: (id: string | null) => void; // Deprecated

  // Logging
  logVisit: (path: string, postId?: string, postTitle?: string) => void;

  // Admin
  refreshAdminData: () => Promise<void>;
  refreshPosts: () => Promise<void>;
  refreshCategories: () => Promise<void>;
  refreshTags: () => Promise<void>;
  refreshFiles: () => Promise<void>;
  refreshVisitLogs: () => Promise<void>;
  refreshDashboardOverview: () => Promise<void>;

  // Admin Management
  adminUsers: User[];
  allComments: any[];
  refreshAdminUsers: () => Promise<void>;
  refreshAllComments: () => Promise<void>;
}

const BlogContext = createContext<BlogContextType | undefined>(undefined);

export const BlogProvider: React.FC<{ children: ReactNode }> = ({ children }) => {
  const [posts, setPosts] = useState<BlogPost[]>([]);
  const [categories, setCategories] = useState<Category[]>([]);
  const [tags, setTags] = useState<Tag[]>([]);
  const [files, setFiles] = useState<MediaFile[]>([]);
  const [user, setUser] = useState<User | null>(null);

  const [isLoading, setIsLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  const [isAuthModalOpen, setAuthModalOpen] = useState(false);
  const [visitLogs, setVisitLogs] = useState<VisitLog[]>([]);
  const [dashboardStats, setDashboardStats] = useState<DashboardOverview | null>(null);

  // Fetch Initial Public Data
  useEffect(() => {
    const fetchData = async () => {
      setIsLoading(true);
      try {
        const [fetchedPosts, fetchedUser] = await Promise.all([
          postService.getPosts(),
          authService.getCurrentUser()
        ]);
        setPosts(fetchedPosts);
        setUser(fetchedUser);
      } catch (err) {
        setError('Failed to load blog data');
        console.error(err);
      } finally {
        setIsLoading(false);
      }
    };

    fetchData();
  }, []);

  // Admin Data Fetching (Lazy Load)
  const refreshAdminData = async () => {
    try {
      const [fetchedFiles, fetchedLogs] = await Promise.all([
        metaService.getFiles(),
        metaService.getVisitLogs()
      ]);
      setFiles(fetchedFiles);
      setVisitLogs(fetchedLogs);
    } catch (err) {
      console.error("Failed to load admin data", err);
    }
  };

  const refreshPosts = async () => {
    try {
      let fetchedPosts;
      if (user?.role === 'admin') {
        fetchedPosts = await postService.getAdminPosts();
      } else {
        fetchedPosts = await postService.getPosts();
      }
      setPosts(fetchedPosts);
    } catch (err) {
      console.error("Failed to refresh posts", err);
    }
  };

  const refreshCategories = async () => {
    try {
      const fetchedCategories = await metaService.getCategories();
      setCategories(fetchedCategories);
    } catch (err) {
      console.error("Failed to refresh categories", err);
    }
  };

  const refreshTags = async () => {
    try {
      const fetchedTags = await metaService.getTags();
      setTags(fetchedTags);
    } catch (err) {
      console.error("Failed to refresh tags", err);
    }
  };

  const refreshFiles = async () => {
    try {
      const fetchedFiles = await metaService.getFiles();
      setFiles(fetchedFiles);
    } catch (err) {
      console.error("Failed to refresh files", err);
    }
  };

  const refreshVisitLogs = async () => {
    try {
      const fetchedLogs = await metaService.getVisitLogs();
      setVisitLogs(fetchedLogs);
    } catch (err) {
      console.error("Failed to refresh visit logs", err);
    }
  };

  const refreshDashboardOverview = async () => {
    try {
      const stats = await metaService.getDashboardOverview();
      setDashboardStats(stats);
    } catch (err) {
      console.error("Failed to refresh dashboard overview", err);
    }
  };
  // --- Admin Users & Comments Logic ---
  const [adminUsers, setAdminUsers] = useState<User[]>([]);
  const [allComments, setAllComments] = useState<any[]>([]); // Using any[] temporarily if Comment type doesn't have post context, but for list view we usually need post title. We will check Comment type.

  const refreshAdminUsers = async () => {
    try {
      const users = await authService.getUsers();
      setAdminUsers(users);
    } catch (err) {
      console.error("Failed to refresh users", err);
    }
  };

  const refreshAllComments = async () => {
    try {
      // @ts-ignore - Assuming postService or commentService has this method
      const comments = await import('../services/commentService').then(m => m.commentService.getAllComments());
      setAllComments(comments);
    } catch (err) {
      console.error("Failed to refresh all comments", err);
    }
  };

  // --- Post Logic ---
  const addPost = async (post: BlogPost) => {
    try {
      // AdminDashboard now sends the correct payload structure (CreatePostDTO)
      // cast to any to avoid strict type checking against BlogPost interface which might differ slightly
      const newPost = await postService.createPost(post);
      setPosts(prev => [newPost, ...prev]);
      // Refresh to ensure consistency
      refreshPosts();
    } catch (err) {
      console.error("Failed to create post", err);
      throw err;
    }
  };

  const updatePost = async (id: string, updatedFields: Partial<BlogPost>) => {
    try {
      const updatedPost = await postService.updatePost(id, updatedFields);
      setPosts(prev => prev.map(p => p.id === id ? updatedPost : p));
    } catch (err) {
      console.error("Failed to update post", err);
      throw err;
    }
  };
  const deletePost = async (id: string) => {
    try {
      await postService.deletePost(id);
      setPosts(prev => prev.filter(post => post.id !== id));
    } catch (err) {
      console.error("Failed to delete post", err);
      throw err;
    }
  };

  // --- Category Logic ---
  const addCategory = async (category: Category) => {
    try {
      await metaService.addCategory(category);
      // Re-fetch categories to get the latest list (including server-generated IDs)
      const updatedCategories = await metaService.getCategories();
      setCategories(updatedCategories);
    } catch (err) {
      console.error("Failed to add category", err);
    }
  }
  const deleteCategory = async (id: string) => {
    try {
      await metaService.deleteCategory(id);
      setCategories(prev => prev.filter(c => c.id !== id));
    } catch (err) {
      console.error("Failed to delete category", err);
    }
  }

  // --- Tag Logic ---
  const addTag = async (tag: Tag) => {
    try {
      await metaService.addTag(tag);
      // Re-fetch tags to get the latest list
      const updatedTags = await metaService.getTags();
      setTags(updatedTags);
    } catch (err) {
      console.error("Failed to add tag", err);
    }
  }
  const deleteTag = async (id: string) => {
    try {
      await metaService.deleteTag(id);
      setTags(prev => prev.filter(t => t.id !== id));
    } catch (err) {
      console.error("Failed to delete tag", err);
    }
  }

  // --- File Logic ---
  const addFile = async (file: MediaFile) => {
    try {
      const newFile = await metaService.addFile(file);
      setFiles(prev => [newFile, ...prev]);
    } catch (err) {
      console.error("Failed to add file", err);
    }
  }
  const deleteFile = async (id: string) => {
    try {
      await metaService.deleteFile(id);
      setFiles(prev => prev.filter(f => f.id !== id));
    } catch (err) {
      console.error("Failed to delete file", err);
    }
  };

  // --- Auth Logic ---
  const login = async (email: string, password?: string) => {
    try {
      const user = await authService.login(email, password);
      if (user) {
        setUser(user);
        return true;
      }
      return false;
    } catch (err) {
      console.error("Login failed", err);
      throw err;
    }
  };

  const register = async (email: string, password: string) => {
    try {
      const user = await authService.register(email, password);
      if (user) {
        setUser(user);
        return true;
      }
      return false;
    } catch (err) {
      console.error("Registration failed", err);
      throw err;
    }
  };

  const logout = () => {
    authService.logout();
    setUser(null);
  };

  const updateUser = async (updatedUser: User) => {
    try {
      const user = await authService.updateProfile(updatedUser);
      setUser(user);
    } catch (err) {
      console.error("Update profile failed", err);
    }
  };

  // --- Visit Logging Logic ---
  const logVisit = async (pagePath: string, postId?: string, postTitle?: string) => {
    // Don't log visits for admins to keep analytics clean
    if (user?.role?.toLowerCase() === 'admin') return;

    try {
      await metaService.logVisit(pagePath, postId, postTitle);
      // Optionally refresh logs if we are on the admin page, but for now we just log it.
      // If we want to see it immediately in the dashboard, we might need to re-fetch logs,
      // but typically analytics are viewed later.
    } catch (err) {
      console.error("Failed to log visit", err);
    }
  };

  return (
    <BlogContext.Provider value={{
      posts, categories, tags, files, user, visitLogs,
      isLoading, error,
      isAuthModalOpen, setAuthModalOpen,
      addPost, updatePost, deletePost,
      addCategory, deleteCategory,
      addTag, deleteTag,
      addFile, deleteFile,
      login, logout, updateUser, register,
      logVisit,
      refreshAdminData,
      refreshPosts,
      refreshCategories,
      refreshTags,
      refreshFiles,
      refreshVisitLogs,
      dashboardStats,
      refreshDashboardOverview,
      // Admin Users & Comments
      adminUsers,
      allComments,
      refreshAdminUsers,
      refreshAllComments
    }}>
      {children}
    </BlogContext.Provider>
  );
};

export const useBlog = () => {
  const context = useContext(BlogContext);
  if (!context) {
    throw new Error('useBlog must be used within a BlogProvider');
  }
  return context;
};
