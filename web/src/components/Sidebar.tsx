
import React from 'react';
import { IconHome, IconGrid, IconLayers, IconTag, IconImage, IconLogOut, IconArrowLeft, IconSparkles, IconUserCircle, IconActivity, IconUser, IconMessageSquare } from '@/components/Icons';
import type { AdminSection } from '@/types';
import { useAuth } from '@/context/AuthContext';

interface SidebarProps {
    currentSection: AdminSection;
    setSection: (section: AdminSection) => void;
    onExit: () => void;
}

const Sidebar: React.FC<SidebarProps> = ({ currentSection, setSection, onExit }) => {
    const { logout, user } = useAuth();

    const menuItems: { id: AdminSection; label: string; icon: React.FC<{ className?: string }> }[] = [
        { id: 'overview', label: 'Command Center', icon: IconHome },
        { id: 'posts', label: 'Journal Entries', icon: IconGrid },
        { id: 'categories', label: 'Categories', icon: IconLayers },
        { id: 'tags', label: 'Topics & Tags', icon: IconTag },
        { id: 'comments', label: 'Discussions', icon: IconMessageSquare },
        { id: 'users', label: 'Sanctuary Members', icon: IconUser },
        { id: 'echoes', label: 'Echoes (Analytics)', icon: IconActivity },
        { id: 'files', label: 'Media Assets', icon: IconImage },
    ];

    return (
        <aside className="fixed left-0 top-0 h-full w-72 bg-[var(--color-surface)] border-r border-[var(--color-border)] flex flex-col z-50">
            {/* Brand - Clickable to Home */}
            <div className="h-24 flex items-center px-8">
                <div
                    className="flex items-center gap-3 cursor-pointer group"
                    onClick={onExit}
                    title="Return to Site"
                >
                    <div className="w-8 h-8 rounded-full bg-[var(--color-surface-alt)] border border-[var(--color-border)] flex items-center justify-center shadow-sm group-hover:scale-110 transition-transform">
                        <IconSparkles className="w-4 h-4 text-gold-600" />
                    </div>
                    <span className="text-2xl font-serif font-bold text-ink tracking-tight group-hover:text-gold-600 transition-colors">Aether<span className="text-gold-500">.</span></span>
                </div>
            </div>

            {/* Nav */}
            <nav className="flex-1 px-4 space-y-2 mt-4">
                <p className="px-4 text-[10px] uppercase tracking-[0.2em] text-[var(--color-text-muted)] font-bold mb-4">Workspace</p>
                {menuItems.map((item) => {
                    const Icon = item.icon;
                    const isActive = currentSection === item.id;
                    return (
                        <button
                            key={item.id}
                            onClick={() => setSection(item.id)}
                            className={`group w-full flex items-center gap-4 px-4 py-3.5 rounded-xl transition-all duration-300 relative overflow-hidden cursor-pointer border focus:outline-none ${isActive
                                ? 'bg-[var(--color-surface-alt)] text-ink shadow-sm border-[var(--color-border)] font-bold'
                                : 'border-transparent text-[var(--color-text-secondary)] hover:text-ink hover:bg-[var(--color-surface-alt)]/50'
                                }`}
                        >
                            <Icon className={`w-5 h-5 transition-colors ${isActive ? 'text-gold-600' : 'text-[var(--color-text-muted)] group-hover:text-[var(--color-text-secondary)]'}`} />
                            <span className={`text-sm font-medium tracking-wide ${isActive ? 'font-bold' : ''}`}>{item.label}</span>
                        </button>
                    )
                })}
            </nav>

            {/* User & Footer */}
            <div className="p-4 border-t border-[var(--color-border)] bg-[var(--color-surface-alt)]/50">
                {/* User Mini Profile */}
                <div className="flex items-center gap-3 mb-6 px-2 group">
                    <div className="w-10 h-10 rounded-full bg-[var(--color-surface)] border border-[var(--color-border)] overflow-hidden shadow-sm transition-transform duration-[600ms] ease-in-out group-hover:rotate-[360deg] cursor-pointer">
                        {user?.avatar ? (
                            <img src={user.avatar} alt="User" className="w-full h-full object-cover" />
                        ) : (
                            <div className="w-full h-full flex items-center justify-center bg-[var(--color-surface-alt)] text-[var(--color-text-muted)]">
                                <IconUserCircle className="w-full h-full opacity-60" />
                            </div>
                        )}
                    </div>
                    <div className="flex-1 min-w-0">
                        <p className="text-sm font-bold text-ink truncate font-serif">{user?.username}</p>
                        <div className="flex items-center gap-1.5">
                            <div className="w-1.5 h-1.5 rounded-full bg-emerald-500"></div>
                            <p className="text-[10px] text-[var(--color-text-secondary)] uppercase tracking-wider">Online</p>
                        </div>
                    </div>
                </div>

                <div className="grid grid-cols-2 gap-2">
                    <button
                        onClick={onExit}
                        className="flex flex-col items-center justify-center gap-1 p-3 rounded-lg bg-[var(--color-surface)] border border-[var(--color-border)] text-[var(--color-text-secondary)] hover:text-ink hover:border-stone-300 hover:shadow-sm transition-all cursor-pointer"
                    >
                        <IconArrowLeft className="w-4 h-4" />
                        <span className="text-[10px] uppercase tracking-wider font-medium">Site</span>
                    </button>
                    <button
                        onClick={() => {
                            logout();
                            onExit();
                        }}
                        className="flex flex-col items-center justify-center gap-1 p-3 rounded-lg bg-[var(--color-surface)] border border-[var(--color-border)] text-[var(--color-text-secondary)] hover:text-red-600 hover:border-red-200 hover:shadow-sm transition-all cursor-pointer"
                    >
                        <IconLogOut className="w-4 h-4" />
                        <span className="text-[10px] uppercase tracking-wider font-medium">Logout</span>
                    </button>
                </div>
            </div>
        </aside>
    );
};

export default Sidebar;
