import type { Comment } from '../types';

interface CommentResponse {
    data: Comment[];
    pagination: {
        total: number;
        page: number;
        limit: number;
        totalPages: number;
    };
}

// Mock Data
const MOCK_COMMENTS: Record<string, Comment[]> = {
    '1': [
        {
            id: 'c1',
            parentId: null,
            content: "This is exactly what I've been looking for. The design principles you mentioned are spot on!",
            createdAt: new Date(Date.now() - 86400000 * 2).toISOString(), // 2 days ago
            user: { username: 'alex_dev', avatar: 'https://api.dicebear.com/7.x/miniavs/svg?seed=1' },
            replies: [
                {
                    id: 'c1-r1',
                    parentId: 'c1',
                    content: "Glad you liked it! Which principle resonated with you the most?",
                    createdAt: new Date(Date.now() - 86400000 * 1.8).toISOString(),
                    user: { username: 'voocel', avatar: 'https://api.dicebear.com/7.x/miniavs/svg?seed=admin' },
                    replyToUser: { username: 'alex_dev', avatar: 'https://api.dicebear.com/7.x/miniavs/svg?seed=1' }
                }
            ]
        },
        {
            id: 'c2',
            parentId: null,
            content: "Great article! I'd love to see a follow-up on the implementation details.",
            createdAt: new Date(Date.now() - 3600000 * 4).toISOString(), // 4 hours ago
            user: { username: 'sarah_design', avatar: 'https://api.dicebear.com/7.x/miniavs/svg?seed=2' },
            replies: []
        }
    ]
};

export const commentService = {
    getComments: async (postId: string, page = 1, limit = 20): Promise<CommentResponse> => {
        // Simulate network delay
        await new Promise(resolve => setTimeout(resolve, 800));

        const comments = MOCK_COMMENTS[postId] || [];
        // In a real app, we would paginate here. For mock, we just return all.

        return {
            data: comments,
            pagination: {
                total: comments.length,
                page,
                limit,
                totalPages: 1
            }
        };
    },

    createComment: async (postId: string, content: string, user: { username: string; avatar?: string }, parentId?: string): Promise<Comment> => {
        await new Promise(resolve => setTimeout(resolve, 600));

        let replyToUser;
        if (parentId) {
            const comments = MOCK_COMMENTS[postId] || [];
            // Simplified search for mock (only 1 level deep search needed for replyToUser logic really, but keeping it simple)
            const parent = comments.find(c => c.id === parentId);
            if (parent) {
                replyToUser = parent.user;
            }
        }

        const newComment: Comment = {
            id: Math.random().toString(36).substr(2, 9),
            parentId: parentId || null,
            content,
            createdAt: new Date().toISOString(),
            user,
            replies: [],
            replyToUser
        };

        if (!MOCK_COMMENTS[postId]) {
            MOCK_COMMENTS[postId] = [];
        }

        if (parentId) {
            const parent = MOCK_COMMENTS[postId].find(c => c.id === parentId);
            if (parent) {
                if (!parent.replies) parent.replies = [];
                parent.replies.push(newComment);
            }
        } else {
            MOCK_COMMENTS[postId].unshift(newComment);
        }

        return newComment;
    }
};
