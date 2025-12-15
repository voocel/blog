import React, { useState, useRef } from 'react';
import { IconUpload, IconCopy, IconTrash, IconX, IconChevronLeft, IconChevronRight } from '../Icons';
import type { MediaFile } from '../../types';
import { uploadImage } from '../../services/uploadService';

interface AdminFilesProps {
    files: MediaFile[];
    onAddFile: (file: MediaFile) => void;
    onDeleteFile: (id: string) => void;
    requestConfirm: (title: string, message: string, onConfirm: () => void) => void;
}

const AdminFiles: React.FC<AdminFilesProps> = ({ files, onAddFile, onDeleteFile, requestConfirm }) => {
    const [newFileUrl, setNewFileUrl] = useState('');
    const [lightboxIndex, setLightboxIndex] = useState<number | null>(null);
    const fileInputRef = useRef<HTMLInputElement>(null);

    const handleFileUpload = async (event: React.ChangeEvent<HTMLInputElement>) => {
        const file = event.target.files?.[0];
        if (file) {
            try {
                const result = await uploadImage(file);
                onAddFile({
                    id: `f-${Date.now()}`,
                    url: result.url,
                    name: result.filename,
                    type: 'image', // In real app, use result.type
                    date: new Date().toLocaleDateString()
                });
            } catch (error) {
                console.error("Upload failed:", error);
            }
        }
    };

    const handleAddUrl = () => {
        if (!newFileUrl) return;
        onAddFile({
            id: `f-${Date.now()}`,
            url: newFileUrl,
            name: 'New Upload',
            type: 'image',
            date: new Date().toLocaleDateString()
        });
        setNewFileUrl('');
    };

    const renderLightbox = () => {
        if (lightboxIndex === null) return null;
        const file = files[lightboxIndex];

        const handleNext = (e: React.MouseEvent) => {
            e.stopPropagation();
            setLightboxIndex((lightboxIndex + 1) % files.length);
        };

        const handlePrev = (e: React.MouseEvent) => {
            e.stopPropagation();
            setLightboxIndex((lightboxIndex - 1 + files.length) % files.length);
        };

        return (
            <div className="fixed inset-0 z-[200] bg-stone-900/95 backdrop-blur-md flex items-center justify-center animate-fade-in" onClick={() => setLightboxIndex(null)}>
                {/* Close Button */}
                <button
                    onClick={() => setLightboxIndex(null)}
                    className="absolute top-6 right-6 p-2 text-white/50 hover:text-white transition-colors cursor-pointer"
                >
                    <IconX className="w-8 h-8" />
                </button>

                {/* Navigation */}
                {files.length > 1 && (
                    <>
                        <button
                            onClick={handlePrev}
                            className="absolute left-6 p-4 text-white/50 hover:text-white transition-colors hover:bg-white/10 rounded-full cursor-pointer"
                        >
                            <IconChevronLeft className="w-8 h-8" />
                        </button>
                        <button
                            onClick={handleNext}
                            className="absolute right-6 p-4 text-white/50 hover:text-white transition-colors hover:bg-white/10 rounded-full cursor-pointer"
                        >
                            <IconChevronRight className="w-8 h-8" />
                        </button>
                    </>
                )}

                {/* Image */}
                <div className="max-w-[90vw] max-h-[90vh] relative" onClick={e => e.stopPropagation()}>
                    <img
                        src={file.url}
                        alt={file.name}
                        className="max-w-full max-h-[90vh] object-contain rounded-lg shadow-2xl"
                    />
                    <div className="absolute -bottom-12 left-0 w-full text-center">
                        <p className="text-white/80 font-serif text-lg">{file.name}</p>
                        <p className="text-white/40 text-xs uppercase tracking-widest">{lightboxIndex + 1} / {files.length}</p>
                    </div>
                </div>
            </div>
        );
    };

    return (
        <div className="p-10 h-full flex flex-col animate-fade-in text-ink w-full max-w-[1600px] mx-auto">
            <div className="flex justify-between items-end mb-10">
                <div>
                    <h2 className="text-4xl font-serif font-bold text-ink mb-2">Media Assets</h2>
                    <p className="text-stone-500">Library of uploaded images and documents.</p>
                </div>
                <div className="flex gap-3 bg-white p-2 rounded-xl border border-stone-200 shadow-sm">
                    <input
                        type="file"
                        ref={fileInputRef}
                        onChange={handleFileUpload}
                        className="hidden"
                        accept="image/*"
                    />
                    <button
                        onClick={() => fileInputRef.current?.click()}
                        className="p-2 bg-stone-100 hover:bg-stone-200 text-stone-600 rounded-lg border border-stone-200 transition-colors cursor-pointer"
                        title="Upload Local File"
                    >
                        <IconUpload className="w-5 h-5" />
                    </button>
                    <div className="w-px bg-stone-200 my-1"></div>
                    <input
                        value={newFileUrl}
                        onChange={e => setNewFileUrl(e.target.value)}
                        placeholder="Paste Image URL..."
                        className="w-64 bg-stone-50 rounded-lg px-4 py-2 text-sm focus:outline-none border border-transparent focus:border-stone-300 transition-colors"
                    />
                    <button onClick={handleAddUrl} className="bg-ink text-white px-6 py-2 rounded-lg text-sm font-bold uppercase tracking-wider hover:bg-stone-800 transition-colors cursor-pointer">Import</button>
                </div>
            </div>

            <div className="grid grid-cols-2 md:grid-cols-4 lg:grid-cols-5 xl:grid-cols-6 gap-6 overflow-y-auto pb-10">
                {files.map((file, index) => (
                    <div key={file.id} className="group relative aspect-square bg-white border border-stone-200 p-3 rounded-2xl shadow-sm hover:shadow-lg transition-all hover:-translate-y-1">
                        <div className="w-full h-full rounded-xl overflow-hidden relative bg-stone-100 cursor-pointer" onClick={() => setLightboxIndex(index)}>
                            <img src={file.url} alt={file.name} className="w-full h-full object-cover" />

                            {/* Overlay Actions */}
                            <div className="absolute inset-0 bg-stone-900/40 opacity-0 group-hover:opacity-100 transition-all duration-300 flex flex-col items-center justify-center gap-3 backdrop-blur-[2px]">
                                <div className="flex gap-2">
                                    <button
                                        onClick={(e) => { e.stopPropagation(); navigator.clipboard.writeText(file.url); }}
                                        className="p-2 bg-white/20 text-white rounded-full hover:bg-white hover:text-stone-900 transition-colors border border-white/30 backdrop-blur-md cursor-pointer"
                                        title="Copy URL"
                                    >
                                        <IconCopy className="w-4 h-4" />
                                    </button>
                                    <button
                                        onClick={(e) => { e.stopPropagation(); requestConfirm('Delete Asset', 'Are you sure you want to delete this file?', () => onDeleteFile(file.id)); }}
                                        className="p-2 bg-red-500/20 text-red-100 rounded-full hover:bg-red-500 hover:text-white transition-colors border border-red-500/30 backdrop-blur-md cursor-pointer"
                                        title="Delete"
                                    >
                                        <IconTrash className="w-4 h-4" />
                                    </button>
                                </div>
                                <span className="text-white text-[10px] uppercase tracking-widest font-bold opacity-80">Click to View</span>
                            </div>
                        </div>
                    </div>
                ))}
            </div>
            {renderLightbox()}
        </div>
    );
};

export default AdminFiles;
