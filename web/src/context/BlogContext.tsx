import React, { createContext, useContext, useState, useEffect, type ReactNode } from 'react';
import type { BlogPost, Category, Tag } from '@/types';
import { postService } from '@/services/postService';
import { metaService } from '@/services/metaService';
import { useAuth } from '@/context/AuthContext';

interface BlogContextType {
  posts: BlogPost[];
  categories: Category[];
  tags: Tag[];
  error: string | null;

  // CRUD Operations
  addPost: (post: BlogPost) => Promise<void>;
  updatePost: (id: number, post: Partial<BlogPost>) => Promise<void>;
  deletePost: (id: number) => Promise<void>;

  addCategory: (category: Category) => Promise<void>;
  deleteCategory: (id: number) => Promise<void>;

  addTag: (tag: Tag) => Promise<void>;
  deleteTag: (id: number) => Promise<void>;

  // Refresh
  refreshPosts: () => Promise<void>;
  refreshCategories: () => Promise<void>;
  refreshTags: () => Promise<void>;

  // Logging
  logVisit: (path: string, postId?: number, postTitle?: string) => void;
}

const BlogContext = createContext<BlogContextType | undefined>(undefined);

export const BlogProvider: React.FC<{ children: ReactNode }> = ({ children }) => {
  const { user } = useAuth();
  const [posts, setPosts] = useState<BlogPost[]>([]);
  const [categories, setCategories] = useState<Category[]>([]);
  const [tags, setTags] = useState<Tag[]>([]);
  const [error, setError] = useState<string | null>(null);

  // Fetch Initial Public Data
  useEffect(() => {
    const fetchData = async () => {
      try {
        const [fetchedCategories, fetchedTags] = await Promise.all([
          metaService.getCategories(),
          metaService.getTags()
        ]);
        setCategories(fetchedCategories);
        setTags(fetchedTags);
      } catch (err) {
        setError('Failed to load blog data');
        console.error(err);
      }
    };
    fetchData();
  }, []);

  // Refresh posts when user role changes (e.g. login/logout)
  useEffect(() => {
    // Only refresh if we have already loaded (not on initial mount)
    // The posts are typically fetched on demand by pages, not globally
    // But admin status affects which posts are visible
  }, [user?.role]);

  const refreshPosts = async () => {
    try {
      let fetchedPosts: BlogPost[] = [];
      if (user?.role === 'admin') {
        fetchedPosts = await postService.getAdminPosts();
      } else {
        const response = await postService.getPosts();
        fetchedPosts = response.data;
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

  // --- Post Logic ---
  const addPost = async (post: BlogPost) => {
    try {
      await postService.createPost(post);
      await refreshPosts();
    } catch (err) {
      console.error("Failed to create post", err);
      throw err;
    }
  };

  const updatePost = async (id: number, updatedFields: Partial<BlogPost>) => {
    try {
      const updatedPost = await postService.updatePost(id, updatedFields);
      setPosts(prev => prev.map(p => p.id === id ? updatedPost : p));
    } catch (err) {
      console.error("Failed to update post", err);
      throw err;
    }
  };

  const deletePost = async (id: number) => {
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
      const updatedCategories = await metaService.getCategories();
      setCategories(updatedCategories);
    } catch (err) {
      console.error("Failed to add category", err);
      throw err;
    }
  };

  const deleteCategory = async (id: number) => {
    try {
      await metaService.deleteCategory(id);
      setCategories(prev => prev.filter(c => c.id !== id));
    } catch (err) {
      console.error("Failed to delete category", err);
      throw err;
    }
  };

  // --- Tag Logic ---
  const addTag = async (tag: Tag) => {
    try {
      await metaService.addTag(tag);
      const updatedTags = await metaService.getTags();
      setTags(updatedTags);
    } catch (err) {
      console.error("Failed to add tag", err);
      throw err;
    }
  };

  const deleteTag = async (id: number) => {
    try {
      await metaService.deleteTag(id);
      setTags(prev => prev.filter(t => t.id !== id));
    } catch (err) {
      console.error("Failed to delete tag", err);
      throw err;
    }
  };

  // --- Visit Logging Logic ---
  const logVisit = async (pagePath: string, postId?: number, postTitle?: string) => {
    // Don't log visits for admins to keep analytics clean
    if (user?.role?.toLowerCase() === 'admin') return;

    try {
      await metaService.logVisit(pagePath, postId, postTitle);
    } catch (err) {
      console.error("Failed to log visit", err);
    }
  };

  return (
    <BlogContext.Provider value={{
      posts, categories, tags, error,
      addPost, updatePost, deletePost,
      addCategory, deleteCategory,
      addTag, deleteTag,
      refreshPosts, refreshCategories, refreshTags,
      logVisit
    }}>
      {children}
    </BlogContext.Provider>
  );
};

// eslint-disable-next-line react-refresh/only-export-components
export const useBlog = () => {
  const context = useContext(BlogContext);
  if (!context) {
    throw new Error('useBlog must be used within a BlogProvider');
  }
  return context;
};
