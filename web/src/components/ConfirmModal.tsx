import React from 'react';

interface ConfirmModalProps {
    isOpen: boolean;
    title: string;
    message: string;
    confirmText?: string;
    cancelText?: string;
    isDestructive?: boolean;
    onConfirm: () => void;
    onCancel: () => void;
}

const ConfirmModal: React.FC<ConfirmModalProps> = ({
    isOpen,
    title,
    message,
    confirmText = 'Confirm',
    cancelText = 'Cancel',
    isDestructive = true,
    onConfirm,
    onCancel
}) => {
    if (!isOpen) return null;

    return (
        <div className="fixed inset-0 z-[300] flex items-center justify-center p-4">
            <div
                className="absolute inset-0 bg-[var(--color-overlay)] backdrop-blur-sm"
                onClick={onCancel}
            />
            <div className="relative bg-[var(--color-surface)] rounded-2xl shadow-2xl max-w-md w-full p-6 animate-slide-up">
                <h3 className="text-xl font-serif font-bold text-ink mb-2">{title}</h3>
                <p className="text-[var(--color-text-secondary)] mb-6">{message}</p>
                <div className="flex justify-end gap-3">
                    <button
                        onClick={onCancel}
                        className="px-4 py-2 text-[var(--color-text-secondary)] hover:text-ink font-medium cursor-pointer"
                    >
                        {cancelText}
                    </button>
                    <button
                        onClick={onConfirm}
                        className={`px-6 py-2 text-white rounded-lg font-bold transition-colors shadow-lg cursor-pointer ${isDestructive
                                ? 'bg-red-600 hover:bg-red-700 shadow-red-100 dark:shadow-red-900/20'
                                : 'bg-emerald-600 hover:bg-emerald-700 shadow-emerald-100 dark:shadow-emerald-900/20'
                            }`}
                    >
                        {confirmText}
                    </button>
                </div>
            </div>
        </div>
    );
};

export default ConfirmModal;
