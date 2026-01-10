
import React, { useState } from 'react';

const MediaWidget: React.FC = () => {
    const [isPlaying, setIsPlaying] = useState(false);

    return (
        <div className="flex items-center justify-between h-full px-2 w-full">
            <div className="flex items-center gap-4">
                <div className={`w-12 h-12 rounded-full flex items-center justify-center bg-orange-400 text-white shadow-md cursor-pointer transition-transform ${isPlaying ? 'scale-95' : 'hover:scale-110'}`}
                    onClick={() => setIsPlaying(!isPlaying)}
                >
                    {isPlaying ? (
                        <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="currentColor" className="w-6 h-6">
                            <path fillRule="evenodd" d="M6.75 5.25a.75.75 0 01.75-.75H9a.75.75 0 01.75.75v13.5a.75.75 0 01-.75.75H7.5a.75.75 0 01-.75-.75V5.25zm7.5 0A.75.75 0 0115 4.5h1.5a.75.75 0 01.75.75v13.5a.75.75 0 01-.75.75H15a.75.75 0 01-.75-.75V5.25z" clipRule="evenodd" />
                        </svg>
                    ) : (
                        <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="currentColor" className="w-6 h-6 pl-1">
                            <path fillRule="evenodd" d="M4.5 5.653c0-1.426 1.529-2.33 2.779-1.643l11.54 6.348c1.295.712 1.295 2.573 0 3.285L7.28 19.991c-1.25.687-2.779-.217-2.779-1.643V5.653z" clipRule="evenodd" />
                        </svg>
                    )}
                </div>

                <div className="flex flex-col">
                    <span className="text-stone-700 font-bold text-xs">Close To You</span>
                    <div className="w-24 h-1 bg-stone-200 rounded-full mt-1.5 overflow-hidden">
                        <div className={`h-full bg-orange-400 rounded-full w-1/2 ${isPlaying ? 'animate-pulse' : ''}`}></div>
                    </div>
                </div>
            </div>

            {/* Animated Notes */}
            <div className="text-lg animate-bounce delay-700">🎵</div>
        </div>
    );
};

export default MediaWidget;
