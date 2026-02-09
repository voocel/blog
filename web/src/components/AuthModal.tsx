
import React, { useState } from 'react';
import { useAuth } from '@/context/AuthContext';
import { IconX, IconLock, IconUser } from '@/components/Icons';

const AuthModal: React.FC = () => {
    const { isAuthModalOpen, setAuthModalOpen, login, register, user } = useAuth();
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const [error, setError] = useState('');
    const [activeTab, setActiveTab] = useState<'login' | 'register'>('login');

    // Auto-close if user is already logged in
    React.useEffect(() => {
        if (user && isAuthModalOpen) {
            setAuthModalOpen(false);
        }
    }, [user, isAuthModalOpen, setAuthModalOpen]);

    if (!isAuthModalOpen) return null;

    const handleSubmit = async (e: React.FormEvent) => {
        e.preventDefault();
        setError('');

        try {
            if (activeTab === 'login') {
                // For login, we use email as the identifier in this system
                await login(email, password);
                // If successful (no error thrown)
                setError('');
                setEmail('');
                setPassword('');
                setAuthModalOpen(false);
            } else {
                // Registration
                if (!email || !password) {
                    setError('All fields are required.');
                    return;
                }

                await register(email, password);
                // If successful
                setError('');
                setEmail('');
                setPassword('');
                setAuthModalOpen(false);
            }
        } catch (err: unknown) {
            console.error("Auth error:", err);
            // Extract error message from backend response if available
            const axiosError = err as { response?: { data?: { error?: string } }; message?: string };
            const backendError = axiosError.response?.data?.error || axiosError.message || 'Authentication failed';
            setError(backendError);
        }
    };

    return (
        <div className="fixed inset-0 z-[60] flex items-center justify-center p-4">
            <div className="absolute inset-0 bg-[var(--color-overlay)] backdrop-blur-sm cursor-pointer" onClick={() => setAuthModalOpen(false)} />

            <div className="relative w-full max-w-md bg-[var(--color-surface)] border border-[var(--color-border)] rounded-2xl shadow-2xl overflow-hidden animate-slide-up">
                {/* Header */}
                <div className="p-6 border-b border-[var(--color-border-subtle)] flex justify-between items-center bg-[var(--color-surface-alt)]">
                    <h2 className="text-xl font-serif text-ink">
                        {activeTab === 'login' ? 'Welcome Back' : 'Join Lumina'}
                    </h2>
                    <button onClick={() => setAuthModalOpen(false)} className="text-[var(--color-text-muted)] hover:text-ink cursor-pointer">
                        <IconX className="w-5 h-5" />
                    </button>
                </div>

                {/* Content */}
                <div className="p-8">
                    <div className="flex gap-4 mb-8 p-1 bg-[var(--color-surface-alt)] rounded-lg">
                        <button
                            onClick={() => { setActiveTab('login'); setError(''); }}
                            className={`flex-1 py-2 text-sm font-medium rounded-md transition-all cursor-pointer ${activeTab === 'login' ? 'bg-[var(--color-surface)] text-ink shadow-sm' : 'text-[var(--color-text-secondary)] hover:text-ink'
                                }`}
                        >
                            Login
                        </button>
                        <button
                            onClick={() => { setActiveTab('register'); setError(''); }}
                            className={`flex-1 py-2 text-sm font-medium rounded-md transition-all cursor-pointer ${activeTab === 'register' ? 'bg-[var(--color-surface)] text-ink shadow-sm' : 'text-[var(--color-text-secondary)] hover:text-ink'
                                }`}
                        >
                            Register
                        </button>
                    </div>

                    <form onSubmit={handleSubmit} className="space-y-4">
                        {/* Email Field */}
                        <div>
                            <label className="block text-xs uppercase tracking-wider text-[var(--color-text-secondary)] mb-2">Email</label>
                            <div className="relative">
                                <IconUser className="absolute left-3 top-1/2 -translate-y-1/2 w-4 h-4 text-[var(--color-text-muted)]" />
                                <input
                                    type="email"
                                    value={email}
                                    onChange={(e) => setEmail(e.target.value)}
                                    className="w-full bg-[var(--color-surface-alt)] border border-[var(--color-border)] rounded-lg py-3 pl-10 pr-4 text-ink focus:border-gold-500 focus:outline-none transition-colors placeholder-[var(--color-text-muted)]"
                                    placeholder="Enter your email"
                                />
                            </div>
                        </div>

                        {/* Password Field */}
                        <div>
                            <label className="block text-xs uppercase tracking-wider text-[var(--color-text-secondary)] mb-2">Password</label>
                            <div className="relative">
                                <IconLock className="absolute left-3 top-1/2 -translate-y-1/2 w-4 h-4 text-[var(--color-text-muted)]" />
                                <input
                                    type="password"
                                    value={password}
                                    onChange={(e) => setPassword(e.target.value)}
                                    className="w-full bg-[var(--color-surface-alt)] border border-[var(--color-border)] rounded-lg py-3 pl-10 pr-4 text-ink focus:border-gold-500 focus:outline-none transition-colors placeholder-[var(--color-text-muted)]"
                                    placeholder="Enter password"
                                />
                            </div>
                        </div>

                        {error && <p className="text-red-500 text-xs text-center">{error}</p>}

                        <button
                            type="submit"
                            className="w-full bg-gradient-to-r from-gold-600 to-gold-500 hover:from-gold-500 hover:to-gold-400 text-white font-bold py-3 rounded-lg shadow-lg shadow-gold-100 dark:shadow-gold-900/20 transition-all transform active:scale-[0.98] cursor-pointer"
                        >
                            {activeTab === 'login' ? 'Access Journal' : 'Create Account'}
                        </button>
                    </form>


                </div>
            </div>
        </div>
    );
};

export default AuthModal;
