import React, { createContext, useContext, useState, useEffect, type ReactNode } from 'react';
import {
  type UserSettings,
  type ThemeMode,
  type MusicSettings,
  loadSettings,
  saveSettings,
  resetSettings as resetSettingsUtil,
  getEffectiveTheme,
  getSystemTheme,
} from '@/config/settings';
import type { Locale } from '@/locales';

interface SettingsContextType {
  settings: UserSettings;
  updateTheme: (theme: ThemeMode) => void;
  updateAnimations: (enabled: boolean) => void;
  updateMusicSettings: (music: Partial<MusicSettings>) => void;
  updateLocale: (locale: Locale) => void;
  resetSettings: () => void;
  effectiveTheme: 'light' | 'dark';
}

const SettingsContext = createContext<SettingsContextType | undefined>(undefined);

// eslint-disable-next-line react-refresh/only-export-components
export const useSettings = () => {
  const context = useContext(SettingsContext);
  if (!context) {
    throw new Error('useSettings must be used within SettingsProvider');
  }
  return context;
};

interface SettingsProviderProps {
  children: ReactNode;
}

export const SettingsProvider: React.FC<SettingsProviderProps> = ({ children }) => {
  const [settings, setSettings] = useState<UserSettings>(() => loadSettings());
  const [effectiveTheme, setEffectiveTheme] = useState<'light' | 'dark'>(() =>
    getEffectiveTheme(loadSettings().appearance.theme)
  );

  // Save settings whenever they change
  useEffect(() => {
    saveSettings(settings);
    setEffectiveTheme(getEffectiveTheme(settings.appearance.theme));
  }, [settings]);

  // Listen for system theme changes when theme is 'auto'
  useEffect(() => {
    if (settings.appearance.theme !== 'auto') return;

    const mediaQuery = window.matchMedia('(prefers-color-scheme: dark)');
    const handleChange = () => {
      setEffectiveTheme(getSystemTheme());
    };

    mediaQuery.addEventListener('change', handleChange);
    return () => mediaQuery.removeEventListener('change', handleChange);
  }, [settings.appearance.theme]);

  // Apply theme to document
  useEffect(() => {
    document.documentElement.classList.remove('light', 'dark');
    document.documentElement.classList.add(effectiveTheme);
  }, [effectiveTheme]);

  const updateTheme = (theme: ThemeMode) => {
    setSettings((prev) => ({
      ...prev,
      appearance: { ...prev.appearance, theme },
    }));
  };

  const updateAnimations = (enabled: boolean) => {
    setSettings((prev) => ({
      ...prev,
      appearance: { ...prev.appearance, enableAnimations: enabled },
    }));
  };

  const updateMusicSettings = (music: Partial<MusicSettings>) => {
    setSettings((prev) => ({
      ...prev,
      music: { ...prev.music, ...music },
    }));
  };

  const updateLocale = (locale: Locale) => {
    setSettings((prev) => ({
      ...prev,
      language: { ...prev.language, locale },
    }));
  };

  const resetSettings = () => {
    const defaults = resetSettingsUtil();
    setSettings(defaults);
  };

  return (
    <SettingsContext.Provider
      value={{
        settings,
        updateTheme,
        updateAnimations,
        updateMusicSettings,
        updateLocale,
        resetSettings,
        effectiveTheme,
      }}
    >
      {children}
    </SettingsContext.Provider>
  );
};
