import React, { useState, Suspense } from 'react';
import { BrowserRouter as Router, Routes, Route, useLocation, Navigate } from 'react-router-dom';
import { AuthProvider } from './context/AuthContext';
import { BlogProvider } from './context/BlogContext';
import { AdminProvider } from './context/AdminContext';
import { SettingsProvider } from './context/SettingsContext';
import type { AdminSection } from './types';
import Sidebar from './components/Sidebar';
import AuthModal from './components/AuthModal';
import AIChat from './components/AIChat';
import ErrorBoundary from './components/ErrorBoundary';
import { ToastProvider } from './components/Toast';

// Lazy Load Pages
const HomePage = React.lazy(() => import('./pages/HomePage'));
const PostPage = React.lazy(() => import('./pages/PostPage'));
const PostsListPage = React.lazy(() => import('./pages/PostsListPage'));
const AdminDashboard = React.lazy(() => import('./pages/AdminDashboard'));
const AboutPage = React.lazy(() => import('./pages/AboutPage'));
const SettingsPage = React.lazy(() => import('./pages/SettingsPage'));
const ClockPage = React.lazy(() => import('./pages/ClockPage'));

const AppContent: React.FC = () => {
  const location = useLocation();
  const isAdminRoute = location.pathname.startsWith('/admin');
  const [adminSection, setAdminSection] = useState<AdminSection>('overview');

  // Admin Layout
  if (isAdminRoute) {
    return (
      <div className="flex h-screen bg-transparent">
        <Sidebar
          currentSection={adminSection}
          setSection={setAdminSection}
          onExit={() => window.location.href = '/'}
        />
        <main className="flex-1 ml-72 bg-transparent overflow-auto">
          <Suspense fallback={<div className="h-full flex items-center justify-center text-stone-400">Loading Dashboard...</div>}>
            <AdminDashboard section={adminSection} onExit={() => window.location.href = '/'} />
          </Suspense>
        </main>
      </div>
    );
  }

  // Public Layout
  return (
    <div className="min-h-screen text-ink font-sans selection:bg-gold-500/30 selection:text-white bg-transparent">
      <AuthModal />



      <Suspense fallback={
        <div className="min-h-screen flex items-center justify-center bg-transparent">
          <div className="animate-pulse flex flex-col items-center gap-4">
            <div className="w-12 h-12 rounded-full bg-stone-200"></div>
            <div className="h-4 w-32 bg-stone-200 rounded"></div>
          </div>
        </div>
      }>
        <Routes location={location} key={location.pathname}>
          <Route path="/" element={<HomePage />} />
          <Route path="/posts" element={<PostsListPage />} />
          <Route path="/post/:id" element={<PostPage />} />
          <Route path="/about" element={<AboutPage />} />
          <Route path="/clock" element={<ClockPage />} />
          <Route path="/settings" element={<SettingsPage onExit={() => window.location.href = '/'} />} />
          <Route path="*" element={<Navigate to="/" replace />} />
        </Routes>
      </Suspense>

      {location.pathname !== '/settings' && <AIChat />}
    </div>
  );
};

const App: React.FC = () => {
  return (
    <ToastProvider>
      <SettingsProvider>
        <AuthProvider>
          <BlogProvider>
            <AdminProvider>
              <Router>
                <ErrorBoundary>
                  <AppContent />
                </ErrorBoundary>
              </Router>
            </AdminProvider>
          </BlogProvider>
        </AuthProvider>
      </SettingsProvider>
    </ToastProvider>
  );
};

export default App;