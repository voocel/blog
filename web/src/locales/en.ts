// English translations
export const en = {
  // Settings Modal
  settings: {
    title: 'Settings',
    appearance: 'Appearance',
    music: 'Music Player',
    language: 'Language',
    close: 'Close',
    save: 'Save Changes',
    reset: 'Reset to Default',
    saved: 'Settings saved!',
    resetConfirm: 'Reset all settings to default?',
  },

  // Appearance Settings
  appearance: {
    themeMode: 'Theme Mode',
    light: 'Light',
    dark: 'Dark',
    auto: 'Auto',
    themeFollows: 'Theme follows your system preference',
    usingTheme: 'Using {theme} theme',
    animations: 'Animations',
    enableAnimations: 'Enable Animations',
    animationsDesc: 'Page transitions and effects',
  },

  // Music Settings
  music: {
    display: 'Display',
    showPlayer: 'Show Music Player',
    showPlayerDesc: 'Display player on homepage',
    defaultVolume: 'Default Volume',
    volume: 'Volume',
    playback: 'Playback',
    autoPlayNext: 'Auto Play Next',
    autoPlayNextDesc: 'Automatically play next song',
    loopPlaylist: 'Loop Playlist',
    loopPlaylistDesc: 'Repeat playlist when finished',
  },

  // Language Settings
  languageSettings: {
    title: 'Language',
    selectLanguage: 'Select Language',
    languageDesc: 'Choose your preferred language',
  },

  // HomePage
  home: {
    latestPost: 'Latest Post',
    randomPick: 'Random Pick',
    readMore: 'Click to read more...',
    noMusic: 'No Music',
    addSongs: 'Add songs to play',
    dashboard: 'Dashboard',
    signIn: 'Sign In',
  },
};

export type TranslationKeys = typeof en;
