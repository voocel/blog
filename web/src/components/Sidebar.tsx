
import React from 'react';
import { IconHome, IconGrid, IconLayers, IconTag, IconImage, IconLogOut, IconArrowLeft, IconSparkles, IconUserCircle, IconActivity } from './Icons';
import type { AdminSection } from '../types';
import { useBlog } from '../context/BlogContext';

interface SidebarProps {
    currentSection: AdminSection;
    setSection: (section: AdminSection) => void;
    onExit: () => void;
}

const Sidebar: React.FC<SidebarProps> = ({ currentSection, setSection, onExit }) => {
    const { logout, user } = useBlog();

    const menuItems: { id: AdminSection; label: string; icon: React.FC<any> }[] = [
        { id: 'overview', label: 'Command Center', icon: IconHome },
        { id: 'echoes', label: 'Echoes (Analytics)', icon: IconActivity },
        { id: 'posts', label: 'Journal Entries', icon: IconGrid },
        { id: 'categories', label: 'Categories', icon: IconLayers },
        { id: 'tags', label: 'Topics & Tags', icon: IconTag },
        { id: 'files', label: 'Media Assets', icon: IconImage },
    ];

    return (
        <aside className="fixed left-0 top-0 h-full w-72 bg-[#F7F5F3] border-r border-[#E7E5E4] flex flex-col z-50">
            {/* Brand - Clickable to Home */}
            <div
                className="h-24 flex items-center px-8 cursor-pointer group"
                onClick={onExit}
                title="Return to Site"
            >
                <div className="flex items-center gap-3">
                    <div className="w-8 h-8 rounded-full bg-white border border-stone-200 flex items-center justify-center shadow-sm group-hover:scale-110 transition-transform">
                        <IconSparkles className="w-4 h-4 text-gold-600" />
                    </div>
                    <span className="text-2xl font-serif font-bold text-ink tracking-tight group-hover:text-gold-600 transition-colors">Aether<span className="text-gold-500">.</span></span>
                </div>
            </div>

            {/* Nav */}
            <nav className="flex-1 px-4 space-y-2 mt-4">
                <p className="px-4 text-[10px] uppercase tracking-[0.2em] text-stone-400 font-bold mb-4">Workspace</p>
                {menuItems.map((item) => {
                    const Icon = item.icon;
                    const isActive = currentSection === item.id;
                    return (
                        <button
                            key={item.id}
                            onClick={() => setSection(item.id)}
                            className={`group w-full flex items-center gap-4 px-4 py-3.5 rounded-xl transition-all duration-300 relative overflow-hidden cursor-pointer border focus:outline-none ${isActive
                                ? 'bg-white text-ink shadow-[0_2px_8px_rgba(0,0,0,0.04)] border-stone-100'
                                : 'border-transparent text-stone-500 hover:text-stone-700 hover:bg-white/50'
                                }`}
                        >
                            <Icon className={`w-5 h-5 transition-colors ${isActive ? 'text-gold-600' : 'text-stone-400 group-hover:text-stone-600'}`} />
                            <span className={`text-sm font-medium tracking-wide ${isActive ? 'font-bold' : ''}`}>{item.label}</span>
                        </button>
                    )
                })}
            </nav>

            {/* User & Footer */}
            <div className="p-4 border-t border-stone-200 bg-[#F5F5F4]/50">
                {/* User Mini Profile */}
                <div className="flex items-center gap-3 mb-6 px-2 group">
                    <div className="w-10 h-10 rounded-full bg-white border border-stone-200 overflow-hidden shadow-sm transition-transform duration-[600ms] ease-in-out group-hover:rotate-[360deg] cursor-pointer">
                        {user?.avatar ? (
                            <img src={user.avatar} alt="User" className="w-full h-full object-cover" />
                        ) : (
                            <div className="w-full h-full flex items-center justify-center bg-stone-100 text-stone-400">
                                <IconUserCircle className="w-full h-full opacity-60" />
                            </div>
                        )}
                    </div>
                    <div className="flex-1 min-w-0">
                        <p className="text-sm font-bold text-ink truncate font-serif">{user?.username}</p>
                        <div className="flex items-center gap-1.5">
                            <div className="w-1.5 h-1.5 rounded-full bg-emerald-500"></div>
                            <p className="text-[10px] text-stone-500 uppercase tracking-wider">Online</p>
                        </div>
                    </div>
                </div>

                <div className="grid grid-cols-2 gap-2">
                    <button
                        onClick={onExit}
                        className="flex flex-col items-center justify-center gap-1 p-3 rounded-lg bg-white border border-stone-200 text-stone-500 hover:text-ink hover:border-stone-300 hover:shadow-sm transition-all cursor-pointer"
                    >
                        <IconArrowLeft className="w-4 h-4" />
                        <span className="text-[10px] uppercase tracking-wider font-medium">Site</span>
                    </button>
                    <button
                        onClick={logout}
                        className="flex flex-col items-center justify-center gap-1 p-3 rounded-lg bg-white border border-stone-200 text-stone-500 hover:text-red-600 hover:border-red-200 hover:shadow-sm transition-all cursor-pointer"
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
