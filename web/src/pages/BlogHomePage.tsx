import React, { useEffect, useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { motion } from 'framer-motion';
import AnimatedNavWidget from '@/components/AnimatedNavWidget';
import SettingsModal from '@/components/SettingsModal';
import SEO from '@/components/SEO';
import { postService } from '@/services/postService';
import type { BlogPost } from '@/types';
import { useAuth } from '@/context/AuthContext';
import { useTranslation } from '@/hooks/useTranslation';

const BlogHomePage: React.FC = () => {
  const navigate = useNavigate();
  const { user, setAuthModalOpen } = useAuth();
  const { t, locale } = useTranslation();
  const [posts, setPosts] = useState<BlogPost[]>([]);
  const [loading, setLoading] = useState(true);
  const [isSettingsOpen, setIsSettingsOpen] = useState(false);

  useEffect(() => {
    postService.getPosts({ page: 1, limit: 30 })
      .then(res => setPosts(res.data))
      .catch(err => console.error('Failed to fetch posts:', err))
      .finally(() => setLoading(false));
  }, []);

  const formatDate = (dateStr: string) => {
    const date = new Date(dateStr);
    return date.toLocaleDateString(locale === 'zh' ? 'zh-CN' : 'en-US', {
      month: 'short',
      day: 'numeric',
      year: 'numeric',
    });
  };

  return (
    <div className="min-h-screen bg-[var(--color-base)]">
      <SEO title="Blog - Voocel" />
      <SettingsModal isOpen={isSettingsOpen} onClose={() => setIsSettingsOpen(false)} />

      {/* Nav - top left */}
      <div className="fixed top-6 left-6 z-50">
        <AnimatedNavWidget isCompact={true} disableFixed={true} />
      </div>

      {/* Actions - top right */}
      <div className="fixed top-6 right-6 z-50 flex items-center gap-2">
        {user && (
          <button
            onClick={() => navigate('/admin')}
            className="bg-red-400 text-white px-3 py-1.5 rounded-xl shadow-red-200 shadow-md font-bold text-xs hover:bg-red-500 transition-colors cursor-pointer"
          >
            {t.home.dashboard}
          </button>
        )}
        {!user && (
          <button
            onClick={() => setAuthModalOpen(true)}
            className="text-[10px] text-[var(--color-text-muted)] font-bold px-2 hover:text-[var(--color-text-secondary)] transition-colors cursor-pointer"
          >
            {t.home.signIn.toUpperCase()}
          </button>
        )}
        <button
          onClick={() => setIsSettingsOpen(true)}
          className="w-8 h-8 rounded-xl bg-[var(--color-elevated)] hover:bg-stone-200/50 dark:hover:bg-stone-700/50 flex items-center justify-center text-[var(--color-text-muted)] transition-colors cursor-pointer shadow-sm"
        >
          <svg width="16" height="16" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg" aria-hidden="true">
            <circle cx="5" cy="5" r="2" fill="currentColor" />
            <circle cx="12" cy="5" r="2" fill="currentColor" />
            <circle cx="19" cy="5" r="2" fill="currentColor" />
            <circle cx="5" cy="12" r="2" fill="currentColor" />
            <circle cx="12" cy="12" r="2" fill="currentColor" />
            <circle cx="19" cy="12" r="2" fill="currentColor" />
            <circle cx="5" cy="19" r="2" fill="currentColor" />
            <circle cx="12" cy="19" r="2" fill="currentColor" />
            <circle cx="19" cy="19" r="2" fill="currentColor" />
          </svg>
        </button>
      </div>

      {/* Main content */}
      <motion.main
        className="max-w-2xl mx-auto px-6 pt-28 pb-20"
        initial={{ opacity: 0, y: 20 }}
        animate={{ opacity: 1, y: 0 }}
        transition={{ duration: 0.4, ease: 'easeOut' }}
      >
        {/* Header */}
        <header className="mb-16">
          <h1 className="text-4xl font-serif font-bold text-ink tracking-tight">
            {t.blogHome.title}
          </h1>
          <p className="text-[var(--color-text-secondary)] text-lg mt-3 font-serif italic">
            {t.blogHome.subtitle}
          </p>
          <div className="w-12 h-[2px] bg-orange-400 mt-6" />
        </header>

        {/* Loading */}
        {loading && (
          <div className="space-y-6">
            {[...Array(5)].map((_, i) => (
              <div key={i} className="animate-pulse">
                <div className="h-3 w-24 bg-stone-200 dark:bg-stone-700 rounded mb-3" />
                <div className="h-4 w-3/4 bg-stone-200 dark:bg-stone-700 rounded mb-2" />
                <div className="h-3 w-1/2 bg-stone-100 dark:bg-stone-800 rounded" />
              </div>
            ))}
          </div>
        )}

        {/* Posts */}
        {!loading && posts.length > 0 && (
          <div>
            {posts.map((post, i) => (
              <motion.article
                key={post.id}
                onClick={() => navigate(`/post/${post.slug}`)}
                className="group cursor-pointer py-5 border-b border-[var(--color-border)] last:border-b-0 -mx-4 px-4 rounded-lg hover:bg-[var(--color-surface-alt)]/40 transition-colors"
                role="button"
                tabIndex={0}
                onKeyDown={(e) => e.key === 'Enter' && navigate(`/post/${post.slug}`)}
                initial={{ opacity: 0, y: 10 }}
                animate={{ opacity: 1, y: 0 }}
                transition={{ delay: i * 0.04, duration: 0.3 }}
              >
                <div className="flex items-start gap-6">
                  {/* Date */}
                  <time className="text-xs text-[var(--color-text-muted)] font-mono shrink-0 w-[100px] pt-0.5 hidden sm:block">
                    {formatDate(post.publishAt)}
                  </time>

                  {/* Content */}
                  <div className="flex-1 min-w-0">
                    <h2 className="text-base font-semibold text-ink group-hover:text-orange-600 transition-colors leading-snug">
                      {post.title}
                    </h2>
                    {post.excerpt && (
                      <p className="text-sm text-[var(--color-text-secondary)] mt-1.5 line-clamp-1">
                        {post.excerpt}
                      </p>
                    )}
                    <div className="flex items-center gap-3 mt-2">
                      {post.category && (
                        <span className="text-[11px] text-orange-500 bg-orange-50 dark:bg-orange-950/40 px-2 py-0.5 rounded-md font-medium">
                          {post.category}
                        </span>
                      )}
                      {post.readTime && (
                        <span className="text-[11px] text-[var(--color-text-muted)]">
                          {post.readTime}
                        </span>
                      )}
                      {/* Mobile date */}
                      <span className="text-[11px] text-[var(--color-text-muted)] font-mono sm:hidden">
                        {formatDate(post.publishAt)}
                      </span>
                    </div>
                  </div>

                  {/* Arrow */}
                  <span className="text-[var(--color-text-muted)] opacity-0 group-hover:opacity-100 transition-opacity shrink-0 pt-0.5">
                    &rarr;
                  </span>
                </div>
              </motion.article>
            ))}

            {/* View all */}
            <motion.div
              className="mt-12 text-center"
              initial={{ opacity: 0 }}
              animate={{ opacity: 1 }}
              transition={{ delay: 0.5 }}
            >
              <button
                onClick={() => navigate('/posts')}
                className="text-sm text-[var(--color-text-secondary)] hover:text-orange-600 transition-colors font-medium cursor-pointer"
              >
                {t.blogHome.viewAll} &rarr;
              </button>
            </motion.div>
          </div>
        )}

        {/* Empty */}
        {!loading && posts.length === 0 && (
          <motion.p
            className="text-center text-[var(--color-text-muted)] py-20"
            initial={{ opacity: 0 }}
            animate={{ opacity: 1 }}
            transition={{ delay: 0.3 }}
          >
            {t.blogHome.noPosts}
          </motion.p>
        )}
      </motion.main>
    </div>
  );
};

export default BlogHomePage;
