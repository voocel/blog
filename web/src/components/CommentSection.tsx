import React, { useState, useEffect } from 'react';
import { useBlog } from '../context/BlogContext';
import { commentService } from '../services/commentService';
import type { Comment } from '../types';
import { IconUserCircle, IconSparkles } from './Icons';

interface CommentSectionProps {
    postId: string;
}

const CommentSection: React.FC<CommentSectionProps> = ({ postId }) => {
    const { user, setAuthModalOpen } = useBlog();
    const [comments, setComments] = useState<Comment[]>([]);
    const [isLoading, setIsLoading] = useState(true);
    const [newComment, setNewComment] = useState('');
    const [replyingTo, setReplyingTo] = useState<Comment | null>(null);
    const [isSubmitting, setIsSubmitting] = useState(false);
    const [isFocused, setIsFocused] = useState(false);

    useEffect(() => {
        const fetchComments = async () => {
            try {
                const response = await commentService.getComments(postId);
                setComments(response.data);
            } catch (error) {
                console.error("Failed to load comments", error);
            } finally {
                setIsLoading(false);
            }
        };
        fetchComments();
    }, [postId]);

    const handleSubmit = async (e: React.FormEvent) => {
        e.preventDefault();
        if (!newComment.trim() || !user) return;

        setIsSubmitting(true);
        try {
            const parentId = replyingTo ? replyingTo.id : undefined;
            const created = await commentService.createComment(postId, newComment, {
                username: user.username,
                avatar: user.avatar
            }, parentId);

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
                <div className="w-8 h-8 rounded-full bg-gold-100 flex items-center justify-center text-gold-600">
                    <IconSparkles className="w-4 h-4" />
                </div>
                <h3 className="text-2xl font-serif text-ink">Discussion</h3>
                <span className="text-sm font-sans text-stone-400 bg-stone-100 px-2 py-0.5 rounded-full">
                    {comments.reduce((acc, c) => acc + 1 + (c.replies?.length || 0), 0)}
                </span>
            </div>

            {/* Input Area */}
            <div className="mb-12 relative group">
                {!user && (
                    <div className="absolute inset-0 z-20 flex flex-col items-center justify-center bg-stone-50/60 backdrop-blur-[2px] rounded-xl border border-stone-100 transition-all duration-500">
                        <p className="text-stone-500 font-serif mb-4 italic">Join the conversation</p>
                        <button
                            onClick={() => setAuthModalOpen(true)}
                            className="bg-ink text-white px-6 py-2.5 rounded-full font-medium shadow-lg shadow-stone-200 hover:shadow-xl hover:-translate-y-0.5 transition-all text-sm cursor-pointer"
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

                <form onSubmit={handleSubmit} className={`relative bg-white border transition-all duration-300 ${replyingTo ? 'rounded-b-xl border-t-0' : 'rounded-xl'} ${isFocused ? 'border-gold-400 shadow-lg shadow-gold-100/50 ring-1 ring-gold-100' : 'border-stone-200 shadow-sm'}`}>
                    <div className="p-4">
                        <textarea
                            value={newComment}
                            onChange={(e) => setNewComment(e.target.value)}
                            onFocus={() => setIsFocused(true)}
                            onBlur={() => setIsFocused(false)}
                            disabled={!user || isSubmitting}
                            placeholder={user ? (replyingTo ? `Write a reply to ${replyingTo.user.username}...` : "Share your thoughts...") : "Sign in to share your thoughts..."}
                            className="w-full min-h-[100px] resize-none outline-none text-ink placeholder:text-stone-300 bg-transparent"
                        />
                    </div>
                    {user && (
                        <div className={`flex justify-end p-3 border-t border-stone-50 bg-stone-50/30 rounded-b-xl transition-opacity duration-300 ${isFocused || newComment ? 'opacity-100' : 'opacity-50'}`}>
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
                            <div className="w-10 h-10 rounded-full bg-stone-200 shrink-0" />
                            <div className="flex-1 space-y-2">
                                <div className="h-4 w-1/4 bg-stone-200 rounded" />
                                <div className="h-4 w-3/4 bg-stone-200 rounded" />
                            </div>
                        </div>
                    ))
                ) : comments.length === 0 ? (
                    <div className="text-center py-10 bg-stone-50 rounded-xl border border-stone-100 border-dashed">
                        <p className="text-stone-400 italic">No comments yet. Be the first to start the discussion.</p>
                    </div>
                ) : (
                    comments.map((comment) => (
                        <div key={comment.id} className="group animate-slide-up">
                            {/* Parent Comment */}
                            <div className="flex gap-4">
                                <div className="shrink-0">
                                    {comment.user.avatar ? (
                                        <img src={comment.user.avatar} alt={comment.user.username} className="w-10 h-10 rounded-full object-cover border border-stone-100 shadow-sm" />
                                    ) : (
                                        <div className="w-10 h-10 rounded-full bg-stone-100 flex items-center justify-center text-stone-400">
                                            <IconUserCircle className="w-6 h-6" />
                                        </div>
                                    )}
                                </div>
                                <div className="flex-1">
                                    <div className="bg-white p-5 rounded-2xl rounded-tl-sm border border-stone-100 shadow-sm group-hover:shadow-md group-hover:border-stone-200 transition-all">
                                        <div className="flex items-center justify-between mb-2">
                                            <span className="font-bold text-ink text-sm">{comment.user.username}</span>
                                            <span className="text-xs text-stone-400">{formatDate(comment.createdAt)}</span>
                                        </div>
                                        <p className="text-stone-600 leading-relaxed text-sm">
                                            {comment.content}
                                        </p>
                                        <div className="mt-3 flex items-center justify-end">
                                            <button
                                                onClick={() => {
                                                    setReplyingTo(comment);
                                                    document.querySelector('textarea')?.focus();
                                                }}
                                                className="text-xs font-medium text-stone-400 hover:text-gold-600 transition-colors cursor-pointer"
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
                                    <div className="absolute left-9 top-0 bottom-6 w-px bg-stone-200"></div>

                                    {comment.replies.map(reply => (
                                        <div key={reply.id} className="flex gap-4 relative">
                                            {/* Horizontal connector line */}
                                            <div className="absolute -left-5 top-5 w-4 h-px bg-stone-200"></div>

                                            <div className="shrink-0">
                                                {reply.user.avatar ? (
                                                    <img src={reply.user.avatar} alt={reply.user.username} className="w-8 h-8 rounded-full object-cover border border-stone-100 shadow-sm" />
                                                ) : (
                                                    <div className="w-8 h-8 rounded-full bg-stone-100 flex items-center justify-center text-stone-400">
                                                        <IconUserCircle className="w-5 h-5" />
                                                    </div>
                                                )}
                                            </div>
                                            <div className="flex-1">
                                                <div className="bg-stone-50 p-4 rounded-xl rounded-tl-sm border border-stone-100">
                                                    <div className="flex items-center justify-between mb-1">
                                                        <span className="font-bold text-ink text-xs">{reply.user.username}</span>
                                                        <span className="text-[10px] text-stone-400">{formatDate(reply.createdAt)}</span>
                                                    </div>
                                                    <p className="text-stone-600 leading-relaxed text-xs">
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
