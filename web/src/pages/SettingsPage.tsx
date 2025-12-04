import React, { useState } from 'react';
import { useBlog } from '../context/BlogContext';
import { IconUser, IconMail, IconCamera, IconArrowLeft, IconCheck, IconSparkles, IconMoon, IconGlobe } from '../components/Icons';
import { uploadImage } from '../services/uploadService';

interface SettingsPageProps {
    onExit: () => void;
}

const SettingsPage: React.FC<SettingsPageProps> = ({ onExit }) => {
    const { user, updateUser } = useBlog();

    // Local state for form
    const [username, setUsername] = useState(user?.username || '');
    const [avatar, setAvatar] = useState(user?.avatar || '');
    const [bio, setBio] = useState('Digital explorer and creator.'); // Mock bio for now
    const [isSaving, setIsSaving] = useState(false);
    const [showSuccess, setShowSuccess] = useState(false);

    // Sync local state with user context when it loads
    React.useEffect(() => {
        if (user) {
            setUsername(user.username);
            setAvatar(user.avatar || '');
        }
    }, [user]);

    // ...

    // File Upload Ref
    const fileInputRef = React.useRef<HTMLInputElement>(null);
    const [isUploading, setIsUploading] = useState(false);

    const handleFileChange = async (event: React.ChangeEvent<HTMLInputElement>) => {
        const file = event.target.files?.[0];
        if (file) {
            try {
                setIsUploading(true);
                const result = await uploadImage(file);
                setAvatar(result.url);
            } catch (error) {
                console.error("Upload failed:", error);
                // Ideally show an error toast here
            } finally {
                setIsUploading(false);
            }
        }
    };

    const handleSave = () => {
        if (!user) return;

        setIsSaving(true);

        // Simulate API delay for effect
        setTimeout(() => {
            updateUser({
                ...user,
                username,
                avatar
            });
            setIsSaving(false);
            setShowSuccess(true);

            // Hide success message after 3s
            setTimeout(() => setShowSuccess(false), 3000);
        }, 800);
    };

    return (
        <div className="min-h-screen bg-[#FDFBF7] animate-fade-in relative overflow-hidden">
            {/* Background Decor */}
            <div className="absolute top-0 right-0 w-[500px] h-[500px] bg-gold-100/20 rounded-full blur-3xl -translate-y-1/2 translate-x-1/2 pointer-events-none"></div>
            <div className="absolute bottom-0 left-0 w-[600px] h-[600px] bg-stone-100/40 rounded-full blur-3xl translate-y-1/2 -translate-x-1/2 pointer-events-none"></div>

            {/* Header */}
            <div className="relative z-10 max-w-4xl mx-auto px-6 py-8 flex justify-between items-center">
                <button
                    onClick={onExit}
                    className="flex items-center gap-2 text-stone-400 hover:text-ink transition-colors group cursor-pointer"
                >
                    <div className="p-2 rounded-full bg-white border border-stone-200 group-hover:border-stone-300 transition-colors">
                        <IconArrowLeft className="w-4 h-4" />
                    </div>
                    <span className="text-xs uppercase tracking-widest font-medium">Return to Journal</span>
                </button>

                <h1 className="font-serif text-2xl font-bold text-ink flex items-center gap-2">
                    <IconSparkles className="w-5 h-5 text-gold-500" />
                    Personal Studio
                </h1>
            </div>

            {/* Main Content */}
            <div className="relative z-10 max-w-3xl mx-auto px-6 pb-20">
                <div className="bg-white/80 backdrop-blur-xl border border-white shadow-[0_20px_40px_-12px_rgba(0,0,0,0.05)] rounded-3xl p-8 md:p-12 animate-slide-up">

                    {/* Identity Section */}
                    <section className="mb-16">
                        <div className="flex items-center gap-4 mb-8">
                            <span className="w-1 h-6 bg-gold-500 rounded-full"></span>
                            <h2 className="text-xl font-serif font-bold text-ink">Identity</h2>
                        </div>

                        <div className="flex flex-col md:flex-row gap-12 items-start">
                            {/* Avatar */}
                            <div className="flex flex-col items-center gap-4">
                                <input
                                    type="file"
                                    ref={fileInputRef}
                                    onChange={handleFileChange}
                                    className="hidden"
                                    accept="image/*"
                                />
                                <div
                                    className="relative group cursor-pointer"
                                    onClick={() => fileInputRef.current?.click()}
                                >
                                    <div className="w-32 h-32 rounded-full overflow-hidden border-4 border-white shadow-lg bg-stone-100 relative">
                                        {avatar ? (
                                            <img src={avatar} alt="Avatar" className="w-full h-full object-cover" />
                                        ) : (
                                            <div className="w-full h-full flex items-center justify-center text-stone-300">
                                                <IconUser className="w-12 h-12" />
                                            </div>
                                        )}
                                        {/* Overlay */}
                                        <div className={`absolute inset-0 bg-black/20 flex items-center justify-center backdrop-blur-[1px] transition-opacity ${isUploading ? 'opacity-100' : 'opacity-0 group-hover:opacity-100'}`}>
                                            {isUploading ? (
                                                <div className="w-8 h-8 border-2 border-white/30 border-t-white rounded-full animate-spin"></div>
                                            ) : (
                                                <IconCamera className="w-8 h-8 text-white drop-shadow-md" />
                                            )}
                                        </div>
                                    </div>
                                    <div className="absolute bottom-0 right-0 p-2 bg-white rounded-full shadow-md border border-stone-100 text-gold-600">
                                        <IconSparkles className="w-4 h-4" />
                                    </div>
                                </div>
                                <p className="text-[10px] uppercase tracking-widest text-stone-400">Profile Image</p>
                            </div>

                            {/* Form Fields */}
                            <div className="flex-1 w-full space-y-6">
                                <div>
                                    <label className="block text-xs uppercase tracking-widest text-stone-500 mb-2 font-bold">Signature (Username)</label>
                                    <input
                                        type="text"
                                        value={username}
                                        onChange={(e) => setUsername(e.target.value)}
                                        className="w-full bg-stone-50 border border-stone-200 rounded-xl px-4 py-3 text-ink font-serif text-lg focus:outline-none focus:border-gold-500 focus:bg-white transition-all"
                                    />
                                </div>

                                <div>
                                    <label className="block text-xs uppercase tracking-widest text-stone-500 mb-2 font-bold">Bio</label>
                                    <textarea
                                        value={bio}
                                        onChange={(e) => setBio(e.target.value)}
                                        rows={3}
                                        className="w-full bg-stone-50 border border-stone-200 rounded-xl px-4 py-3 text-sm text-stone-600 focus:outline-none focus:border-gold-500 focus:bg-white transition-all resize-none"
                                        placeholder="Tell your story..."
                                    />
                                </div>

                                <div>
                                    <label className="block text-xs uppercase tracking-widest text-stone-500 mb-2 font-bold">Avatar URL / Upload</label>
                                    <div className="flex gap-2">
                                        <input
                                            type="text"
                                            value={avatar}
                                            onChange={(e) => setAvatar(e.target.value)}
                                            placeholder="https://..."
                                            className="flex-1 bg-stone-50 border border-stone-200 rounded-xl px-4 py-3 text-xs font-mono text-stone-500 focus:outline-none focus:border-gold-500 focus:bg-white transition-all"
                                        />
                                        <button
                                            onClick={() => fileInputRef.current?.click()}
                                            className="bg-stone-100 hover:bg-stone-200 text-stone-600 px-4 rounded-xl border border-stone-200 transition-colors"
                                            title="Upload Image"
                                        >
                                            <IconCamera className="w-4 h-4" />
                                        </button>
                                    </div>
                                </div>
                            </div>
                        </div>
                    </section>

                    <div className="w-full h-px bg-stone-100 mb-16"></div>

                    {/* Aesthetics Section */}
                    <section className="mb-16">
                        <div className="flex items-center gap-4 mb-8">
                            <span className="w-1 h-6 bg-stone-300 rounded-full"></span>
                            <h2 className="text-xl font-serif font-bold text-ink">Aesthetics</h2>
                        </div>

                        <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
                            <div className="p-6 rounded-2xl border border-stone-200 bg-stone-50/50 flex items-center justify-between group hover:border-gold-200 transition-colors cursor-pointer">
                                <div className="flex items-center gap-4">
                                    <div className="p-3 bg-white rounded-xl shadow-sm text-stone-600">
                                        <IconMoon className="w-5 h-5" />
                                    </div>
                                    <div>
                                        <h3 className="font-bold text-ink">Theme</h3>
                                        <p className="text-xs text-stone-500">Daylight (Default)</p>
                                    </div>
                                </div>
                                <div className="w-12 h-6 bg-stone-200 rounded-full relative">
                                    <div className="absolute left-1 top-1 w-4 h-4 bg-white rounded-full shadow-sm"></div>
                                </div>
                            </div>

                            <div className="p-6 rounded-2xl border border-stone-200 bg-stone-50/50 flex items-center justify-between group hover:border-gold-200 transition-colors cursor-pointer">
                                <div className="flex items-center gap-4">
                                    <div className="p-3 bg-white rounded-xl shadow-sm text-stone-600">
                                        <IconGlobe className="w-5 h-5" />
                                    </div>
                                    <div>
                                        <h3 className="font-bold text-ink">Language</h3>
                                        <p className="text-xs text-stone-500">English (US)</p>
                                    </div>
                                </div>
                                <span className="text-xs font-bold text-stone-400 uppercase">Change</span>
                            </div>
                        </div>
                    </section>

                    {/* Footer Actions */}
                    <div className="flex items-center justify-between pt-8 border-t border-stone-100">
                        <div className="flex items-center gap-2 text-stone-400 text-sm">
                            <IconMail className="w-4 h-4" />
                            <span>{user?.email}</span>
                        </div>

                        <button
                            onClick={handleSave}
                            disabled={isSaving}
                            className={`
                                relative overflow-hidden px-8 py-3 rounded-xl font-bold tracking-wide transition-all duration-300 shadow-lg cursor-pointer
                                ${showSuccess
                                    ? 'bg-emerald-500 text-white shadow-emerald-200'
                                    : 'bg-ink text-white hover:bg-gold-600 shadow-stone-300'
                                }
                            `}
                        >
                            <div className="relative z-10 flex items-center gap-2">
                                {isSaving ? (
                                    <>
                                        <div className="w-4 h-4 border-2 border-white/30 border-t-white rounded-full animate-spin"></div>
                                        <span>Saving...</span>
                                    </>
                                ) : showSuccess ? (
                                    <>
                                        <IconCheck className="w-4 h-4" />
                                        <span>Saved</span>
                                    </>
                                ) : (
                                    <span>Save Changes</span>
                                )}
                            </div>
                        </button>
                    </div>

                </div>
            </div>
        </div>
    );
};

export default SettingsPage;
