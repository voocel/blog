import { useState, useEffect, useRef, useCallback } from 'react';
import { useSettings } from '@/context/SettingsContext';

export interface Song {
  id: string;
  title: string;
  artist?: string;
  url: string;
  cover?: string;
}

export interface UseAudioPlayerReturn {
  // Playback state
  isPlaying: boolean;
  currentTime: number;
  duration: number;
  progress: number;
  volume: number;
  currentSong: Song | null;
  error: string | null;

  // Controls
  play: () => void;
  pause: () => void;
  togglePlay: () => void;
  seek: (time: number) => void;
  setVolume: (volume: number) => void;
  nextSong: () => void;
  prevSong: () => void;
  loadSong: (song: Song) => void;
}

export const useAudioPlayer = (playlist: Song[] = []): UseAudioPlayerReturn => {
  const { settings } = useSettings();
  const audioRef = useRef<HTMLAudioElement | null>(null);
  const [isPlaying, setIsPlaying] = useState(false);
  const [currentTime, setCurrentTime] = useState(0);
  const [duration, setDuration] = useState(0);
  const [volume, setVolumeState] = useState(settings.music.defaultVolume);
  const [currentSongIndex, setCurrentSongIndex] = useState(0);
  const [error, setError] = useState<string | null>(null);

  // Initialize audio element
  useEffect(() => {
    // Don't initialize audio if playlist is empty
    if (playlist.length === 0) {
      return;
    }

    const audio = new Audio();
    audio.volume = volume;
    audioRef.current = audio;

    // Load first song
    audio.src = playlist[0].url;
    audio.load();

    return () => {
      audio.pause();
      audio.src = '';
    };
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  // Update audio source when song changes
  useEffect(() => {
    if (audioRef.current && playlist.length > 0) {
      const wasPlaying = isPlaying;
      const currentSong = playlist[currentSongIndex];

      audioRef.current.src = currentSong.url;
      audioRef.current.load();

      if (wasPlaying) {
        audioRef.current.play().catch(err => {
          console.error('Failed to play audio:', err);
          setError(`Failed to play: ${currentSong.title}`);
          setIsPlaying(false);
        });
      }
    }
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [currentSongIndex, playlist]);

  // Setup event listeners
  useEffect(() => {
    const audio = audioRef.current;
    if (!audio) return;

    const handleTimeUpdate = () => {
      setCurrentTime(audio.currentTime);
    };

    const handleDurationChange = () => {
      setDuration(audio.duration);
    };

    const handleEnded = () => {
      setIsPlaying(false);
      // Auto play next song if enabled
      if (settings.music.autoPlayNext && playlist.length > 1) {
        const nextIndex = (currentSongIndex + 1) % playlist.length;
        // If loop is disabled and we're at the end, don't play
        if (!settings.music.loop && nextIndex === 0) {
          return;
        }
        setCurrentSongIndex(nextIndex);
        setTimeout(() => {
          audio.play().catch(err => {
            console.error('Auto-play failed:', err);
            setError('Auto-play failed');
          });
        }, 100);
      }
    };

    const handlePlay = () => {
      setIsPlaying(true);
      setError(null);
    };

    const handlePause = () => setIsPlaying(false);

    const handleError = (e: Event) => {
      console.error('Audio error:', e);
      setError('Failed to load audio file');
      setIsPlaying(false);
    };

    const handleCanPlay = () => {
      setError(null);
    };

    audio.addEventListener('timeupdate', handleTimeUpdate);
    audio.addEventListener('durationchange', handleDurationChange);
    audio.addEventListener('ended', handleEnded);
    audio.addEventListener('play', handlePlay);
    audio.addEventListener('pause', handlePause);
    audio.addEventListener('error', handleError);
    audio.addEventListener('canplay', handleCanPlay);

    return () => {
      audio.removeEventListener('timeupdate', handleTimeUpdate);
      audio.removeEventListener('durationchange', handleDurationChange);
      audio.removeEventListener('ended', handleEnded);
      audio.removeEventListener('play', handlePlay);
      audio.removeEventListener('pause', handlePause);
      audio.removeEventListener('error', handleError);
      audio.removeEventListener('canplay', handleCanPlay);
    };
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [playlist.length, currentSongIndex]);

  const play = useCallback(() => {
    if (audioRef.current) {
      audioRef.current.play().catch(err => {
        console.error('Failed to play audio:', err);
        setError('Playback failed. Click to retry.');
        setIsPlaying(false);
      });
    }
  }, []);

  const pause = useCallback(() => {
    if (audioRef.current) {
      audioRef.current.pause();
    }
  }, []);

  const togglePlay = useCallback(() => {
    if (isPlaying) {
      pause();
    } else {
      play();
    }
  }, [isPlaying, play, pause]);

  const seek = useCallback((time: number) => {
    if (audioRef.current) {
      audioRef.current.currentTime = time;
      setCurrentTime(time);
    }
  }, []);

  const setVolume = useCallback((newVolume: number) => {
    const clampedVolume = Math.max(0, Math.min(1, newVolume));
    setVolumeState(clampedVolume);
    if (audioRef.current) {
      audioRef.current.volume = clampedVolume;
    }
  }, []);

  const nextSong = useCallback(() => {
    if (playlist.length > 0) {
      setCurrentSongIndex((prev) => (prev + 1) % playlist.length);
    }
  }, [playlist.length]);

  const prevSong = useCallback(() => {
    if (playlist.length > 0) {
      setCurrentSongIndex((prev) => (prev - 1 + playlist.length) % playlist.length);
    }
  }, [playlist.length]);

  const loadSong = useCallback((song: Song) => {
    const index = playlist.findIndex(s => s.id === song.id);
    if (index !== -1) {
      setCurrentSongIndex(index);
    }
  }, [playlist]);

  const progress = duration > 0 ? (currentTime / duration) * 100 : 0;
  const currentSong = playlist.length > 0 ? playlist[currentSongIndex] : null;

  return {
    isPlaying,
    currentTime,
    duration,
    progress,
    volume,
    currentSong,
    error,
    play,
    pause,
    togglePlay,
    seek,
    setVolume,
    nextSong,
    prevSong,
    loadSong,
  };
};
