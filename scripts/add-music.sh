#!/bin/bash

# Music File Helper Script
# Usage: ./add-music.sh /path/to/your/music/file.mp3 "Song Title" "Artist Name"

if [ $# -lt 3 ]; then
    echo "Usage: ./add-music.sh <music-file-path> <song-title> <artist-name>"
    echo "Example: ./add-music.sh ~/Music/song.mp3 \"Shape of You\" \"Ed Sheeran\""
    exit 1
fi

MUSIC_FILE="$1"
SONG_TITLE="$2"
ARTIST_NAME="$3"

# Check if file exists
if [ ! -f "$MUSIC_FILE" ]; then
    echo "Error: File not found: $MUSIC_FILE"
    exit 1
fi

# Get filename
FILENAME=$(basename "$MUSIC_FILE")

# Target directory
TARGET_DIR="web/public/music"

# Create directory if it doesn't exist
mkdir -p "$TARGET_DIR"

# Copy file
echo "Copying file..."
cp "$MUSIC_FILE" "$TARGET_DIR/$FILENAME"

if [ $? -eq 0 ]; then
    echo "✅ File copied to: $TARGET_DIR/$FILENAME"
    echo ""
    echo "📝 Add the following configuration to web/src/config/musicConfig.ts:"
    echo ""
    echo "{"
    echo "  id: '$(date +%s)',"
    echo "  title: '$SONG_TITLE',"
    echo "  artist: '$ARTIST_NAME',"
    echo "  url: '/music/$FILENAME',"
    echo "},"
    echo ""
else
    echo "❌ Copy failed"
    exit 1
fi
