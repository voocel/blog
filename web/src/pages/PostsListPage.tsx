
import React, { useEffect, useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { motion } from 'framer-motion';
import AnimatedNavWidget from '../components/AnimatedNavWidget';
import SEO from '../components/SEO';
import { postService } from '../services/postService';
import type { BlogPost } from '../types';

interface PostGroup {
    label: string;
    count: number;
    posts: BlogPost[];
}

// Animation variants
const containerVariants = {
    hidden: { opacity: 0 },
    visible: {
        opacity: 1,
        transition: {
            staggerChildren: 0.15,
            delayChildren: 0.1,
        },
    },
};

const filterVariants = {
    hidden: { opacity: 0, y: -20 },
    visible: {
        opacity: 1,
        y: 0,
        transition: {
            type: 'spring' as const,
            stiffness: 100,
            damping: 15,
        },
    },
};



const PostsListPage: React.FC = () => {
    const navigate = useNavigate();
    const [posts, setPosts] = useState<BlogPost[]>([]);
    const [loading, setLoading] = useState(true);
    const [activeFilter, setActiveFilter] = useState<'day' | 'week' | 'month' | 'year' | 'category'>('year');

    useEffect(() => {
        postService.getPosts({ page: 1, limit: 100 })
            .then(res => {
                setPosts(res.data);
            })
            .catch(err => {
                console.error('Failed to fetch posts:', err);
            })
            .finally(() => {
                setLoading(false);
            });
    }, []);

    // Group posts based on active filter
    const groupedPosts: PostGroup[] = React.useMemo(() => {
        if (posts.length === 0) {
            return [];
        }

        const groups: Record<string, BlogPost[]> = {};

        posts.forEach(post => {
            const date = new Date(post.publishAt);
            let key: string;

            switch (activeFilter) {
                case 'day':
                    // Group by specific date: "Jan 09, 2026"
                    key = `${date.toLocaleDateString('en-US', { month: 'short', day: '2-digit', year: 'numeric' })}`;
                    break;
                case 'week':
                    // Group by week number
                    const weekNum = Math.ceil((date.getDate()) / 7);
                    key = `${date.toLocaleDateString('en-US', { month: 'short', year: 'numeric' })} Week ${weekNum}`;
                    break;
                case 'month':
                    // Group by month: "Jan 2026"
                    key = `${date.toLocaleDateString('en-US', { month: 'short', year: 'numeric' })}`;
                    break;
                case 'year':
                    // Group by year: "2026"
                    key = `${date.getFullYear()}`;
                    break;
                case 'category':
                    // Group by category
                    key = post.category || 'Uncategorized';
                    break;
                default:
                    key = `${date.getFullYear()}`;
            }

            if (!groups[key]) {
                groups[key] = [];
            }
            groups[key].push(post);
        });

        const result = Object.entries(groups)
            .map(([label, posts]) => ({
                label,
                count: posts.length,
                posts: posts.sort((a, b) => new Date(b.publishAt).getTime() - new Date(a.publishAt).getTime()),
            }))
            .sort((a, b) => {
                // Sort descending by extracting numbers or alphabetically
                if (a.label > b.label) return -1;
                if (a.label < b.label) return 1;
                return 0;
            });

        return result;
    }, [posts, activeFilter]);

    const formatDate = (dateStr: string) => {
        const date = new Date(dateStr);
        return `${String(date.getMonth() + 1).padStart(2, '0')}-${String(date.getDate()).padStart(2, '0')}`;
    };

    const filters = [
        { key: 'day', label: 'Day' },
        { key: 'week', label: 'Week' },
        { key: 'month', label: 'Month' },
        { key: 'year', label: 'Year' },
        { key: 'category', label: 'Category' },
    ] as const;

    return (
        <div className="min-h-screen w-full bg-[#fdfaf6] bg-[radial-gradient(ellipse_at_top_left,_var(--tw-gradient-stops))] from-orange-100/40 via-rose-100/20 to-transparent">
            <SEO title="Posts - Voocel Blog" />

            {/* Back Button and Nav */}
            {/* Back Button and Nav */}
            <div className="fixed top-6 left-6 z-50 flex items-center gap-4">
                {/* Animated Navigation - Compact mode with Embedded Back Button */}
                <AnimatedNavWidget isCompact={true} disableFixed={true} showBackButton={true} onBackClick={() => navigate('/')} />
            </div>

            {/* Background Decorations */}
            <div className="fixed top-20 left-10 w-64 h-64 bg-purple-200 rounded-full mix-blend-multiply filter blur-3xl opacity-20 animate-blob pointer-events-none" aria-hidden="true"></div>
            <div className="fixed top-20 right-10 w-64 h-64 bg-orange-200 rounded-full mix-blend-multiply filter blur-3xl opacity-20 animate-blob animation-delay-2000 pointer-events-none" aria-hidden="true"></div>
            <div className="fixed -bottom-8 left-20 w-64 h-64 bg-pink-200 rounded-full mix-blend-multiply filter blur-3xl opacity-20 animate-blob animation-delay-4000 pointer-events-none" aria-hidden="true"></div>

            {/* Main Content - Staggered Animation Container */}
            <motion.div
                className="pt-24 pb-12 px-8 max-w-4xl mx-auto"
                variants={containerVariants}
                initial="hidden"
                animate="visible"
            >

                {/* Filter Tabs - Appears First */}
                <motion.div className="flex justify-center mb-8" variants={filterVariants}>
                    <div className="inline-flex items-center gap-1 bg-white/50 backdrop-blur-md rounded-full p-1 shadow-sm border border-white/50">
                        {filters.map(filter => (
                            <button
                                key={filter.key}
                                onClick={() => setActiveFilter(filter.key)}
                                className={`px-4 py-1.5 rounded-full text-sm font-medium cursor-pointer transition-all duration-200 ${activeFilter === filter.key
                                    ? 'bg-orange-500 text-white shadow-md'
                                    : 'text-stone-500 hover:bg-orange-100/60 hover:text-orange-600'
                                    }`}
                            >
                                {filter.label}
                            </button>
                        ))}
                    </div>
                </motion.div>

                {/* Loading State */}
                {loading && (
                    <motion.div
                        className="flex justify-center py-12"
                        initial={{ opacity: 0 }}
                        animate={{ opacity: 1 }}
                    >
                        <div className="text-stone-400">Loading...</div>
                    </motion.div>
                )}

                {/* Posts List - Grouped by Filter */}
                {!loading && groupedPosts.map((group, groupIndex) => (
                    <motion.div
                        key={group.label}
                        className="mb-8 bg-white/40 backdrop-blur-xl rounded-3xl p-6 shadow-sm border border-white/50"
                        initial={{ opacity: 0, y: 30 }}
                        animate={{ opacity: 1, y: 0 }}
                        transition={{ delay: 0.2 + groupIndex * 0.1, type: 'spring' as const, stiffness: 80, damping: 12 }}
                    >
                        {/* Group Header */}
                        <div className="flex items-center gap-3 mb-6">
                            <h2 className="text-xl font-bold text-stone-800">{group.label}</h2>
                            <span className="text-xs text-stone-400">☆ {group.count} posts</span>
                        </div>

                        {/* Posts */}
                        <div className="space-y-1">
                            {group.posts.map((post, postIndex) => (
                                <motion.div
                                    key={post.id}
                                    onClick={() => navigate(`/post/${post.slug}`)}
                                    className="flex items-start gap-4 py-2.5 px-3 -mx-3 rounded-xl hover:bg-white/70 cursor-pointer transition-all duration-200 group"
                                    role="button"
                                    tabIndex={0}
                                    onKeyDown={(e) => e.key === 'Enter' && navigate(`/post/${post.slug}`)}
                                    initial={{ opacity: 0, x: -10 }}
                                    animate={{ opacity: 1, x: 0 }}
                                    transition={{ delay: 0.3 + postIndex * 0.05 }}
                                    whileHover={{ x: 4, backgroundColor: 'rgba(255,255,255,0.8)' }}
                                >
                                    {/* Date */}
                                    <span className="text-sm text-stone-400 group-hover:text-stone-500 font-mono w-14 shrink-0 transition-colors">
                                        {formatDate(post.publishAt)}
                                    </span>

                                    {/* Indicator */}
                                    <span className="text-orange-400 mt-0.5 group-hover:text-orange-500 transition-colors">•</span>

                                    {/* Title */}
                                    <div className="flex-1 min-w-0">
                                        <h3 className="text-sm text-stone-700 group-hover:text-orange-600 transition-colors truncate font-medium">
                                            {post.title}
                                        </h3>
                                    </div>

                                    {/* Tags */}
                                    <div className="flex gap-1.5 shrink-0">
                                        {post.tags?.slice(0, 2).map(tag => (
                                            <span
                                                key={tag}
                                                className="text-[10px] text-stone-400 group-hover:text-orange-500 bg-stone-100 group-hover:bg-orange-50 px-2 py-0.5 rounded cursor-pointer transition-colors"
                                            >
                                                #{tag}
                                            </span>
                                        ))}
                                    </div>
                                </motion.div>
                            ))}
                        </div>
                    </motion.div>
                ))}

                {/* Empty State */}
                {!loading && posts.length === 0 && (
                    <motion.div
                        className="text-center py-12 text-stone-400"
                        initial={{ opacity: 0 }}
                        animate={{ opacity: 1 }}
                        transition={{ delay: 0.5 }}
                    >
                        No posts yet...
                    </motion.div>
                )}
            </motion.div>
        </div>
    );
};

export default PostsListPage;
