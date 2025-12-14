
import React, { useState, useRef, useEffect, type ReactNode } from 'react';
import { useNavigate } from 'react-router-dom';
import { useBlog } from '../context/BlogContext';
import { HERO_CONTENT } from '../constants';
import PostCard from '../components/PostCard';
import { IconArrowDown, IconArrowLeft } from '../components/Icons';
import SEO from '../components/SEO';

// --- Reusable Scroll Reveal Component ---
interface RevealProps {
  children: ReactNode;
  className?: string;
  delay?: number;
  threshold?: number;
}

const Reveal: React.FC<RevealProps> = ({ children, className = "", delay = 0, threshold = 0.1 }) => {
  const ref = useRef<HTMLDivElement>(null);
  const [isVisible, setIsVisible] = useState(false);

  useEffect(() => {
    const observer = new IntersectionObserver(
      ([entry]) => {
        if (entry.isIntersecting) {
          setIsVisible(true);
          observer.disconnect(); // Only trigger once
        }
      },
      { threshold }
    );

    if (ref.current) {
      observer.observe(ref.current);
    }

    return () => observer.disconnect();
  }, [threshold]);

  return (
    <div
      ref={ref}
      className={`${isVisible ? 'animate-blur-in' : 'opacity-0'} ${className}`}
      style={{ animationDelay: `${delay}ms`, animationFillMode: 'both' }}
    >
      {children}
    </div>
  );
};

import { postService } from '../services/postService';
import type { BlogPost } from '../types';

// ... (Reveal Component stays same)

