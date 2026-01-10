
import React, { useState, useEffect } from 'react';

const ClockWidget: React.FC = () => {
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
        <div className="flex items-center justify-center h-full w-full bg-stone-200/30 rounded-2xl inner-shadow px-8">
            <h1 className="text-4xl md:text-5xl font-mono font-light text-stone-700 tracking-tighter tabular-nums selection:bg-transparent">
                {formatTime(time)}
            </h1>
        </div>
    );
};

export default ClockWidget;
