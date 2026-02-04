import React, { useState, useEffect, useRef } from 'react';
import { motion, AnimatePresence } from 'framer-motion';
import { useNavigate } from 'react-router-dom';
import AnimatedNavWidget from '@/components/AnimatedNavWidget';


type Mode = 'stopwatch' | 'timer';

const ClockPage: React.FC = () => {
    const navigate = useNavigate();
    const [mode, setMode] = useState<Mode>('stopwatch');

    // --- Stopwatch State ---
    const [swTime, setSwTime] = useState(0); // in ms
    const [swRunning, setSwRunning] = useState(false);
    const [laps, setLaps] = useState<number[]>([]);

    // --- Timer State ---
    const [timerDuration, setTimerDuration] = useState(5 * 60 * 1000); // default 5 min
    const [timerTimeLeft, setTimerTimeLeft] = useState(5 * 60 * 1000);
    const [timerRunning, setTimerRunning] = useState(false);

    // Refs for intervals
    const swRef = useRef<number | null>(null);
    const timerRef = useRef<number | null>(null);

    // --- Stopwatch Logic ---
    useEffect(() => {
        if (swRunning) {
            const startTime = Date.now() - swTime;
            swRef.current = window.setInterval(() => {
                setSwTime(Date.now() - startTime);
            }, 10);
        } else {
            if (swRef.current) clearInterval(swRef.current);
        }
        return () => { if (swRef.current) clearInterval(swRef.current); };
        // eslint-disable-next-line react-hooks/exhaustive-deps
    }, [swRunning]);

    const handleSwToggle = () => setSwRunning(!swRunning);
    const handleSwReset = () => {
        setSwRunning(false);
        setSwTime(0);
        setLaps([]);
    };
    const handleSwLap = () => {
        setLaps([swTime, ...laps]);
    };

    // --- Timer Logic ---
    useEffect(() => {
        if (timerRunning && timerTimeLeft > 0) {
            const startTime = Date.now();
            const initialLeft = timerTimeLeft;
            timerRef.current = window.setInterval(() => {
                const elapsed = Date.now() - startTime;
                const remaining = initialLeft - elapsed;
                if (remaining <= 0) {
                    setTimerTimeLeft(0);
                    setTimerRunning(false);
                    if (timerRef.current) clearInterval(timerRef.current);
                    // Could add sound or notification here
                } else {
                    setTimerTimeLeft(remaining);
                }
            }, 10);
        } else {
            if (timerRef.current) clearInterval(timerRef.current);
        }
        return () => { if (timerRef.current) clearInterval(timerRef.current); };
        // eslint-disable-next-line react-hooks/exhaustive-deps
    }, [timerRunning]);

    const handleTimerToggle = () => {
        if (timerTimeLeft <= 0) setTimerTimeLeft(timerDuration);
        setTimerRunning(!timerRunning);
    };
    const handleTimerReset = () => {
        setTimerRunning(false);
        setTimerTimeLeft(timerDuration);
    };

    // Format helpers
    const formatMs = (ms: number) => {
        const min = Math.floor(ms / 60000);
        const sec = Math.floor((ms % 60000) / 1000);
        const centi = Math.floor((ms % 1000) / 10);
        return `${min.toString().padStart(2, '0')}:${sec.toString().padStart(2, '0')}:${centi.toString().padStart(2, '0')}`;
    };

    return (
        <div className="min-h-screen bg-[#fdfaf6] bg-[radial-gradient(ellipse_at_center,_var(--tw-gradient-stops))] from-orange-50 via-stone-50 to-stone-100 flex flex-col items-center justify-center relative overflow-hidden text-stone-700 font-sans selection:bg-orange-200">

            {/* Background Blobs */}
            <div className="absolute top-[-10%] left-[-10%] w-[500px] h-[500px] bg-purple-200/30 rounded-full mix-blend-multiply blur-3xl animate-blob pointer-events-none" />
            <div className="absolute bottom-[-10%] right-[-10%] w-[500px] h-[500px] bg-orange-200/30 rounded-full mix-blend-multiply blur-3xl animate-blob animation-delay-2000 pointer-events-none" />

            {/* Nav */}
            <div className="fixed top-8 left-8 z-50">
                <AnimatedNavWidget isCompact={true} disableFixed={true} showBackButton={true} onBackClick={() => navigate('/')} />
            </div>

            <motion.div
                initial={{ opacity: 0, y: 20 }}
                animate={{ opacity: 1, y: 0 }}
                className="relative z-10 w-full max-w-2xl px-6 flex flex-col items-center gap-12"
            >

                {/* Toggle Switch */}
                <div className="flex bg-stone-200/50 p-1.5 rounded-full backdrop-blur-sm shadow-inner">
                    <button
                        onClick={() => setMode('stopwatch')}
                        className={`px-8 py-2 rounded-full text-sm font-bold transition-all duration-300 ${mode === 'stopwatch' ? 'bg-white shadow-sm text-orange-500' : 'text-stone-500 hover:text-stone-700'}`}
                    >
                        Stopwatch
                    </button>
                    <button
                        onClick={() => setMode('timer')}
                        className={`px-8 py-2 rounded-full text-sm font-bold transition-all duration-300 ${mode === 'timer' ? 'bg-white shadow-sm text-orange-500' : 'text-stone-500 hover:text-stone-700'}`}
                    >
                        Timer
                    </button>
                </div>

                {/* Display Container */}
                <motion.div
                    layout
                    className="w-full bg-white/40 backdrop-blur-xl border border-white/60 shadow-xl rounded-[2.5rem] p-6 md:p-10 flex flex-col items-center justify-center gap-6 relative overflow-hidden"
                >
                    <AnimatePresence mode="wait">
                        <motion.div
                            key={mode}
                            initial={{ opacity: 0, scale: 0.9 }}
                            animate={{ opacity: 1, scale: 1 }}
                            exit={{ opacity: 0, scale: 0.9 }}
                            transition={{ duration: 0.2 }}
                            className="flex flex-col items-center gap-8 w-full"
                        >
                            {/* Time Display */}
                            <div className="text-5xl md:text-7xl font-sans font-light tracking-tight tabular-nums text-stone-800 drop-shadow-sm">
                                {mode === 'stopwatch' ? formatMs(swTime) : formatMs(timerTimeLeft)}
                            </div>

                            {/* Controls */}
                            <div className="flex items-center gap-8 md:gap-12 mt-4">
                                {mode === 'stopwatch' && (
                                    <>
                                        <button
                                            onClick={handleSwLap}
                                            disabled={!swRunning}
                                            className="w-14 h-14 rounded-full bg-stone-100 hover:bg-white text-stone-500 font-bold text-sm shadow-sm hover:shadow-md transition-all flex items-center justify-center disabled:opacity-50 disabled:cursor-not-allowed"
                                        >
                                            LAP
                                        </button>
                                        <button
                                            onClick={handleSwToggle}
                                            className={`w-24 h-24 rounded-full flex items-center justify-center text-white text-3xl shadow-lg hover:shadow-xl hover:scale-105 transition-all active:scale-95 ${swRunning ? 'bg-orange-400' : 'bg-green-500'}`}
                                        >
                                            {swRunning ? <span className="ml-1">II</span> : <span className="ml-1">▶</span>}
                                        </button>
                                        <button
                                            onClick={handleSwReset}
                                            className="w-14 h-14 rounded-full bg-stone-100 hover:bg-white text-stone-500 font-bold text-sm shadow-sm hover:shadow-md transition-all flex items-center justify-center"
                                        >
                                            CLR
                                        </button>
                                    </>
                                )}

                                {mode === 'timer' && (
                                    <>
                                        <div className="w-14" /> {/* Spacer */}
                                        <button
                                            onClick={handleTimerToggle}
                                            className={`w-24 h-24 rounded-full flex items-center justify-center text-white text-3xl shadow-lg hover:shadow-xl hover:scale-105 transition-all active:scale-95 ${timerRunning ? 'bg-orange-400' : 'bg-green-500'}`}
                                        >
                                            {timerRunning ? <span className="ml-1">II</span> : <span className="ml-1">▶</span>}
                                        </button>
                                        <button
                                            onClick={handleTimerReset}
                                            className="w-14 h-14 rounded-full bg-stone-100 hover:bg-white text-stone-500 font-bold text-sm shadow-sm hover:shadow-md transition-all flex items-center justify-center"
                                        >
                                            RST
                                        </button>
                                    </>
                                )}
                            </div>
                        </motion.div>
                    </AnimatePresence>

                    {/* Laps Display (Bottom Sheet style inside card) */}
                    {mode === 'stopwatch' && laps.length > 0 && (
                        <motion.div
                            initial={{ opacity: 0, height: 0 }}
                            animate={{ opacity: 1, height: 'auto' }}
                            className="w-full max-h-40 overflow-y-auto mt-4 px-4 custom-scrollbar"
                        >
                            <div className="flex flex-col gap-2 w-full">
                                {laps.map((lap, i) => (
                                    <div key={i} className="flex justify-between text-stone-500 text-sm font-mono border-b border-stone-100 pb-1">
                                        <span>Lap {laps.length - i}</span>
                                        <span>{formatMs(lap)}</span>
                                    </div>
                                ))}
                            </div>
                        </motion.div>
                    )}

                    {/* Timer Presets */}
                    {mode === 'timer' && !timerRunning && (
                        <div className="flex gap-2 mt-4">
                            {[1, 5, 10, 25].map(min => (
                                <button
                                    key={min}
                                    onClick={() => { setTimerDuration(min * 60 * 1000); setTimerTimeLeft(min * 60 * 1000); }}
                                    className={`px-3 py-1 rounded-lg text-xs font-bold transition-all ${timerDuration === min * 60 * 1000 ? 'bg-orange-100 text-orange-600' : 'bg-stone-100 text-stone-500 hover:bg-stone-200'}`}
                                >
                                    {min}m
                                </button>
                            ))}
                        </div>
                    )}

                </motion.div>

            </motion.div>
        </div>
    );
};

export default ClockPage;
