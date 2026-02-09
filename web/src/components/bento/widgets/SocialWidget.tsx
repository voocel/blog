
import React from 'react';
import { IconGithub } from '@/components/Icons';

const SocialWidget: React.FC = () => {
    return (
        <div className="flex items-center justify-center gap-3 h-full w-full px-3">
            {/* Github - Pill Shape */}
            <a
                href="https://github.com/voocel"
                target="_blank"
                rel="noopener noreferrer"
                className="flex items-center justify-center bg-black text-white px-5 py-3 rounded-full gap-2 hover:scale-105 transition-transform duration-300 shadow-md group"
            >
                <IconGithub className="w-5 h-5" />
                <span className="font-bold text-sm">Github</span>
            </a>

            {/* Twitter (X) - Pill Shape */}
            <a
                href="https://twitter.com"
                target="_blank"
                rel="noopener noreferrer"
                className="flex items-center justify-center bg-stone-900 text-white px-5 py-3 rounded-full gap-2 hover:scale-105 transition-transform duration-300 shadow-md bg-gradient-to-r from-stone-800 to-stone-900"
            >
                <span className="font-bold text-lg leading-none">𝕏</span>
                <span className="font-bold text-sm">Twitter</span>
            </a>

            {/* Email - Square/Rounded */}
            <div className="bg-orange-100 dark:bg-orange-900/40 text-orange-500 rounded-2xl flex items-center justify-center text-xl shadow-sm cursor-pointer hover:scale-105 transition-transform hover:bg-orange-200 dark:hover:bg-orange-800/40 px-4 py-3">
                ✉️
            </div>
        </div>
    );
};

export default SocialWidget;
