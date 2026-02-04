import React, { useState } from 'react';
import { IconActivity, IconGlobe, IconMapPin } from '@/components/Icons';
import type { VisitLog, BlogPost } from '@/types';

interface AdminEchoesProps {
    visitLogs: VisitLog[];
    posts: BlogPost[];
}

const AdminEchoes: React.FC<AdminEchoesProps> = ({ visitLogs, posts }) => {
    const [visibleLogs, setVisibleLogs] = useState(20);

    // --- Stats Calculation ---
    const now = new Date();
    const startOfDay = new Date(now.getFullYear(), now.getMonth(), now.getDate()).getTime();
    const startOfMonth = new Date(now.getFullYear(), now.getMonth(), 1).getTime();

    // Resonances (Page Views)
    const totalResonances = visitLogs.length;
    const monthResonances = visitLogs.filter(l => l.timestamp * 1000 >= startOfMonth).length;
    const dayResonances = visitLogs.filter(l => l.timestamp * 1000 >= startOfDay).length;

    // Wanderers (Unique Visitors)
    const uniqueVisitorsTotal = new Set(visitLogs.map(log => log.ip)).size;
    const uniqueVisitorsMonth = new Set(visitLogs.filter(l => l.timestamp * 1000 >= startOfMonth).map(l => l.ip)).size;
    const uniqueVisitorsDay = new Set(visitLogs.filter(l => l.timestamp * 1000 >= startOfDay).map(l => l.ip)).size;

    // Top Posts (Today Only)
    const today = new Date().toDateString();
    const topPosts = visitLogs
        .filter(log => log.postId && new Date(log.timestamp * 1000).toDateString() === today)
        .reduce((acc, log) => {
            const id = log.postId!;
            acc[id] = (acc[id] || 0) + 1;
            return acc;
        }, {} as Record<number, number>);

    const sortedTopPosts = Object.entries(topPosts)
        .sort(([, a], [, b]) => b - a)
        .slice(0, 5)
        .map(([id, count]) => {
            const post = posts.find(p => p.id === Number(id));
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
                                                Reading: <span className="text-indigo-600 font-medium not-italic">{log.postTitle ? `"${log.postTitle}"` : log.pagePath === '/' ? "The Homepage" : log.pagePath}</span>
                                            </p>
                                        </div>
                                        <div className="text-xs text-stone-400 font-mono text-right">
                                            <div className="font-bold text-stone-500">{log.ip}</div>
                                            <div>{new Date(log.timestamp * 1000).toLocaleString(undefined, { month: 'numeric', day: 'numeric', hour: '2-digit', minute: '2-digit' })}</div>
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

export default AdminEchoes;
