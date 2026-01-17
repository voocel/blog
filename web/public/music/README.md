# Music Files

Place your music files here to use with the MediaWidget player.

## Supported Formats
- MP3 (recommended)
- WAV
- OGG
- M4A

## Usage

1. Add your music files to this directory
2. Update the playlist in `web/src/components/bento/widgets/MediaWidget.tsx`:

```typescript
const DEFAULT_PLAYLIST: Song[] = [
    {
        id: '1',
        title: 'Your Song Title',
        artist: 'Artist Name',
        url: '/music/your-song.mp3',
    },
    // Add more songs...
];
```

## Example Files

You can download free music from:
- [Free Music Archive](https://freemusicarchive.org/)
- [Incompetech](https://incompetech.com/music/royalty-free/)
- [YouTube Audio Library](https://www.youtube.com/audiolibrary)

## File Size Recommendations

- Keep files under 10MB for better loading performance
- Use MP3 format with 128-192 kbps bitrate for web
- Consider using a CDN for production deployments

## Notes

- The player will automatically loop through the playlist
- Progress bar shows real-time playback position
- Previous/Next buttons navigate through the playlist
- Volume is set to 70% by default
