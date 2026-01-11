
import React from 'react';
import { useNavigate, useLocation } from 'react-router-dom';

interface NavItem {
    icon: string;
    label: string;
    path: string;
}

const CompactNav: React.FC = () => {
    const navigate = useNavigate();
    const location = useLocation();

    const navItems: NavItem[] = [
        { icon: '📄', label: 'Posts', path: '/posts' },
        { icon: '🔲', label: 'Projects', path: '/projects' },
        { icon: '💬', label: 'About', path: '/about' },
        { icon: '☆', label: 'Favorites', path: '/favorites' },
        { icon: '🌐', label: 'Blog', path: '/' },
    ];

    const isActive = (path: string) => location.pathname === path;

    return (
        <div className="fixed top-6 left-6 z-50 flex items-center gap-1 bg-white/60 backdrop-blur-xl rounded-2xl p-1.5 shadow-lg border border-white/50">
            {/* Avatar */}
            <button
                onClick={() => navigate('/')}
                className="w-10 h-10 rounded-xl bg-orange-100 flex items-center justify-center text-lg hover:scale-105 transition-transform"
                aria-label="Go to homepage"
            >
                🐱
            </button>

            {/* Navigation Icons */}
            {navItems.map((item) => (
                <button
                    key={item.path}
                    onClick={() => navigate(item.path)}
                    className={`w-10 h-10 rounded-xl flex items-center justify-center text-lg transition-all duration-200 ${isActive(item.path)
                            ? 'bg-orange-500 text-white shadow-md shadow-orange-200'
                            : 'hover:bg-white/80 text-stone-500'
                        }`}
                    aria-label={item.label}
                    title={item.label}
                >
                    {item.icon}
                </button>
            ))}
        </div>
    );
};

export default CompactNav;
