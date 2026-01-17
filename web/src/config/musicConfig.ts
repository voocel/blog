// Music Player Configuration
// Add your music files to public/music/ directory and configure them here

export interface Song {
  id: string;
  title: string;
  artist?: string;
  url: string;
  cover?: string;
}

// Default playlist configuration
// Using online music for quick testing (no download required)
// Replace with your own music files for production use

export const MUSIC_PLAYLIST: Song[] = [
  // Online test music (works out of the box)
  {
    id: '1',
    title: 'SoundHelix Song 1',
    artist: 'SoundHelix',
    url: 'https://www.soundhelix.com/examples/mp3/SoundHelix-Song-1.mp3',
  },
  {
    id: '2',
    title: 'SoundHelix Song 2',
    artist: 'SoundHelix',
    url: 'https://www.soundhelix.com/examples/mp3/SoundHelix-Song-2.mp3',
  },
  {
    id: '3',
    title: 'SoundHelix Song 3',
    artist: 'SoundHelix',
    url: 'https://www.soundhelix.com/examples/mp3/SoundHelix-Song-3.mp3',
  },

  // To use your own local music files:
  // 1. Place MP3 files in web/public/music/
  // 2. Uncomment and configure below:
  // {
  //   id: '4',
  //   title: 'Your Song',
  //   artist: 'Your Artist',
  //   url: '/music/your-song.mp3',
  // },
];

// Player settings
export const PLAYER_CONFIG = {
  defaultVolume: 0.7, // 0.0 to 1.0
  autoPlayNext: true, // Auto play next song when current ends
  loop: true, // Loop playlist
};
