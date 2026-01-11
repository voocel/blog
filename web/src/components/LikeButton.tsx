
import React, { useState, useCallback } from 'react';
import { motion, AnimatePresence } from 'framer-motion';

interface FloatingHeart {
    id: number;
    x: number;
    color: string;
    scale: number;
    rotation: number;
}

interface LikeButtonProps {
    initialCount?: number;
    onLike?: () => Promise<void>;
}

const heartColors = [
    '#FF6B8A', // Pink
    '#FF4D6D', // Coral
    '#FF8FA3', // Light pink
    '#FF0A54', // Hot pink
    '#FF85A1', // Rose
];

const LikeButton: React.FC<LikeButtonProps> = ({ initialCount = 0, onLike }) => {
    const [count, setCount] = useState(initialCount);
    const [floatingHearts, setFloatingHearts] = useState<FloatingHeart[]>([]);
    const [isAnimating, setIsAnimating] = useState(false);

    const handleClick = useCallback(async () => {
        // Trigger button animation
        setIsAnimating(true);
        setTimeout(() => setIsAnimating(false), 150);

        // Optimistically update count
        setCount(prev => prev + 1);

        // Spawn floating hearts (3-5 random hearts)
        const numHearts = 3 + Math.floor(Math.random() * 3);
        const newHearts: FloatingHeart[] = [];

        for (let i = 0; i < numHearts; i++) {
            newHearts.push({
                id: Date.now() + i,
                x: (Math.random() - 0.5) * 60, // Random horizontal offset
                color: heartColors[Math.floor(Math.random() * heartColors.length)],
                scale: 0.6 + Math.random() * 0.6, // Random size 0.6-1.2
                rotation: (Math.random() - 0.5) * 40, // Random rotation
            });
        }

        setFloatingHearts(prev => [...prev, ...newHearts]);

        // Remove hearts after animation completes
        setTimeout(() => {
            setFloatingHearts(prev => prev.filter(h => !newHearts.find(nh => nh.id === h.id)));
        }, 2000);

        // Call API if provided
        if (onLike) {
            try {
                await onLike();
            } catch (error) {
                console.error('Failed to like:', error);
                // Revert on failure
                setCount(prev => prev - 1);
            }
        }
    }, [onLike]);

    // Update count when initialCount changes
    React.useEffect(() => {
        setCount(initialCount);
    }, [initialCount]);

    return (
        <div className="relative flex items-center gap-2">
            {/* Floating Hearts Container */}
            <div className="absolute inset-0 pointer-events-none overflow-visible" style={{ zIndex: 100 }}>
                <AnimatePresence>
                    {floatingHearts.map(heart => (
                        <motion.div
                            key={heart.id}
                            className="absolute"
                            style={{
                                left: '50%',
                                bottom: '100%',
                            }}
                            initial={{
                                opacity: 1,
                                scale: 0,
                                x: heart.x,
                                y: 0,
                                rotate: heart.rotation,
                            }}
                            animate={{
                                opacity: [1, 1, 0],
                                scale: [0, heart.scale, heart.scale * 0.5],
                                y: -120 - Math.random() * 40,
                                x: heart.x + (Math.random() - 0.5) * 30,
                                rotate: heart.rotation + (Math.random() - 0.5) * 20,
                            }}
                            exit={{ opacity: 0 }}
                            transition={{
                                duration: 1.5 + Math.random() * 0.5,
                                ease: 'easeOut',
                            }}
                        >
                            <svg
                                width="24"
                                height="24"
                                viewBox="0 0 24 24"
                                fill={heart.color}
                                style={{ filter: `drop-shadow(0 0 4px ${heart.color}40)` }}
                            >
                                <path d="M12 21.35l-1.45-1.32C5.4 15.36 2 12.28 2 8.5 2 5.42 4.42 3 7.5 3c1.74 0 3.41.81 4.5 2.09C13.09 3.81 14.76 3 16.5 3 19.58 3 22 5.42 22 8.5c0 3.78-3.4 6.86-8.55 11.54L12 21.35z" />
                            </svg>
                        </motion.div>
                    ))}
                </AnimatePresence>
            </div>

            {/* Like Button */}
            <motion.button
                onClick={handleClick}
                className="relative w-14 h-14 rounded-full bg-white/60 backdrop-blur-xl border border-white/50 shadow-lg flex items-center justify-center cursor-pointer hover:bg-white/80 transition-colors"
                animate={isAnimating ? { scale: 1.2 } : { scale: 1 }}
                whileHover={{ scale: 1.05 }}
                whileTap={{ scale: 0.95 }}
                aria-label="Like this page"
            >
                <svg
                    width="28"
                    height="28"
                    viewBox="0 0 24 24"
                    fill="#FF6B8A"
                    className="transition-transform"
                    style={{ filter: 'drop-shadow(0 0 6px rgba(255,107,138,0.4))' }}
                >
                    <path d="M12 21.35l-1.45-1.32C5.4 15.36 2 12.28 2 8.5 2 5.42 4.42 3 7.5 3c1.74 0 3.41.81 4.5 2.09C13.09 3.81 14.76 3 16.5 3 19.58 3 22 5.42 22 8.5c0 3.78-3.4 6.86-8.55 11.54L12 21.35z" />
                </svg>
            </motion.button>

            {/* Like Count Badge */}
            <motion.div
                className="absolute -top-2 -right-2 bg-pink-400 text-white text-xs font-bold px-2 py-0.5 rounded-full shadow-md min-w-[32px] text-center"
                key={count}
                initial={{ scale: 1.3 }}
                animate={{ scale: 1 }}
                transition={{ type: 'spring', stiffness: 400, damping: 15 }}
            >
                {count >= 1000 ? `${(count / 1000).toFixed(1)}k` : count}
            </motion.div>
        </div>
    );
};

export default LikeButton;
