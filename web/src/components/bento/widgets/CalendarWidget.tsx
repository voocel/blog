
import React from 'react';

const CalendarWidget: React.FC = () => {
    const today = new Date();
    const currentDay = today.getDate();
    const currentMonth = today.toLocaleString('default', { month: 'numeric' });
    const currentYear = today.getFullYear();
    const weekDay = today.toLocaleString('default', { weekday: 'long' });

    // Generate a simple calendar grid (mock for visual, or simple logic)
    // For aesthetics, let's keep it simple: Just show the current month grid roughly

    const daysInMonth = new Date(currentYear, today.getMonth() + 1, 0).getDate();
    const firstDayOfMonth = new Date(currentYear, today.getMonth(), 1).getDay(); // 0 = Sun

    // Adjust to start Monday? Standard JS 0=Sun. Let's do Mon=0 for the grid if we want "Mon Tue..."
    // The screenshot shows "Yi Er San..." (Chinese Mon-Sun).
    // Let's stick to simple numbers.

    const days = Array.from({ length: daysInMonth }, (_, i) => i + 1);
    const blanks = Array.from({ length: (firstDayOfMonth === 0 ? 6 : firstDayOfMonth - 1) }, (_, i) => i); // Mon start adjustment

    const weekHeaders = ['一', '二', '三', '四', '五', '六', '日'];

    return (
        <div className="h-full flex flex-col p-2">
            <div className="mb-4 text-[var(--color-text-secondary)] text-sm font-medium pl-1">
                {currentYear}/{currentMonth}/{currentDay} {weekDay}
            </div>

            <div className="grid grid-cols-7 gap-1 text-center text-xs text-[var(--color-text-muted)] mb-2">
                {weekHeaders.map(d => <div key={d}>{d}</div>)}
            </div>

            <div className="grid grid-cols-7 gap-1 text-center text-sm font-medium text-[var(--color-text-secondary)]">
                {blanks.map((_, i) => <div key={`blank-${i}`} />)}
                {days.map(d => (
                    <div
                        key={d}
                        className={`aspect-square flex items-center justify-center rounded-lg cursor-pointer hover:bg-[var(--color-elevated)] transition-colors ${d === currentDay ? 'bg-orange-500 text-white shadow-md' : ''
                            }`}
                    >
                        {d}
                    </div>
                ))}
            </div>
        </div>
    );
};

export default CalendarWidget;
