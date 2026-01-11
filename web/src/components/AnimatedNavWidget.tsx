
import React, { useState } from 'react';
import { motion } from 'framer-motion';
import { useNavigate, useLocation } from 'react-router-dom';

interface NavItem {
    icon: string;
    label: string;
    labelCn: string;
    path: string;
}

interface AnimatedNavWidgetProps {
    isCompact?: boolean;
}

const navItems: NavItem[] = [
    { icon: '📄', label: 'Posts', labelCn: '近期文章', path: '/posts' },
    { icon: '🔲', label: 'Projects', labelCn: '我的项目', path: '/projects' },
    { icon: '💬', label: 'About', labelCn: '关于网站', path: '/about' },
    { icon: '☆', label: 'Favorites', labelCn: '推荐分享', path: '/favorites' },
    { icon: '🌐', label: 'Blog', labelCn: '优秀博客', path: '/blogs' },
];

const AnimatedNavWidget: React.FC<AnimatedNavWidgetProps> = ({ isCompact = false }) => {
    const navigate = useNavigate();
    const location = useLocation();
    const [hoveredIndex, setHoveredIndex] = useState<number | null>(null);

    const isActive = (path: string) => location.pathname === path;
    const activeIndex = navItems.findIndex(item => isActive(item.path));

    const handleNavClick = (path: string) => {
        navigate(path);
    };

    // Get the indicator position (hover takes priority, then active, then first item)
    const indicatorIndex = hoveredIndex !== null ? hoveredIndex : (activeIndex >= 0 ? activeIndex : 0);

    // Compact horizontal layout
    if (isCompact) {
        return (
            <motion.div
                layoutId="nav-container"
                className="fixed top-6 left-6 z-50 flex items-center gap-1 bg-white/60 backdrop-blur-xl rounded-2xl p-1.5 shadow-lg border border-white/50"
                initial={{ opacity: 0, scale: 0.8 }}
                animate={{ opacity: 1, scale: 1 }}
                transition={{ duration: 0.4, ease: [0.4, 0, 0.2, 1] }}
            >
                {/* Avatar */}
                <motion.button
                    layoutId="nav-avatar"
                    onClick={() => navigate('/')}
                    className="w-10 h-10 rounded-xl bg-orange-100 flex items-center justify-center text-lg cursor-pointer hover:bg-orange-200 transition-colors relative z-10"
                    whileHover={{ scale: 1.05 }}
                    whileTap={{ scale: 0.95 }}
                    aria-label="Go to homepage"
                >
                    🐱
                </motion.button>

                {/* Navigation Icons with Sliding Indicator */}
                <div className="relative flex items-center gap-1">
                    {/* Sliding White Background Indicator */}
                    <motion.div
                        className="absolute w-10 h-10 bg-white rounded-xl shadow-sm"
                        initial={false}
                        animate={{
                            x: indicatorIndex * 44, // 40px width + 4px gap
                        }}
                        transition={{
                            type: 'spring',
                            stiffness: 400,
                            damping: 30,
                        }}
                    />

                    {navItems.map((item, index) => (
                        <motion.button
                            key={item.path}
                            layoutId={`nav-item-${index}`}
                            onClick={() => handleNavClick(item.path)}
                            onMouseEnter={() => setHoveredIndex(index)}
                            onMouseLeave={() => setHoveredIndex(null)}
                            className={`w-10 h-10 rounded-xl flex items-center justify-center text-lg cursor-pointer transition-colors duration-200 relative z-10 ${isActive(item.path)
                                ? 'text-orange-500'
                                : 'text-stone-500 hover:text-orange-500'
                                }`}
                            whileHover={{ scale: 1.05 }}
                            whileTap={{ scale: 0.95 }}
                            aria-label={item.label}
                            title={item.label}
                        >
                            {item.icon}
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
            className="h-auto shadow-sm hover:shadow-md transition-shadow bg-white/40 backdrop-blur-xl border border-white/50 rounded-3xl p-6"
            initial={{ opacity: 0, y: 20 }}
            animate={{ opacity: 1, y: 0 }}
            transition={{ duration: 0.4, ease: [0.4, 0, 0.2, 1] }}
        >
            {/* Header with Avatar */}
            <motion.div layoutId="nav-header" className="flex items-center gap-3 mb-4">
                <motion.div
                    layoutId="nav-avatar"
                    className="w-10 h-10 rounded-full bg-orange-100 flex items-center justify-center text-lg cursor-pointer hover:bg-orange-200 transition-colors"
                    onClick={() => navigate('/')}
                    whileHover={{ scale: 1.05 }}
                    whileTap={{ scale: 0.95 }}
                >
                    🐱
                </motion.div>
                <div>
                    <h3 className="font-bold text-stone-800 text-sm">Voocel</h3>
                    <span className="text-[9px] bg-orange-100 text-orange-600 px-1.5 py-0.5 rounded font-bold">开发中</span>
                </div>
            </motion.div>

            {/* Section Label */}
            <motion.div
                className="text-[9px] text-stone-400 font-bold mb-2 uppercase tracking-wider"
                initial={{ opacity: 0 }}
                animate={{ opacity: 1 }}
                transition={{ delay: 0.2 }}
            >
                General
            </motion.div>

            {/* Navigation Items with Sliding Indicator */}
            <div className="relative flex flex-col gap-1">
                {/* Sliding White Background Indicator */}
                <motion.div
                    className="absolute left-0 right-0 bg-white rounded-xl shadow-sm pointer-events-none"
                    style={{ height: 40 }}
                    initial={false}
                    animate={{
                        y: indicatorIndex * 44, // 40px button height + 4px gap
                    }}
                    transition={{
                        type: 'spring',
                        stiffness: 400,
                        damping: 30,
                    }}
                />

                {navItems.map((item, index) => (
                    <motion.button
                        key={item.path}
                        layoutId={`nav-item-${index}`}
                        onClick={() => handleNavClick(item.path)}
                        onMouseEnter={() => setHoveredIndex(index)}
                        onMouseLeave={() => setHoveredIndex(null)}
                        className="w-full h-10 text-left px-3 rounded-xl flex items-center gap-3 cursor-pointer transition-colors duration-200 relative z-10 group"
                        whileTap={{ scale: 0.98 }}
                    >
                        <div
                            className={`p-1.5 rounded-lg transition-all duration-200 ${isActive(item.path) || hoveredIndex === index
                                ? 'bg-orange-400 text-white'
                                : 'bg-stone-100 text-stone-500'
                                }`}
                        >
                            <span className="text-sm">{item.icon}</span>
                        </div>
                        <span
                            className={`text-sm font-medium transition-colors duration-200 ${isActive(item.path) || hoveredIndex === index
                                ? 'text-stone-800'
                                : 'text-stone-500'
                                }`}
                        >
                            {item.labelCn}
                        </span>
                    </motion.button>
                ))}
            </div>
        </motion.div>
    );
};

export default AnimatedNavWidget;
