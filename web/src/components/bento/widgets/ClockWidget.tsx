import React, { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';

const ClockWidget: React.FC = () => {
    const navigate = useNavigate();
    const [time, setTime] = useState(new Date());

    useEffect(() => {
        const timer = setInterval(() => setTime(new Date()), 1000);
        return () => clearInterval(timer);
    }, []);

    const formatTime = (date: Date) => {
        return date.toLocaleTimeString('en-US', {
            hour12: false,
            hour: '2-digit',
            minute: '2-digit',
        });
    };

    return (
        <div
            onClick={() => navigate('/clock')}
            className="flex flex-col items-center justify-center h-full w-full bg-stone-100/50 backdrop-blur-sm rounded-2xl cursor-pointer group hover:bg-white/60 transition-all duration-300 relative overflow-hidden"
        >
            {/* Soft decorative blob for aesthetics */}
            <div className="absolute -top-10 -right-10 w-24 h-24 bg-orange-200/20 rounded-full blur-xl group-hover:bg-orange-200/40 transition-colors" />

            <h1 className="text-4xl md:text-5xl font-mono font-light text-stone-700 tracking-tighter tabular-nums selection:bg-transparent relative z-10 group-hover:scale-105 transition-transform duration-500">
                {formatTime(time)}
            </h1>
            <span className="text-[10px] uppercase tracking-[0.3em] text-stone-400 font-bold mt-2 opacity-0 group-hover:opacity-100 transition-opacity transform translate-y-2 group-hover:translate-y-0 duration-300">
                Studio
            </span>
        </div>
    );
};

export default ClockWidget;
