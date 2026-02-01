import React, { useState, useCallback } from 'react';
import MDEditor from '@uiw/react-md-editor';
import { IconX, IconGrid, IconClock } from '../Icons';
import { useDraftAutoSave } from '../../hooks/useDraftAutoSave';
import ConfirmModal from '../ConfirmModal';
import type { BlogPost, Category, Tag } from '../../types';

interface PostEditorProps {
    post: Partial<BlogPost>;
    categories: Category[];
    tags: Tag[];
    onSave: (post: Partial<BlogPost>) => Promise<void>;
    onClose: () => void;
    onDraftSaved?: () => void;
}

const PostEditor: React.FC<PostEditorProps> = ({
    post: initialPost,
    categories,
    tags,
    onSave,
    onClose,
    onDraftSaved
}) => {
    const [editingPost, setEditingPost] = useState<Partial<BlogPost>>(initialPost);
    const [isSaving, setIsSaving] = useState(false);
    const [showCloseConfirm, setShowCloseConfirm] = useState(false);

    // Draft auto-save hook
    const {
        lastSavedContentRef,
        showDraftRecovery,
        setShowDraftRecovery,
        recoveryDraft,
        setRecoveryDraft,
        saveDraft,
        clearDraft,
        checkForDraft,
        hasUnsavedChanges
    } = useDraftAutoSave({
        post: editingPost,
        isEnabled: true,
        intervalMs: 5000,
        onSave: onDraftSaved
    });

    // Check for existing draft on mount
    React.useEffect(() => {
        const existingDraft = checkForDraft(initialPost.id);
        if (existingDraft && JSON.stringify(existingDraft) !== JSON.stringify(initialPost)) {
            setRecoveryDraft(existingDraft);
            setShowDraftRecovery(true);
        }
        // Initialize lastSavedContentRef
        lastSavedContentRef.current = JSON.stringify(initialPost);
    }, []);

    // Handle save
    const handleSave = async () => {
        if (!editingPost.title) return;
        setIsSaving(true);
        try {
            await onSave(editingPost);
            clearDraft(editingPost.id);
            onClose();
        } catch (error) {
            console.error('Failed to save:', error);
        } finally {
            setIsSaving(false);
        }
    };

    // Handle close with unsaved changes check
    const handleClose = useCallback(() => {
        if (hasUnsavedChanges(editingPost)) {
            setShowCloseConfirm(true);
        } else {
            onClose();
        }
    }, [editingPost, hasUnsavedChanges, onClose]);

    // Handle close confirm - save draft and close
    const handleSaveAndClose = () => {
        saveDraft(editingPost);
        setShowCloseConfirm(false);
        onClose();
    };

    // Handle close confirm - discard and close
    const handleDiscardAndClose = () => {
        clearDraft(editingPost.id);
        setShowCloseConfirm(false);
        onClose();
    };

    // Handle draft recovery
    const handleRecoverDraft = () => {
        if (recoveryDraft) {
            setEditingPost(recoveryDraft);
        }
        setShowDraftRecovery(false);
        setRecoveryDraft(null);
    };

    const handleDiscardDraft = () => {
        clearDraft(editingPost.id);
        setShowDraftRecovery(false);
        setRecoveryDraft(null);
    };

    return (
        <div className="fixed inset-0 z-[100] bg-[#FDFBF7] flex flex-col animate-slide-up">
            {/* Toolbar */}
            <div className="h-20 border-b border-stone-200 flex justify-between items-center px-8 bg-white/90 backdrop-blur-md shadow-sm z-20">
                <div className="flex items-center gap-6">
                    <button onClick={handleClose} className="text-stone-400 hover:text-ink flex items-center gap-2 transition-colors cursor-pointer">
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
                                onClick={() => setEditingPost({ ...editingPost, status: s as 'draft' | 'published' })}
                                className={`px-4 py-2 rounded-md text-xs uppercase tracking-wider font-bold transition-all cursor-pointer ${editingPost.status === s ? 'bg-white shadow-sm text-ink' : 'text-stone-400 hover:text-stone-600'
                                    }`}
                            >
                                {s}
                            </button>
                        ))}
                    </div>
                    <button
                        onClick={handleSave}
                        disabled={isSaving}
                        className="bg-ink text-white px-8 py-3 rounded-xl font-bold tracking-wide hover:bg-gold-600 transition-colors shadow-lg flex items-center gap-2 cursor-pointer disabled:opacity-50"
                    >
                        {isSaving ? 'Saving...' : 'Save Changes'}
                    </button>
                </div>
            </div>

            {/* Main Editor Area */}
            <div className="flex-1 flex overflow-hidden">
                {/* Meta Sidebar */}
                <div className="w-96 border-r border-stone-200 bg-stone-50 p-8 overflow-y-auto hidden lg:block custom-scrollbar">
                    <h3 className="font-serif font-bold text-ink mb-8 text-xl">Entry Metadata</h3>

                    <div className="space-y-8">
                        {/* Publish Time */}
                        <div>
                            <label className="block text-xs uppercase tracking-widest text-stone-500 mb-3 font-bold">Publish Time</label>
                            <div
                                className="relative group cursor-pointer"
                                onClick={() => {
                                    try {
                                        // @ts-ignore
                                        document.getElementById('publish-date-picker')?.showPicker();
                                    } catch (e) {
                                        document.getElementById('publish-date-picker')?.focus();
                                    }
                                }}
                            >
                                <div className="w-full bg-white border border-stone-200 group-hover:border-gold-400 rounded-xl p-3 flex items-center justify-between transition-all shadow-sm">
                                    <div className="flex items-center gap-3">
                                        <div className="w-8 h-8 rounded-lg bg-stone-50 flex items-center justify-center text-stone-400 group-hover:text-gold-500 transition-colors">
                                            <IconClock className="w-4 h-4" />
                                        </div>
                                        <div className="flex flex-col">
                                            <span className="text-xs text-stone-400 font-medium uppercase tracking-wider">Scheduled For</span>
                                            <span className={`text-sm font-serif font-medium ${editingPost.publishAt ? 'text-ink' : 'text-stone-300 italic'}`}>
                                                {editingPost.publishAt
                                                    ? new Date(editingPost.publishAt).toLocaleString('en-US', {
                                                        month: 'short', day: 'numeric', year: 'numeric',
                                                        hour: 'numeric', minute: 'numeric', hour12: true
                                                    })
                                                    : 'Set publish time...'}
                                            </span>
                                        </div>
                                    </div>
                                </div>
                                <input
                                    id="publish-date-picker"
                                    type="datetime-local"
                                    className="absolute inset-0 w-full h-full opacity-0 pointer-events-none"
                                    value={editingPost.publishAt ? new Date(editingPost.publishAt).toISOString().slice(0, 16) : ''}
                                    onChange={e => {
                                        const v = e.target.value;
                                        const iso = v ? new Date(v).toISOString() : '';
                                        setEditingPost({ ...editingPost, publishAt: iso });
                                    }}
                                />
                            </div>
                        </div>

                        {/* Slug */}
                        <div>
                            <label className="block text-xs uppercase tracking-widest text-stone-500 mb-3 font-bold">URL Slug</label>
                            <input
                                className="w-full bg-white border border-stone-200 rounded-xl p-3 text-sm focus:outline-none focus:border-gold-500 shadow-sm font-mono"
                                value={editingPost.slug || ''}
                                onChange={e => setEditingPost({ ...editingPost, slug: e.target.value.toLowerCase().replace(/\s+/g, '-') })}
                                placeholder="auto-generated-from-title"
                            />
                            <p className="text-xs text-stone-400 mt-2">Leave empty to auto-generate from title. URL: /post/{editingPost.slug || 'your-slug'}</p>
                        </div>

                        {/* Category */}
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

                        {/* Tags */}
                        <div>
                            <label className="block text-xs uppercase tracking-widest text-stone-500 mb-3 font-bold">Tags</label>
                            <div className="flex flex-wrap gap-2">
                                {tags.map(tag => {
                                    const isSelected = editingPost.tags?.includes(tag.id) || editingPost.tags?.includes(tag.name);
                                    return (
                                        <button
                                            key={tag.id}
                                            onClick={() => {
                                                const currentTags = editingPost.tags || [];
                                                const alreadySelected = currentTags.includes(tag.id) || currentTags.includes(tag.name);
                                                const newTags = alreadySelected
                                                    ? currentTags.filter(t => t !== tag.id && t !== tag.name)
                                                    : [...currentTags, tag.id];
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
                                {tags.length === 0 && <span className="text-xs text-stone-400 italic">No tags available.</span>}
                            </div>
                        </div>

                        {/* Excerpt */}
                        <div>
                            <label className="block text-xs uppercase tracking-widest text-stone-500 mb-3 font-bold">Excerpt / Summary</label>
                            <textarea
                                className="w-full bg-white border border-stone-200 rounded-xl p-4 text-sm h-40 resize-none focus:outline-none focus:border-gold-500 shadow-sm leading-relaxed"
                                value={editingPost.excerpt}
                                onChange={e => setEditingPost({ ...editingPost, excerpt: e.target.value })}
                                placeholder="Write a short summary for the feed display..."
                            />
                        </div>

                        {/* Cover Image */}
                        <div>
                            <label className="block text-xs uppercase tracking-widest text-stone-500 mb-3 font-bold">Cover Image URL</label>
                            <div className="w-full h-48 bg-stone-200 rounded-xl mb-3 overflow-hidden border border-stone-300 relative group">
                                <img src={editingPost.cover} className="w-full h-full object-cover opacity-90 group-hover:scale-105 transition-transform duration-700" alt="Cover" />
                            </div>
                            <input
                                className="w-full bg-white border border-stone-200 rounded-xl p-3 text-xs font-mono focus:outline-none focus:border-gold-500 shadow-sm"
                                value={editingPost.cover}
                                onChange={e => setEditingPost({ ...editingPost, cover: e.target.value })}
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

                        <div className="flex-1" data-color-mode="light">
                            <MDEditor
                                value={editingPost.content}
                                onChange={(val) => setEditingPost({ ...editingPost, content: val || '' })}
                                preview="edit"
                                height={600}
                                visibleDragbar={false}
                                highlightEnable={false}
                                textareaProps={{
                                    placeholder: "Start writing your thoughts..."
                                }}
                            />
                        </div>
                    </div>
                </div>
            </div>

            {/* Close Confirmation Modal */}
            <ConfirmModal
                isOpen={showCloseConfirm}
                title="Unsaved Changes"
                message="You have unsaved changes. Do you want to save before leaving?"
                confirmText="Save Draft"
                cancelText="Discard"
                isDestructive={false}
                onConfirm={handleSaveAndClose}
                onCancel={handleDiscardAndClose}
            />

            {/* Draft Recovery Modal */}
            {showDraftRecovery && recoveryDraft && (
                <div className="fixed inset-0 z-[400] flex items-center justify-center p-4">
                    <div className="absolute inset-0 bg-stone-900/50 backdrop-blur-sm" />
                    <div className="relative bg-white rounded-2xl shadow-2xl max-w-md w-full p-6 animate-slide-up">
                        <h3 className="text-xl font-serif font-bold text-ink mb-2">Recover Draft?</h3>
                        <p className="text-stone-500 mb-4">
                            We found an unsaved draft from your previous session. Would you like to recover it?
                        </p>
                        <div className="bg-stone-50 rounded-lg p-3 mb-6 text-sm text-stone-600">
                            <p className="font-medium truncate">{recoveryDraft.title || '(Untitled)'}</p>
                            <p className="text-stone-400 text-xs mt-1">
                                {recoveryDraft.content?.slice(0, 100)}...
                            </p>
                        </div>
                        <div className="flex justify-end gap-3">
                            <button
                                onClick={handleDiscardDraft}
                                className="px-4 py-2 text-stone-500 hover:text-ink font-medium cursor-pointer"
                            >
                                Discard Draft
                            </button>
                            <button
                                onClick={handleRecoverDraft}
                                className="px-6 py-2 bg-emerald-600 hover:bg-emerald-700 text-white rounded-lg font-bold transition-colors shadow-lg cursor-pointer"
                            >
                                Recover Draft
                            </button>
                        </div>
                    </div>
                </div>
            )}
        </div>
    );
};

export default PostEditor;
