import React, { useState } from 'react';
import type { Comment } from '@/types';
import { IconSearch, IconTrash, IconMessageSquare } from '@/components/Icons';

interface AdminCommentsProps {
    comments: Comment[];
    onDeleteComment: (id: number) => void;
}

const AdminComments: React.FC<AdminCommentsProps> = ({ comments, onDeleteComment }) => {
    const [searchTerm, setSearchTerm] = useState('');

    // Ensure comments is an array before filtering
    const safeComments = Array.isArray(comments) ? comments : [];

    const filteredComments = safeComments.filter(comment => {
        const contentMatch = comment.content?.toLowerCase().includes(searchTerm.toLowerCase());
        const userMatch = comment.user?.username?.toLowerCase().includes(searchTerm.toLowerCase());
        return contentMatch || userMatch;
    });

    const formatDate = (dateString: string) => {
        if (!dateString) return '-';
        return new Date(dateString).toLocaleDateString('en-US', {
            month: 'short', day: 'numeric', hour: '2-digit', minute: '2-digit'
        });
    };

    return (
        <div className="p-10 h-full flex flex-col animate-fade-in text-ink max-w-[1600px] mx-auto w-full">
            <div className="flex justify-between items-end mb-10">
                <div>
                    <h2 className="text-4xl font-serif font-bold text-ink mb-2">Comments</h2>
                    <p className="text-[var(--color-text-secondary)]">Moderate discussions and community interactions.</p>
                </div>
                <div className="relative">
                    <IconSearch className="absolute left-4 top-1/2 -translate-y-1/2 w-5 h-5 text-[var(--color-text-muted)]" />
                    <input
                        type="text"
                        placeholder="Search content or user..."
                        value={searchTerm}
                        onChange={(e) => setSearchTerm(e.target.value)}
                        className="pl-12 pr-4 py-3 bg-[var(--color-surface)] border border-[var(--color-border)] rounded-xl shadow-sm focus:outline-none focus:border-gold-400 focus:ring-1 focus:ring-gold-100 transition-all w-64"
                    />
                </div>
            </div>

            <div className="bg-[var(--color-surface)] border border-[var(--color-border)] rounded-2xl shadow-sm overflow-hidden flex-1 flex flex-col">
                <div className="overflow-x-auto custom-scrollbar flex-1">
                    <table className="w-full text-left border-collapse">
                        <thead className="bg-[var(--color-surface-alt)] border-b border-[var(--color-border-subtle)] sticky top-0 z-10">
                            <tr>
                                <th className="py-4 px-6 text-xs font-bold uppercase tracking-wider text-[var(--color-text-muted)] w-1/4">Author</th>
                                <th className="py-4 px-6 text-xs font-bold uppercase tracking-wider text-[var(--color-text-muted)] w-1/2">Comment</th>
                                <th className="py-4 px-6 text-xs font-bold uppercase tracking-wider text-[var(--color-text-muted)]">Context</th>
                                <th className="py-4 px-6 text-xs font-bold uppercase tracking-wider text-[var(--color-text-muted)] text-right">Actions</th>
                            </tr>
                        </thead>
                        <tbody className="divide-y divide-[var(--color-border-subtle)]">
                            {filteredComments.length > 0 ? (
                                filteredComments.map((comment) => (
                                    <tr key={comment.id} className="hover:bg-[var(--color-surface-alt)]/50 transition-colors group">
                                        <td className="py-4 px-6 align-top">
                                            <div className="flex items-center gap-3">
                                                <div className="w-8 h-8 rounded-full bg-[var(--color-surface-alt)] border border-[var(--color-border)] overflow-hidden shrink-0">
                                                    {comment.user?.avatar ? (
                                                        <img src={comment.user.avatar} alt={comment.user.username} className="w-full h-full object-cover" />
                                                    ) : (
                                                        <div className="w-full h-full flex items-center justify-center text-[var(--color-text-muted)]">
                                                            <div className="w-4 h-4 bg-[var(--color-text-muted)] rounded-full" />
                                                        </div>
                                                    )}
                                                </div>
                                                <div>
                                                    <div className="font-bold text-ink text-sm">{comment.user?.username || 'Unknown'}</div>
                                                    <div className="text-[10px] text-[var(--color-text-muted)]">{formatDate(comment.createdAt)}</div>
                                                </div>
                                            </div>
                                        </td>
                                        <td className="py-4 px-6 align-top">
                                            <p className="text-sm text-[var(--color-text-secondary)] leading-relaxed line-clamp-3">
                                                {comment.content}
                                            </p>
                                        </td>
                                        <td className="py-4 px-6 align-top">
                                            {comment.postTitle ? (
                                                <span className="inline-flex items-center gap-1.5 px-2 py-1 rounded bg-[var(--color-surface-alt)] text-[var(--color-text-secondary)] text-xs truncate max-w-[150px]">
                                                    <IconMessageSquare className="w-3 h-3 shrink-0" />
                                                    <span className="truncate">{comment.postTitle}</span>
                                                </span>
                                            ) : (
                                                <span className="text-xs text-[var(--color-text-muted)] italic">Unknown Post</span>
                                            )}
                                        </td>
                                        <td className="py-4 px-6 align-top text-right">
                                            <button
                                                onClick={() => onDeleteComment(comment.id)}
                                                className="text-[var(--color-text-muted)] hover:text-red-600 p-2 rounded-lg hover:bg-red-50 dark:hover:bg-red-900/40 transition-all opacity-0 group-hover:opacity-100 cursor-pointer"
                                                title="Delete Comment"
                                            >
                                                <IconTrash className="w-4 h-4" />
                                            </button>
                                        </td>
                                    </tr>
                                ))
                            ) : (
                                <tr>
                                    <td colSpan={4} className="py-20 text-center text-[var(--color-text-muted)] italic">
                                        No comments found.
                                    </td>
                                </tr>
                            )}
                        </tbody>
                    </table>
                </div>
                <div className="p-4 border-t border-[var(--color-border-subtle)] bg-[var(--color-surface-alt)] flex justify-between items-center text-xs text-[var(--color-text-muted)] uppercase tracking-widest font-bold">
                    <span>Total Comments: {safeComments.length}</span>
                </div>
            </div>
        </div>
    );
};

export default AdminComments;
