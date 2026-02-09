import React, { useState } from 'react';
import { useSettings } from '@/context/SettingsContext';
import { useTranslation } from '@/hooks/useTranslation';
import type { ThemeMode } from '@/config/settings';
import { type Locale, localeNames } from '@/locales';

interface SettingsModalProps {
    isOpen: boolean;
    onClose: () => void;
}

const SettingsModal: React.FC<SettingsModalProps> = ({ isOpen, onClose }) => {
    const { settings, updateTheme, updateAnimations, updateMusicSettings, updateLocale, resetSettings } = useSettings();
    const { t } = useTranslation();
    const [activeTab, setActiveTab] = useState<'appearance' | 'music' | 'language'>('appearance');
    const [showSaveToast, setShowSaveToast] = useState(false);

    if (!isOpen) return null;

    const handleSave = () => {
        setShowSaveToast(true);
        setTimeout(() => setShowSaveToast(false), 2000);
    };

    const handleReset = () => {
        if (confirm(t.settings.resetConfirm)) {
            resetSettings();
            setShowSaveToast(true);
            setTimeout(() => setShowSaveToast(false), 2000);
        }
    };

    return (
        <div className="fixed inset-0 z-50 flex items-center justify-center p-4 bg-[var(--color-overlay)] backdrop-blur-sm">
            <div className="bg-[var(--color-base)] w-full max-w-2xl max-h-[85vh] rounded-3xl shadow-2xl overflow-hidden flex flex-col animate-fade-in-up">

                {/* Header */}
                <div className="px-8 py-6 border-b border-[var(--color-border-subtle)] flex items-center justify-between bg-[var(--color-elevated)] backdrop-blur-md">
                    <div className="flex gap-8 text-sm font-bold text-[var(--color-text-muted)]">
                        <button
                            onClick={() => setActiveTab('appearance')}
                            className={`${activeTab === 'appearance' ? 'text-red-500 relative after:absolute after:-bottom-6 after:left-0 after:w-full after:h-0.5 after:bg-red-500' : 'hover:text-[var(--color-text-secondary)]'} transition-colors cursor-pointer`}
                        >
                            {t.settings.appearance}
                        </button>
                        <button
                            onClick={() => setActiveTab('music')}
                            className={`${activeTab === 'music' ? 'text-red-500 relative after:absolute after:-bottom-6 after:left-0 after:w-full after:h-0.5 after:bg-red-500' : 'hover:text-[var(--color-text-secondary)]'} transition-colors cursor-pointer`}
                        >
                            {t.settings.music}
                        </button>
                        <button
                            onClick={() => setActiveTab('language')}
                            className={`${activeTab === 'language' ? 'text-red-500 relative after:absolute after:-bottom-6 after:left-0 after:w-full after:h-0.5 after:bg-red-500' : 'hover:text-[var(--color-text-secondary)]'} transition-colors cursor-pointer`}
                        >
                            {t.settings.language}
                        </button>
                    </div>

                    <div className="flex items-center gap-4">
                        <button
                            className="text-[var(--color-text-muted)] hover:text-[var(--color-text-secondary)] text-sm font-medium cursor-pointer"
                            onClick={onClose}
                        >
                            {t.settings.close}
                        </button>
                    </div>
                </div>

                {/* Content */}
                <div className="flex-1 overflow-y-auto p-8 bg-[var(--color-surface-alt)]/30">
                    {activeTab === 'appearance' && (
                        <div className="max-w-xl mx-auto space-y-8">
                            {/* Theme Mode */}
                            <div>
                                <h3 className="text-sm font-bold text-ink mb-4">{t.appearance.themeMode}</h3>
                                <div className="grid grid-cols-3 gap-3">
                                    {(['light', 'dark', 'auto'] as ThemeMode[]).map((theme) => (
                                        <button
                                            key={theme}
                                            onClick={() => updateTheme(theme)}
                                            className={`p-4 rounded-xl border-2 transition-all cursor-pointer ${
                                                settings.appearance.theme === theme
                                                    ? 'border-red-400 bg-red-50 dark:bg-red-950 dark:border-red-600'
                                                    : 'border-[var(--color-border)] bg-[var(--color-surface)] hover:border-stone-300'
                                            }`}
                                        >
                                            <div className="text-2xl mb-2">
                                                {theme === 'light' && '☀️'}
                                                {theme === 'dark' && '🌙'}
                                                {theme === 'auto' && '🌓'}
                                            </div>
                                            <div className="text-sm font-medium text-ink capitalize">
                                                {t.appearance[theme]}
                                            </div>
                                        </button>
                                    ))}
                                </div>
                                <p className="text-xs text-[var(--color-text-secondary)] mt-2">
                                    {settings.appearance.theme === 'auto'
                                        ? t.appearance.themeFollows
                                        : t.appearance.usingTheme.replace('{theme}', t.appearance[settings.appearance.theme])}
                                </p>
                            </div>

                            {/* Animations */}
                            <div>
                                <h3 className="text-sm font-bold text-ink mb-4">{t.appearance.animations}</h3>
                                <label className="flex items-center justify-between p-4 bg-[var(--color-surface)] rounded-xl border-2 border-[var(--color-border)] cursor-pointer hover:border-stone-300 transition-colors">
                                    <div>
                                        <div className="text-sm font-medium text-ink">{t.appearance.enableAnimations}</div>
                                        <div className="text-xs text-[var(--color-text-secondary)] mt-1">{t.appearance.animationsDesc}</div>
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
                                <h3 className="text-sm font-bold text-ink mb-4">{t.music.display}</h3>
                                <label className="flex items-center justify-between p-4 bg-[var(--color-surface)] rounded-xl border-2 border-[var(--color-border)] cursor-pointer hover:border-stone-300 transition-colors">
                                    <div>
                                        <div className="text-sm font-medium text-ink">{t.music.showPlayer}</div>
                                        <div className="text-xs text-[var(--color-text-secondary)] mt-1">{t.music.showPlayerDesc}</div>
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
                                <h3 className="text-sm font-bold text-ink mb-4">{t.music.defaultVolume}</h3>
                                <div className="p-4 bg-[var(--color-surface)] rounded-xl border-2 border-[var(--color-border)]">
                                    <div className="flex items-center justify-between mb-3">
                                        <span className="text-sm text-[var(--color-text-secondary)]">{t.music.volume}</span>
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
                                        className="w-full h-2 bg-[var(--color-muted)] rounded-lg appearance-none cursor-pointer accent-red-500"
                                    />
                                </div>
                            </div>

                            {/* Playback Options */}
                            <div>
                                <h3 className="text-sm font-bold text-ink mb-4">{t.music.playback}</h3>
                                <div className="space-y-3">
                                    <label className="flex items-center justify-between p-4 bg-[var(--color-surface)] rounded-xl border-2 border-[var(--color-border)] cursor-pointer hover:border-stone-300 transition-colors">
                                        <div>
                                            <div className="text-sm font-medium text-ink">{t.music.autoPlayNext}</div>
                                            <div className="text-xs text-[var(--color-text-secondary)] mt-1">{t.music.autoPlayNextDesc}</div>
                                        </div>
                                        <input
                                            type="checkbox"
                                            checked={settings.music.autoPlayNext}
                                            onChange={(e) => updateMusicSettings({ autoPlayNext: e.target.checked })}
                                            className="w-5 h-5 text-red-500 rounded focus:ring-2 focus:ring-red-200"
                                        />
                                    </label>

                                    <label className="flex items-center justify-between p-4 bg-[var(--color-surface)] rounded-xl border-2 border-[var(--color-border)] cursor-pointer hover:border-stone-300 transition-colors">
                                        <div>
                                            <div className="text-sm font-medium text-ink">{t.music.loopPlaylist}</div>
                                            <div className="text-xs text-[var(--color-text-secondary)] mt-1">{t.music.loopPlaylistDesc}</div>
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

                    {activeTab === 'language' && (
                        <div className="max-w-xl mx-auto space-y-8">
                            {/* Language Selection */}
                            <div>
                                <h3 className="text-sm font-bold text-ink mb-4">{t.languageSettings.selectLanguage}</h3>
                                <div className="grid grid-cols-2 gap-3">
                                    {(['en', 'zh'] as Locale[]).map((locale) => (
                                        <button
                                            key={locale}
                                            onClick={() => updateLocale(locale)}
                                            className={`p-4 rounded-xl border-2 transition-all cursor-pointer ${
                                                settings.language.locale === locale
                                                    ? 'border-red-400 bg-red-50 dark:bg-red-950 dark:border-red-600'
                                                    : 'border-[var(--color-border)] bg-[var(--color-surface)] hover:border-stone-300'
                                            }`}
                                        >
                                            <div className="text-2xl mb-2">
                                                {locale === 'en' && '🇺🇸'}
                                                {locale === 'zh' && '🇨🇳'}
                                            </div>
                                            <div className="text-sm font-medium text-ink">
                                                {localeNames[locale]}
                                            </div>
                                        </button>
                                    ))}
                                </div>
                                <p className="text-xs text-[var(--color-text-secondary)] mt-2">
                                    {t.languageSettings.languageDesc}
                                </p>
                            </div>
                        </div>
                    )}
                </div>

                {/* Footer */}
                <div className="px-8 py-4 border-t border-[var(--color-border-subtle)] bg-[var(--color-elevated)] backdrop-blur-md flex items-center justify-between">
                    <button
                        onClick={handleReset}
                        className="text-sm text-[var(--color-text-secondary)] hover:text-red-500 font-medium transition-colors cursor-pointer"
                    >
                        {t.settings.reset}
                    </button>
                    <button
                        onClick={handleSave}
                        className="px-6 py-2 bg-red-500 text-white rounded-xl font-medium text-sm hover:bg-red-600 transition-colors shadow-sm cursor-pointer"
                    >
                        {t.settings.save}
                    </button>
                </div>

                {/* Save Toast */}
                {showSaveToast && (
                    <div className="absolute top-4 right-4 bg-green-500 text-white px-4 py-2 rounded-lg shadow-lg animate-fade-in-up">
                        ✓ {t.settings.saved}
                    </div>
                )}
            </div>
        </div>
    );
};

export default SettingsModal;
