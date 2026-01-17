// Settings Types and Default Configuration

export type ThemeMode = 'light' | 'dark' | 'auto';

export interface MusicSettings {
  defaultVolume: number; // 0.0 to 1.0
  autoPlayNext: boolean;
  loop: boolean;
  showPlayer: boolean;
}

export interface AppearanceSettings {
  theme: ThemeMode;
  enableAnimations: boolean;
}

export interface UserSettings {
  appearance: AppearanceSettings;
  music: MusicSettings;
}

// Default settings
export const DEFAULT_SETTINGS: UserSettings = {
  appearance: {
    theme: 'light',
    enableAnimations: true,
  },
  music: {
    defaultVolume: 0.7,
    autoPlayNext: true,
    loop: true,
    showPlayer: true,
  },
};

// LocalStorage key
export const SETTINGS_STORAGE_KEY = 'blog-user-settings';

// Helper functions
export const loadSettings = (): UserSettings => {
  try {
    const stored = localStorage.getItem(SETTINGS_STORAGE_KEY);
    if (stored) {
      const parsed = JSON.parse(stored);
      // Merge with defaults to handle new settings
      return {
        appearance: { ...DEFAULT_SETTINGS.appearance, ...parsed.appearance },
        music: { ...DEFAULT_SETTINGS.music, ...parsed.music },
      };
    }
  } catch (error) {
    console.error('Failed to load settings:', error);
  }
  return DEFAULT_SETTINGS;
};

export const saveSettings = (settings: UserSettings): void => {
  try {
    localStorage.setItem(SETTINGS_STORAGE_KEY, JSON.stringify(settings));
  } catch (error) {
    console.error('Failed to save settings:', error);
  }
};

export const resetSettings = (): UserSettings => {
  try {
    localStorage.removeItem(SETTINGS_STORAGE_KEY);
  } catch (error) {
    console.error('Failed to reset settings:', error);
  }
  return DEFAULT_SETTINGS;
};

// Get system theme preference
export const getSystemTheme = (): 'light' | 'dark' => {
  if (typeof window === 'undefined') return 'light';
  return window.matchMedia('(prefers-color-scheme: dark)').matches ? 'dark' : 'light';
};

// Get effective theme (resolves 'auto' to actual theme)
export const getEffectiveTheme = (theme: ThemeMode): 'light' | 'dark' => {
  if (theme === 'auto') {
    return getSystemTheme();
  }
  return theme;
};
