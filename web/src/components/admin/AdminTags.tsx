import React, { useState } from 'react';
import { IconX } from '../Icons';
import type { Tag } from '../../types';

interface AdminTagsProps {
    tags: Tag[];
    onAddTag: (name: string) => void;
    onDeleteTag: (id: string) => void;
    requestConfirm: (title: string, message: string, onConfirm: () => void) => void;
}

const AdminTags: React.FC<AdminTagsProps> = ({ tags, onAddTag, onDeleteTag, requestConfirm }) => {
    const [newTagName, setNewTagName] = useState('');

    const handleAdd = () => {
        if (!newTagName) return;
        onAddTag(newTagName);
        setNewTagName('');
    };

    return (
        <div className="p-10 animate-fade-in text-ink max-w-[1200px] mx-auto w-full">
            <h2 className="text-4xl font-serif font-bold text-ink mb-2">Topics & Tags</h2>
            <p className="text-stone-500 mb-10">Micro-categorization for your posts.</p>

            <div className="bg-white border border-stone-200 rounded-2xl p-10 shadow-sm mb-10">
                <div className="flex gap-4 mb-10 max-w-xl border-b border-stone-100 pb-10">
                    <input
                        value={newTagName}
                        onChange={e => setNewTagName(e.target.value)}
                        placeholder="New Tag Name"
                        className="flex-1 bg-stone-50 border border-stone-200 rounded-xl px-5 py-3 text-ink focus:outline-none focus:border-teal-500"
                    />
                    <button onClick={handleAdd} className="bg-teal-600 text-white px-8 py-3 rounded-xl font-bold hover:bg-teal-700 transition-colors shadow-lg shadow-teal-100 cursor-pointer">Add Tag</button>
                </div>

                <div className="flex flex-wrap gap-3">
                    {tags.map(tag => (
                        <div key={tag.id} className="group flex items-center gap-2 bg-white border border-stone-200 px-4 py-2.5 rounded-full hover:border-teal-400 hover:shadow-sm transition-all cursor-default">
                            <span className="text-stone-400">#</span>
                            <span className="text-stone-700 font-medium">{tag.name}</span>
                            <button onClick={() => requestConfirm('Delete Tag', `Are you sure you want to delete #${tag.name}?`, () => onDeleteTag(tag.id))} className="text-stone-300 hover:text-red-500 ml-2 p-0.5 rounded-full hover:bg-red-50 cursor-pointer">
                                <IconX className="w-3 h-3" />
                            </button>
                        </div>
                    ))}
                </div>
            </div>
        </div>
    );
};

export default AdminTags;
