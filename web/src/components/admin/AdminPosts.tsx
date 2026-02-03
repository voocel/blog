import React, { useState } from 'react';
import { IconPlus, IconX, IconClock, IconEye, IconEdit, IconTrash, IconArrowLeft, IconArrowDown } from '../Icons';
import type { BlogPost } from '../../types';

interface AdminPostsProps {
    posts: BlogPost[];
    onEditPost: (post?: BlogPost) => void;
    onDeletePost: (id: number) => void;
    onPublishPost: (id: number) => void;
    onViewPost: (slug: string) => void;
    requestConfirm: (title: string, message: string, onConfirm: () => void, options?: { confirmText?: string; isDestructive?: boolean }) => void;
}

const AdminPosts: React.FC<AdminPostsProps> = ({ posts, onEditPost, onDeletePost, onPublishPost, onViewPost, requestConfirm }) => {
    const [postSearch, setPostSearch] = useState('');
    const [postFilter, setPostFilter] = useState<'all' | 'published' | 'draft'>('all');
    const [currentPage, setCurrentPage] = useState(1);
    const itemsPerPage = 5;

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
                    onClick={() => onEditPost()}
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
                            <div className="w-28 h-20 rounded-lg bg-stone-100 overflow-hidden shrink-0 relative border border-stone-100 cursor-pointer" onClick={() => onViewPost(post.slug)}>
                                <img src={post.cover} className="w-full h-full object-cover group-hover:scale-105 transition-transform duration-500" alt="" />
                            </div>

                            {/* Content */}
                            <div className="flex-1 min-w-0 py-1">
                                <div className="flex items-center gap-3 mb-1.5">
                                    <h3 onClick={() => onViewPost(post.slug)} className="text-xl font-serif font-bold text-ink truncate hover:text-gold-600 transition-colors cursor-pointer">{post.title}</h3>
                                    {post.status === 'draft' && (
                                        <span className="text-[10px] uppercase tracking-wider font-bold text-amber-600 bg-amber-50 px-2 py-0.5 rounded border border-amber-200 animate-pulse">
                                            Draft
                                        </span>
                                    )}
                                    {post.status === 'published' && (
                                        <span className="text-[10px] uppercase tracking-wider text-emerald-600 bg-emerald-50 px-2 py-0.5 rounded border border-emerald-200">
                                            Published
                                        </span>
                                    )}
                                    <span className="text-[10px] uppercase tracking-wider text-stone-500 bg-stone-100 px-2 py-0.5 rounded border border-stone-200">{post.category}</span>
                                </div>
                                <div className="flex items-center gap-6 text-xs text-stone-400 font-medium">
                                    <span className="flex items-center gap-1.5"><IconClock className="w-3.5 h-3.5" /> {new Date(post.publishAt).toLocaleString()}</span>
                                    <span className="flex items-center gap-1.5"><IconEye className="w-3.5 h-3.5" /> {post.views.toLocaleString()} reads</span>
                                    <div className="flex gap-2">
                                        {post.tags.slice(0, 3).map(t => <span key={t} className="text-stone-300">#{t}</span>)}
                                    </div>
                                </div>
                            </div>

                            {/* Actions */}
                            <div className="flex items-center gap-2 opacity-0 group-hover:opacity-100 transition-opacity translate-x-4 group-hover:translate-x-0 duration-300">
                                {post.status === 'draft' && (
                                    <button
                                        onClick={(e) => {
                                            e.stopPropagation();
                                            requestConfirm(
                                                'Publish Entry',
                                                'Are you sure you want to publish this entry?',
                                                () => onPublishPost(post.id),
                                                { confirmText: 'Publish', isDestructive: false }
                                            );
                                        }}
                                        className="flex items-center gap-2 px-4 py-2 bg-emerald-50 text-emerald-600 hover:text-emerald-700 hover:bg-emerald-100 rounded-lg transition-colors border border-emerald-200 cursor-pointer"
                                        title="Publish"
                                    >
                                        <span className="text-xs font-bold uppercase tracking-wider">Publish</span>
                                    </button>
                                )}
                                <button onClick={(e) => { e.stopPropagation(); onEditPost(post); }} className="flex items-center gap-2 px-4 py-2 bg-stone-50 text-stone-600 hover:text-gold-600 hover:bg-gold-50 rounded-lg transition-colors border border-stone-200 cursor-pointer" title="Edit">
                                    <IconEdit className="w-4 h-4" />
                                    <span className="text-xs font-bold uppercase tracking-wider">Edit</span>
                                </button>
                                <button onClick={(e) => { e.stopPropagation(); requestConfirm('Delete Entry', 'Are you sure you want to delete this entry? This action cannot be undone.', () => onDeletePost(post.id)); }} className="p-2.5 text-stone-400 hover:text-red-600 hover:bg-red-50 rounded-lg transition-colors border border-transparent hover:border-red-100 cursor-pointer" title="Delete">
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

export default AdminPosts;
