import React, { createContext, useContext, useState, useCallback } from 'react';
import { IconCheck, IconX, IconSparkles } from '@/components/Icons';

export type ToastType = 'success' | 'error' | 'info';

interface Toast {
    id: string;
    message: string;
    details?: string;
    type: ToastType;
}

interface ToastContextType {
    showToast: (message: string, type?: ToastType, details?: string) => void;
}

const ToastContext = createContext<ToastContextType | undefined>(undefined);

// eslint-disable-next-line react-refresh/only-export-components
export const useToast = () => {
    const context = useContext(ToastContext);
    if (!context) {
        throw new Error('useToast must be used within a ToastProvider');
    }
    return context;
};

export const ToastProvider: React.FC<{ children: React.ReactNode }> = ({ children }) => {
    const [toasts, setToasts] = useState<Toast[]>([]);

    const removeToast = useCallback((id: string) => {
        setToasts((prev) => prev.filter((toast) => toast.id !== id));
    }, []);

    const showToast = useCallback((message: string, type: ToastType = 'info', details?: string) => {
        const id = Math.random().toString(36).substr(2, 9);
        setToasts((prev) => [...prev, { id, message, type, details }]);

        // Auto remove
        const duration = details ? 6000 : 3000; // Longer if there are details
        setTimeout(() => {
            removeToast(id);
        }, duration);
    }, [removeToast]);

    return (
        <ToastContext.Provider value={{ showToast }}>
            {children}
            <div className="fixed bottom-6 right-6 z-[9999] flex flex-col gap-3 pointer-events-none">
                {toasts.map((toast) => (
                    <div
                        key={toast.id}
                        className={`pointer-events-auto min-w-[320px] max-w-sm rounded-xl p-4 shadow-2xl backdrop-blur-xl border animate-slide-up flex items-start gap-3 transition-all ${toast.type === 'success'
                            ? 'bg-white/90 border-emerald-100 text-emerald-900'
                            : toast.type === 'error'
                                ? 'bg-white/90 border-red-100 text-red-900'
                                : 'bg-white/90 border-stone-100 text-ink'
                            }`}
                    >
                        <div className={`mt-0.5 p-1 rounded-full shrink-0 ${toast.type === 'success' ? 'bg-emerald-100 text-emerald-600' :
                            toast.type === 'error' ? 'bg-red-100 text-red-600' :
                                'bg-stone-100 text-stone-500'
                            }`}>
                            {toast.type === 'success' && <IconCheck className="w-4 h-4" />}
                            {toast.type === 'error' && <IconX className="w-4 h-4" />}
                            {toast.type === 'info' && <IconSparkles className="w-4 h-4" />}
                        </div>
                        <div className="flex-1 min-w-0">
                            <h4 className="font-bold text-sm leading-tight">{toast.message}</h4>
                            {toast.details && (
                                <p className="text-xs mt-1 opacity-80 break-words leading-relaxed font-mono bg-black/5 p-1.5 rounded">
                                    {toast.details}
                                </p>
                            )}
                        </div>
                        <button
                            onClick={() => removeToast(toast.id)}
                            className="text-black/20 hover:text-black/50 transition-colors -mr-1 -mt-1 p-1"
                        >
                            <IconX className="w-4 h-4" />
                        </button>
                    </div>
                ))}
            </div>
        </ToastContext.Provider>
    );
};
