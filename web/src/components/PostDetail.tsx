
import React, { useState, useEffect, type PropsWithChildren } from 'react';
import { useNavigate } from 'react-router-dom';
import MDEditor from '@uiw/react-md-editor';
import type { BlogPost } from '../types';
import { getAssetUrl } from '../utils/urlUtils';
import { IconSparkles, IconBrain } from './Icons';
import { generateSummary, generateInsight } from '../services/geminiService';
import { postService } from '../services/postService';
import SEO from './SEO';
import AnimatedNavWidget from './AnimatedNavWidget';
import LikeButton from './LikeButton';

interface PostDetailProps {
    post: BlogPost;
}

const PostDetail: React.FC<PropsWithChildren<PostDetailProps>> = ({ post, children }) => {
    const navigate = useNavigate();
    const [summary, setSummary] = useState<string | null>(null);
    const [insight, setInsight] = useState<string | null>(null);
    const [loadingSummary, setLoadingSummary] = useState(false);
    const [loadingInsight, setLoadingInsight] = useState(false);
    const [articleLikes, setArticleLikes] = useState(0);
    const [tocItems, setTocItems] = useState<{ id: string; text: string; level: number }[]>([]);

    useEffect(() => {
        window.scrollTo(0, 0);

        // Fetch article likes
        postService.getLikes(`post-${post.slug}`)
            .then(count => setArticleLikes(count))
            .catch(err => console.error('Failed to fetch article likes:', err));

        // Generate TOC from content
        const headingRegex = /^(#{1,3})\s+(.+)$/gm;
        const matches = [...post.content.matchAll(headingRegex)];
        const items = matches.map((match, index) => ({
            id: `heading-${index}`,
            text: match[2],
            level: match[1].length,
        }));
        setTocItems(items);
    }, [post.slug, post.content]);

    const handleGenerateSummary = async () => {
        setLoadingSummary(true);
        const result = await generateSummary(post.content);
        setSummary(result);
        setLoadingSummary(false);
    };

    const handleGenerateInsight = async () => {
        setLoadingInsight(true);
        const result = await generateInsight(post.content);
        setInsight(result);
        setLoadingInsight(false);
    };

    const handleBack = () => {
        if (window.history.length > 1) {
            navigate(-1);
        } else {
            navigate('/');
        }
    };

    return (
        <div className="min-h-screen bg-[#F3F0E9] bg-[radial-gradient(circle_at_top_left,_#FFD6D6_0%,_transparent_40%),radial-gradient(circle_at_bottom_right,_#FFE8AB_0%,_transparent_40%)] text-stone-700">
            <SEO
                title={post.title}
                description={post.excerpt}
                image={getAssetUrl(post.cover)}
                type="article"
            />

            {/* Top Navigation Area */}
            <div className="fixed top-6 left-6 z-50 flex items-center gap-4">
                {/* Compact Nav Widget with Embedded Back Button */}
                <AnimatedNavWidget isCompact={true} disableFixed={true} showBackButton={true} onBackClick={handleBack} />
            </div>

            {/* Main Layout Grid */}
            <div className="max-w-7xl mx-auto pt-32 pb-20 px-6 flex flex-col lg:flex-row gap-8">

                {/* Left Column: Main Content Card & Comments */}
                <main className="flex-1 min-w-0">
                    <div className="bg-white/90 backdrop-blur-sm rounded-[2rem] shadow-sm p-12 md:p-16">

                        {/* Article Header */}
                        <header className="text-center mb-16">
                            <h1 className="text-3xl md:text-4xl lg:text-5xl font-bold text-stone-800 mb-6 leading-tight font-serif">
                                {post.title}
                            </h1>

                            <div className="flex flex-col items-center gap-3 text-sm">
                                <div className="flex items-center gap-2">
                                    {post.tags.map(tag => (
                                        <span key={tag} className="text-stone-400 font-medium">
                                            #{tag}
                                        </span>
                                    ))}
                                </div>
                                <time className="text-stone-400">
                                    {new Date(post.publishAt).toLocaleDateString('zh-CN', {
                                        year: 'numeric',
                                        month: 'long',
                                        day: 'numeric'
                                    })}
                                </time>
                            </div>
                        </header>

                        {/* AI Features (Moved to Top) */}
                        <div className="mb-16 grid grid-cols-1 md:grid-cols-2 gap-6">
                            {/* Summary Block */}
                            <div className="bg-orange-50/50 rounded-2xl p-6 border border-orange-100/50 hover:border-orange-200 transition-colors">
                                <div className="flex items-center gap-2 mb-3 text-orange-600">
                                    <IconBrain className="w-5 h-5" />
                                    <h3 className="font-bold uppercase tracking-widest text-xs">AI Synopsis</h3>
                                </div>
                                {summary ? (
                                    <div className="text-sm text-stone-700 leading-relaxed animate-fade-in font-serif">
                                        <MDEditor.Markdown source={summary} style={{ background: 'transparent', color: 'inherit' }} />
                                    </div>
                                ) : (
                                    <button
                                        onClick={handleGenerateSummary}
                                        disabled={loadingSummary}
                                        className="text-sm text-orange-500 hover:text-orange-700 font-medium underline decoration-orange-300 underline-offset-4 disabled:opacity-50 cursor-pointer"
                                    >
                                        {loadingSummary ? 'Running Analysis...' : 'Generate AI Summary'}
                                    </button>
                                )}
                            </div>

                            {/* Insight Block */}
                            <div className="bg-teal-50/50 rounded-2xl p-6 border border-teal-100/50 hover:border-teal-200 transition-colors">
                                <div className="flex items-center gap-2 mb-3 text-teal-600">
                                    <IconSparkles className="w-5 h-5" />
                                    <h3 className="font-bold uppercase tracking-widest text-xs">Deep Insight</h3>
                                </div>
                                {insight ? (
                                    <div className="text-sm text-stone-700 leading-relaxed animate-fade-in italic font-serif">
                                        <MDEditor.Markdown source={insight} style={{ background: 'transparent', color: 'inherit' }} />
                                    </div>
                                ) : (
                                    <button
                                        onClick={handleGenerateInsight}
                                        disabled={loadingInsight}
                                        className="text-sm text-teal-600 hover:text-teal-800 font-medium underline decoration-teal-300 underline-offset-4 disabled:opacity-50 cursor-pointer"
                                    >
                                        {loadingInsight ? 'Thinking...' : 'Reveal Deep Insight'}
                                    </button>
                                )}
                            </div>
                        </div>

                        {/* Article Content */}
                        <div
                            className="prose prose-lg max-w-none prose-stone prose-headings:font-bold prose-headings:text-stone-800 prose-p:text-stone-600 prose-p:leading-relaxed prose-a:text-orange-500 hover:prose-a:text-orange-600 prose-blockquote:border-l-4 prose-blockquote:border-orange-200 prose-blockquote:bg-orange-50/30 prose-blockquote:py-2 prose-blockquote:px-6 prose-img:rounded-2xl"
                            data-color-mode="light"
                        >
                            <MDEditor.Markdown
                                source={post.content}
                                style={{ backgroundColor: 'transparent', color: 'inherit', fontFamily: 'inherit' }}
                            />
                        </div>

                    </div>

                    {/* Comments Section (Rendered as children) */}
                    <div className="mt-12">
                        {children}
                    </div>
                </main>

                {/* Right Column: Sidebar */}
                <aside className="w-full lg:w-64 shrink-0 space-y-4">
                    {/* Cover Image / Monitor Illustration */}
                    <div className="bg-white/60 backdrop-blur-md rounded-[1.2rem] p-2.5 shadow-sm border border-white/50">
                        <div className="aspect-[4/3] rounded-xl overflow-hidden bg-stone-100 flex items-center justify-center">
                            {post.cover ? (
                                <img
                                    src={getAssetUrl(post.cover)}
                                    alt={post.title}
                                    className="w-full h-full object-cover hover:scale-105 transition-transform duration-500"
                                />
                            ) : (
                                <div className="text-stone-300">
                                    <svg className="w-12 h-12" fill="currentColor" viewBox="0 0 24 24"><path d="M21 19V5c0-1.1-.9-2-2-2H5c-1.1 0-2 .9-2 2v14c0 1.1.9 2 2 2h14c1.1 0 2-.9 2-2zM8.5 13.5l2.5 3.01L14.5 12l4.5 6H5l3.5-4.5z" /></svg>
                                </div>
                            )}
                        </div>
                    </div>

                    {/* Summary Sidebar Card */}
                    <div className="bg-white/60 backdrop-blur-md rounded-[1.2rem] p-5 shadow-sm border border-white/50">
                        <h4 className="text-[10px] font-bold text-stone-400 mb-3 tracking-wider uppercase">Summary</h4>
                        <div className="text-xs text-stone-600 leading-relaxed font-sans">
                            {post.excerpt || "No summary available for this article."}
                        </div>
                    </div>

                    {/* TOC Sidebar Card */}
                    {tocItems.length > 0 && (
                        <div className="bg-white/60 backdrop-blur-md rounded-[1.2rem] p-5 shadow-sm border border-white/50">
                            <h4 className="text-[10px] font-bold text-stone-400 mb-3 tracking-wider uppercase">Contents</h4>
                            <nav className="space-y-2">
                                {tocItems.map((item, index) => (
                                    <a
                                        key={index}
                                        href={`#${item.id}`}
                                        className={`block text-xs transition-colors ${item.level === 1 ? 'text-stone-800 font-medium hover:text-orange-500' :
                                            'text-stone-500 hover:text-stone-800 pl-3 border-l-[1.5px] border-transparent hover:border-orange-300'
                                            }`}
                                    >
                                        {item.text}
                                    </a>
                                ))}
                            </nav>
                        </div>
                    )}

                    {/* Stats & Like Button (Floating Badge style) */}
                    <div className="relative h-20">
                        <div className="absolute left-0 top-0 scale-90 origin-top-left">
                            <LikeButton
                                initialCount={articleLikes}
                                onLike={async () => { await postService.like(`post-${post.slug}`); }}
                            />
                        </div>
                    </div>
                </aside>
            </div>
        </div>
    );
};

export default PostDetail;
