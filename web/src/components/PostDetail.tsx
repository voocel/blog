
import React, { useState, useEffect } from 'react';
import MDEditor from '@uiw/react-md-editor';
import type { BlogPost } from '../types';
import { getAssetUrl } from '../utils/urlUtils';
import { IconArrowLeft, IconSparkles, IconBrain } from './Icons';
import { useBlog } from '../context/BlogContext';
import { generateSummary, generateInsight } from '../services/geminiService';
import SEO from './SEO';

interface PostDetailProps {
    post: BlogPost;
    onBack: () => void;
}

const PostDetail: React.FC<PostDetailProps> = ({ post, onBack }) => {
    const { logVisit } = useBlog(); // Added useBlog hook call
    const [summary, setSummary] = useState<string | null>(null);
    const [insight, setInsight] = useState<string | null>(null);
    const [loadingSummary, setLoadingSummary] = useState(false);
    const [loadingInsight, setLoadingInsight] = useState(false);

    useEffect(() => {
        logVisit(`/ post / ${post.id} `, post.id, post.title); // Added new useEffect
    }, [post.id]);

    useEffect(() => {
        window.scrollTo(0, 0);
    }, []);

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

    return (
        <div className="animate-fade-in min-h-screen pl-0 bg-transparent text-ink">
            <SEO
                title={post.title}
                description={post.excerpt}
                image={getAssetUrl(post.cover)}
                type="article"
            />
            {/* Hero Image with Title Overlay */}
            <div className="relative h-[50vh] md:h-[60vh] w-full">
                <div className="absolute inset-0">
                    <img src={getAssetUrl(post.cover)} alt={post.title} className="w-full h-full object-cover" />
                </div>
                {/* Light overlay */}
                <div className="absolute inset-0 bg-gradient-to-b from-stone-900/10 via-stone-900/20 to-[#FDFBF7]" />

                <div className="absolute inset-0 flex flex-col justify-end px-6 md:px-12 pb-12 max-w-5xl mx-auto">
                    <button
                        onClick={onBack}
                        className="fixed top-24 left-6 md:left-12 flex items-center gap-2 bg-white/80 backdrop-blur px-4 py-2 rounded-full shadow-sm text-stone-600 hover:text-gold-600 transition-colors z-20 cursor-pointer"
                    >
                        <IconArrowLeft className="w-4 h-4" />
                        <span className="text-xs tracking-widest uppercase">Return</span>
                    </button>

                    <div className="flex items-center gap-4 mb-4">
                        <span className="text-gold-600 uppercase tracking-widest text-sm font-bold bg-white/80 px-2 py-1">{post.category}</span>
                        <div className="w-px h-4 bg-ink/30" />
                        <span className="text-ink font-medium text-sm">{new Date(post.publishAt).toLocaleString()}</span>
                    </div>

                    <h1 className="text-4xl md:text-6xl lg:text-7xl font-serif font-bold text-ink leading-tight mb-6 drop-shadow-sm">
                        {post.title}
                    </h1>

                    <div className="flex items-center gap-3">
                        <div className="w-10 h-10 rounded-full bg-gold-600 flex items-center justify-center text-white font-serif font-bold text-lg">
                            {post.author.charAt(0)}
                        </div>
                        <div>
                            <p className="text-ink font-medium">{post.author}</p>
                            <p className="text-xs text-stone-500">{post.readTime}</p>
                        </div>
                    </div>
                </div>
            </div>

            <div className="max-w-4xl mx-auto px-6 md:px-12 py-12">

                {/* AI Intelligence Grid */}
                <div className="grid grid-cols-1 md:grid-cols-2 gap-6 mb-16">
                    {/* Summary Block */}
                    <div className="glass-panel p-6 rounded-lg border-l-2 border-gold-500">
                        <div className="flex items-center gap-2 mb-4 text-gold-600">
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
                                className="w-full py-3 border border-stone-300 rounded text-sm text-stone-500 hover:border-gold-500 hover:text-gold-600 transition-all disabled:opacity-50"
                            >
                                {loadingSummary ? 'Analyzing...' : 'Generate Summary'}
                            </button>
                        )}
                    </div>

                    {/* Insight Block */}
                    <div className="glass-panel p-6 rounded-lg border-l-2 border-teal-500">
                        <div className="flex items-center gap-2 mb-4 text-teal-600">
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
                                className="w-full py-3 border border-stone-300 rounded text-sm text-stone-500 hover:border-teal-500 hover:text-teal-600 transition-all disabled:opacity-50"
                            >
                                {loadingInsight ? 'Thinking...' : 'Reveal Hidden Meaning'}
                            </button>
                        )}
                    </div>
                </div>

                {/* Main Content */}
                <div
                    className="prose prose-lg max-w-none prose-headings:font-serif prose-headings:text-ink prose-p:text-stone-700 prose-a:text-gold-600 prose-blockquote:border-gold-500 prose-blockquote:bg-stone-50 prose-blockquote:py-2 prose-blockquote:px-6 prose-strong:text-ink"
                    data-color-mode="light"
                >
                    <MDEditor.Markdown
                        source={post.content}
                        style={{ backgroundColor: 'transparent', color: 'inherit', fontFamily: 'inherit' }}
                    />
                </div>

                {/* Tags */}
                <div className="mt-16 pt-8 border-t border-stone-200 flex flex-wrap gap-3">
                    {post.tags.map(tag => (
                        <span key={tag} className="px-3 py-1 text-xs uppercase tracking-wider border border-stone-300 text-stone-500 rounded-full hover:border-gold-500 hover:text-gold-600 transition-colors cursor-default">
                            #{tag}
                        </span>
                    ))}
                </div>
            </div>
        </div>
    );
};

export default PostDetail;
