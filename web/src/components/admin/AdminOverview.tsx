import React from 'react';
import { IconGrid, IconLayers, IconTag, IconImage, IconClock, IconEye } from '@/components/Icons';
import type { BlogPost, DashboardOverview, User } from '@/types';

interface AdminOverviewProps {
    user: User | null;
    dashboardStats: DashboardOverview | null;
    onEditPost: (post?: BlogPost) => void;
}

const AdminOverview: React.FC<AdminOverviewProps> = ({ user, dashboardStats, onEditPost }) => {
    const recentPosts = dashboardStats?.recentPosts || [];

    return (
        <div className="p-10 animate-fade-in text-ink max-w-[1600px] mx-auto">
            {/* Welcome Header */}
            <div className="mb-12 flex justify-between items-end">
                <div>
                    <h1 className="text-4xl md:text-5xl font-serif font-bold text-ink mb-2">Good Morning, {user?.username.split(' ')[0]}.</h1>
                    <p className="text-[var(--color-text-secondary)] text-lg font-serif italic">The aether awaits your thoughts today.</p>
                </div>
                <div className="hidden md:block text-right">
                    <p className="text-xs uppercase tracking-widest text-[var(--color-text-muted)] mb-1">Current Date</p>
                    <p className="text-xl font-serif font-bold text-ink">{new Date().toLocaleDateString('en-US', { weekday: 'long', month: 'long', day: 'numeric' })}</p>
                </div>
            </div>

            {/* Stats Grid */}
            <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6 mb-12">
                {[
                    { label: 'Total Entries', value: dashboardStats?.counts.posts || 0, icon: IconGrid, color: 'text-gold-600', bg: 'bg-gold-50 dark:bg-gold-950/30', border: 'border-gold-100 dark:border-gold-800' },
                    { label: 'Categories', value: dashboardStats?.counts.categories || 0, icon: IconLayers, color: 'text-emerald-600', bg: 'bg-emerald-50 dark:bg-emerald-950/30', border: 'border-emerald-100 dark:border-emerald-800' },
                    { label: 'Active Tags', value: dashboardStats?.counts.tags || 0, icon: IconTag, color: 'text-teal-600', bg: 'bg-teal-50 dark:bg-teal-950/30', border: 'border-teal-100 dark:border-teal-800' },
                    { label: 'Media Assets', value: dashboardStats?.counts.files || 0, icon: IconImage, color: 'text-[var(--color-text-secondary)]', bg: 'bg-[var(--color-surface-alt)]', border: 'border-[var(--color-border)]' }
                ].map((stat, i) => (
                    <div key={i} className={`bg-[var(--color-surface)] border ${stat.border} p-6 rounded-2xl shadow-sm hover:shadow-md transition-shadow relative overflow-hidden group cursor-default`}>
                        <div className="flex justify-between items-start mb-4 relative z-10">
                            <div className={`p-3 rounded-xl ${stat.bg}`}>
                                <stat.icon className={`w-6 h-6 ${stat.color}`} />
                            </div>
                        </div>
                        <div className="text-4xl font-serif font-bold text-ink mb-1 relative z-10">{stat.value}</div>
                        <div className="text-xs uppercase tracking-widest text-[var(--color-text-muted)] relative z-10">{stat.label}</div>

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
                        <button onClick={() => onEditPost()} className="text-sm text-gold-600 hover:text-gold-700 font-bold uppercase tracking-wider cursor-pointer">+ Quick Draft</button>
                    </div>
                    <div className="space-y-4">
                        {recentPosts.map(post => (
                            <div key={post.id} className="group flex items-center gap-5 bg-[var(--color-surface)] p-4 rounded-xl border border-[var(--color-border)] hover:border-gold-300 shadow-sm transition-all cursor-pointer" onClick={() => onEditPost(post)}>
                                <div className="w-20 h-20 rounded-lg overflow-hidden bg-[var(--color-surface-alt)] shrink-0 relative">
                                    <img src={post.cover} alt="" className="w-full h-full object-cover group-hover:scale-105 transition-transform duration-500" />
                                </div>
                                <div className="flex-1 min-w-0">
                                    <h4 className="font-serif font-bold text-lg text-ink truncate group-hover:text-gold-600 transition-colors mb-1">{post.title}</h4>
                                    <p className="text-sm text-[var(--color-text-secondary)] truncate mb-2">{post.excerpt}</p>
                                    <div className="flex items-center gap-3 text-xs text-[var(--color-text-muted)]">
                                        <span className="flex items-center gap-1"><IconClock className="w-3 h-3" /> {new Date(post.publishAt).toLocaleString()}</span>
                                        <span className="flex items-center gap-1"><IconEye className="w-3 h-3" /> {post.views}</span>
                                    </div>
                                </div>
                                <div className="text-right shrink-0">
                                    <span className={`inline-flex items-center px-2.5 py-1 rounded-md text-[10px] uppercase tracking-wider font-bold ${post.status === 'published' ? 'bg-emerald-50 dark:bg-emerald-950/30 text-emerald-700 dark:text-emerald-400 border border-emerald-100 dark:border-emerald-800' : 'bg-[var(--color-surface-alt)] text-[var(--color-text-secondary)] border border-[var(--color-border)]'
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
                        <span className="w-2 h-8 bg-[var(--color-text-muted)] rounded-full"></span>
                        System Status
                    </h3>
                    <div className="bg-[var(--color-surface)] p-8 rounded-2xl border border-[var(--color-border)] shadow-sm space-y-8">
                        <div>
                            <div className="flex justify-between text-sm mb-2">
                                <span className="text-[var(--color-text-secondary)] font-medium">Storage Usage</span>
                                <span className="text-ink font-bold">{dashboardStats?.systemStatus.storageUsage || 0}%</span>
                            </div>
                            <div className="w-full h-2.5 bg-[var(--color-surface-alt)] rounded-full overflow-hidden">
                                <div className="h-full bg-gradient-to-r from-gold-400 to-gold-600 rounded-full" style={{ width: `${dashboardStats?.systemStatus.storageUsage || 0}%` }}></div>
                            </div>
                        </div>
                        <div>
                            <div className="flex justify-between text-sm mb-2">
                                <span className="text-[var(--color-text-secondary)] font-medium">AI Quota (Gemini)</span>
                                <span className="text-ink font-bold">{dashboardStats?.systemStatus.aiQuota || 0}%</span>
                            </div>
                            <div className="w-full h-2.5 bg-[var(--color-surface-alt)] rounded-full overflow-hidden">
                                <div className="h-full bg-gradient-to-r from-teal-400 to-teal-600 rounded-full" style={{ width: `${dashboardStats?.systemStatus.aiQuota || 0}%` }}></div>
                            </div>
                        </div>

                        <div className="p-4 bg-[var(--color-surface-alt)] rounded-lg border border-[var(--color-border-subtle)]">
                            <div className="flex items-center gap-2 text-emerald-600 mb-1">
                                <div className="w-2 h-2 rounded-full bg-emerald-500 animate-pulse"></div>
                                <span className="text-xs font-bold uppercase tracking-wider">Operational</span>
                            </div>
                            <p className="text-xs text-[var(--color-text-secondary)] leading-relaxed">
                                System is operating normally. Next automated backup scheduled for 02:00 AM.
                            </p>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    );
};

export default AdminOverview;
