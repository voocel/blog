import React, { useState } from 'react';
import { useSettings } from '../context/SettingsContext';
import { ThemeMode } from '../config/settings';

interface SettingsModalProps {
    isOpen: boolean;
    onClose: () => void;
}

const SettingsModal: React.FC<SettingsModalProps> = ({ isOpen, onClose }) => {
    if (!isOpen) return null;

    const { settings, updateTheme, updateAnimations, updateMusicSettings, resetSettings } = useSettings();
    const [activeTab, setActiveTab] = useState<'appearance' | 'music'>('appearance');
    const [showSaveToast, setShowSaveToast] = useState(false);

    const handleSave = () => {
        setShowSaveToast(true);
        setTimeout(() => setShowSaveToast(false), 2000);
    };

    const handleReset = () => {
        if (confirm('Reset all settings to default?')) {
            resetSettings();
            setShowSaveToast(true);
            setTimeout(() => setShowSaveToast(false), 2000);
        }
    };

    return (
        <div className="fixed inset-0 z-50 flex items-center justify-center p-4 bg-black/20 backdrop-blur-sm">
            <div className="bg-[#fdfaf6] w-full max-w-2xl max-h-[85vh] rounded-3xl shadow-2xl overflow-hidden flex flex-col animate-fade-in-up">

                {/* Header */}
                <div className="px-8 py-6 border-b border-stone-100 flex items-center justify-between bg-white/50 backdrop-blur-md">
                    <div className="flex gap-8 text-sm font-bold text-stone-400">
                        <button
                            onClick={() => setActiveTab('appearance')}
                            className={`${activeTab === 'appearance' ? 'text-red-500 relative after:absolute after:-bottom-6 after:left-0 after:w-full after:h-0.5 after:bg-red-500' : 'hover:text-stone-600'} transition-colors cursor-pointer`}
                        >
                            Appearance
                        </button>
                        <button
                            onClick={() => setActiveTab('music')}
                            className={`${activeTab === 'music' ? 'text-red-500 relative after:absolute after:-bottom-6 after:left-0 after:w-full after:h-0.5 after:bg-red-500' : 'hover:text-stone-600'} transition-colors cursor-pointer`}
                        >
                            Music Player
                        </button>
                    </div>

                    <div className="flex items-center gap-4">
                        <button
                            className="text-stone-400 hover:text-stone-600 text-sm font-medium cursor-pointer"
                            onClick={onClose}
                        >
                            Close
                        </button>
                    </div>
                </div>

                {/* Content */}
                <div className="flex-1 overflow-y-auto p-8 bg-stone-50/30">
                    {activeTab === 'appearance' && (
                        <div className="max-w-xl mx-auto space-y-8">
                            {/* Theme Mode */}
                            <div>
                                <h3 className="text-sm font-bold text-stone-700 mb-4">Theme Mode</h3>
                                <div className="grid grid-cols-3 gap-3">
                                    {(['light', 'dark', 'auto'] as ThemeMode[]).map((theme) => (
                                        <button
                                            key={theme}
                                            onClick={() => updateTheme(theme)}
                                            className={`p-4 rounded-xl border-2 transition-all cursor-pointer ${
                                                settings.appearance.theme === theme
                                                    ? 'border-red-400 bg-red-50'
                                                    : 'border-stone-200 bg-white hover:border-stone-300'
                                            }`}
                                        >
                                            <div className="text-2xl mb-2">
                                                {theme === 'light' && '☀️'}
                                                {theme === 'dark' && '🌙'}
                                                {theme === 'auto' && '🌓'}
                                            </div>
                                            <div className="text-sm font-medium text-stone-700 capitalize">
                                                {theme}
                                            </div>
                                        </button>
                                    ))}
                                </div>
                                <p className="text-xs text-stone-500 mt-2">
                                    {settings.appearance.theme === 'auto'
                                        ? 'Theme follows your system preference'
                                        : `Using ${settings.appearance.theme} theme`}
                                </p>
                            </div>

                            {/* Animations */}
                            <div>
                                <h3 className="text-sm font-bold text-stone-700 mb-4">Animations</h3>
                                <label className="flex items-center justify-between p-4 bg-white rounded-xl border-2 border-stone-200 cursor-pointer hover:border-stone-300 transition-colors">
                                    <div>
                                        <div className="text-sm font-medium text-stone-700">Enable Animations</div>
                                        <div className="text-xs text-stone-500 mt-1">Page transitions and effects</div>
                                    </div>
                                    <input
                                        type="checkbox"
                                        checked={settings.appearance.enableAnimations}
                                        onChange={(e) => updateAnimations(e.target.checked)}
                                        className="w-5 h-5 text-red-500 rounded focus:ring-2 focus:ring-red-200"
                                    />
                                </label>
                            </div>
                        </div>
                    )}

                    {activeTab === 'music' && (
                        <div className="max-w-xl mx-auto space-y-8">
                            {/* Show Player */}
                            <div>
                                <h3 className="text-sm font-bold text-stone-700 mb-4">Display</h3>
                                <label className="flex items-center justify-between p-4 bg-white rounded-xl border-2 border-stone-200 cursor-pointer hover:border-stone-300 transition-colors">
                                    <div>
                                        <div className="text-sm font-medium text-stone-700">Show Music Player</div>
                                        <div className="text-xs text-stone-500 mt-1">Display player on homepage</div>
                                    </div>
                                    <input
                                        type="checkbox"
                                        checked={settings.music.showPlayer}
                                        onChange={(e) => updateMusicSettings({ showPlayer: e.target.checked })}
                                        className="w-5 h-5 text-red-500 rounded focus:ring-2 focus:ring-red-200"
                                    />
                                </label>
                            </div>

                            {/* Volume */}
                            <div>
                                <h3 className="text-sm font-bold text-stone-700 mb-4">Default Volume</h3>
                                <div className="p-4 bg-white rounded-xl border-2 border-stone-200">
                                    <div className="flex items-center justify-between mb-3">
                                        <span className="text-sm text-stone-600">Volume</span>
                                        <span className="text-sm font-bold text-red-500">
                                            {Math.round(settings.music.defaultVolume * 100)}%
                                        </span>
                                    </div>
                                    <input
                                        type="range"
                                        min="0"
                                        max="100"
                                        value={settings.music.defaultVolume * 100}
                                        onChange={(e) => updateMusicSettings({ defaultVolume: Number(e.target.value) / 100 })}
                                        className="w-full h-2 bg-stone-200 rounded-lg appearance-none cursor-pointer accent-red-500"
                                    />
                                </div>
                            </div>

                            {/* Playback Options */}
                            <div>
                                <h3 className="text-sm font-bold text-stone-700 mb-4">Playback</h3>
                                <div className="space-y-3">
                                    <label className="flex items-center justify-between p-4 bg-white rounded-xl border-2 border-stone-200 cursor-pointer hover:border-stone-300 transition-colors">
                                        <div>
                                            <div className="text-sm font-medium text-stone-700">Auto Play Next</div>
                                            <div className="text-xs text-stone-500 mt-1">Automatically play next song</div>
                                        </div>
                                        <input
                                            type="checkbox"
                                            checked={settings.music.autoPlayNext}
                                            onChange={(e) => updateMusicSettings({ autoPlayNext: e.target.checked })}
                                            className="w-5 h-5 text-red-500 rounded focus:ring-2 focus:ring-red-200"
                                        />
                                    </label>

                                    <label className="flex items-center justify-between p-4 bg-white rounded-xl border-2 border-stone-200 cursor-pointer hover:border-stone-300 transition-colors">
                                        <div>
                                            <div className="text-sm font-medium text-stone-700">Loop Playlist</div>
                                            <div className="text-xs text-stone-500 mt-1">Repeat playlist when finished</div>
                                        </div>
                                        <input
                                            type="checkbox"
                                            checked={settings.music.loop}
                                            onChange={(e) => updateMusicSettings({ loop: e.target.checked })}
                                            className="w-5 h-5 text-red-500 rounded focus:ring-2 focus:ring-red-200"
                                        />
                                    </label>
                                </div>
                            </div>
                        </div>
                    )}
                </div>

                {/* Footer */}
                <div className="px-8 py-4 border-t border-stone-100 bg-white/50 backdrop-blur-md flex items-center justify-between">
                    <button
                        onClick={handleReset}
                        className="text-sm text-stone-500 hover:text-red-500 font-medium transition-colors cursor-pointer"
                    >
                        Reset to Default
                    </button>
                    <button
                        onClick={handleSave}
                        className="px-6 py-2 bg-red-500 text-white rounded-xl font-medium text-sm hover:bg-red-600 transition-colors shadow-sm cursor-pointer"
                    >
                        Save Changes
                    </button>
                </div>

                {/* Save Toast */}
                {showSaveToast && (
                    <div className="absolute top-4 right-4 bg-green-500 text-white px-4 py-2 rounded-lg shadow-lg animate-fade-in-up">
                        ✓ Settings saved!
                    </div>
                )}
            </div>
        </div>
    );
};

export default SettingsModal;
