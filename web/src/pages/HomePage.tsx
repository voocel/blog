
import React, { useEffect, useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { motion } from 'framer-motion';
import BentoItem from '../components/bento/BentoItem';
import ProfileWidget from '../components/bento/widgets/ProfileWidget';
import AnimatedNavWidget from '../components/AnimatedNavWidget';
import ClockWidget from '../components/bento/widgets/ClockWidget';
import CalendarWidget from '../components/bento/widgets/CalendarWidget';
import MediaWidget from '../components/bento/widgets/MediaWidget';
import SocialWidget from '../components/bento/widgets/SocialWidget';
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

const HomePage: React.FC = () => {
  const navigate = useNavigate();
  const { user, setAuthModalOpen } = useAuth();
  const [latestPost, setLatestPost] = useState<BlogPost | null>(null);
  const [isSettingsOpen, setIsSettingsOpen] = useState(false);

  // Fetch latest post on mount
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
  }, []);

  const catImage = "https://images.unsplash.com/photo-1514888286974-6c03e2ca1dba?q=80&w=2643&auto=format&fit=crop";

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

      {/* Main Layout Container - Scattered organic positioning with staggered animations */}
      <motion.div
        className="relative w-full max-w-[1100px] mx-auto h-screen p-8"
        variants={containerVariants}
        initial="hidden"
        animate="visible"
      >

        {/* ==================== LEFT ZONE ==================== */}

        {/* Navigation Widget - Site navigation menu with animation */}
        <motion.div className="absolute left-[10%] top-[8%] w-[280px]" variants={itemVariants}>
          <AnimatedNavWidget isCompact={false} />
        </motion.div>

        {/* Recent Post Widget - Latest blog post preview */}
        <motion.div className="absolute left-[12%] top-[62%] w-[210px]" variants={itemVariants}>
          <BentoItem className="h-[170px] hover:!bg-white cursor-pointer relative group p-0 overflow-hidden shadow-sm">
            {latestPost ? (
              <div
                onClick={() => navigate(`/post/${latestPost.id}`)}
                className="h-full flex flex-col"
                role="button"
                tabIndex={0}
                onKeyDown={(e) => e.key === 'Enter' && navigate(`/post/${latestPost.id}`)}
                aria-label={`Read article: ${latestPost.title}`}
              >
                <div className="h-20 w-full bg-stone-100 relative">
                  <img src={latestPost.cover} className="w-full h-full object-cover opacity-90 group-hover:opacity-100 transition-opacity" alt={latestPost.title} />
                  <div className="absolute top-2 left-2 bg-white/90 backdrop-blur px-2 py-0.5 rounded text-[9px] font-bold text-stone-600">
                    最新文章
                  </div>
                </div>
                <div className="p-2.5 bg-white/50 backdrop-blur-sm flex-1">
                  <h4 className="font-bold text-stone-700 text-xs leading-tight mb-1 line-clamp-2">{latestPost.title}</h4>
                  <div className="text-[9px] text-stone-400">{new Date(latestPost.publishAt).toLocaleDateString()}</div>
                </div>
              </div>
            ) : (
              <div className="h-full flex items-center justify-center text-stone-300 text-xs">It's quiet here...</div>
            )}
          </BentoItem>
        </motion.div>

        {/* ==================== CENTER ZONE ==================== */}

        {/* Cat Image Widget - Decorative hero image */}
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

        {/* Profile Widget - User greeting and introduction */}
        <motion.div className="absolute left-[50%] -translate-x-1/2 top-[26%] w-[280px]" variants={itemVariants}>
          <BentoItem className="h-[220px] flex items-center justify-center bg-gradient-to-b from-white/80 to-orange-50/50 shadow-xl border-white/60">
            <ProfileWidget />
          </BentoItem>
        </motion.div>

        {/* Social Widget - Social media links */}
        <motion.div className="absolute left-[50%] -translate-x-1/2 top-[58%] w-[250px]" variants={itemVariants}>
          <div className="h-[50px]">
            <SocialWidget />
          </div>
        </motion.div>

        {/* Random Pick Widget - Random content recommendation */}
        <motion.div className="absolute left-[34%] top-[70%] w-[170px]" variants={itemVariants}>
          <BentoItem className="h-[100px] bg-gradient-to-br from-orange-200/80 to-rose-200/80 !text-stone-700 !border-none !p-3 shadow-md">
            <div className="flex flex-col h-full justify-between">
              <span className="text-[8px] opacity-60 uppercase tracking-wider">Random Pick</span>
              <div className="flex items-center gap-2">
                <div className="text-xl" aria-hidden="true">🌤️</div>
                <div>
                  <div className="font-bold text-xs leading-tight">Summer</div>
                  <div className="text-[8px] opacity-60">Afternoon</div>
                </div>
              </div>
            </div>
          </BentoItem>
        </motion.div>

        {/* Music Widget - Now playing music */}
        <motion.div className="absolute left-[54%] top-[70%] w-[220px]" variants={itemVariants}>
          <BentoItem className="h-[60px] !bg-orange-50/60 !p-1.5 !border-none shadow-sm flex items-center">
            <MediaWidget />
          </BentoItem>
        </motion.div>

        {/* ==================== RIGHT ZONE ==================== */}

        {/* Auth Buttons - Login/Dashboard and Settings */}
        <motion.div className="absolute right-[12%] top-[6%] flex gap-2" variants={itemVariants}>
          {user && (
            <button
              onClick={handleDashboardClick}
              className="bg-red-400 text-white px-4 py-2 rounded-xl shadow-red-200 shadow-md font-bold text-sm hover:bg-red-500 transition-colors"
              aria-label="Go to dashboard"
            >
              Dashboard
            </button>
          )}

          <div className="flex items-center gap-1.5 bg-white/50 backdrop-blur-md rounded-xl px-2 shadow-sm">
            {!user && (
              <button
                onClick={handleLoginClick}
                className="text-[10px] text-stone-400 font-bold px-1.5 hover:text-stone-600 transition-colors"
                aria-label="Sign in to your account"
              >
                SIGN IN
              </button>
            )}
            <button
              onClick={() => setIsSettingsOpen(true)}
              className="w-7 h-7 rounded-lg hover:bg-stone-200/50 flex items-center justify-center text-stone-400 transition-colors"
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

        {/* Clock Widget - Current time display */}
        <motion.div className="absolute right-[10%] top-[14%] w-[170px]" variants={itemVariants}>
          <BentoItem className="h-[100px] !p-0 !bg-stone-100/70 !border-white/30 shadow-inner">
            <ClockWidget />
          </BentoItem>
        </motion.div>

        {/* Calendar Widget - Monthly calendar view */}
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