const HomePage: React.FC = () => {
  const { categories, logVisit } = useBlog(); // Removed posts from here
  const [posts, setPosts] = useState<BlogPost[]>([]);
  const [isLoadingPosts, setIsLoadingPosts] = useState(true);
  const [selectedCategory, setSelectedCategory] = useState<string>('All');
  const [pagination, setPagination] = useState({ page: 1, totalPages: 1, total: 0 });
  const navigate = useNavigate();

  const itemsPerPage = 5;

  // Track visit
  useEffect(() => {
    logVisit('/');
  }, []);

  // Fetch Posts Server-Side
  useEffect(() => {
    const fetchPosts = async () => {
      setIsLoadingPosts(true);
      try {
        const params: any = {
          page: pagination.page,
          limit: itemsPerPage
        };

        if (selectedCategory !== 'All') {
          // Find category ID from name
          const cat = categories.find(c => c.name === selectedCategory);
          if (cat) params.category = cat.id;
        }

        const { data, pagination: meta } = await postService.getPosts(params);
        setPosts(data);
        if (meta) {
          setPagination(prev => ({ ...prev, totalPages: meta.totalPages, total: meta.total }));
        }
      } catch (error) {
        console.error("Failed to fetch posts", error);
      } finally {
        setIsLoadingPosts(false);
      }
    };

    fetchPosts();
  }, [selectedCategory, pagination.page, categories]); // Depends on categories to resolve ID

  // Extract unique categories for filter bar (Using Context Categories is better/source of truth, but user might want only categories with posts? API returns all categories. Let's use Context Categories + All)
  const categoryList = ['All', ...categories.map(c => c.name)];

  const scrollToContent = () => {
    const element = document.getElementById('journal-feed');
    if (element) {
      const y = element.getBoundingClientRect().top + window.scrollY - 80;
      window.scrollTo({ top: y, behavior: 'smooth' });
    }
  };

  const toggleCategory = (cat: string) => {
    if (selectedCategory === cat) return; // No op if same
    setSelectedCategory(cat);
    setPagination(prev => ({ ...prev, page: 1 })); // Reset to page 1
  };

  const handlePageChange = (newPage: number) => {
    if (newPage >= 1 && newPage <= pagination.totalPages) {
      setPagination(prev => ({ ...prev, page: newPage }));
      scrollToContent();
    }
  };

  // Skip global loading check, handle local loading
  // if (isLoading) ... (removed)

  // Hero Post (Static Content)
  const heroPost = HERO_CONTENT;

  return (
    <div className="min-h-screen pb-20 bg-transparent">
      <SEO title="Voocel Journal" />

      {/* Magazine Cover Hero */}
      <section className="relative h-screen w-full flex items-center justify-center overflow-hidden">
        <div className="absolute inset-0">
          <img
            src={heroPost.cover}
            alt="Hero"
            className="w-full h-full object-cover filter saturate-[0.8] sepia-[0.15] opacity-90 animate-slow-pan"
          />
        </div>
        {/* Light gradient fade to match global bg */}
        <div className="absolute inset-0 bg-gradient-to-t from-[#FDFBF7] via-[#FDFBF7]/40 to-transparent"></div>

        <div className="relative z-10 text-center max-w-4xl px-6 flex flex-col items-center">

          <Reveal delay={0}>
            <p className="text-ink text-xs uppercase tracking-[0.4em] mb-6 px-4 py-2 rounded-full border border-ink/10 bg-white/50 backdrop-blur-md inline-block">
              Featured Story
            </p>
          </Reveal>

          <Reveal delay={200}>
            <h1
              className="text-5xl md:text-8xl font-serif font-bold text-ink mb-8 leading-[1.1] cursor-default hover:text-gold-600 transition-colors drop-shadow-sm"
            >
              {heroPost.title}
            </h1>
          </Reveal>

          <Reveal delay={400}>
            <p className="text-xl text-stone-600 font-light font-serif italic mb-10 max-w-2xl mx-auto">
              {heroPost.excerpt}
            </p>
          </Reveal>

          {/* Artistic Scroll Down Button with Border Beam */}
          <Reveal delay={600}>
            <button
              onClick={scrollToContent}
              className="group cursor-pointer flex flex-col items-center gap-4 text-stone-500 hover:text-ink transition-all duration-500 mt-8"
            >
              <span className="text-[10px] uppercase tracking-[0.4em] font-light group-hover:tracking-[0.5em] transition-all">Begin Journey</span>

              {/* Button Container with Hover Effect */}
              <div className="relative rounded-full overflow-hidden p-[1px] transform transition-transform duration-500 group-hover:scale-105">
                {/* Border Beam - Conic Gradient Animation */}
                <span className="absolute inset-[-100%] bg-[conic-gradient(from_90deg_at_50%_50%,#FDFBF7_0%,#FDFBF7_50%,#CA8A04_100%)] opacity-0 group-hover:opacity-100 animate-spin-slow" />

                {/* Inner Button */}
                <div className="relative p-3 bg-white/80 rounded-full border border-stone-300 group-hover:border-transparent transition-colors bg-white/50 backdrop-blur-sm z-10">
                  <IconArrowDown className="w-5 h-5 group-hover:animate-bounce group-hover:text-gold-600" />
                </div>
              </div>
            </button>
          </Reveal>
        </div>
      </section>

      {/* Filter Bar (Curated Index) */}
      <div id="journal-feed" className="sticky top-[55px] z-30 bg-[#FDFBF7]/90 backdrop-blur-md border-b border-stone-200 py-4 mb-12 shadow-sm transition-all">
        <div className="max-w-6xl mx-auto px-6 overflow-x-auto">
          <div className="flex gap-8 md:gap-12 min-w-max justify-center items-center">
            <span className="text-[10px] uppercase text-stone-400 tracking-widest mr-4 border-r border-stone-200 pr-6 hidden md:block">Filter By</span>
            <div className="flex justify-center gap-8 animate-fade-in-up delay-200">
              {categoryList.map((cat) => (
                <button
                  key={cat}
                  onClick={() => toggleCategory(cat)}
                  className={`text-xs uppercase tracking-[0.2em] transition-all duration-300 relative py-2 cursor-pointer ${selectedCategory === cat
                    ? 'text-ink font-bold after:content-[""] after:absolute after:bottom-0 after:left-0 after:w-full after:h-px after:bg-gold-500'
                    : 'text-stone-400 hover:text-stone-600'
                    }`}
                >
                  {cat}
                  {selectedCategory === cat && (
                    <div className="absolute top-0 left-1/2 -translate-x-1/2 w-0 h-0 border-l-[4px] border-l-transparent border-r-[4px] border-r-transparent border-t-[5px] border-t-gold-500"></div>
                  )}
                </button>
              ))}
            </div>
          </div>
        </div>
      </div>

      {/* Content Grid (Wider Layout) */}
      <main className="max-w-6xl mx-auto px-6 min-h-[500px]">
        {isLoadingPosts ? (
          <div className="flex flex-col items-center justify-center py-32 opacity-50">
            <div className="w-8 h-8 border-2 border-stone-200 border-t-gold-600 rounded-full animate-spin mb-4"></div>
            <span className="text-stone-400 font-serif text-sm tracking-widest italic">Fetching entries...</span>
          </div>
        ) : (
          <div className="flex flex-col gap-16 md:gap-20">
            {posts.map((post, index) => (
              <Reveal key={post.id} delay={index * 150} threshold={0.05}>
                <PostCard post={post} onClick={(id) => navigate(`/post/${id}`)} />
              </Reveal>
            ))}
          </div>
        )}

        {!isLoadingPosts && posts.length === 0 && (
          <div className="py-20 text-center text-stone-400 italic font-serif text-lg">
            No entries found in this collection.
          </div>
        )}

        {/* Pagination Controls */}
        {pagination.totalPages > 1 && (
          <div className="mt-24 flex items-center justify-center gap-8 text-sm font-serif border-t border-stone-200 pt-12 max-w-lg mx-auto">
            <button
              onClick={() => handlePageChange(pagination.page - 1)}
              disabled={pagination.page === 1}
              className="flex items-center gap-2 text-stone-400 hover:text-ink disabled:opacity-30 disabled:cursor-not-allowed cursor-pointer transition-colors uppercase tracking-widest text-xs"
            >
              <IconArrowLeft className="w-4 h-4" />
              Previous
            </button>

            <div className="flex items-center gap-2">
              {Array.from({ length: pagination.totalPages }, (_, i) => i + 1).map(page => (
                <button
                  key={page}
                  onClick={() => handlePageChange(page)}
                  className={`w-8 h-8 flex items-center justify-center rounded-full transition-all leading-none cursor-pointer ${pagination.page === page
                    ? 'bg-gold-600 text-white shadow-sm'
                    : 'text-stone-500 hover:bg-stone-100'
                    }`}
                >
                  {page}
                </button>
              ))}
            </div>

            <button
              onClick={() => handlePageChange(pagination.page + 1)}
              disabled={pagination.page === pagination.totalPages}
              className="flex items-center gap-2 text-stone-400 hover:text-ink disabled:opacity-30 disabled:cursor-not-allowed cursor-pointer transition-colors uppercase tracking-widest text-xs"
            >
              Next
              <IconArrowDown className="w-4 h-4 -rotate-90" />
            </button>
          </div>
        )}
      </main>

      {/* Footer */}
      <footer className="mt-32 border-t border-stone-200 py-20 text-center bg-transparent">
        <h2 className="text-2xl font-serif text-ink mb-6">Voocel.</h2>
        <p className="text-stone-500 text-sm mb-8 max-w-md mx-auto">
          A digital sanctuary for thoughts, aesthetics, and the silent rhythm of algorithms.
        </p>
        <div className="text-stone-400 text-xs uppercase tracking-widest">
          © 2024 Voocel Journal. All Rights Reserved.
        </div>
      </footer>
    </div >
  );
};

export default HomePage;

