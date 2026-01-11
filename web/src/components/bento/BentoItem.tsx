
import React, { type ReactNode } from 'react';

interface BentoItemProps {
    children: ReactNode;
    className?: string; // For grid-col-span or height adjustments
    title?: string; // Optional title for the widget
}

const BentoItem: React.FC<BentoItemProps> = ({ children, className = "", title }) => {
    return (
        <div className={`bg-white/40 backdrop-blur-xl border border-white/50 shadow-sm rounded-3xl p-6 transition-all duration-300 hover:scale-[1.03] hover:-translate-y-1 hover:shadow-lg nav-no-drag ${className}`}>
            {title && (
                <h3 className="text-xs font-bold text-stone-400 uppercase tracking-widest mb-4">
                    {title}
                </h3>
            )}
            <div className="h-full w-full">{children}</div>
        </div>
    );
};

export default BentoItem;
