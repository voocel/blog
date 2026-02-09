import React from 'react';
import { useAudioPlayer } from '@/hooks/useAudioPlayer';
import { MUSIC_PLAYLIST } from '@/config/musicConfig';

const MediaWidget: React.FC = () => {
    const {
        isPlaying,
        progress,
        currentSong,
        togglePlay,
        nextSong,
        prevSong
    } = useAudioPlayer(MUSIC_PLAYLIST);

    const displayTitle = currentSong?.title || 'No Music';
    const displayArtist = currentSong?.artist || '';
    const hasPlaylist = MUSIC_PLAYLIST.length > 0;

    // Show placeholder when no music is configured
    if (!hasPlaylist) {
        return (
            <div className="flex items-center justify-between h-full px-2 w-full gap-2">
                <div className="flex items-center gap-2 flex-1 min-w-0">
                    {/* Disabled Play Button */}
                    <div
                        className="w-12 h-12 flex-shrink-0 rounded-full flex items-center justify-center bg-orange-200 text-white shadow-md opacity-50 cursor-not-allowed"
                        title="No music configured"
                    >
                        <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="currentColor" className="w-6 h-6 pl-1">
                            <path fillRule="evenodd" d="M4.5 5.653c0-1.426 1.529-2.33 2.779-1.643l11.54 6.348c1.295.712 1.295 2.573 0 3.285L7.28 19.991c-1.25.687-2.779-.217-2.779-1.643V5.653z" clipRule="evenodd" />
                        </svg>
                    </div>

                    {/* Placeholder Info */}
                    <div className="flex flex-col flex-1 min-w-0 ml-1">
                        <span className="text-[var(--color-text-secondary)] font-bold text-xs">No Music</span>
                        <span className="text-[var(--color-text-muted)] text-[10px]">Add songs to play</span>
                        {/* Empty Progress Bar */}
                        <div className="w-full max-w-[120px] h-1 bg-[var(--color-muted)] rounded-full mt-1 overflow-hidden">
                            <div className="h-full bg-orange-300 rounded-full w-0"></div>
                        </div>
                    </div>
                </div>

                {/* Static Note Icon */}
                <div className="text-lg flex-shrink-0 opacity-30">🎵</div>
            </div>
        );
    }

    return (
        <div className="flex items-center justify-between h-full px-2 w-full gap-2">
            <div className="flex items-center gap-2 flex-1 min-w-0">
                {/* Previous Button */}
                <button
                    onClick={prevSong}
                    className="w-7 h-7 flex-shrink-0 rounded-full flex items-center justify-center bg-orange-300/50 text-white hover:bg-orange-300 transition-colors cursor-pointer"
                    title="Previous song"
                    disabled={MUSIC_PLAYLIST.length <= 1}
                >
                    <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="currentColor" className="w-4 h-4">
                        <path d="M9.195 18.44c1.25.713 2.805-.19 2.805-1.629v-2.34l6.945 3.968c1.25.714 2.805-.188 2.805-1.628V8.688c0-1.44-1.555-2.342-2.805-1.628L12 11.03v-2.34c0-1.44-1.555-2.343-2.805-1.629l-7.108 4.062c-1.26.72-1.26 2.536 0 3.256l7.108 4.061z" />
                    </svg>
                </button>

                {/* Play/Pause Button */}
                <button
                    className={`w-12 h-12 flex-shrink-0 rounded-full flex items-center justify-center bg-orange-400 text-white shadow-md cursor-pointer transition-transform ${isPlaying ? 'scale-95' : 'hover:scale-110'}`}
                    onClick={togglePlay}
                    title={isPlaying ? 'Pause' : 'Play'}
                >
                    {isPlaying ? (
                        <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="currentColor" className="w-6 h-6">
                            <path fillRule="evenodd" d="M6.75 5.25a.75.75 0 01.75-.75H9a.75.75 0 01.75.75v13.5a.75.75 0 01-.75.75H7.5a.75.75 0 01-.75-.75V5.25zm7.5 0A.75.75 0 0115 4.5h1.5a.75.75 0 01.75.75v13.5a.75.75 0 01-.75.75H15a.75.75 0 01-.75-.75V5.25z" clipRule="evenodd" />
                        </svg>
                    ) : (
                        <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="currentColor" className="w-6 h-6 pl-1">
                            <path fillRule="evenodd" d="M4.5 5.653c0-1.426 1.529-2.33 2.779-1.643l11.54 6.348c1.295.712 1.295 2.573 0 3.285L7.28 19.991c-1.25.687-2.779-.217-2.779-1.643V5.653z" clipRule="evenodd" />
                        </svg>
                    )}
                </button>

                {/* Next Button */}
                <button
                    onClick={nextSong}
                    className="w-7 h-7 flex-shrink-0 rounded-full flex items-center justify-center bg-orange-300/50 text-white hover:bg-orange-300 transition-colors cursor-pointer"
                    title="Next song"
                    disabled={MUSIC_PLAYLIST.length <= 1}
                >
                    <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="currentColor" className="w-4 h-4">
                        <path d="M5.055 7.06c-1.25-.714-2.805.189-2.805 1.628v8.123c0 1.44 1.555 2.342 2.805 1.628L12 14.471v2.34c0 1.44 1.555 2.342 2.805 1.628l7.108-4.061c1.26-.72 1.26-2.536 0-3.256L14.805 7.06C13.555 6.346 12 7.25 12 8.688v2.34L5.055 7.06z" />
                    </svg>
                </button>

                {/* Song Info - Flexible width */}
                <div className="flex flex-col flex-1 min-w-0 ml-1">
                    <span className="text-ink font-bold text-xs truncate" title={displayTitle}>
                        {displayTitle}
                    </span>
                    {displayArtist && (
                        <span className="text-[var(--color-text-secondary)] text-[10px] truncate" title={displayArtist}>
                            {displayArtist}
                        </span>
                    )}
                    {/* Real-time Progress Bar */}
                    <div className="w-full max-w-[120px] h-1 bg-[var(--color-muted)] rounded-full mt-1 overflow-hidden">
                        <div
                            className="h-full bg-orange-400 rounded-full transition-all duration-300"
                            style={{ width: `${progress}%` }}
                        ></div>
                    </div>
                </div>
            </div>

            {/* Animated Notes */}
            <div className={`text-lg flex-shrink-0 transition-all ${isPlaying ? 'animate-bounce' : 'opacity-50'}`}>
                🎵
            </div>
        </div>
    );
};

export default MediaWidget;
