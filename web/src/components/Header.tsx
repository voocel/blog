

import React, { useState, useEffect, useRef } from 'react';
import { Link, useNavigate } from 'react-router-dom';
import { useBlog } from '../context/BlogContext';
import { IconMenu, IconX, IconUser, IconSettings, IconLogOut, IconGrid, IconUserCircle, IconLock, IconEye } from './Icons';

interface HeaderProps {
  // Props are no longer needed for navigation, but keeping interface clean if we need other props later
}

const Header: React.FC<HeaderProps> = () => {
  const { user, setAuthModalOpen, logout } = useBlog();
  const navigate = useNavigate();
  const [scrolled, setScrolled] = useState(false);
  const [mobileMenuOpen, setMobileMenuOpen] = useState(false);
  const [userMenuOpen, setUserMenuOpen] = useState(false);
  const dropdownRef = useRef<HTMLDivElement>(null);

  useEffect(() => {
    const handleScroll = () => {
      setScrolled(window.scrollY > 50);
    };
    window.addEventListener('scroll', handleScroll);
    return () => window.removeEventListener('scroll', handleScroll);
  }, []);

  // Close dropdown when clicking outside
  useEffect(() => {
    const handleClickOutside = (event: MouseEvent) => {
      if (dropdownRef.current && !dropdownRef.current.contains(event.target as Node)) {
        setUserMenuOpen(false);
      }
    };
    document.addEventListener('mousedown', handleClickOutside);
    return () => document.removeEventListener('mousedown', handleClickOutside);
  }, []);

  const handleLogout = () => {
    setUserMenuOpen(false);
    logout();
    navigate('/');
  };

  return (
    <>
      <header className={`fixed top-0 left-0 w-full z-40 transition-all duration-500 ${scrolled
        ? 'bg-paper/80 backdrop-blur-lg border-b border-stone-200 py-3 shadow-sm'
        : 'bg-transparent border-transparent py-6'
        }`}>
        <div className="max-w-7xl mx-auto px-6 flex justify-between items-center">
          {/* Logo */}
          <Link
            to="/"
            className={`text-2xl md:text-3xl font-serif font-bold tracking-tighter cursor-pointer z-50 relative flex items-center gap-1 transition-colors ${scrolled ? 'text-ink' : 'text-ink'}`}
          >
            Voocel<span className="text-gold-500">.</span>
          </Link>

          <nav className="hidden md:flex items-center gap-8">
            <Link to="/" className="text-xs uppercase tracking-[0.2em] font-medium text-stone-600 hover:text-ink transition-colors cursor-pointer">Journal</Link>
            <button className="text-xs uppercase tracking-[0.2em] font-medium text-stone-600 hover:text-ink transition-colors cursor-pointer">Collections</button>
            <Link to="/about" className="text-xs uppercase tracking-[0.2em] font-medium text-stone-600 hover:text-ink transition-colors cursor-pointer">About</Link>
          </nav>

          {/* Auth & Actions */}
          <div className="hidden md:flex items-center gap-6">
            {user ? (
              <div className="relative flex items-center" ref={dropdownRef}>

                {/* Role Badge */}
                <div className="flex items-center gap-2 mr-5">
                  {user.role === 'admin' ? (
                    <Link
                      to="/admin"
                      className="flex items-center gap-2 group cursor-pointer"
                      title="Go to Dashboard"
                    >
                      <IconLock className="w-3 h-3 text-gold-500" />
                      <span className="text-[10px] font-bold uppercase tracking-[0.2em] text-gold-600 group-hover:text-gold-700 transition-colors">
                        Admin
                      </span>
                    </Link>
                  ) : (
                    <div className="flex items-center gap-2 select-none opacity-80">
                      <IconEye className="w-3 h-3 text-stone-400" />
                      <span className="text-[10px] font-bold uppercase tracking-[0.2em] text-stone-400">
                        Reader
                      </span>
                    </div>
                  )}
                </div>

                {/* Vertical Divider */}
                <div className="h-8 w-px bg-stone-200 mr-5" />

                {/* User Info Block */}
                <div className="flex flex-col items-start mr-5">
                  <span className="text-[9px] font-medium uppercase tracking-widest text-stone-400 leading-tight mb-0.5">
                    Signed in as
                  </span>
                  <span className="font-serif text-base font-bold text-[#96742E] leading-none tracking-wide">
                    {user.username}
                  </span>
                </div>

                {/* Avatar Trigger */}
                <button
                  onClick={() => setUserMenuOpen(!userMenuOpen)}
                  className="flex items-center gap-3 focus:outline-none group cursor-pointer"
                >
                  <div className={`w-10 h-10 rounded-full p-[2px] transition-all duration-[600ms] ease-in-out group-hover:rotate-[360deg] ${user.role === 'admin'
                    ? 'bg-gradient-to-tr from-gold-300 to-gold-600 shadow-md group-hover:shadow-lg'
                    : 'bg-stone-200 group-hover:bg-stone-300'
                    }`}>
                    <div className="w-full h-full rounded-full overflow-hidden bg-white border border-white">
                      {user.avatar ? (
                        <img src={user.avatar} alt={user.username} className="w-full h-full object-cover" />
                      ) : (
                        <div className="w-full h-full flex items-center justify-center bg-stone-50 text-stone-400">
                          <IconUserCircle className="w-full h-full opacity-80" />
                        </div>
                      )}
                    </div>
                  </div>
                </button>

                {/* Dropdown Menu */}
                {userMenuOpen && (
                  <div className="absolute right-0 top-full mt-3 w-56 bg-white/90 backdrop-blur-xl border border-white/50 rounded-xl shadow-[0_8px_32px_rgba(0,0,0,0.08)] py-2 animate-blur-in origin-top-right overflow-hidden ring-1 ring-black/5">
                    <div className="px-4 py-3 border-b border-stone-100 md:hidden">
                      <p className="text-sm font-medium text-ink truncate">{user.username}</p>
                      <p className="text-xs text-stone-500 truncate capitalize">{user.role}</p>
                    </div>

                    <div className="py-1">
                      {user.role === 'admin' && (
                        <Link
                          to="/admin"
                          onClick={() => setUserMenuOpen(false)}
                          className="w-full text-left px-4 py-2.5 text-sm text-stone-600 hover:bg-gold-50 hover:text-gold-700 flex items-center gap-3 transition-colors cursor-pointer"
                        >
                          <IconGrid className="w-4 h-4" />
                          <span>Dashboard</span>
                        </Link>
                      )}
                      <Link
                        to="/settings"
                        onClick={() => setUserMenuOpen(false)}
                        className="w-full text-left px-4 py-2.5 text-sm text-stone-600 hover:bg-stone-50 hover:text-ink flex items-center gap-3 transition-colors cursor-pointer"
                      >
                        <IconSettings className="w-4 h-4" />
                        <span>Settings</span>
                      </Link>
                    </div>

                    <div className="border-t border-stone-100 my-1"></div>

                    <button
                      onClick={handleLogout}
                      className="w-full text-left px-4 py-2.5 text-sm text-red-500 hover:bg-red-50 hover:text-red-600 flex items-center gap-3 transition-colors cursor-pointer"
                    >
                      <IconLogOut className="w-4 h-4" />
                      <span>Sign out</span>
                    </button>
                  </div>
                )}
              </div>
            ) : (
              <button
                onClick={() => setAuthModalOpen(true)}
                className="flex items-center gap-2 text-xs uppercase tracking-[0.15em] hover:text-gold-600 transition-colors text-ink cursor-pointer"
              >
                <IconUser className="w-4 h-4" />
                <span>Sign In</span>
              </button>
            )}
          </div>

          {/* Mobile Menu Toggle */}
          <button
            onClick={() => setMobileMenuOpen(!mobileMenuOpen)}
            className="md:hidden text-ink z-50 focus:outline-none cursor-pointer"
          >
            {mobileMenuOpen ? <IconX className="w-6 h-6" /> : <IconMenu className="w-6 h-6" />}
          </button>
        </div>
      </header>

      {/* Mobile Menu Overlay */}
      {mobileMenuOpen && (
        <div className="fixed inset-0 bg-paper z-30 flex flex-col items-center justify-center animate-fade-in md:hidden">
          <nav className="flex flex-col gap-8 text-center w-full px-8">
            <Link to="/" onClick={() => setMobileMenuOpen(false)} className="text-3xl font-serif text-ink cursor-pointer">Journal</Link>
            <button onClick={() => setMobileMenuOpen(false)} className="text-3xl font-serif text-stone-500 cursor-pointer">Collections</button>
            <Link to="/about" onClick={() => setMobileMenuOpen(false)} className="text-3xl font-serif text-stone-500 cursor-pointer">About</Link>

            <div className="w-12 h-px bg-stone-300 mx-auto my-4"></div>

            {user ? (
              <div className="flex flex-col gap-4 items-center animate-slide-up">
                <div className="flex items-center gap-3 mb-4">
                  <div className="w-12 h-12 rounded-full overflow-hidden border-2 border-stone-200">
                    {user.avatar ? (
                      <img src={user.avatar} alt="Avatar" className="w-full h-full object-cover" />
                    ) : (
                      <div className="w-full h-full bg-stone-100 flex items-center justify-center text-stone-400">
                        <IconUserCircle className="w-8 h-8" />
                      </div>
                    )}
                  </div>
                  <div className="text-left">
                    <div className="font-serif text-xl text-ink">{user.username}</div>
                    <div className="text-xs uppercase tracking-widest text-stone-500 flex items-center gap-1">
                      {user.role === 'admin' && <IconLock className="w-3 h-3 text-gold-500" />}
                      {user.role}
                    </div>
                  </div>
                </div>

                {user.role === 'admin' && (
                  <Link
                    to="/admin"
                    onClick={() => setMobileMenuOpen(false)}
                    className="w-full max-w-xs py-3 border border-gold-200 bg-gold-50 text-gold-700 rounded-lg flex items-center justify-center gap-2 cursor-pointer"
                  >
                    <IconGrid className="w-4 h-4" /> Dashboard
                  </Link>
                )}

                <Link
                  to="/settings"
                  onClick={() => setMobileMenuOpen(false)}
                  className="text-stone-600 flex items-center gap-2"
                >
                  <IconSettings className="w-4 h-4" /> Settings
                </Link>

                <button
                  onClick={() => { logout(); setMobileMenuOpen(false); }}
                  className="text-red-500 uppercase tracking-widest text-sm font-bold mt-4 cursor-pointer"
                >
                  Sign Out
                </button>
              </div>
            ) : (
              <button onClick={() => { setAuthModalOpen(true); setMobileMenuOpen(false); }} className="text-xl text-ink font-serif">Sign In</button>
            )}
          </nav>
        </div>
      )}
    </>
  );
};

export default Header;

