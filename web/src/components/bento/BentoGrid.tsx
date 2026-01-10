
import React, { type ReactNode } from 'react';

interface BentoGridProps {
    children: ReactNode;
}

const BentoGrid: React.FC<BentoGridProps> = ({ children }) => {
    return (
        <div className="grid grid-cols-1 md:grid-cols-4 lg:grid-cols-4 gap-6 auto-rows-[minmax(180px,auto)] max-w-7xl mx-auto p-4 md:p-8">
            {children}
        </div>
    );
};

export default BentoGrid;
