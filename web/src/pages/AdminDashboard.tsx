import React, { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import { useBlog } from '../context/BlogContext';
import type { AdminSection, BlogPost } from '../types';
import { useToast } from '../components/Toast';
import { AUTHOR_NAME } from '../constants';
import ConfirmModal from '../components/ConfirmModal';
import PostEditor from '../components/admin/PostEditor';

import AdminOverview from '../components/admin/AdminOverview';
import AdminPosts from '../components/admin/AdminPosts';
import AdminCategories from '../components/admin/AdminCategories';
import AdminTags from '../components/admin/AdminTags';
import AdminFiles from '../components/admin/AdminFiles';
import AdminEchoes from '../components/admin/AdminEchoes';
import AdminUsers from '../components/admin/AdminUsers';
import AdminComments from '../components/admin/AdminComments';

interface AdminDashboardProps {
    section: AdminSection;
    onExit: () => void;
}

const AdminDashboard: React.FC<AdminDashboardProps> = ({ section, onExit: _onExit }) => {
    const navigate = useNavigate();
    const {
        posts, categories, tags, files, user,
        addPost, updatePost, deletePost,
        addCategory, deleteCategory,
        addTag, deleteTag,
        addFile, deleteFile, visitLogs, dashboardStats,
        refreshPosts, refreshCategories, refreshTags, refreshFiles, refreshVisitLogs, refreshDashboardOverview,
        adminUsers, refreshAdminUsers, allComments, refreshAllComments
    } = useBlog();

    const { showToast } = useToast();

    // Fetch data based on section
    useEffect(() => {
        if (section === 'overview') {
            refreshDashboardOverview();
            refreshPosts();
        } else if (section === 'posts') {
            refreshPosts();
            refreshCategories();
            refreshTags();
        } else if (section === 'categories') {
            refreshCategories();
        } else if (section === 'tags') {
            refreshTags();
        } else if (section === 'files') {
            refreshFiles();
        } else if (section === 'echoes') {
            refreshVisitLogs();
        } else if (section === 'users') {
            refreshAdminUsers();
        } else if (section === 'comments') {
            refreshAllComments();
        }
    }, [section]);

    // --- Editor State ---
    const [isEditorOpen, setIsEditorOpen] = useState(false);
    const [editingPost, setEditingPost] = useState<Partial<BlogPost> | null>(null);

    // --- Confirmation Modal State ---
    const [confirmModal, setConfirmModal] = useState<{
        isOpen: boolean;
        title: string;
        message: string;
        confirmText: string;
        isDestructive: boolean;
        onConfirm: () => void;
    }>({
        isOpen: false,
        title: '',
        message: '',
        confirmText: 'Confirm',
        isDestructive: true,
        onConfirm: () => { }
    });

    const requestConfirm = (
        title: string,
        message: string,
        onConfirm: () => void,
        options: { confirmText?: string; isDestructive?: boolean } = {}
    ) => {
        setConfirmModal({
            isOpen: true,
            title,
            message,
            onConfirm,
            confirmText: options.confirmText || 'Confirm Delete',
            isDestructive: options.isDestructive !== undefined ? options.isDestructive : true
        });
    };

    // --- Post Handlers ---
    const handleEditPost = (post?: BlogPost) => {
        if (post) {
            const tagIds = (post.tags || []).map(t => {
                const tagByName = tags.find(tag => tag.name === t);
                return tagByName ? tagByName.id : t;
            });
            setEditingPost({ ...post, tags: tagIds });
        } else {
            setEditingPost({
                title: '',
                excerpt: '',
                content: '',
                author: AUTHOR_NAME,
                publishAt: new Date().toISOString(),
                category: categories[0]?.name || 'General',
                readTime: '5 min read',
                cover: 'https://images.unsplash.com/photo-1618005182384-a83a8bd57fbe?auto=format&fit=crop&w=800&q=80',
                tags: [],
                status: 'draft',
                views: 0
            });
        }
        refreshCategories();
        refreshTags();
        setIsEditorOpen(true);
    };

    const handleSavePost = async (post: Partial<BlogPost>) => {
        if (!post.title) return;

        const payload: any = {
            ...post,
            categoryId: post.categoryId,
            tags: (post.tags || []).map(t => {
                const tagByName = tags.find(tag => tag.name === t);
                return tagByName ? tagByName.id : t;
            })
        };

        if (!payload.categoryId && categories.length > 0) {
            const matchedCat = categories.find(c => c.name === post.category);
            payload.categoryId = matchedCat ? matchedCat.id : categories[0].id;
        }

        const existing = posts.find(p => p.id === post.id);
        if (existing && post.id) {
            await updatePost(post.id, payload);
            showToast("Entry updated successfully", "success");
        } else {
            await addPost(payload as BlogPost);
            showToast("New entry created", "success");
        }
    };

    const handleCloseEditor = () => {
        setIsEditorOpen(false);
        setEditingPost(null);
    };

    const handlePublishPost = async (id: string) => {
        try {
            await updatePost(id, { status: 'published' });
            refreshPosts();
            showToast("Entry published successfully", "success");
        } catch (error) {
            console.error("Failed to publish post:", error);
            showToast("Failed to publish post", "error");
        }
    };

    // --- Category & Tag Handlers ---
    const handleAddCategory = (name: string) => {
        addCategory({
            id: `c-${Date.now()}`,
            name: name,
            slug: name.toLowerCase().replace(/ /g, '-'),
            count: 0
        });
    };

    const handleAddTag = (name: string) => {
        addTag({
            id: `t-${Date.now()}`,
            name: name
        });
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
                    onPublishPost={handlePublishPost}
                    onViewPost={(id) => { navigate(`/post/${id}`); }}
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
            {section === 'users' && (
                <AdminUsers users={adminUsers} requestConfirm={requestConfirm} />
            )}
            {section === 'comments' && (
                <AdminComments
                    comments={allComments}
                    onDeleteComment={(id) => requestConfirm(
                        'Delete Comment?',
                        'Are you sure you want to remove this comment? This action cannot be undone.',
                        async () => {
                            try {
                                const { commentService } = await import('../services/commentService');
                                await commentService.deleteComment(id);
                                await refreshAllComments();
                                showToast("Comment deleted", "success");
                            } catch (e) {
                                console.error(e);
                                showToast("Failed to delete comment", "error");
                            }
                        }
                    )}
                />
            )}

            {/* Post Editor */}
            {isEditorOpen && editingPost && (
                <PostEditor
                    post={editingPost}
                    categories={categories}
                    tags={tags}
                    onSave={handleSavePost}
                    onClose={handleCloseEditor}
                    onDraftSaved={() => showToast('Draft saved', 'success')}
                />
            )}

            {/* Confirmation Modal */}
            <ConfirmModal
                isOpen={confirmModal.isOpen}
                title={confirmModal.title}
                message={confirmModal.message}
                confirmText={confirmModal.confirmText}
                isDestructive={confirmModal.isDestructive}
                onConfirm={() => {
                    confirmModal.onConfirm();
                    setConfirmModal({ ...confirmModal, isOpen: false });
                }}
                onCancel={() => setConfirmModal({ ...confirmModal, isOpen: false })}
            />
        </div>
    );
};

export default AdminDashboard;
