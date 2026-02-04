import { useState, useRef, useCallback, useEffect } from 'react';
import type { EditingPost } from '@/types';

const DRAFT_KEY_PREFIX = 'blog_draft_';

interface UseDraftAutoSaveOptions {
    post: EditingPost | null;
    isEnabled: boolean;
    intervalMs?: number;
    onSave?: () => void;
}

interface UseDraftAutoSaveReturn {
    lastSavedContentRef: React.MutableRefObject<string>;
    showDraftRecovery: boolean;
    setShowDraftRecovery: (show: boolean) => void;
    recoveryDraft: EditingPost | null;
    setRecoveryDraft: (draft: EditingPost | null) => void;
    saveDraft: (post: EditingPost) => void;
    loadDraft: (postId?: number) => EditingPost | null;
    clearDraft: (postId?: number) => void;
    checkForDraft: (postId?: number) => EditingPost | null;
    hasUnsavedChanges: (post: EditingPost | null) => boolean;
}

export function useDraftAutoSave({
    post,
    isEnabled,
    intervalMs = 5000,
    onSave
}: UseDraftAutoSaveOptions): UseDraftAutoSaveReturn {
    const lastSavedContentRef = useRef<string>('');
    const [showDraftRecovery, setShowDraftRecovery] = useState(false);
    const [recoveryDraft, setRecoveryDraft] = useState<EditingPost | null>(null);

    // Get draft key based on post ID
    const getDraftKey = useCallback((postId?: number) => {
        return `${DRAFT_KEY_PREFIX}${postId ?? 'new'}`;
    }, []);

    // Clear draft from localStorage
    const clearDraft = useCallback((postId?: number) => {
        const key = getDraftKey(postId);
        localStorage.removeItem(key);
        lastSavedContentRef.current = '';
    }, [getDraftKey]);

    // Load draft from localStorage
    const loadDraft = useCallback((postId?: number): EditingPost | null => {
        const key = getDraftKey(postId);
        const saved = localStorage.getItem(key);
        if (saved) {
            try {
                return JSON.parse(saved);
            } catch {
                return null;
            }
        }
        return null;
    }, [getDraftKey]);

    // Save draft to localStorage
    const saveDraft = useCallback((postToSave: EditingPost) => {
        const key = getDraftKey(postToSave.id);
        const content = JSON.stringify(postToSave);
        if (content !== lastSavedContentRef.current) {
            localStorage.setItem(key, content);
            lastSavedContentRef.current = content;
            onSave?.();
        }
    }, [getDraftKey, onSave]);

    // Check if a draft exists for a post
    const checkForDraft = useCallback((postId?: number): EditingPost | null => {
        return loadDraft(postId);
    }, [loadDraft]);

    // Check if there are unsaved changes
    const hasUnsavedChanges = useCallback((currentPost: EditingPost | null): boolean => {
        if (!currentPost) return false;
        const currentContent = JSON.stringify(currentPost);
        return currentContent !== lastSavedContentRef.current &&
            Boolean(currentPost.title || currentPost.content);
    }, []);

    // Auto-save effect
    useEffect(() => {
        if (!isEnabled || !post) return;

        const interval = setInterval(() => {
            if (post.title || post.content) {
                saveDraft(post);
            }
        }, intervalMs);

        return () => clearInterval(interval);
    }, [isEnabled, post, saveDraft, intervalMs]);

    return {
        lastSavedContentRef,
        showDraftRecovery,
        setShowDraftRecovery,
        recoveryDraft,
        setRecoveryDraft,
        saveDraft,
        loadDraft,
        clearDraft,
        checkForDraft,
        hasUnsavedChanges
    };
}
