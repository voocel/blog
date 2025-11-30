import React, { useState } from 'react';
import MDEditor from '@uiw/react-md-editor';
import { useBlog } from '../context/BlogContext';
import type { AdminSection, BlogPost } from '../types';
import { IconX, IconGrid } from '../components/Icons';
import { AUTHOR_NAME } from '../constants';

import AdminOverview from '../components/admin/AdminOverview';
import AdminPosts from '../components/admin/AdminPosts';
import AdminCategories from '../components/admin/AdminCategories';
import AdminTags from '../components/admin/AdminTags';
import AdminFiles from '../components/admin/AdminFiles';
import AdminEchoes from '../components/admin/AdminEchoes';

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

    // --- Post State (Editor) ---
    const [editingPost, setEditingPost] = useState<Partial<BlogPost> | null>(null);

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
    const handleAddCategory = (name: string) => {
        addCategory({
            id: `c-${Date.now()}`,
            name: name,
            slug: name.toLowerCase().replace(/ /g, '-'),
            count: 0
        });
    };

    // TAGS
    const handleAddTag = (name: string) => {
        addTag({
            id: `t-${Date.now()}`,
            name: name
        });
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

    return (
        <div className="min-h-full bg-[#FDFBF7]">
            {section === 'overview' && (
                <AdminOverview
                    user={user}
                    dashboardStats={dashboardStats}
                    onEditPost={handleEditPost}
                />
            )}
            {section === 'posts' && (
                <AdminPosts
                    posts={posts}
                    onEditPost={handleEditPost}
                    onDeletePost={deletePost}
                    onViewPost={(id) => { setActivePostId(id); onExit(); }}
                    requestConfirm={requestConfirm}
                />
            )}
            {section === 'categories' && (
                <AdminCategories
                    categories={categories}
                    onAddCategory={handleAddCategory}
                    onDeleteCategory={deleteCategory}
                    requestConfirm={requestConfirm}
                />
            )}
            {section === 'tags' && (
                <AdminTags
                    tags={tags}
                    onAddTag={handleAddTag}
                    onDeleteTag={deleteTag}
                    requestConfirm={requestConfirm}
                />
            )}
            {section === 'files' && (
                <AdminFiles
                    files={files}
                    onAddFile={addFile}
                    onDeleteFile={deleteFile}
                    requestConfirm={requestConfirm}
                />
            )}
            {section === 'echoes' && (
                <AdminEchoes
                    visitLogs={visitLogs}
                    posts={posts}
                />
            )}

            {isEditorOpen && renderEditor()}

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
