import { Component, type ErrorInfo, type ReactNode } from 'react';
import { IconSparkles } from '@/components/Icons';

interface Props {
    children: ReactNode;
}

interface State {
    hasError: boolean;
    error: Error | null;
}

class ErrorBoundary extends Component<Props, State> {
    public state: State = {
        hasError: false,
        error: null
    };

    public static getDerivedStateFromError(error: Error): State {
        return { hasError: true, error };
    }

    public componentDidCatch(error: Error, errorInfo: ErrorInfo) {
        console.error('Uncaught error:', error, errorInfo);
    }

    render() {
        if (this.state.hasError) {
            return (
                <div className="min-h-screen bg-stone-50 flex flex-col items-center justify-center p-4">
                    <div className="text-center max-w-md">
                        <div className="inline-flex items-center justify-center w-16 h-16 rounded-full bg-red-100 text-red-500 mb-6">
                            <IconSparkles className="w-8 h-8 rotate-180" /> {/* Rotating sparkles for broken magic effect */}
                        </div>
                        <h1 className="text-3xl font-serif font-bold text-ink mb-4">Something went wrong</h1>
                        <p className="text-stone-500 mb-8">
                            We encountered an unexpected error. Please try refreshing the page.
                        </p>
                        <button
                            onClick={() => window.location.reload()}
                            className="px-6 py-3 bg-ink text-white rounded-lg font-medium hover:bg-stone-800 transition-colors cursor-pointer shadow-lg active:scale-95 transform duration-150"
                        >
                            Reload Application
                        </button>
                        {/* Optional: Show error details in dev */}
                        {import.meta.env.DEV && this.state.error && (
                            <div className="mt-8 p-4 bg-red-50 border border-red-100 rounded text-left overflow-auto text-xs text-red-800 font-mono">
                                {this.state.error.toString()}
                            </div>
                        )}
                    </div>
                </div>
            );
        }

        return this.props.children;
    }
}

export default ErrorBoundary;
