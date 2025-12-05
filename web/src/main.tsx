import React, { useState } from 'react';
import { BrowserRouter as Router, Routes, Route, useLocation, Navigate } from 'react-router-dom';
import { HelmetProvider } from 'react-helmet-async';
import { BlogProvider } from './context/BlogContext';
import type { AdminSection } from './types';
import Header from './components/Header';
import Sidebar from './components/Sidebar';
import AuthModal from './components/AuthModal';
import HomePage from './pages/HomePage';
import PostPage from './pages/PostPage';
import AdminDashboard from './pages/AdminDashboard';
import AboutPage from './pages/AboutPage';
import AIChat from './components/AIChat';
import SettingsPage from './pages/SettingsPage';

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
          <AdminDashboard section={adminSection} onExit={() => window.location.href = '/'} />
        </main>
      </div>
    );
  }

  // Public Layout
  return (
    <div className="min-h-screen text-ink font-sans selection:bg-gold-500/30 selection:text-white bg-transparent">
      <AuthModal />

      {/* Header is shown on all public pages except settings if we wanted, but let's keep it consistent */}
      {location.pathname !== '/settings' && (
        <Header />
      )}

      <Routes>
        <Route path="/" element={<HomePage />} />
        <Route path="/post/:id" element={<PostPage />} />
        <Route path="/about" element={<AboutPage />} />
        <Route path="/settings" element={<SettingsPage onExit={() => window.location.href = '/'} />} />
        {/* Fallback */}
        <Route path="*" element={<Navigate to="/" replace />} />
      </Routes>

      {/* AI Chat is available on the public site */}
      {location.pathname !== '/settings' && <AIChat />}
    </div>
  );
};

const App: React.FC = () => {
  return (
    <HelmetProvider>
      <BlogProvider>
        <Router>
          <AppContent />
        </Router>
      </BlogProvider>
    </HelmetProvider>
  );
};

export default App;