
import React from 'react';
import { useNavigate } from 'react-router-dom';
import { IconUser, IconHome as IconLayout, IconMenu as IconHashtag } from '../../Icons';

interface NavItem {
    label: string;
    icon: React.ReactNode;
    path?: string;
    action?: () => void;
    isDev?: boolean;
}

const NavWidget: React.FC = () => {
    const navigate = useNavigate();

    const navItems: NavItem[] = [
        { label: '近期文章', icon: <IconLayout className="w-4 h-4" />, path: '/posts' },
        { label: '我的项目', icon: <IconHashtag className="w-4 h-4" />, path: '/projects', isDev: true },
        { label: '关于网站', icon: <IconUser className="w-4 h-4" />, path: '/about' },
    ];

    return (
        <div className="flex flex-col h-full justify-center">
            <div className="flex items-center gap-3 mb-6">
                <div className="w-10 h-10 rounded-full bg-orange-100 flex items-center justify-center text-xl">
                    🐱
                </div>
                <div>
                    <h3 className="font-bold text-stone-700">Voocel</h3>
                    <span className="text-[10px] text-orange-500 bg-orange-100 px-2 py-0.5 rounded-full font-bold">开发中</span>
                </div>
            </div>

            <div className="space-y-1">
                <p className="text-[10px] text-stone-400 uppercase tracking-widest mb-2 pl-2">General</p>
                {navItems.map((item, index) => (
                    <button
                        key={index}
                        onClick={() => item.path && navigate(item.path)}
                        className={`w-full text-left px-3 py-2 rounded-xl flex items-center gap-3 transition-all duration-300 group ${item.label === '我的项目' ? 'bg-white shadow-sm' : 'hover:bg-white/50'
                            }`}
                    >
                        <div className={`p-1.5 rounded-lg ${item.label === '我的项目' ? 'bg-orange-500 text-white' : 'bg-stone-100 text-stone-500 group-hover:bg-white group-hover:text-orange-500'
                            } transition-colors`}>
                            {item.icon}
                        </div>
                        <span className={`text-sm font-medium ${item.label === '我的项目' ? 'text-stone-800' : 'text-stone-500'}`}>
                            {item.label}
                        </span>
                    </button>
                ))}
            </div>
        </div>
    );
};

export default NavWidget;
