import React from 'react';
import type { BlogPost } from '../types';
import { getAssetUrl } from '../utils/urlUtils';
import { IconClock, IconEye } from './Icons';

interface PostCardProps {
  post: BlogPost;
  onClick: (id: string) => void;
}

const PostCard: React.FC<PostCardProps> = ({ post, onClick }) => {
  // Simple time ago calculation
  const getTimeAgo = (dateString: string) => {
    const date = new Date(dateString);
    const now = new Date();
    const diffInSeconds = Math.floor((now.getTime() - date.getTime()) / 1000);

    if (diffInSeconds < 60) return 'Just now';
    if (diffInSeconds < 3600) return `${Math.floor(diffInSeconds / 60)}m ago`;
    if (diffInSeconds < 86400) return `${Math.floor(diffInSeconds / 3600)}h ago`;
    if (diffInSeconds < 604800) return `${Math.floor(diffInSeconds / 86400)}d ago`;
    return date.toLocaleDateString();
  };



  return (
    <article
      onClick={() => onClick(post.id)}
      className="group cursor-pointer bg-white rounded-xl shadow-sm hover:shadow-2xl transition-all duration-500 border border-stone-100 overflow-hidden flex flex-col md:flex-row h-auto md:h-[22rem]"
    >
      {/* Image Section - Takes up top half on mobile, left 50% on desktop */}
      <div className="w-full md:w-[55%] h-64 md:h-full relative overflow-hidden bg-stone-100">
        {/* Placeholder - Always rendered, sits behind image */}
        <div className="absolute inset-0 bg-stone-200 animate-pulse" />

        <img
          src={getAssetUrl(post.cover)}
          alt={post.title}
          loading="lazy"
          className="w-full h-full object-cover transform group-hover:scale-105 transition-transform duration-700 relative z-10"
        />

        {/* Overlay - Sits on top */}
        <div className="absolute inset-0 bg-black/5 group-hover:bg-transparent transition-colors pointer-events-none z-20" />
      </div>

      {/* Content Section */}
      <div className="flex-1 p-8 md:p-10 flex flex-col justify-between relative bg-white">
        <div>
          {/* Top Metadata */}
          <div className="flex items-center gap-4 mb-4">
            <span className="px-3 py-1 text-[10px] uppercase tracking-[0.2em] bg-stone-100 text-stone-600 font-bold rounded-sm">
              {post.category}
            </span>
            <div className="flex items-center gap-2 text-stone-400 text-xs tracking-wider">
              <IconClock className="w-3.5 h-3.5" />
              <span>{getTimeAgo(post.publishAt)}</span>
            </div>
          </div>

          {/* Title */}
          <h2 className="text-2xl md:text-3xl font-serif font-bold text-ink mb-4 leading-tight group-hover:text-gold-600 transition-colors">
            {post.title}
          </h2>

          {/* Excerpt */}
          <p className="text-stone-500 text-sm md:text-base leading-relaxed line-clamp-3 md:line-clamp-4 font-serif italic">
            {post.excerpt}
          </p>
        </div>

        {/* Bottom Metadata */}
        <div className="flex flex-wrap items-center justify-between gap-4 pt-6 border-t border-stone-100 mt-6">
          <div className="flex gap-2">
            {post.tags.slice(0, 3).map(tag => (
              <span key={tag} className="text-[10px] text-stone-400 uppercase tracking-wider hover:text-gold-600 transition-colors">
                #{tag}
              </span>
            ))}
          </div>

          <div className="flex items-center gap-6 text-xs font-medium text-stone-400">
            <div className="flex items-center gap-2" title="Views">
              <IconEye className="w-4 h-4 text-stone-300" />
              <span>{post.views.toLocaleString()}</span>
            </div>
            <span>{post.readTime}</span>
          </div>
        </div>
      </div>
    </article>
  );
};

export default PostCard;