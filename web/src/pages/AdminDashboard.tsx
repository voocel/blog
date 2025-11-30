
import React, { useState } from 'react';
import MDEditor from '@uiw/react-md-editor';
import { useBlog } from '../context/BlogContext';
import type { AdminSection, BlogPost } from '../types';
import { uploadImage } from '../services/uploadService';
import { IconEdit, IconTrash, IconPlus, IconX, IconGrid, IconTag, IconLayers, IconImage, IconClock, IconEye, IconArrowLeft, IconArrowDown, IconChevronLeft, IconChevronRight, IconCopy, IconActivity, IconGlobe, IconMapPin, IconUpload } from '../components/Icons';
import { AUTHOR_NAME } from '../constants';

interface AdminDashboardProps {
    section: AdminSection;
    onExit: () => void;
}

const AdminDashboard: React.FC<AdminDashboardProps> = ({ section, onExit }) => {
    const {
        posts, categories, tags, files, user,
        addPost, updatePost, deletePost,
        addCategory, deleteCategory,
        addTag, deleteTag,
        addFile, deleteFile,
        setActivePostId,
        visitLogs,
        refreshAdminData,
        refreshPosts,
        refreshCategories,
        refreshTags,
        refreshFiles,
        refreshVisitLogs,
        dashboardStats,
        refreshDashboardOverview
    } = useBlog();


    // Refresh data when switching sections
    React.useEffect(() => {
        if (section === 'overview') {
            refreshDashboardOverview();
        } else if (section === 'posts') {
            refreshPosts();
        } else if (section === 'categories') {
            refreshCategories();
        } else if (section === 'tags') {
            refreshTags();
        } else if (section === 'files') {
            refreshFiles();
        } else if (section === 'echoes') {
            refreshVisitLogs();
        }
    }, [section]);

    // --- Shared State ---
    const [isEditorOpen, setIsEditorOpen] = useState(false);

    // --- Post State ---
    const [editingPost, setEditingPost] = useState<Partial<BlogPost> | null>(null);
    const [postSearch, setPostSearch] = useState('');
    const [postFilter, setPostFilter] = useState<'all' | 'published' | 'draft'>('all');
    const [currentPage, setCurrentPage] = useState(1);
    const itemsPerPage = 5;

    // --- Category State ---
    const [newCatName, setNewCatName] = useState('');

    // --- Tag State ---
    const [newTagName, setNewTagName] = useState('');

    // --- File State ---
    const [newFileUrl, setNewFileUrl] = useState('');
    const [lightboxIndex, setLightboxIndex] = useState<number | null>(null);
    const fileInputRef = React.useRef<HTMLInputElement>(null);

    const handleFileUpload = async (event: React.ChangeEvent<HTMLInputElement>) => {
        const file = event.target.files?.[0];
        if (file) {
            try {
                const result = await uploadImage(file);
                addFile({
                    id: `f-${Date.now()}`,
                    url: result.url,
                    name: result.filename,
                    type: 'image', // In real app, use result.type
                    date: new Date().toLocaleDateString()
                });
            } catch (error) {
                console.error("Upload failed:", error);
            }
        }
    };

    // --- Echoes State ---
    const [visibleLogs, setVisibleLogs] = useState(20);

    // --- Confirmation State ---
    const [confirmModal, setConfirmModal] = useState<{
        isOpen: boolean;
        title: string;
        message: string;
        onConfirm: () => void;
    }>({ isOpen: false, title: '', message: '', onConfirm: () => { } });

    const requestConfirm = (title: string, message: string, onConfirm: () => void) => {
        setConfirmModal({ isOpen: true, title, message, onConfirm });
    };

    // --- Handlers ---

    // POSTS
    const handleEditPost = (post?: BlogPost) => {
        if (post) {
            // Ensure tags are IDs when loading a post
            const tagIds = (post.tags || []).map(t => {
                const tagByName = tags.find(tag => tag.name === t);
                return tagByName ? tagByName.id : t;
            });
            setEditingPost({ ...post, tags: tagIds });
        } else {
            setEditingPost({
                // id: Date.now().toString(), // Let backend generate ID
                title: '',
                excerpt: '',
                content: '',
                author: AUTHOR_NAME,
                date: new Date().toLocaleDateString(),
                category: categories[0]?.name || 'General',
                readTime: '5 min read',
                imageUrl: 'https://images.unsplash.com/photo-1618005182384-a83a8bd57fbe?auto=format&fit=crop&w=800&q=80',
                tags: [],
                status: 'draft',
                views: 0
            });
        }
        // Ensure categories and tags are fresh
        refreshCategories();
        refreshTags();
        setIsEditorOpen(true);
    };

    const savePost = async () => {
        if (!editingPost || !editingPost.title) return;

        // Construct payload with IDs
        const payload: any = {
            ...editingPost,
            categoryId: editingPost.categoryId, // Ensure this is set
            tags: (editingPost.tags || []).map(t => {
                const tagByName = tags.find(tag => tag.name === t);
                return tagByName ? tagByName.id : t;
            })
        };

        // If categoryId is missing but we have category name (legacy/fallback), try to find ID
        if (!payload.categoryId && editingPost.category) {
            const cat = categories.find(c => c.name === editingPost.category);
            if (cat) payload.categoryId = cat.id;
        }

        // If categoryId is still missing, default to first category
        if (!payload.categoryId && categories.length > 0) {
            payload.categoryId = categories[0].id;
        }

        try {
            const existing = posts.find(p => p.id === editingPost.id);
            if (existing && editingPost.id) {
                await updatePost(editingPost.id, payload);
            } else {
                await addPost(payload as BlogPost);
            }
            setIsEditorOpen(false);
            setEditingPost(null);
        } catch (error) {
            console.error("Failed to save post:", error);
            alert("Failed to save post. Please try again.\n" + (error instanceof Error ? error.message : String(error)));
        }
    };

    // CATEGORIES
    const handleAddCategory = () => {
        if (!newCatName) return;
        addCategory({
            id: `c-${Date.now()}`,
            name: newCatName,
            slug: newCatName.toLowerCase().replace(/ /g, '-'),
            count: 0
        });
        setNewCatName('');
    };

    // TAGS
    const handleAddTag = () => {
        if (!newTagName) return;
        addTag({
            id: `t-${Date.now()}`,
            name: newTagName
        });
        setNewTagName('');
    };

    // FILES
    const handleAddFile = () => {
        if (!newFileUrl) return;
        addFile({
            id: `f-${Date.now()}`,
            url: newFileUrl,
            name: 'New Upload',
            type: 'image',
            date: new Date().toLocaleDateString()
        });
        setNewFileUrl('');
    };


    // --- Render Functions ---

    const renderOverview = () => {
        const recentPosts = dashboardStats?.recentPosts || [];

        return (
            <div className="p-10 animate-fade-in text-ink max-w-[1600px] mx-auto">
                {/* Welcome Header */}
                <div className="mb-12 flex justify-between items-end">
                    <div>
                        <h1 className="text-4xl md:text-5xl font-serif font-bold text-ink mb-2">Good Morning, {user?.username.split(' ')[0]}.</h1>
                        <p className="text-stone-500 text-lg font-serif italic">The aether awaits your thoughts today.</p>
                    </div>
                    <div className="hidden md:block text-right">
                        <p className="text-xs uppercase tracking-widest text-stone-400 mb-1">Current Date</p>
                        <p className="text-xl font-serif font-bold text-ink">{new Date().toLocaleDateString('en-US', { weekday: 'long', month: 'long', day: 'numeric' })}</p>
                    </div>
                </div>

                {/* Stats Grid */}
                <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6 mb-12">
                    {[
                        { label: 'Total Entries', value: dashboardStats?.counts.posts || 0, icon: IconGrid, color: 'text-gold-600', bg: 'bg-gold-50', border: 'border-gold-100' },
                        { label: 'Categories', value: dashboardStats?.counts.categories || 0, icon: IconLayers, color: 'text-emerald-600', bg: 'bg-emerald-50', border: 'border-emerald-100' },
                        { label: 'Active Tags', value: dashboardStats?.counts.tags || 0, icon: IconTag, color: 'text-teal-600', bg: 'bg-teal-50', border: 'border-teal-100' },
                        { label: 'Media Assets', value: dashboardStats?.counts.files || 0, icon: IconImage, color: 'text-stone-600', bg: 'bg-stone-50', border: 'border-stone-200' }
                    ].map((stat, i) => (
                        <div key={i} className={`bg-white border ${stat.border} p-6 rounded-2xl shadow-sm hover:shadow-md transition-shadow relative overflow-hidden group cursor-default`}>
                            <div className="flex justify-between items-start mb-4 relative z-10">
                                <div className={`p-3 rounded-xl ${stat.bg}`}>
                                    <stat.icon className={`w-6 h-6 ${stat.color}`} />
                                </div>
                            </div>
                            <div className="text-4xl font-serif font-bold text-ink mb-1 relative z-10">{stat.value}</div>
                            <div className="text-xs uppercase tracking-widest text-stone-400 relative z-10">{stat.label}</div>

                            {/* Decorative Background Blob */}
                            <div className={`absolute -right-6 -bottom-6 w-24 h-24 rounded-full opacity-10 group-hover:scale-110 transition-transform ${stat.bg.replace('bg-', 'bg-')}`}></div>
                        </div>
                    ))}
                </div>

                {/* Recent Activity */}
                <div className="grid grid-cols-1 lg:grid-cols-3 gap-8">
                    <div className="lg:col-span-2">
                        <div className="flex justify-between items-center mb-6">
                            <h3 className="text-xl font-serif font-bold text-ink flex items-center gap-2">
                                <span className="w-2 h-8 bg-ink rounded-full"></span>
                                Recent Entries
                            </h3>
                            <button onClick={() => handleEditPost()} className="text-sm text-gold-600 hover:text-gold-700 font-bold uppercase tracking-wider cursor-pointer">+ Quick Draft</button>
                        </div>
                        <div className="space-y-4">
                            {recentPosts.map(post => (
                                <div key={post.id} className="group flex items-center gap-5 bg-white p-4 rounded-xl border border-stone-200 hover:border-gold-300 shadow-sm transition-all cursor-pointer" onClick={() => handleEditPost(post)}>
                                    <div className="w-20 h-20 rounded-lg overflow-hidden bg-stone-100 shrink-0 relative">
                                        <img src={post.imageUrl} alt="" className="w-full h-full object-cover group-hover:scale-105 transition-transform duration-500" />
                                    </div>
                                    <div className="flex-1 min-w-0">
                                        <h4 className="font-serif font-bold text-lg text-ink truncate group-hover:text-gold-600 transition-colors mb-1">{post.title}</h4>
                                        <p className="text-sm text-stone-500 truncate mb-2">{post.excerpt}</p>
                                        <div className="flex items-center gap-3 text-xs text-stone-400">
                                            <span className="flex items-center gap-1"><IconClock className="w-3 h-3" /> {post.date}</span>
                                            <span className="flex items-center gap-1"><IconEye className="w-3 h-3" /> {post.views}</span>
                                        </div>
                                    </div>
                                    <div className="text-right shrink-0">
                                        <span className={`inline-flex items-center px-2.5 py-1 rounded-md text-[10px] uppercase tracking-wider font-bold ${post.status === 'published' ? 'bg-emerald-50 text-emerald-700 border border-emerald-100' : 'bg-stone-100 text-stone-500 border border-stone-200'
                                            }`}>
                                            {post.status}
                                        </span>
                                    </div>
                                </div>
                            ))}
                        </div>
                    </div>

                    <div className="lg:col-span-1">
                        <h3 className="text-xl font-serif font-bold text-ink mb-6 flex items-center gap-2">
                            <span className="w-2 h-8 bg-stone-300 rounded-full"></span>
                            System Status
                        </h3>
                        <div className="bg-white p-8 rounded-2xl border border-stone-200 shadow-sm space-y-8">
                            <div>
                                <div className="flex justify-between text-sm mb-2">
                                    <span className="text-stone-500 font-medium">Storage Usage</span>
                                    <span className="text-ink font-bold">{dashboardStats?.systemStatus.storageUsage || 0}%</span>
                                </div>
                                <div className="w-full h-2.5 bg-stone-100 rounded-full overflow-hidden">
                                    <div className="h-full bg-gradient-to-r from-gold-400 to-gold-600 rounded-full" style={{ width: `${dashboardStats?.systemStatus.storageUsage || 0}%` }}></div>
                                </div>
                            </div>
                            <div>
                                <div className="flex justify-between text-sm mb-2">
                                    <span className="text-stone-500 font-medium">AI Quota (Gemini)</span>
                                    <span className="text-ink font-bold">{dashboardStats?.systemStatus.aiQuota || 0}%</span>
                                </div>
                                <div className="w-full h-2.5 bg-stone-100 rounded-full overflow-hidden">
                                    <div className="h-full bg-gradient-to-r from-teal-400 to-teal-600 rounded-full" style={{ width: `${dashboardStats?.systemStatus.aiQuota || 0}%` }}></div>
                                </div>
                            </div>

                            <div className="p-4 bg-stone-50 rounded-lg border border-stone-100">
                                <div className="flex items-center gap-2 text-emerald-600 mb-1">
                                    <div className="w-2 h-2 rounded-full bg-emerald-500 animate-pulse"></div>
                                    <span className="text-xs font-bold uppercase tracking-wider">Operational</span>
                                </div>
                                <p className="text-xs text-stone-500 leading-relaxed">
                                    System is operating normally. Next automated backup scheduled for 02:00 AM.
                                </p>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        );
    };

    const renderPosts = () => {
        const filteredPosts = posts.filter(post => {
            const matchesSearch = post.title.toLowerCase().includes(postSearch.toLowerCase()) || post.excerpt.toLowerCase().includes(postSearch.toLowerCase());
            const matchesStatus = postFilter === 'all' || post.status === postFilter;
            return matchesSearch && matchesStatus;
        });

        const totalPages = Math.ceil(filteredPosts.length / itemsPerPage);
        const paginatedPosts = filteredPosts.slice((currentPage - 1) * itemsPerPage, currentPage * itemsPerPage);

        const handlePageChange = (page: number) => {
            if (page >= 1 && page <= totalPages) {
                setCurrentPage(page);
                window.scrollTo({ top: 0, behavior: 'smooth' });
            }
        };

        return (
            <div className="p-10 h-full flex flex-col animate-fade-in text-ink max-w-[1600px] mx-auto w-full">
                <div className="flex flex-col md:flex-row justify-between items-start md:items-center mb-10 gap-6">
                    <div>
                        <h2 className="text-4xl font-serif font-bold text-ink mb-2">Journal Entries</h2>
                        <p className="text-stone-500">Curate and manage your published works.</p>
                    </div>
                    <button
                        onClick={() => handleEditPost()}
                        className="bg-ink text-white px-6 py-3 rounded-xl hover:bg-gold-600 transition-colors shadow-lg shadow-stone-200 flex items-center gap-2 group cursor-pointer"
                    >
                        <IconPlus className="w-5 h-5 group-hover:rotate-90 transition-transform" />
                        <span className="font-medium tracking-wide">Create New Entry</span>
                    </button>
                </div>

                {/* Toolbar */}
                <div className="bg-white p-4 rounded-xl border border-stone-200 shadow-sm mb-6 flex flex-col md:flex-row gap-4 justify-between items-center">
                    {/* Search */}
                    <div className="relative w-full md:w-96">
                        <input
                            type="text"
                            placeholder="Search by title or content..."
                            value={postSearch}
                            onChange={(e) => setPostSearch(e.target.value)}
                            className="w-full pl-4 pr-10 py-2.5 bg-stone-50 border border-stone-200 rounded-lg text-sm focus:outline-none focus:border-gold-500 focus:ring-1 focus:ring-gold-500 transition-all"
                        />
                        {postSearch && (
                            <button onClick={() => setPostSearch('')} className="absolute right-3 top-1/2 -translate-y-1/2 text-stone-400 hover:text-ink cursor-pointer">
                                <IconX className="w-4 h-4" />
                            </button>
                        )}
                    </div>

                    {/* Filter Tabs */}
                    <div className="flex bg-stone-100 p-1 rounded-lg">
                        {(['all', 'published', 'draft'] as const).map(status => (
                            <button
                                key={status}
                                onClick={() => setPostFilter(status)}
                                className={`px-4 py-1.5 rounded-md text-xs uppercase tracking-wider font-bold transition-all cursor-pointer ${postFilter === status
                                    ? 'bg-white text-ink shadow-sm'
                                    : 'text-stone-500 hover:text-stone-700'
                                    }`}
                            >
                                {status}
                            </button>
                        ))}
                    </div>
                </div>

                {/* List */}
                <div className="space-y-4 pb-20">
                    {paginatedPosts.length > 0 ? (
                        paginatedPosts.map(post => (
                            <div key={post.id} className="group bg-white rounded-xl border border-stone-200 p-5 flex items-center gap-6 hover:shadow-md hover:border-gold-300 transition-all">
                                {/* Status Indicator */}
                                <div className={`w-1.5 h-full min-h-[4rem] rounded-full ${post.status === 'published' ? 'bg-emerald-400' : 'bg-stone-300'}`} />

                                {/* Image */}
                                <div className="w-28 h-20 rounded-lg bg-stone-100 overflow-hidden shrink-0 relative border border-stone-100 cursor-pointer" onClick={() => { setActivePostId(post.id); onExit(); }}>
                                    <img src={post.imageUrl} className="w-full h-full object-cover group-hover:scale-105 transition-transform duration-500" alt="" />
                                </div>

                                {/* Content */}
                                <div className="flex-1 min-w-0 py-1">
                                    <div className="flex items-center gap-3 mb-1.5">
                                        <h3 onClick={() => { setActivePostId(post.id); onExit(); }} className="text-xl font-serif font-bold text-ink truncate hover:text-gold-600 transition-colors cursor-pointer">{post.title}</h3>
                                        <span className="text-[10px] uppercase tracking-wider text-stone-500 bg-stone-100 px-2 py-0.5 rounded border border-stone-200">{post.category}</span>
                                    </div>
                                    <div className="flex items-center gap-6 text-xs text-stone-400 font-medium">
                                        <span className="flex items-center gap-1.5"><IconClock className="w-3.5 h-3.5" /> {post.date}</span>
                                        <span className="flex items-center gap-1.5"><IconEye className="w-3.5 h-3.5" /> {post.views.toLocaleString()} reads</span>
                                        <div className="flex gap-2">
                                            {post.tags.slice(0, 3).map(t => <span key={t} className="text-stone-300">#{t}</span>)}
                                        </div>
                                    </div>
                                </div>

                                {/* Actions */}
                                <div className="flex items-center gap-2 opacity-0 group-hover:opacity-100 transition-opacity translate-x-4 group-hover:translate-x-0 duration-300">
                                    <button onClick={(e) => { e.stopPropagation(); handleEditPost(post); }} className="flex items-center gap-2 px-4 py-2 bg-stone-50 text-stone-600 hover:text-gold-600 hover:bg-gold-50 rounded-lg transition-colors border border-stone-200 cursor-pointer" title="Edit">
                                        <IconEdit className="w-4 h-4" />
                                        <span className="text-xs font-bold uppercase tracking-wider">Edit</span>
                                    </button>
                                    <button onClick={(e) => { e.stopPropagation(); requestConfirm('Delete Entry', 'Are you sure you want to delete this entry? This action cannot be undone.', () => deletePost(post.id)); }} className="p-2.5 text-stone-400 hover:text-red-600 hover:bg-red-50 rounded-lg transition-colors border border-transparent hover:border-red-100 cursor-pointer" title="Delete">
                                        <IconTrash className="w-4 h-4" />
                                    </button>
                                </div>
                            </div>
                        ))
                    ) : (
                        <div className="text-center py-20 bg-stone-50 rounded-xl border border-dashed border-stone-300">
                            <p className="text-stone-400 font-serif italic text-lg">No entries match your search criteria.</p>
                            <button onClick={() => { setPostSearch(''); setPostFilter('all'); }} className="mt-4 text-sm text-gold-600 hover:underline cursor-pointer">Clear Filters</button>
                        </div>
                    )}
                </div>

                {/* Pagination */}
                {totalPages > 1 && (
                    <div className="flex justify-center items-center gap-4 pb-20">
                        <button
                            onClick={() => handlePageChange(currentPage - 1)}
                            disabled={currentPage === 1}
                            className="flex items-center gap-2 text-stone-400 hover:text-ink disabled:opacity-30 disabled:cursor-not-allowed cursor-pointer transition-colors uppercase tracking-widest text-xs"
                        >
                            <IconArrowLeft className="w-4 h-4" />
                            Previous
                        </button>

                        <div className="flex items-center gap-2">
                            {Array.from({ length: totalPages }, (_, i) => i + 1).map(page => (
                                <button
                                    key={page}
                                    onClick={() => handlePageChange(page)}
                                    className={`w-8 h-8 flex items-center justify-center rounded-full transition-all leading-none cursor-pointer ${currentPage === page
                                        ? 'bg-gold-600 text-white shadow-sm'
                                        : 'text-stone-500 hover:bg-stone-100'
                                        }`}
                                >
                                    {page}
                                </button>
                            ))}
                        </div>

                        <button
                            onClick={() => handlePageChange(currentPage + 1)}
                            disabled={currentPage === totalPages}
                            className="flex items-center gap-2 text-stone-400 hover:text-ink disabled:opacity-30 disabled:cursor-not-allowed cursor-pointer transition-colors uppercase tracking-widest text-xs"
                        >
                            Next
                            <IconArrowDown className="w-4 h-4 -rotate-90" />
                        </button>
                    </div>
                )}
            </div>
        );
    };

    const renderCategories = () => (
        <div className="p-10 animate-fade-in text-ink max-w-[1200px] mx-auto w-full">
            <h2 className="text-4xl font-serif font-bold text-ink mb-2">Taxonomy</h2>
            <p className="text-stone-500 mb-10">Organize your content structure.</p>

            <div className="grid grid-cols-1 md:grid-cols-3 gap-12">
                <div className="md:col-span-1">
                    <div className="bg-white border border-stone-200 p-8 rounded-2xl shadow-sm sticky top-10">
                        <h3 className="font-serif font-bold text-xl mb-6">Add Category</h3>
                        <div className="space-y-4">
                            <div>
                                <label className="block text-xs uppercase tracking-widest text-stone-400 mb-2">Name</label>
                                <input
                                    value={newCatName}
                                    onChange={e => setNewCatName(e.target.value)}
                                    placeholder="e.g., Philosophy"
                                    className="w-full bg-stone-50 border border-stone-200 rounded-lg px-4 py-3 text-ink focus:outline-none focus:border-gold-500 transition-colors"
                                />
                            </div>
                            <div className="p-4 bg-stone-50 rounded-lg border border-stone-100">
                                <p className="text-xs text-stone-500 leading-relaxed">
                                    <span className="font-bold">Tip:</span> Use broad, high-level topics for categories. Use tags for specific details.
                                </p>
                            </div>
                            <button
                                onClick={handleAddCategory}
                                disabled={!newCatName}
                                className="w-full bg-emerald-600 text-white py-3.5 rounded-xl font-bold tracking-wide hover:bg-emerald-700 transition-colors disabled:opacity-50 shadow-emerald-100 shadow-lg cursor-pointer"
                            >
                                Create Category
                            </button>
                        </div>
                    </div>
                </div>

                <div className="md:col-span-2 space-y-3">
                    {categories.map(cat => (
                        <div key={cat.id} className="group flex justify-between items-center bg-white border border-stone-200 p-6 rounded-xl hover:border-emerald-400 hover:shadow-md transition-all">
                            <div className="flex items-center gap-4">
                                <div className="w-10 h-10 rounded-full bg-emerald-50 flex items-center justify-center text-emerald-600 font-serif font-bold text-lg border border-emerald-100">
                                    {cat.name.charAt(0)}
                                </div>
                                <div>
                                    <span className="text-lg font-serif font-bold text-ink block group-hover:text-emerald-700 transition-colors">{cat.name}</span>
                                    <span className="text-xs text-stone-400 font-mono tracking-tight">{cat.slug} • {cat.count} posts</span>
                                </div>
                            </div>
                            <button onClick={() => requestConfirm('Delete Category', `Are you sure you want to delete "${cat.name}"?`, () => deleteCategory(cat.id))} className="p-2 text-stone-300 hover:text-red-500 hover:bg-red-50 rounded-lg opacity-0 group-hover:opacity-100 transition-all cursor-pointer">
                                <IconTrash className="w-5 h-5" />
                            </button>
                        </div>
                    ))}
                </div>
            </div>
        </div>
    );

    const renderTags = () => (
        <div className="p-10 animate-fade-in text-ink max-w-[1200px] mx-auto w-full">
            <h2 className="text-4xl font-serif font-bold text-ink mb-2">Topics & Tags</h2>
            <p className="text-stone-500 mb-10">Micro-categorization for your posts.</p>

            <div className="bg-white border border-stone-200 rounded-2xl p-10 shadow-sm mb-10">
                <div className="flex gap-4 mb-10 max-w-xl border-b border-stone-100 pb-10">
                    <input
                        value={newTagName}
                        onChange={e => setNewTagName(e.target.value)}
                        placeholder="New Tag Name"
                        className="flex-1 bg-stone-50 border border-stone-200 rounded-xl px-5 py-3 text-ink focus:outline-none focus:border-teal-500"
                    />
                    <button onClick={handleAddTag} className="bg-teal-600 text-white px-8 py-3 rounded-xl font-bold hover:bg-teal-700 transition-colors shadow-lg shadow-teal-100 cursor-pointer">Add Tag</button>
                </div>

                <div className="flex flex-wrap gap-3">
                    {tags.map(tag => (
                        <div key={tag.id} className="group flex items-center gap-2 bg-white border border-stone-200 px-4 py-2.5 rounded-full hover:border-teal-400 hover:shadow-sm transition-all cursor-default">
                            <span className="text-stone-400">#</span>
                            <span className="text-stone-700 font-medium">{tag.name}</span>
                            <button onClick={() => requestConfirm('Delete Tag', `Are you sure you want to delete #${tag.name}?`, () => deleteTag(tag.id))} className="text-stone-300 hover:text-red-500 ml-2 p-0.5 rounded-full hover:bg-red-50 cursor-pointer">
                                <IconX className="w-3 h-3" />
                            </button>
                        </div>
                    ))}
                </div>
            </div>
        </div>
    );

    const renderFiles = () => (
        <div className="p-10 h-full flex flex-col animate-fade-in text-ink w-full max-w-[1600px] mx-auto">
            <div className="flex justify-between items-end mb-10">
                <div>
                    <h2 className="text-4xl font-serif font-bold text-ink mb-2">Media Assets</h2>
                    <p className="text-stone-500">Library of uploaded images and documents.</p>
                </div>
                <div className="flex gap-3 bg-white p-2 rounded-xl border border-stone-200 shadow-sm">
                    <input
                        type="file"
                        ref={fileInputRef}
                        onChange={handleFileUpload}
                        className="hidden"
                        accept="image/*"
                    />
                    <button
                        onClick={() => fileInputRef.current?.click()}
                        className="p-2 bg-stone-100 hover:bg-stone-200 text-stone-600 rounded-lg border border-stone-200 transition-colors"
                        title="Upload Local File"
                    >
                        <IconUpload className="w-5 h-5" />
                    </button>
                    <div className="w-px bg-stone-200 my-1"></div>
                    <input
                        value={newFileUrl}
                        onChange={e => setNewFileUrl(e.target.value)}
                        placeholder="Paste Image URL..."
                        className="w-64 bg-stone-50 rounded-lg px-4 py-2 text-sm focus:outline-none border border-transparent focus:border-stone-300 transition-colors"
                    />
                    <button onClick={handleAddFile} className="bg-ink text-white px-6 py-2 rounded-lg text-sm font-bold uppercase tracking-wider hover:bg-stone-800 transition-colors cursor-pointer">Import</button>
                </div>
            </div>

            <div className="grid grid-cols-2 md:grid-cols-4 lg:grid-cols-5 xl:grid-cols-6 gap-6 overflow-y-auto pb-10">
                {files.map((file, index) => (
                    <div key={file.id} className="group relative aspect-square bg-white border border-stone-200 p-3 rounded-2xl shadow-sm hover:shadow-lg transition-all hover:-translate-y-1">
                        <div className="w-full h-full rounded-xl overflow-hidden relative bg-stone-100 cursor-pointer" onClick={() => setLightboxIndex(index)}>
                            <img src={file.url} alt={file.name} className="w-full h-full object-cover" />

                            {/* Overlay Actions */}
                            <div className="absolute inset-0 bg-stone-900/40 opacity-0 group-hover:opacity-100 transition-all duration-300 flex flex-col items-center justify-center gap-3 backdrop-blur-[2px]">
                                <div className="flex gap-2">
                                    <button
                                        onClick={(e) => { e.stopPropagation(); navigator.clipboard.writeText(file.url); }}
                                        className="p-2 bg-white/20 text-white rounded-full hover:bg-white hover:text-stone-900 transition-colors border border-white/30 backdrop-blur-md cursor-pointer"
                                        title="Copy URL"
                                    >
                                        <IconCopy className="w-4 h-4" />
                                    </button>
                                    <button
                                        onClick={(e) => { e.stopPropagation(); requestConfirm('Delete Asset', 'Are you sure you want to delete this file?', () => deleteFile(file.id)); }}
                                        className="p-2 bg-red-500/20 text-red-100 rounded-full hover:bg-red-500 hover:text-white transition-colors border border-red-500/30 backdrop-blur-md cursor-pointer"
                                        title="Delete"
                                    >
                                        <IconTrash className="w-4 h-4" />
                                    </button>
                                </div>
                                <span className="text-white text-[10px] uppercase tracking-widest font-bold opacity-80">Click to View</span>
                            </div>
                        </div>
                    </div>
                ))}
            </div>
        </div>
    );

    const renderEchoes = () => {
        // --- Stats Calculation ---
        const now = new Date();
        const startOfDay = new Date(now.getFullYear(), now.getMonth(), now.getDate()).getTime();
        const startOfMonth = new Date(now.getFullYear(), now.getMonth(), 1).getTime();

        // Resonances (Page Views)
        const totalResonances = visitLogs.length;
        const monthResonances = visitLogs.filter(l => l.timestamp >= startOfMonth).length;
        const dayResonances = visitLogs.filter(l => l.timestamp >= startOfDay).length;

        // Wanderers (Unique Visitors)
        const uniqueVisitorsTotal = new Set(visitLogs.map(log => log.ip)).size;
        const uniqueVisitorsMonth = new Set(visitLogs.filter(l => l.timestamp >= startOfMonth).map(l => l.ip)).size;
        const uniqueVisitorsDay = new Set(visitLogs.filter(l => l.timestamp >= startOfDay).map(l => l.ip)).size;

        // Top Posts (Today Only)
        const today = new Date().toDateString();
        const topPosts = visitLogs
            .filter(log => log.postId && new Date(log.timestamp).toDateString() === today)
            .reduce((acc, log) => {
                const id = log.postId!;
                acc[id] = (acc[id] || 0) + 1;
                return acc;
            }, {} as Record<string, number>);

        const sortedTopPosts = Object.entries(topPosts)
            .sort(([, a], [, b]) => b - a)
            .slice(0, 5)
            .map(([id, count]) => {
                const post = posts.find(p => p.id === id);
                return { ...post, visitCount: count };
            });

        // Infinite Scroll Handler
        const handleStreamScroll = (e: React.UIEvent<HTMLDivElement>) => {
            const { scrollTop, scrollHeight, clientHeight } = e.currentTarget;
            if (scrollHeight - scrollTop <= clientHeight + 50) {
                setVisibleLogs(prev => Math.min(prev + 20, visitLogs.length));
            }
        };

        return (
            <div className="p-10 h-full flex flex-col animate-fade-in text-ink max-w-[1600px] mx-auto w-full">
                <div className="mb-10">
                    <h2 className="text-4xl font-serif font-bold text-ink mb-2">Echoes</h2>
                    <p className="text-stone-500">Resonances left by wanderers in the aether.</p>
                </div>

                {/* Hero Stats (Redesigned) */}
                <div className="grid grid-cols-1 md:grid-cols-3 gap-6 mb-12">
                    {/* Resonances Card */}
                    <div className="bg-white border border-indigo-100 p-8 rounded-2xl shadow-sm relative overflow-hidden group">
                        <div className="flex justify-between items-start mb-6 relative z-10">
                            <div className="p-3 rounded-xl bg-indigo-50 text-indigo-600">
                                <IconActivity className="w-6 h-6" />
                            </div>
                            <span className="text-xs font-bold uppercase tracking-wider text-indigo-200">Views</span>
                        </div>
                        <div className="relative z-10">
                            <div className="text-4xl font-serif font-bold text-ink mb-1">{totalResonances.toLocaleString()}</div>
                            <div className="text-xs uppercase tracking-widest text-stone-400 mb-6">Total Resonances</div>

                            <div className="grid grid-cols-2 gap-4 border-t border-indigo-50 pt-4">
                                <div>
                                    <div className="text-lg font-bold text-ink">{monthResonances.toLocaleString()}</div>
                                    <div className="text-[10px] uppercase tracking-wider text-stone-400">This Month</div>
                                </div>
                                <div>
                                    <div className="text-lg font-bold text-indigo-600">{dayResonances.toLocaleString()}</div>
                                    <div className="text-[10px] uppercase tracking-wider text-stone-400">Today</div>
                                </div>
                            </div>
                        </div>
                        <div className="absolute -right-6 -bottom-6 w-32 h-32 rounded-full opacity-5 bg-indigo-600 group-hover:scale-110 transition-transform"></div>
                    </div>

                    {/* Wanderers Card */}
                    <div className="bg-white border border-violet-100 p-8 rounded-2xl shadow-sm relative overflow-hidden group">
                        <div className="flex justify-between items-start mb-6 relative z-10">
                            <div className="p-3 rounded-xl bg-violet-50 text-violet-600">
                                <IconGlobe className="w-6 h-6" />
                            </div>
                            <span className="text-xs font-bold uppercase tracking-wider text-violet-200">Visitors</span>
                        </div>
                        <div className="relative z-10">
                            <div className="text-4xl font-serif font-bold text-ink mb-1">{uniqueVisitorsTotal.toLocaleString()}</div>
                            <div className="text-xs uppercase tracking-widest text-stone-400 mb-6">Unique Wanderers</div>

                            <div className="grid grid-cols-2 gap-4 border-t border-violet-50 pt-4">
                                <div>
                                    <div className="text-lg font-bold text-ink">{uniqueVisitorsMonth.toLocaleString()}</div>
                                    <div className="text-[10px] uppercase tracking-wider text-stone-400">This Month</div>
                                </div>
                                <div>
                                    <div className="text-lg font-bold text-violet-600">{uniqueVisitorsDay.toLocaleString()}</div>
                                    <div className="text-[10px] uppercase tracking-wider text-stone-400">Today</div>
                                </div>
                            </div>
                        </div>
                        <div className="absolute -right-6 -bottom-6 w-32 h-32 rounded-full opacity-5 bg-violet-600 group-hover:scale-110 transition-transform"></div>
                    </div>

                    {/* Most Visited Card (Trending) */}
                    <div className="bg-white border border-fuchsia-100 p-8 rounded-2xl shadow-sm relative overflow-hidden group">
                        <div className="flex justify-between items-start mb-6 relative z-10">
                            <div className="p-3 rounded-xl bg-fuchsia-50 text-fuchsia-600">
                                <IconMapPin className="w-6 h-6" />
                            </div>
                            <span className="text-xs font-bold uppercase tracking-wider text-fuchsia-200">Trending</span>
                        </div>
                        <div className="relative z-10">
                            <div className="text-xl font-serif font-bold text-ink mb-1 truncate leading-tight">
                                {sortedTopPosts[0]?.title || 'No Activity Yet'}
                            </div>
                            <div className="text-xs uppercase tracking-widest text-stone-400 mb-6">Most Resonant Today</div>

                            <div className="flex items-center gap-2 border-t border-fuchsia-50 pt-4">
                                <div className="text-3xl font-bold text-fuchsia-600">{sortedTopPosts[0]?.visitCount || 0}</div>
                                <div className="text-xs text-stone-400 leading-tight">
                                    views<br />today
                                </div>
                            </div>
                        </div>
                        <div className="absolute -right-6 -bottom-6 w-32 h-32 rounded-full opacity-5 bg-fuchsia-600 group-hover:scale-110 transition-transform"></div>
                    </div>
                </div>

                <div className="grid grid-cols-1 lg:grid-cols-3 gap-8">
                    {/* The Stream */}
                    <div className="lg:col-span-2">
                        <h3 className="text-xl font-serif font-bold text-ink mb-6 flex items-center gap-2">
                            <span className="w-2 h-8 bg-indigo-500 rounded-full"></span>
                            The Stream
                        </h3>
                        <div
                            className="bg-white rounded-2xl border border-stone-200 shadow-sm overflow-hidden flex flex-col max-h-[600px]"
                        >
                            {visitLogs.length > 0 ? (
                                <div
                                    className="divide-y divide-stone-100 overflow-y-auto custom-scrollbar"
                                    onScroll={handleStreamScroll}
                                >
                                    {visitLogs.slice(0, visibleLogs).map((log) => (
                                        <div key={log.id} className="p-5 hover:bg-stone-50 transition-colors flex items-center gap-4 group animate-fade-in">
                                            <div className="w-10 h-10 rounded-full bg-stone-100 flex items-center justify-center text-stone-400 shrink-0">
                                                <IconGlobe className="w-5 h-5" />
                                            </div>
                                            <div className="flex-1 min-w-0">
                                                <p className="text-sm text-ink">
                                                    <span className="font-bold text-indigo-600">Visitor</span> from <span className="font-medium">{log.location}</span>
                                                </p>
                                                <p className="text-sm font-serif italic text-stone-500 truncate">
                                                    Reading: {log.postTitle ? `"${log.postTitle}"` : log.pagePath === '/' ? "The Homepage" : log.pagePath}
                                                </p>
                                            </div>
                                            <div className="text-xs text-stone-400 font-mono text-right">
                                                <div className="font-bold text-stone-500">{log.ip}</div>
                                                <div>{new Date(log.timestamp).toLocaleString(undefined, { month: 'numeric', day: 'numeric', hour: '2-digit', minute: '2-digit' })}</div>
                                            </div>
                                        </div>
                                    ))}
                                    {visibleLogs < visitLogs.length && (
                                        <div className="p-4 text-center text-xs text-stone-400 italic">
                                            Loading more echoes...
                                        </div>
                                    )}
                                </div>
                            ) : (
                                <div className="p-10 text-center text-stone-400 italic font-serif">
                                    No echoes recorded yet...
                                </div>
                            )}
                        </div>
                    </div>

                    {/* Top Stories */}
                    <div className="lg:col-span-1">
                        <h3 className="text-xl font-serif font-bold text-ink mb-6 flex items-center gap-2">
                            <span className="w-2 h-8 bg-fuchsia-500 rounded-full"></span>
                            Most Resonant
                        </h3>
                        <div className="space-y-4">
                            {sortedTopPosts.map((post, index) => post.id ? (
                                <div key={post.id} className="bg-white p-4 rounded-xl border border-stone-200 shadow-sm flex items-center gap-4 group hover:border-fuchsia-300 transition-colors">
                                    <div className="text-2xl font-serif font-bold text-stone-200 group-hover:text-fuchsia-500 transition-colors w-8 text-center">
                                        {index + 1}
                                    </div>
                                    <div className="flex-1 min-w-0">
                                        <h4 className="font-bold text-ink truncate text-sm mb-1">{post.title}</h4>
                                        <p className="text-xs text-stone-500">{post.visitCount} views</p>
                                    </div>
                                </div>
                            ) : null)}
                        </div>
                    </div>
                </div>
            </div>
        );
    };

    // --- ZEN EDITOR ---
    const renderEditor = () => {
        if (!editingPost) return null;
        return (
            <div className="fixed inset-0 z-[100] bg-[#FDFBF7] flex flex-col animate-slide-up">
                {/* Toolbar */}
                <div className="h-20 border-b border-stone-200 flex justify-between items-center px-8 bg-white/90 backdrop-blur-md shadow-sm z-20">
                    <div className="flex items-center gap-6">
                        <button onClick={() => setIsEditorOpen(false)} className="text-stone-400 hover:text-ink flex items-center gap-2 transition-colors cursor-pointer">
                            <IconX className="w-6 h-6" />
                        </button>
                        <div className="h-6 w-px bg-stone-200"></div>
                        <span className="font-serif italic text-stone-400 text-lg">
                            {editingPost.id ? 'Editing Entry' : 'New Entry'}
                        </span>
                    </div>

                    <div className="flex gap-4 items-center">
                        <div className="flex items-center gap-2 bg-stone-100 rounded-lg p-1.5 mr-4">
                            {['draft', 'published'].map(s => (
                                <button
                                    key={s}
                                    onClick={() => setEditingPost({ ...editingPost, status: s as any })}
                                    className={`px-4 py-2 rounded-md text-xs uppercase tracking-wider font-bold transition-all cursor-pointer ${editingPost.status === s ? 'bg-white shadow-sm text-ink' : 'text-stone-400 hover:text-stone-600'
                                        }`}
                                >
                                    {s}
                                </button>
                            ))}
                        </div>
                        <button
                            onClick={savePost}
                            className="bg-ink text-white px-8 py-3 rounded-xl font-bold tracking-wide hover:bg-gold-600 transition-colors shadow-lg flex items-center gap-2 cursor-pointer"
                        >
                            Save Changes
                        </button>
                    </div>
                </div>

                {/* Main Editor Area */}
                <div className="flex-1 flex overflow-hidden">
                    {/* Meta Sidebar */}
                    <div className="w-96 border-r border-stone-200 bg-stone-50 p-8 overflow-y-auto hidden lg:block custom-scrollbar">
                        <h3 className="font-serif font-bold text-ink mb-8 text-xl">Entry Metadata</h3>

                        <div className="space-y-8">
                            <div>
                                <label className="block text-xs uppercase tracking-widest text-stone-500 mb-3 font-bold">Category</label>
                                <div className="relative">
                                    <select
                                        className="w-full bg-white border border-stone-200 rounded-xl p-3 text-sm focus:outline-none focus:border-gold-500 appearance-none shadow-sm font-serif"
                                        value={editingPost.categoryId || (categories.find(c => c.name === editingPost.category)?.id) || ''}
                                        onChange={e => {
                                            const cat = categories.find(c => c.id === e.target.value);
                                            setEditingPost({
                                                ...editingPost,
                                                categoryId: e.target.value,
                                                category: cat ? cat.name : ''
                                            });
                                        }}
                                    >
                                        <option value="" disabled>Select a category</option>
                                        {categories.map(c => <option key={c.id} value={c.id}>{c.name}</option>)}
                                    </select>
                                    <div className="absolute right-3 top-1/2 -translate-y-1/2 pointer-events-none">
                                        <IconGrid className="w-4 h-4 text-stone-400" />
                                    </div>
                                </div>
                            </div>

                            <div>
                                <label className="block text-xs uppercase tracking-widest text-stone-500 mb-3 font-bold">Tags</label>
                                <div className="flex flex-wrap gap-2">
                                    {tags.map(tag => {
                                        const isSelected = editingPost.tags?.includes(tag.id) || editingPost.tags?.includes(tag.name); // Handle both ID and Name for compatibility
                                        return (
                                            <button
                                                key={tag.id}
                                                onClick={() => {
                                                    const currentTags = editingPost.tags || [];
                                                    // Check if tag is already selected (by ID or Name)
                                                    const alreadySelected = currentTags.includes(tag.id) || currentTags.includes(tag.name);

                                                    let newTags;
                                                    if (alreadySelected) {
                                                        // Remove
                                                        newTags = currentTags.filter(t => t !== tag.id && t !== tag.name);
                                                    } else {
                                                        // Add ID
                                                        newTags = [...currentTags, tag.id];
                                                    }
                                                    setEditingPost({ ...editingPost, tags: newTags });
                                                }}
                                                className={`px-3 py-1.5 rounded-full text-xs font-medium transition-all border ${isSelected
                                                    ? 'bg-teal-50 border-teal-200 text-teal-700'
                                                    : 'bg-white border-stone-200 text-stone-500 hover:border-stone-300'
                                                    }`}
                                            >
                                                {tag.name}
                                            </button>
                                        );
                                    })}
                                    {tags.length === 0 && <span className="text-xs text-stone-400 italic">No tags available. Create some in the Tags section.</span>}
                                </div>
                            </div>

                            <div>
                                <label className="block text-xs uppercase tracking-widest text-stone-500 mb-3 font-bold">Excerpt / Summary</label>
                                <textarea
                                    className="w-full bg-white border border-stone-200 rounded-xl p-4 text-sm h-40 resize-none focus:outline-none focus:border-gold-500 shadow-sm leading-relaxed"
                                    value={editingPost.excerpt}
                                    onChange={e => setEditingPost({ ...editingPost, excerpt: e.target.value })}
                                    placeholder="Write a short summary for the feed display..."
                                />
                            </div>

                            <div>
                                <label className="block text-xs uppercase tracking-widest text-stone-500 mb-3 font-bold">Cover Image URL</label>
                                <div className="w-full h-48 bg-stone-200 rounded-xl mb-3 overflow-hidden border border-stone-300 relative group">
                                    <img src={editingPost.imageUrl} className="w-full h-full object-cover opacity-90 group-hover:scale-105 transition-transform duration-700" alt="Cover" />
                                </div>
                                <input
                                    className="w-full bg-white border border-stone-200 rounded-xl p-3 text-xs font-mono focus:outline-none focus:border-gold-500 shadow-sm"
                                    value={editingPost.imageUrl}
                                    onChange={e => setEditingPost({ ...editingPost, imageUrl: e.target.value })}
                                    placeholder="https://..."
                                />
                            </div>
                        </div>
                    </div>

                    {/* Writing Canvas */}
                    <div className="flex-1 overflow-y-auto bg-[#FDFBF7]">
                        <div className="max-w-4xl mx-auto py-20 px-12 h-full flex flex-col">
                            <input
                                className="w-full text-3xl md:text-5xl font-serif font-bold text-ink bg-transparent border-none focus:outline-none focus:ring-0 placeholder-stone-300 leading-tight mb-8 tracking-tight caret-gold-500"
                                placeholder="Untitled Entry"
                                value={editingPost.title}
                                onChange={e => setEditingPost({ ...editingPost, title: e.target.value })}
                            />

                            <div className="flex-1 min-h-[500px]" data-color-mode="light">
                                <MDEditor
                                    value={editingPost.content}
                                    onChange={(val) => setEditingPost({ ...editingPost, content: val || '' })}
                                    preview="edit"
                                    height="100%"
                                    visibleDragbar={false}
                                    textareaProps={{
                                        placeholder: "Start writing your thoughts..."
                                    }}
                                    style={{
                                        backgroundColor: 'transparent',
                                        color: '#44403c', // stone-700
                                        fontFamily: 'Inter, sans-serif',
                                        boxShadow: 'none'
                                    }}
                                />
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        );
    };

    const renderLightbox = () => {
        if (lightboxIndex === null) return null;
        const file = files[lightboxIndex];

        const handleNext = (e: React.MouseEvent) => {
            e.stopPropagation();
            setLightboxIndex((lightboxIndex + 1) % files.length);
        };

        const handlePrev = (e: React.MouseEvent) => {
            e.stopPropagation();
            setLightboxIndex((lightboxIndex - 1 + files.length) % files.length);
        };

        return (
            <div className="fixed inset-0 z-[200] bg-stone-900/95 backdrop-blur-md flex items-center justify-center animate-fade-in" onClick={() => setLightboxIndex(null)}>
                {/* Close Button */}
                <button
                    onClick={() => setLightboxIndex(null)}
                    className="absolute top-6 right-6 p-2 text-white/50 hover:text-white transition-colors cursor-pointer"
                >
                    <IconX className="w-8 h-8" />
                </button>

                {/* Navigation */}
                {files.length > 1 && (
                    <>
                        <button
                            onClick={handlePrev}
                            className="absolute left-6 p-4 text-white/50 hover:text-white transition-colors hover:bg-white/10 rounded-full cursor-pointer"
                        >
                            <IconChevronLeft className="w-8 h-8" />
                        </button>
                        <button
                            onClick={handleNext}
                            className="absolute right-6 p-4 text-white/50 hover:text-white transition-colors hover:bg-white/10 rounded-full cursor-pointer"
                        >
                            <IconChevronRight className="w-8 h-8" />
                        </button>
                    </>
                )}

                {/* Image */}
                <div className="max-w-[90vw] max-h-[90vh] relative" onClick={e => e.stopPropagation()}>
                    <img
                        src={file.url}
                        alt={file.name}
                        className="max-w-full max-h-[90vh] object-contain rounded-lg shadow-2xl"
                    />
                    <div className="absolute -bottom-12 left-0 w-full text-center">
                        <p className="text-white/80 font-serif text-lg">{file.name}</p>
                        <p className="text-white/40 text-xs uppercase tracking-widest">{lightboxIndex + 1} / {files.length}</p>
                    </div>
                </div>
            </div>
        );
    };

    return (
        <div className="min-h-full bg-[#FDFBF7]">
            {section === 'overview' && renderOverview()}
            {section === 'echoes' && renderEchoes()}
            {section === 'posts' && renderPosts()}
            {section === 'categories' && renderCategories()}
            {section === 'tags' && renderTags()}
            {section === 'files' && renderFiles()}
            {isEditorOpen && renderEditor()}
            {renderLightbox()}

            {/* Confirmation Modal */}
            {confirmModal.isOpen && (
                <div className="fixed inset-0 z-[300] flex items-center justify-center p-4">
                    <div className="absolute inset-0 bg-stone-900/50 backdrop-blur-sm" onClick={() => setConfirmModal({ ...confirmModal, isOpen: false })} />
                    <div className="relative bg-white rounded-2xl shadow-2xl max-w-md w-full p-6 animate-slide-up">
                        <h3 className="text-xl font-serif font-bold text-ink mb-2">{confirmModal.title}</h3>
                        <p className="text-stone-500 mb-6">{confirmModal.message}</p>
                        <div className="flex justify-end gap-3">
                            <button
                                onClick={() => setConfirmModal({ ...confirmModal, isOpen: false })}
                                className="px-4 py-2 text-stone-500 hover:text-ink font-medium cursor-pointer"
                            >
                                Cancel
                            </button>
                            <button
                                onClick={() => { confirmModal.onConfirm(); setConfirmModal({ ...confirmModal, isOpen: false }); }}
                                className="px-6 py-2 bg-red-600 text-white rounded-lg font-bold hover:bg-red-700 transition-colors shadow-lg shadow-red-100 cursor-pointer"
                            >
                                Confirm Delete
                            </button>
                        </div>
                    </div>
                </div>
            )}
        </div>
    );
};

export default AdminDashboard;
