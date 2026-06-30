
import React, { useState } from 'react';
import { motion } from 'framer-motion';
import { useNavigate, useLocation } from 'react-router-dom';
import { useTranslation } from '@/hooks/useTranslation';
import { useAuth } from '@/context/AuthContext';
import { IconUser } from '@/components/Icons';

interface NavItem {
    icon: string | React.FC<{ className?: string }>;
    label: string;
    labelCn: string;
    path: string;
    requiresAuth?: boolean;
}

interface AnimatedNavWidgetProps {
    isCompact?: boolean;
    disableFixed?: boolean;
    showBackButton?: boolean;
    onBackClick?: () => void;
}

const navItems: NavItem[] = [
    { icon: '📄', label: 'Posts', labelCn: '近期文章', path: '/posts' },
    { icon: '🔲', label: 'Projects', labelCn: '我的项目', path: '/projects' },
    { icon: 'ⓘ', label: 'About', labelCn: '关于网站', path: '/about' },
    { icon: '☆', label: 'Favorites', labelCn: '推荐分享', path: '/favorites' },
    { icon: '🌐', label: 'Blog', labelCn: '优秀博客', path: '/blogs' },
    { icon: IconUser, label: 'Profile', labelCn: '个人资料', path: '/settings', requiresAuth: true },
];

