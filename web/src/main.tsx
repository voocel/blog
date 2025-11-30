import React, { useState } from 'react';
import { HashRouter as Router } from 'react-router-dom';
import { BlogProvider } from './context/BlogContext';
import type { AdminSection } from './types';
import Header from './components/Header';
import Sidebar from './components/Sidebar';
import AuthModal from './components/AuthModal';
import HomePage from './pages/HomePage';
import AdminDashboard from './pages/AdminDashboard';
import AboutPage from './pages/AboutPage';
import AIChat from './components/AIChat';
import SettingsPage from './pages/SettingsPage';

const AppContent: React.FC = () => {
  const [view, setView] = useState<'public' | 'admin' | 'about' | 'settings'>('public');
  const [adminSection, setAdminSection] = useState<AdminSection>('overview');

  const handleAdminAccess = () => {
    setView('admin');
  };

  const handleSettingsAccess = () => {
    setView('settings');
  };

  const handleExitAdmin = () => {
    setView('public');
  };

  return (
    <div className="min-h-screen text-ink font-sans selection:bg-gold-500/30 selection:text-white bg-transparent">

      <AuthModal />

      {/* Main Layout Switcher */}
      {view === 'public' || view === 'about' || view === 'settings' ? (
        <>
          {view !== 'settings' && (
            <Header
              onAdminClick={handleAdminAccess}
              onAboutClick={() => setView('about')}
              onHomeClick={() => setView('public')}
              onSettingsClick={handleSettingsAccess}
            />
          )}

          {view === 'public' && <HomePage />}
          {view === 'about' && <AboutPage />}
          {view === 'settings' && <SettingsPage onExit={() => setView('public')} />}

          {/* AI Chat is available on the public site */}
          {view !== 'settings' && <AIChat />}
        </>
      ) : (
        <div className="flex h-screen bg-transparent">
          <Sidebar
            currentSection={adminSection}
            setSection={setAdminSection}
            onExit={handleExitAdmin}
          />
          <main className="flex-1 ml-64 bg-transparent overflow-auto">
            <AdminDashboard section={adminSection} onExit={handleExitAdmin} />
          </main>
        </div>
      )}

    </div>
  );
};

const App: React.FC = () => {
  return (
    <BlogProvider>
      <Router>
        <AppContent />
      </Router>
    </BlogProvider>
  );
};

export default App;