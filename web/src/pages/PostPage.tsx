import React, { useEffect, useState } from 'react';
import { useParams, useNavigate } from 'react-router-dom';
import { useAuth } from '@/context/AuthContext';
import PostDetail from '@/components/PostDetail';
import CommentSection from '@/components/CommentSection';
import { postService } from '@/services/postService';
import type { BlogPost } from '@/types';

const PostPage: React.FC = () => {
    const { slug } = useParams<{ slug: string }>();
    const navigate = useNavigate();
    const { isLoading } = useAuth();
    const [post, setPost] = useState<BlogPost | null>(null);
    const [loading, setLoading] = useState(true);

    useEffect(() => {
        const fetchPost = async () => {
            if (!slug) return;

            // Always fetch from API to trigger view count increment
            try {
                const fetchedPost = await postService.getPost(slug);
                setPost(fetchedPost || null);
            } catch (error) {
                console.error("Failed to load post", error);
            } finally {
                setLoading(false);
            }
        };

        fetchPost();
    }, [slug]);

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
            <PostDetail post={post}>
                <CommentSection postSlug={post.slug} />
            </PostDetail>
        </>
    );
};

export default PostPage;
