import React, { useState, useEffect } from 'react';
import { useAuth } from '@/context/AuthContext';
import { commentService } from '@/services/commentService';
import type { Comment } from '@/types';
import { IconUserCircle, IconSparkles } from '@/components/Icons';

interface CommentSectionProps {
    postSlug: string;
}

const CommentSection: React.FC<CommentSectionProps> = ({ postSlug }) => {
    const { user, setAuthModalOpen } = useAuth();
    const [comments, setComments] = useState<Comment[]>([]);
    const [isLoading, setIsLoading] = useState(true);
    const [newComment, setNewComment] = useState('');
    const [replyingTo, setReplyingTo] = useState<Comment | null>(null);
    const [isSubmitting, setIsSubmitting] = useState(false);
    const [isFocused, setIsFocused] = useState(false);

    useEffect(() => {
        const fetchComments = async () => {
            try {
                const response = await commentService.getComments(postSlug);
                setComments(response.data);
            } catch (error) {
                console.error("Failed to load comments", error);
            } finally {
                setIsLoading(false);
            }
        };
        fetchComments();
    }, [postSlug]);

    const handleSubmit = async (e: React.FormEvent) => {
        e.preventDefault();
        if (!newComment.trim() || !user) return;

        setIsSubmitting(true);
        try {
            const parentId = replyingTo ? replyingTo.id : undefined;
            const created = await commentService.createComment(postSlug, newComment, parentId);

            if (parentId) {
                // Optimistically update replies
                setComments(prev => prev.map(c => {
                    if (c.id === parentId) {
                        return {
                            ...c,
                            replies: [...(c.replies || []), created]
                        };
                    }
                    return c;
                }));
            } else {
                setComments(prev => [created, ...prev]);
            }

            setNewComment('');
            setReplyingTo(null);
        } catch (error) {
            console.error("Failed to post comment", error);
        } finally {
            setIsSubmitting(false);
        }
    };

    const formatDate = (dateString: string) => {
        const date = new Date(dateString);
        return new Intl.DateTimeFormat('en-US', { month: 'short', day: 'numeric', hour: 'numeric', minute: 'numeric' }).format(date);
    };

    return (
        <div className="mt-16 max-w-3xl mx-auto">
            <div className="flex items-center gap-3 mb-8">
                <div className="w-8 h-8 rounded-full bg-gold-100 dark:bg-gold-900/30 flex items-center justify-center text-gold-600">
                    <IconSparkles className="w-4 h-4" />
                </div>
                <h3 className="text-2xl font-serif text-ink">Discussion</h3>
                <span className="text-sm font-sans text-[var(--color-text-muted)] bg-[var(--color-surface-alt)] px-2 py-0.5 rounded-full">
                    {comments.reduce((acc, c) => acc + 1 + (c.replies?.length || 0), 0)}
                </span>
            </div>

            {/* Input Area */}
            <div className="mb-12 relative group">
                {!user && (
                    <div className="absolute inset-0 z-20 flex flex-col items-center justify-center bg-[var(--color-surface-alt)]/60 backdrop-blur-[2px] rounded-xl border border-[var(--color-border-subtle)] transition-all duration-500">
                        <p className="text-[var(--color-text-secondary)] font-serif mb-4 italic">Join the conversation</p>
                        <button
                            onClick={() => setAuthModalOpen(true)}
                            className="bg-ink text-[var(--color-base)] px-6 py-2.5 rounded-full font-medium shadow-lg shadow-[var(--color-muted)] hover:shadow-xl hover:-translate-y-0.5 transition-all text-sm cursor-pointer"
                        >
                            Sign in to Comment
                        </button>
                    </div>
                )}

                {/* Replying Context Banner */}
                {replyingTo && (
                    <div className="flex items-center justify-between bg-gold-50 px-4 py-2 rounded-t-xl border border-gold-100 border-b-0 text-gold-800 text-sm animate-fade-in">
                        <span>Replying to <strong>@{replyingTo.user.username}</strong></span>
                        <button
                            onClick={() => setReplyingTo(null)}
                            className="text-gold-600 hover:text-gold-800 hover:underline text-xs"
                        >
                            Cancel
                        </button>
                    </div>
                )}

                <form onSubmit={handleSubmit} className={`relative bg-[var(--color-surface)] border transition-all duration-300 ${replyingTo ? 'rounded-b-xl border-t-0' : 'rounded-xl'} ${isFocused ? 'border-gold-400 shadow-lg shadow-gold-100/50 ring-1 ring-gold-100' : 'border-[var(--color-border)] shadow-sm'}`}>
                    <div className="p-4">
                        <textarea
                            value={newComment}
                            onChange={(e) => setNewComment(e.target.value)}
                            onKeyDown={(e) => {
                                if (e.key === 'Enter' && !e.shiftKey) {
                                    e.preventDefault();
                                    handleSubmit(e);
                                }
                            }}
                            onFocus={() => setIsFocused(true)}
                            onBlur={() => setIsFocused(false)}
                            disabled={!user || isSubmitting}
                            placeholder={user ? (replyingTo ? `Write a reply to ${replyingTo.user.username}...` : "Share your thoughts...") : "Sign in to share your thoughts..."}
                            className="w-full min-h-[100px] resize-none outline-none text-ink placeholder:text-[var(--color-text-muted)] bg-transparent"
                        />
                    </div>
                    {user && (
                        <div className={`flex justify-end p-3 border-t border-[var(--color-border-subtle)] bg-[var(--color-surface-alt)]/30 rounded-b-xl transition-opacity duration-300 ${isFocused || newComment ? 'opacity-100' : 'opacity-50'}`}>
                            <button
                                type="submit"
                                disabled={!newComment.trim() || isSubmitting}
                                className="bg-gold-500 text-white px-5 py-2 rounded-lg text-sm font-medium hover:bg-gold-600 disabled:opacity-50 disabled:cursor-not-allowed transition-colors cursor-pointer"
                            >
                                {isSubmitting ? 'Posting...' : (replyingTo ? 'Post Reply' : 'Post Comment')}
                            </button>
                        </div>
                    )}
                </form>
            </div>

            {/* Comments List */}
            <div className="space-y-8">
                {isLoading ? (
                    // Loading Skeleton
                    [1, 2].map(i => (
                        <div key={i} className="animate-pulse flex gap-4">
                            <div className="w-10 h-10 rounded-full bg-[var(--color-muted)] shrink-0" />
                            <div className="flex-1 space-y-2">
                                <div className="h-4 w-1/4 bg-[var(--color-muted)] rounded" />
                                <div className="h-4 w-3/4 bg-[var(--color-muted)] rounded" />
                            </div>
                        </div>
                    ))
                ) : comments.length === 0 ? (
                    <div className="text-center py-10 bg-[var(--color-surface-alt)] rounded-xl border border-[var(--color-border-subtle)] border-dashed">
                        <p className="text-[var(--color-text-muted)] italic">No comments yet. Be the first to start the discussion.</p>
                    </div>
                ) : (
                    comments.map((comment) => (
                        <div key={comment.id} className="group animate-slide-up">
                            {/* Parent Comment */}
                            <div className="flex gap-4">
                                <div className="shrink-0">
                                    {comment.user.avatar ? (
                                        <img src={comment.user.avatar} alt={comment.user.username} className="w-10 h-10 rounded-full object-cover border border-[var(--color-border-subtle)] shadow-sm" />
                                    ) : (
                                        <div className="w-10 h-10 rounded-full bg-[var(--color-surface-alt)] flex items-center justify-center text-[var(--color-text-muted)]">
                                            <IconUserCircle className="w-6 h-6" />
                                        </div>
                                    )}
                                </div>
                                <div className="flex-1">
                                    <div className="bg-[var(--color-surface)] p-5 rounded-2xl rounded-tl-sm border border-[var(--color-border-subtle)] shadow-sm group-hover:shadow-md group-hover:border-[var(--color-border)] transition-all">
                                        <div className="flex items-center justify-between mb-2">
                                            <span className="font-bold text-ink text-sm">{comment.user.username}</span>
                                            <span className="text-xs text-[var(--color-text-muted)]">{formatDate(comment.createdAt)}</span>
                                        </div>
                                        <p className="text-[var(--color-text-secondary)] leading-relaxed text-sm">
                                            {comment.content}
                                        </p>
                                        <div className="mt-3 flex items-center justify-end">
                                            <button
                                                onClick={() => {
                                                    setReplyingTo(comment);
                                                    document.querySelector('textarea')?.focus();
                                                }}
                                                className="text-xs font-medium text-[var(--color-text-muted)] hover:text-gold-600 transition-colors cursor-pointer"
                                            >
                                                Reply
                                            </button>
                                        </div>
                                    </div>
                                </div>
                            </div>

                            {/* Nested Replies */}
                            {comment.replies && comment.replies.length > 0 && (
                                <div className="mt-3 pl-14 space-y-3 relative">
                                    {/* Vertical logic line */}
                                    <div className="absolute left-9 top-0 bottom-6 w-px bg-[var(--color-border)]"></div>

                                    {comment.replies.map(reply => (
                                        <div key={reply.id} className="flex gap-4 relative">
                                            {/* Horizontal connector line */}
                                            <div className="absolute -left-5 top-5 w-4 h-px bg-[var(--color-border)]"></div>

                                            <div className="shrink-0">
                                                {reply.user.avatar ? (
                                                    <img src={reply.user.avatar} alt={reply.user.username} className="w-8 h-8 rounded-full object-cover border border-[var(--color-border-subtle)] shadow-sm" />
                                                ) : (
                                                    <div className="w-8 h-8 rounded-full bg-[var(--color-surface-alt)] flex items-center justify-center text-[var(--color-text-muted)]">
                                                        <IconUserCircle className="w-5 h-5" />
                                                    </div>
                                                )}
                                            </div>
                                            <div className="flex-1">
                                                <div className="bg-[var(--color-surface-alt)] p-4 rounded-xl rounded-tl-sm border border-[var(--color-border-subtle)]">
                                                    <div className="flex items-center justify-between mb-1">
                                                        <span className="font-bold text-ink text-xs">{reply.user.username}</span>
                                                        <span className="text-[10px] text-[var(--color-text-muted)]">{formatDate(reply.createdAt)}</span>
                                                    </div>
                                                    <p className="text-[var(--color-text-secondary)] leading-relaxed text-xs">
                                                        {reply.replyToUser && (
                                                            <span className="text-gold-600 font-medium mr-1">@{reply.replyToUser.username}</span>
                                                        )}
                                                        {reply.content}
                                                    </p>
                                                </div>
                                            </div>
                                        </div>
                                    ))}
                                </div>
                            )}
                        </div>
                    ))
                )}
            </div>
        </div>
    );
};

export default CommentSection;
