
import React from 'react';
import { AUTHOR_NAME } from '../../../constants';

const ProfileWidget: React.FC = () => {
    // Using a soft gradient text effect
    return (
        <div className="flex flex-col items-center justify-center h-full text-center p-2">
            <div className="w-20 h-20 rounded-full overflow-hidden mb-4 border-4 border-white shadow-sm hover:scale-105 transition-transform duration-500">
                <img
                    src="https://images.unsplash.com/photo-1534528741775-53994a69daeb?auto=format&fit=crop&w=200&q=80"
                    alt="Profile"
                    className="w-full h-full object-cover"
                />
            </div>
            <h2 className="text-xl font-serif text-stone-700 mb-1">
                Good Afternoon
            </h2>
            <p className="text-base text-stone-500 font-light">
                I'm <span className="font-bold text-orange-500/90">{AUTHOR_NAME}</span>, Nice to meet you!
            </p>
        </div>
    );
};

export default ProfileWidget;
