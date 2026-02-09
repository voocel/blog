import React from 'react';
import { IconSparkles } from '@/components/Icons';

const QUOTES = [
    { text: "Code is poetry, and design is the silent language of connection.", author: "Digital Soul" },
    { text: "Simplicity is the ultimate sophistication.", author: "Leonardo da Vinci" },
    { text: "Design is not just what it looks like. Design is how it works.", author: "Steve Jobs" },
    { text: "Digital spaces should feel like warm homes, not cold machines.", author: "Manifesto" },
    { text: "The details are not the details. They make the design.", author: "Charles Eames" },
    { text: "Creativity is intelligence having fun.", author: "Albert Einstein" },
    { text: "Less is more.", author: "Mies van der Rohe" },
    { text: "Good design is as little design as possible.", author: "Dieter Rams" },
    { text: "In a world of noise, silence is a luxury.", author: "Minimalism" },
    { text: "Technology best serves us when it's invisible.", author: "Human Interface" },
];

const QuoteWidget: React.FC = () => {
    // Select quote based on the current date to rotate daily
    const dailyQuote = React.useMemo(() => {
        const today = new Date();
        const dateString = today.toISOString().split('T')[0]; // YYYY-MM-DD

        // Simple hash function to get a stable number from the date string
        let hash = 0;
        for (let i = 0; i < dateString.length; i++) {
            hash = dateString.charCodeAt(i) + ((hash << 5) - hash);
        }

        const index = Math.abs(hash) % QUOTES.length;
        return QUOTES[index];
    }, []);

    return (
        <div className="flex flex-col items-center justify-center h-full text-center p-6 relative overflow-hidden group">
            {/* Background Decor */}
            <div className="absolute top-0 right-0 p-4 opacity-10 group-hover:opacity-20 transition-opacity duration-500">
                <IconSparkles className="w-12 h-12 text-orange-400" />
            </div>

            <div className="relative z-10 animate-fade-in">
                <div className="text-3xl font-serif text-[var(--color-text-muted)] mb-2 leading-none">"</div>
                <p className="text-lg font-serif text-ink italic leading-relaxed mb-3 line-clamp-3">
                    {dailyQuote.text}
                </p>
                <div className="h-px w-12 bg-orange-200 mx-auto mb-3"></div>
                <p className="text-[10px] uppercase tracking-[0.2em] text-[var(--color-text-muted)] font-bold">
                    {dailyQuote.author}
                </p>
            </div>
        </div>
    );
};

export default QuoteWidget;
