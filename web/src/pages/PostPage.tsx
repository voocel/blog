import React, { useEffect, useState } from 'react';
import { useParams, useNavigate } from 'react-router-dom';
import { useBlog } from '../context/BlogContext';
import PostDetail from '../components/PostDetail';
import CommentSection from '../components/CommentSection';
import { postService } from '../services/postService';
import type { BlogPost } from '../types';

const PostPage: React.FC = () => {
    const { id } = useParams<{ id: string }>();
    const navigate = useNavigate();
    const { posts, isLoading } = useBlog();
    const [post, setPost] = useState<BlogPost | null>(null);
    const [loading, setLoading] = useState(true);

    useEffect(() => {
        const fetchPost = async () => {
            if (!id) return;

            // First check if it's already in context
            const existingPost = posts.find(p => p.id === id);
            if (existingPost) {
                setPost(existingPost);
                setLoading(false);
                return;
            }

            // If not in context (e.g. direct link), fetch it
            try {
                const fetchedPost = await postService.getPost(id);
                setPost(fetchedPost || null);
            } catch (error) {
                console.error("Failed to load post", error);
                // navigate('/'); // Optional: redirect to home on error
            } finally {
                setLoading(false);
            }
        };

        if (!isLoading) {
            fetchPost();
        }
    }, [id, posts, isLoading]);

    if (loading || isLoading) {
        return (
            <div className="min-h-screen flex items-center justify-center bg-[#FDFBF7]">
                <div className="flex flex-col items-center gap-4">
                    <div className="w-12 h-12 border-4 border-stone-200 border-t-gold-600 rounded-full animate-spin"></div>
                </div>
            </div>
        );
    }

    if (!post) {
        return (
            <div className="min-h-screen flex items-center justify-center bg-[#FDFBF7]">
                <div className="text-center">
                    <h2 className="text-2xl font-serif text-ink mb-4">Entry Not Found</h2>
                    <button
                        onClick={() => navigate('/')}
                        className="text-gold-600 hover:text-gold-700 underline"
                    >
                        Return to Journal
                    </button>
                </div>
            </div>
        );
    }

    return (
        <>
            <PostDetail post={post} onBack={() => navigate('/')} />
            <div className="bg-[#FDFBF7] pb-24 px-4">
                <CommentSection postId={post.id} />
            </div>
        </>
    );
};

export default PostPage;
