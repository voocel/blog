import React, { useState } from 'react';
import { IconTrash } from '../Icons';
import type { Category } from '../../types';

interface AdminCategoriesProps {
    categories: Category[];
    onAddCategory: (name: string) => void;
    onDeleteCategory: (id: number) => void;
    requestConfirm: (title: string, message: string, onConfirm: () => void) => void;
}

const AdminCategories: React.FC<AdminCategoriesProps> = ({ categories, onAddCategory, onDeleteCategory, requestConfirm }) => {
    const [newCatName, setNewCatName] = useState('');

    const handleAdd = () => {
        if (!newCatName) return;
        onAddCategory(newCatName);
        setNewCatName('');
    };

    return (
        <div className="p-10 animate-fade-in text-ink max-w-[1200px] mx-auto w-full">
            <h2 className="text-4xl font-serif font-bold text-ink mb-2">Taxonomy</h2>
            <p className="text-stone-500 mb-10">Organize your content structure.</p>

            <div className="grid grid-cols-1 md:grid-cols-3 gap-12">
                <div className="md:col-span-1">
                    <div className="bg-white border border-stone-200 p-8 rounded-2xl shadow-sm sticky top-10">
                        <h3 className="font-serif font-bold text-xl mb-6">Add Category</h3>
                        <div className="space-y-4">
                            <div>
                                <label className="block text-xs uppercase tracking-widest text-stone-400 mb-2">Name</label>
                                <input
                                    value={newCatName}
                                    onChange={e => setNewCatName(e.target.value)}
                                    placeholder="e.g., Philosophy"
                                    className="w-full bg-stone-50 border border-stone-200 rounded-lg px-4 py-3 text-ink focus:outline-none focus:border-gold-500 transition-colors"
                                />
                            </div>
                            <div className="p-4 bg-stone-50 rounded-lg border border-stone-100">
                                <p className="text-xs text-stone-500 leading-relaxed">
                                    <span className="font-bold">Tip:</span> Use broad, high-level topics for categories. Use tags for specific details.
                                </p>
                            </div>
                            <button
                                onClick={handleAdd}
                                disabled={!newCatName}
                                className="w-full bg-emerald-600 text-white py-3.5 rounded-xl font-bold tracking-wide hover:bg-emerald-700 transition-colors disabled:opacity-50 shadow-emerald-100 shadow-lg cursor-pointer"
                            >
                                Create Category
                            </button>
                        </div>
                    </div>
                </div>

                <div className="md:col-span-2 space-y-3">
                    {categories.map(cat => (
                        <div key={cat.id} className="group flex justify-between items-center bg-white border border-stone-200 p-6 rounded-xl hover:border-emerald-400 hover:shadow-md transition-all">
                            <div className="flex items-center gap-4">
                                <div className="w-10 h-10 rounded-full bg-emerald-50 flex items-center justify-center text-emerald-600 font-serif font-bold text-lg border border-emerald-100">
                                    {cat.name.charAt(0)}
                                </div>
                                <div>
                                    <span className="text-lg font-serif font-bold text-ink block group-hover:text-emerald-700 transition-colors">{cat.name}</span>
                                    <span className="text-xs text-stone-400 font-mono tracking-tight">{cat.slug} • {cat.count} posts</span>
                                </div>
                            </div>
                            <button onClick={() => requestConfirm('Delete Category', `Are you sure you want to delete "${cat.name}"?`, () => onDeleteCategory(cat.id))} className="p-2 text-stone-300 hover:text-red-500 hover:bg-red-50 rounded-lg opacity-0 group-hover:opacity-100 transition-all cursor-pointer">
                                <IconTrash className="w-5 h-5" />
                            </button>
                        </div>
                    ))}
                </div>
            </div>
        </div>
    );
};

export default AdminCategories;