const AnimatedNavWidget: React.FC<AnimatedNavWidgetProps> = ({ isCompact = false, disableFixed = false, showBackButton = false, onBackClick }) => {
    const navigate = useNavigate();
    const location = useLocation();
    const { locale } = useTranslation();
    const { user, setAuthModalOpen } = useAuth();
    const [hoveredIndex, setHoveredIndex] = useState<number | null>(null);
    const [lastHoveredIndex, setLastHoveredIndex] = useState<number>(0);

    const isActive = (path: string) => location.pathname === path;
    const activeIndex = navItems.findIndex(item => isActive(item.path));

    const handleNavClick = (item: NavItem) => {
        if (item.requiresAuth && !user) {
            setAuthModalOpen(true);
            return;
        }
        navigate(item.path);
    };

    const getLabel = (item: NavItem) => locale === 'zh' ? item.labelCn : item.label;
    const renderIcon = (item: NavItem, className: string) => {
        if (typeof item.icon === 'string') {
            return <span className={className}>{item.icon}</span>;
        }
        const Icon = item.icon;
        return <Icon className={className} />;
    };

    const indicatorIndex = hoveredIndex !== null
        ? hoveredIndex
        : (activeIndex >= 0 ? activeIndex : lastHoveredIndex);

    // Compact horizontal layout
    if (isCompact) {
        const positionClasses = disableFixed
            ? ""
            : "fixed top-6 left-1/2 -translate-x-1/2 md:left-6 md:translate-x-0";

        return (
            <motion.div
                layoutId="nav-container"
                className={`${positionClasses} z-50 flex items-center gap-1 bg-[var(--color-elevated)] backdrop-blur-xl rounded-2xl p-1.5 shadow-lg border border-[var(--color-elevated-border)]`}
                initial={{ opacity: 0, scale: 0.8 }}
                animate={{ opacity: 1, scale: 1 }}
                exit={{ opacity: 0, scale: 0.8 }}
                transition={{ duration: 0.4, ease: [0.4, 0, 0.2, 1] }}
            >
                {showBackButton && (
                    <div className="flex items-center gap-1 mr-1 pr-2 border-r border-[var(--color-border)]/60">
                        <motion.button
                            whileHover={{ scale: 1.05, x: -2 }}
                            whileTap={{ scale: 0.95 }}
                            onClick={() => {
                                if (onBackClick) onBackClick();
                                else navigate(-1);
                            }}
                            className="w-10 h-10 rounded-xl flex items-center justify-center text-[var(--color-text-secondary)] hover:text-orange-500 hover:bg-[var(--color-muted)]/30 cursor-pointer transition-all"
                            title="Go Back"
                        >
                            <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round">
                                <path d="M19 12H5M12 19l-7-7 7-7" />
                            </svg>
                        </motion.button>
                    </div>
                )}

                {/* Avatar */}
                <motion.button
                    layoutId="nav-avatar"
                    onClick={() => navigate('/')}
                    className="w-10 h-10 rounded-full overflow-hidden border border-[var(--color-border)]/50 relative z-10"
                    whileHover={{ scale: 1.05 }}
                    whileTap={{ scale: 0.95 }}
                    aria-label="Go to homepage"
                >
                    <img src="/images/avatar.jpg" alt="Home" className="w-full h-full object-cover" />
                </motion.button>

                {/* Navigation Icons with Sliding Indicator */}
                <div className="relative flex items-center gap-1">
                    <motion.div
                        className="absolute w-10 h-10 bg-[var(--color-surface)] rounded-xl shadow-sm"
                        initial={false}
                        animate={{ x: indicatorIndex * 44 }}
                        transition={{ type: 'spring', stiffness: 400, damping: 30 }}
                    />

                    {navItems.map((item, index) => (
                        <motion.button
                            key={item.path}
                            layoutId={`nav-item-${index}`}
                            onClick={() => handleNavClick(item)}
                            onMouseEnter={() => {
                                setHoveredIndex(index);
                                setLastHoveredIndex(index);
                            }}
                            onMouseLeave={() => setHoveredIndex(null)}
                            className={`w-10 h-10 rounded-xl flex items-center justify-center text-lg cursor-pointer transition-colors duration-200 relative z-10 ${isActive(item.path)
                                ? 'text-orange-500'
                                : 'text-[var(--color-text-secondary)] hover:text-orange-500'
                                }`}
                            whileHover={{ scale: 1.05 }}
                            whileTap={{ scale: 0.95 }}
                            aria-label={getLabel(item)}
                            title={getLabel(item)}
                        >
                            {renderIcon(item, typeof item.icon === 'string' ? 'text-lg leading-none' : 'w-5 h-5')}
                        </motion.button>
                    ))}
                </div>
            </motion.div>
        );
    }

    // Expanded vertical layout
    return (
        <motion.div
            layoutId="nav-container"
            className="h-auto shadow-sm hover:shadow-md transition-shadow bg-[var(--color-elevated)] backdrop-blur-xl border border-[var(--color-elevated-border)] rounded-3xl p-7"
            initial={{ opacity: 0, y: 20 }}
            animate={{ opacity: 1, y: 0 }}
            exit={{ opacity: 0, y: 20 }}
            transition={{ duration: 0.4, ease: [0.4, 0, 0.2, 1] }}
        >
            <motion.div layoutId="nav-header" className="flex items-center gap-3 mb-4">
                <motion.div
                    layoutId="nav-avatar"
                    className="w-12 h-12 rounded-full overflow-hidden border border-[var(--color-border)] cursor-pointer shadow-sm group-hover:shadow-md transition-all"
                    onClick={() => navigate('/')}
                    whileHover={{ scale: 1.05 }}
                    whileTap={{ scale: 0.95 }}
                >
                    <img src="/images/avatar.jpg" alt="Home" className="w-full h-full object-cover" />
                </motion.div>
                <div>
                    <h3 className="font-bold text-ink text-base">Voocel</h3>
                    <span className="text-[10px] bg-orange-100 dark:bg-orange-900/40 text-orange-600 dark:text-orange-400 px-1.5 py-0.5 rounded font-bold uppercase tracking-wider">LAB</span>
                </div>
            </motion.div>

            <motion.div
                className="text-[10px] text-[var(--color-text-muted)] font-bold mb-3 uppercase tracking-wider"
                initial={{ opacity: 0 }}
                animate={{ opacity: 1 }}
                transition={{ delay: 0.2 }}
            >
                General
            </motion.div>

            <div className="relative flex flex-col gap-1.5">
                <motion.div
                    className="absolute left-0 right-0 bg-[var(--color-surface)] rounded-xl shadow-sm pointer-events-none"
                    style={{ height: 48 }}
                    initial={false}
                    animate={{ y: indicatorIndex * 54 }}
                    transition={{ type: 'spring', stiffness: 400, damping: 30 }}
                />

                {navItems.map((item, index) => (
                    <motion.button
                        key={item.path}
                        layoutId={`nav-item-${index}`}
                        onClick={() => handleNavClick(item)}
                        onMouseEnter={() => {
                            setHoveredIndex(index);
                            setLastHoveredIndex(index);
                        }}
                        onMouseLeave={() => setHoveredIndex(null)}
                        className="w-full h-12 text-left px-4 rounded-xl flex items-center gap-4 cursor-pointer transition-colors duration-200 relative z-10 group"
                        whileTap={{ scale: 0.98 }}
                    >
                        <div
                            className={`p-2 rounded-lg transition-all duration-200 ${isActive(item.path) || hoveredIndex === index
                                ? 'bg-orange-400 text-white'
                                : 'bg-[var(--color-surface-alt)] text-[var(--color-text-secondary)]'
                                }`}
                        >
                            {renderIcon(item, typeof item.icon === 'string' ? 'text-base leading-none' : 'w-4 h-4')}
                        </div>
                        <span
                            className={`text-base font-medium transition-colors duration-200 ${isActive(item.path) || hoveredIndex === index
                                ? 'text-ink'
                                : 'text-[var(--color-text-secondary)]'
                                }`}
                        >
                            {getLabel(item)}
                        </span>
                    </motion.button>
                ))}
            </div>
        </motion.div>
    );
};

export default AnimatedNavWidget;
