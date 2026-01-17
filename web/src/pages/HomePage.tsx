
import React, { useEffect, useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { motion } from 'framer-motion';
import BentoItem from '../components/bento/BentoItem';
import QuoteWidget from '../components/bento/widgets/QuoteWidget';
import AnimatedNavWidget from '../components/AnimatedNavWidget';
import ClockWidget from '../components/bento/widgets/ClockWidget';
import CalendarWidget from '../components/bento/widgets/CalendarWidget';
import MediaWidget from '../components/bento/widgets/MediaWidget';
import SocialWidget from '../components/bento/widgets/SocialWidget';
import LikeButton from '../components/LikeButton';
import SettingsModal from '../components/SettingsModal';
import SEO from '../components/SEO';
import { postService } from '../services/postService';
import type { BlogPost } from '../types';
import { useAuth } from '../context/AuthContext';

// Staggered animation variants
const containerVariants = {
  hidden: { opacity: 0 },
  visible: {
    opacity: 1,
    transition: {
      staggerChildren: 0.1,
      delayChildren: 0.2,
    },
  },
};

const itemVariants = {
  hidden: { opacity: 0, y: 30, scale: 0.95 },
  visible: {
    opacity: 1,
    y: 0,
    scale: 1,
    transition: {
      type: 'spring' as const,
      stiffness: 100,
      damping: 15,
    },
  },
};

const catImage = "/images/cute_cat.png";

const HomePage: React.FC = () => {
  const navigate = useNavigate();
  const { user, setAuthModalOpen } = useAuth();
  const [latestPost, setLatestPost] = useState<BlogPost | null>(null);
  const [randomPost, setRandomPost] = useState<BlogPost | null>(null);
  const [isSettingsOpen, setIsSettingsOpen] = useState(false);
  const [homepageLikes, setHomepageLikes] = useState(0);

  // Fetch latest post, random post and homepage likes on mount
  useEffect(() => {
    postService.getPosts({ page: 1, limit: 1 })
      .then(res => {
        if (res.data.length > 0) {
          setLatestPost(res.data[0]);
        }
      })
      .catch(err => {
        console.error('Failed to fetch latest post:', err);
      });

    // Fetch posts for random selection (get more posts to pick from)
    postService.getPosts({ page: 1, limit: 20 })
      .then(res => {
        if (res.data.length > 0) {
          // Pick a random post from the list
          const randomIndex = Math.floor(Math.random() * res.data.length);
          setRandomPost(res.data[randomIndex]);
        }
      })
      .catch(err => {
        console.error('Failed to fetch random post:', err);
      });

    postService.getLikes('home')
      .then(count => setHomepageLikes(count))
      .catch(err => console.error('Failed to fetch homepage likes:', err));
  }, []);

  // Open auth modal for login
  const handleLoginClick = () => {
    setAuthModalOpen(true);
  };

  // Navigate to admin dashboard
  const handleDashboardClick = () => {
    navigate('/admin');
  };

  return (
    <div className="min-h-screen w-full bg-[#fdfaf6] bg-[radial-gradient(ellipse_at_top_left,_var(--tw-gradient-stops))] from-orange-100/40 via-rose-100/20 to-transparent overflow-hidden relative">
      <SEO title="Home - Voocel Dashboard" />
      <SettingsModal isOpen={isSettingsOpen} onClose={() => setIsSettingsOpen(false)} />

      {/* Background Decorations - Animated blobs */}
      <div className="fixed top-20 left-10 w-64 h-64 bg-purple-200 rounded-full mix-blend-multiply filter blur-3xl opacity-20 animate-blob pointer-events-none" aria-hidden="true"></div>
      <div className="fixed top-20 right-10 w-64 h-64 bg-orange-200 rounded-full mix-blend-multiply filter blur-3xl opacity-20 animate-blob animation-delay-2000 pointer-events-none" aria-hidden="true"></div>
      <div className="fixed -bottom-8 left-20 w-64 h-64 bg-pink-200 rounded-full mix-blend-multiply filter blur-3xl opacity-20 animate-blob animation-delay-4000 pointer-events-none" aria-hidden="true"></div>


      {/* Mobile Navigation - Fixed at top, outside animation container */}
      <div className="md:hidden">
        <AnimatedNavWidget isCompact={true} />
      </div>

      {/* ==================== MOBILE LAYOUT (< md) ==================== */}
      <motion.div
        className="md:hidden flex flex-col gap-4 p-4 pt-24"
        variants={containerVariants}
        initial="hidden"
        animate="visible"
      >

        {/* Cat Image */}
        <motion.div variants={itemVariants}>
          <BentoItem className="h-[180px] !p-0 overflow-hidden group relative shadow-lg !rounded-[2rem]">
            <div className="absolute inset-0">
              <img src={catImage} alt="Decorative cat" className="w-full h-full object-cover object-[center_35%]" />
            </div>
            <div className="absolute top-3 right-3 bg-white/80 backdrop-blur-md px-2.5 py-1 rounded-full text-[9px] font-bold text-stone-500 shadow-sm">
              Do not disturb
            </div>
          </BentoItem>
        </motion.div>

        {/* Profile */}
        <motion.div variants={itemVariants}>
          <BentoItem className="py-8 flex items-center justify-center bg-gradient-to-b from-white/80 to-orange-50/50 shadow-xl border-white/60">
            <QuoteWidget />
          </BentoItem>
        </motion.div>

        {/* Social Links */}
        <motion.div variants={itemVariants}>
          <SocialWidget />
        </motion.div>

        {/* Latest Post - Mobile Redesigned */}
        <motion.div variants={itemVariants}>
          <BentoItem className="h-auto !bg-[#FFFBF0] hover:!bg-[#FFFBF0] cursor-pointer relative group !p-4 overflow-hidden shadow-sm border border-orange-100/30">
            {latestPost ? (
              <div
                onClick={() => navigate(`/post/${latestPost.id}`)}
                className="flex flex-col gap-3"
                role="button"
                tabIndex={0}
                onKeyDown={(e) => e.key === 'Enter' && navigate(`/post/${latestPost.id}`)}
              >
                {/* Header */}
                <div className="text-sm font-bold text-stone-500 tracking-wide">最新文章</div>

                <div className="flex gap-4 items-start">
                  {/* Small Image */}
                  <div className="w-16 h-16 rounded-xl overflow-hidden shrink-0 shadow-sm border border-white/50">
                    <img src={latestPost.cover} className="w-full h-full object-cover" alt={latestPost.title} />
                  </div>

                  {/* Content */}
                  <div className="flex-1 min-w-0 flex flex-col gap-1">
                    <h4 className="font-bold text-stone-700 text-sm leading-snug line-clamp-2">{latestPost.title}</h4>
                    <p className="text-xs text-stone-500 line-clamp-1">{latestPost.excerpt || "点击阅读更多内容..."}</p>
                    <div className="text-[10px] text-stone-400 mt-1">{new Date(latestPost.publishAt).toLocaleDateString()}</div>
                  </div>
                </div>
              </div>
            ) : (
              <div className="h-24 flex items-center justify-center text-stone-300 text-xs">It's quiet here...</div>
            )}
          </BentoItem>
        </motion.div>

        {/* Music + Like */}
        <motion.div className="flex items-center gap-8" variants={itemVariants}>
          <BentoItem className="h-[60px] flex-1 !bg-orange-50/60 !p-1.5 !border-none shadow-sm flex items-center">
            <MediaWidget />
          </BentoItem>
          <LikeButton
            initialCount={homepageLikes}
            onLike={async () => { await postService.like('home'); }}
          />
        </motion.div>

        {/* Login/Settings on mobile */}
        <motion.div className="flex justify-center gap-2 pt-4" variants={itemVariants}>
          {user && (
            <button
              onClick={handleDashboardClick}
              className="bg-red-400 text-white px-4 py-2 rounded-xl shadow-red-200 shadow-md font-bold text-sm cursor-pointer"
            >
              Dashboard
            </button>
          )}
          {!user && (
            <button
              onClick={handleLoginClick}
              className="bg-white/70 text-stone-500 px-4 py-2 rounded-xl shadow-sm font-bold text-sm cursor-pointer"
            >
              Sign In
            </button>
          )}
          <button
            onClick={() => setIsSettingsOpen(true)}
            className="w-10 h-10 rounded-xl bg-white/50 flex items-center justify-center text-stone-400"
          >
            ⚙️
          </button>
        </motion.div>
      </motion.div>

      {/* ==================== DESKTOP LAYOUT (>= md) ==================== */}
      <motion.div
        className="hidden md:block relative w-full max-w-[1100px] mx-auto h-screen p-8"
        variants={containerVariants}
        initial="hidden"
        animate="visible"
      >

        {/* ==================== LEFT ZONE ==================== */}

        {/* Navigation Widget */}
        <motion.div className="absolute left-[10%] top-[8%] w-[280px]" variants={itemVariants}>
          <AnimatedNavWidget isCompact={false} />
        </motion.div>

        {/* Recent Post Widget - Redesigned */}
        <motion.div className="absolute left-[12%] top-[62%] w-[220px]" variants={itemVariants}>
          <BentoItem className="h-auto !bg-[#FFFBF0] hover:!bg-[#FFFBF0] cursor-pointer relative group !p-5 overflow-hidden shadow-sm border border-orange-100/30 !rounded-[1.5rem]">
            {latestPost ? (
              <div
                onClick={() => navigate(`/post/${latestPost.id}`)}
                className="flex flex-col gap-4"
                role="button"
                tabIndex={0}
                onKeyDown={(e) => e.key === 'Enter' && navigate(`/post/${latestPost.id}`)}
                aria-label={`Read article: ${latestPost.title}`}
              >
                {/* Header */}
                <div className="text-sm font-bold text-stone-500 tracking-wide">最新文章</div>

                <div className="flex gap-4">
                  {/* Image */}
                  <div className="w-14 h-14 rounded-[0.8rem] overflow-hidden shrink-0 shadow-sm border border-white/50 bg-stone-100">
                    <img src={latestPost.cover} className="w-full h-full object-cover group-hover:scale-105 transition-transform duration-500" alt={latestPost.title} />
                  </div>

                  {/* Content */}
                  <div className="flex-1 min-w-0 flex flex-col justify-center gap-1">
                    <h4 className="font-bold text-stone-700 text-[13px] leading-snug line-clamp-2 group-hover:text-orange-600 transition-colors">{latestPost.title}</h4>
                    <p className="text-[10px] text-stone-500 line-clamp-1 opacity-80">{latestPost.excerpt || "点击阅读..."}</p>
                    <div className="text-[10px] text-stone-400 font-mono">{new Date(latestPost.publishAt).toLocaleDateString()}</div>
                  </div>
                </div>
              </div>
            ) : (
              <div className="h-24 flex items-center justify-center text-stone-300 text-xs">It's quiet here...</div>
            )}
          </BentoItem>
        </motion.div>

        {/* ==================== CENTER ZONE ==================== */}

        {/* Cat Image Widget */}
        <motion.div className="absolute left-[50%] -translate-x-1/2 top-[4%] w-[250px]" variants={itemVariants}>
          <BentoItem className="h-[155px] !p-0 overflow-hidden group relative shadow-lg !rounded-[2rem]">
            <div className="absolute inset-0">
              <img src={catImage} alt="Decorative cat" className="w-full h-full object-cover object-[center_35%] transition-transform duration-700 group-hover:scale-105" />
            </div>
            <div className="absolute top-3 right-3 bg-white/80 backdrop-blur-md px-2.5 py-1 rounded-full text-[9px] font-bold text-stone-500 shadow-sm">
              Do not disturb
            </div>
          </BentoItem>
        </motion.div>

        {/* Profile Widget */}
        <motion.div className="absolute left-[50%] -translate-x-1/2 top-[26%] w-[280px]" variants={itemVariants}>
          <BentoItem className="h-[220px] flex items-center justify-center bg-gradient-to-b from-white/80 to-orange-50/50 shadow-xl border-white/60">
            <QuoteWidget />
          </BentoItem>
        </motion.div>

        {/* Social Widget */}
        <motion.div className="absolute left-[50%] -translate-x-1/2 top-[58%] w-[250px]" variants={itemVariants}>
          <div className="h-[50px]">
            <SocialWidget />
          </div>
        </motion.div>

        {/* Random Pick Widget - Same design as Latest Post */}
        <motion.div className="absolute left-[34%] top-[70%] w-[220px]" variants={itemVariants}>
          <BentoItem className="h-auto !bg-[#FFF5F0] hover:!bg-[#FFF5F0] cursor-pointer relative group !p-5 overflow-hidden shadow-sm border border-rose-100/30 !rounded-[1.5rem]">
            {randomPost ? (
              <div
                onClick={() => navigate(`/post/${randomPost.id}`)}
                className="flex flex-col gap-4"
                role="button"
                tabIndex={0}
                onKeyDown={(e) => e.key === 'Enter' && navigate(`/post/${randomPost.id}`)}
                aria-label={`Read article: ${randomPost.title}`}
              >
                {/* Header */}
                <div className="text-sm font-bold text-rose-400 tracking-wide">随机推荐</div>

                <div className="flex gap-4">
                  {/* Image */}
                  <div className="w-14 h-14 rounded-[0.8rem] overflow-hidden shrink-0 shadow-sm border border-white/50 bg-stone-100">
                    <img src={randomPost.cover} className="w-full h-full object-cover group-hover:scale-105 transition-transform duration-500" alt={randomPost.title} />
                  </div>

                  {/* Content */}
                  <div className="flex-1 min-w-0 flex flex-col justify-center gap-1">
                    <h4 className="font-bold text-stone-700 text-[13px] leading-snug line-clamp-2 group-hover:text-rose-500 transition-colors">{randomPost.title}</h4>
                    <p className="text-[10px] text-stone-500 line-clamp-1 opacity-80">{randomPost.excerpt || "点击阅读..."}</p>
                    <div className="text-[10px] text-stone-400 font-mono">{new Date(randomPost.publishAt).toLocaleDateString()}</div>
                  </div>
                </div>
              </div>
            ) : (
              <div className="h-24 flex items-center justify-center text-stone-300 text-xs">暂无推荐...</div>
            )}
          </BentoItem>
        </motion.div>

        {/* Music Widget */}
        <motion.div className="absolute left-[60%] top-[73%]" variants={itemVariants}>
          <BentoItem className="h-[60px] w-[220px] !bg-orange-50/60 !p-1.5 !border-none shadow-sm flex items-center">
            <MediaWidget />
          </BentoItem>
        </motion.div>

        {/* Like Button */}
        <motion.div className="absolute left-[82%] top-[70%]" variants={itemVariants}>
          <LikeButton
            initialCount={homepageLikes}
            onLike={async () => { await postService.like('home'); }}
          />
        </motion.div>

        {/* ==================== RIGHT ZONE ==================== */}

        {/* Auth Buttons */}
        <motion.div className="absolute right-[14%] top-[6%] flex gap-2" variants={itemVariants}>
          {user && (
            <button
              onClick={handleDashboardClick}
              className="bg-red-400 text-white px-4 py-2 rounded-xl shadow-red-200 shadow-md font-bold text-sm hover:bg-red-500 transition-colors cursor-pointer"
              aria-label="Go to dashboard"
            >
              Dashboard
            </button>
          )}

          <div className="flex items-center gap-1.5 bg-white/50 backdrop-blur-md rounded-xl px-2 shadow-sm">
            {!user && (
              <button
                onClick={handleLoginClick}
                className="text-[10px] text-stone-400 font-bold px-1.5 hover:text-stone-600 transition-colors cursor-pointer"
                aria-label="Sign in to your account"
              >
                SIGN IN
              </button>
            )}
            <button
              onClick={() => setIsSettingsOpen(true)}
              className="w-7 h-7 rounded-lg hover:bg-stone-200/50 flex items-center justify-center text-stone-400 transition-colors cursor-pointer"
              aria-label="Open settings menu"
            >
              <svg width="16" height="16" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg" aria-hidden="true">
                <circle cx="5" cy="5" r="2" fill="currentColor" />
                <circle cx="12" cy="5" r="2" fill="currentColor" />
                <circle cx="19" cy="5" r="2" fill="currentColor" />
                <circle cx="5" cy="12" r="2" fill="currentColor" />
                <circle cx="12" cy="12" r="2" fill="currentColor" />
                <circle cx="19" cy="12" r="2" fill="currentColor" />
                <circle cx="5" cy="19" r="2" fill="currentColor" />
                <circle cx="12" cy="19" r="2" fill="currentColor" />
                <circle cx="19" cy="19" r="2" fill="currentColor" />
              </svg>
            </button>
          </div>
        </motion.div>

        {/* Clock Widget */}
        <motion.div className="absolute right-[10%] top-[14%] w-[170px]" variants={itemVariants}>
          <BentoItem className="h-[100px] !p-0 !bg-stone-100/70 !border-white/30 shadow-inner">
            <ClockWidget />
          </BentoItem>
        </motion.div>

        {/* Calendar Widget */}
        <motion.div className="absolute right-[8%] top-[32%] w-[240px]" variants={itemVariants}>
          <BentoItem className="h-[260px] shadow-sm">
            <CalendarWidget />
          </BentoItem>
        </motion.div>

        {/* Decorative Elements */}
        <motion.div
          className="absolute right-[20%] top-[78%] text-xl opacity-30"
          aria-hidden="true"
          initial={{ opacity: 0, scale: 0 }}
          animate={{ opacity: 0.3, scale: 1 }}
          transition={{ delay: 1.2, type: 'spring' }}
        >
          💕
        </motion.div>
        <motion.div
          className="absolute left-[25%] top-[82%] text-base opacity-40"
          aria-hidden="true"
          initial={{ opacity: 0, scale: 0 }}
          animate={{ opacity: 0.4, scale: 1 }}
          transition={{ delay: 1.4, type: 'spring' }}
        >
          ✨
        </motion.div>

      </motion.div>
    </div>
  );
};

export default HomePage;
