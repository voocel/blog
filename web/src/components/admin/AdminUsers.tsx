import React, { useState } from 'react';
import { useBlog } from '../../context/BlogContext';
import { IconUserCircle, IconSearch, IconShield, IconLock, IconCheck, IconGoogle, IconGithub, IconMail } from '../Icons';
import { authService } from '../../services/authService';
import type { User } from '../../types';

interface AdminUsersProps {
    users: User[];
    requestConfirm: (title: string, message: string, onConfirm: () => void, options?: { confirmText?: string; isDestructive?: boolean }) => void;
}

import { useToast } from '../../components/Toast';

// ...

const AdminUsers: React.FC<AdminUsersProps> = ({ users, requestConfirm }) => {
    const { refreshAdminUsers } = useBlog();
    const { showToast } = useToast();
    const [searchTerm, setSearchTerm] = useState('');
    const [processingId, setProcessingId] = useState<string | null>(null);

    const filteredUsers = users.filter(user =>
        user.username.toLowerCase().includes(searchTerm.toLowerCase()) ||
        user.email.toLowerCase().includes(searchTerm.toLowerCase())
    );

    const handleToggleStatus = (user: User) => {
        if (!user.id) return;
        // Don't ban admins themselves for safety
        if (user.role === 'admin') {
            showToast("Cannot ban administrators", "error");
            return;
        }

        const newStatus = user.status === 'banned' ? 'active' : 'banned';
        const confirmMsg = newStatus === 'banned'
            ? `Are you sure you want to BAN ${user.username}? They will no longer be able to log in.`
            : `Unban ${user.username}?`;

        requestConfirm(
            newStatus === 'banned' ? 'Ban User?' : 'Unban User?',
            confirmMsg,
            async () => {
                setProcessingId(user.id);
                try {
                    await authService.updateUserStatus(user.id!, newStatus);
                    await refreshAdminUsers();
                    showToast(`User ${user.username} has been ${newStatus === 'banned' ? 'banned' : 'unbanned'}`, "success");
                } catch (error) {
                    console.error("Failed to update status", error);
                    showToast("Failed to update user status", "error");
                } finally {
                    setProcessingId(null);
                }
            },
            {
                confirmText: newStatus === 'banned' ? 'Ban User' : 'Unban',
                isDestructive: newStatus === 'banned'
            }
        );
    };

    return (
        <div className="p-10 h-full flex flex-col animate-fade-in text-ink max-w-[1600px] mx-auto w-full">
            <div className="flex justify-between items-end mb-10">
                <div>
                    <h2 className="text-4xl font-serif font-bold text-ink mb-2">Users</h2>
                    <p className="text-stone-500">Manage registered members of the sanctuary.</p>
                </div>
                <div className="relative">
                    <IconSearch className="absolute left-4 top-1/2 -translate-y-1/2 w-5 h-5 text-stone-400" />
                    <input
                        type="text"
                        placeholder="Search users..."
                        value={searchTerm}
                        onChange={(e) => setSearchTerm(e.target.value)}
                        className="pl-12 pr-4 py-3 bg-white border border-stone-200 rounded-xl shadow-sm focus:outline-none focus:border-gold-400 focus:ring-1 focus:ring-gold-100 transition-all w-64"
                    />
                </div>
            </div>

            <div className="bg-white border border-stone-200 rounded-2xl shadow-sm overflow-hidden flex-1 flex flex-col">
                <div className="overflow-x-auto custom-scrollbar flex-1">
                    <table className="w-full text-left border-collapse">
                        <thead className="bg-stone-50 border-b border-stone-100 sticky top-0 z-10">
                            <tr>
                                <th className="py-4 px-6 text-xs font-bold uppercase tracking-wider text-stone-400">User</th>
                                <th className="py-4 px-6 text-xs font-bold uppercase tracking-wider text-stone-400">Provider</th>
                                <th className="py-4 px-6 text-xs font-bold uppercase tracking-wider text-stone-400">Role</th>
                                <th className="py-4 px-6 text-xs font-bold uppercase tracking-wider text-stone-400">Status</th>
                                <th className="py-4 px-6 text-xs font-bold uppercase tracking-wider text-stone-400 text-right">Actions</th>
                            </tr>
                        </thead>
                        <tbody className="divide-y divide-stone-50">
                            {filteredUsers.length > 0 ? (
                                filteredUsers.map((user) => (
                                    <tr key={user.email} className={`transition-colors group ${user.status === 'banned' ? 'bg-red-50/30' : 'hover:bg-stone-50/50'}`}>
                                        <td className="py-4 px-6">
                                            <div className="flex items-center gap-4">
                                                <div className={`w-10 h-10 rounded-full border overflow-hidden shrink-0 ${user.status === 'banned' ? 'border-red-200 grayscale' : 'bg-stone-100 border-stone-200'}`}>
                                                    {user.avatar ? (
                                                        <img src={user.avatar} alt={user.username} className="w-full h-full object-cover" />
                                                    ) : (
                                                        <div className="w-full h-full flex items-center justify-center text-stone-400">
                                                            <IconUserCircle className="w-6 h-6" />
                                                        </div>
                                                    )}
                                                </div>
                                                <div className={user.status === 'banned' ? 'opacity-50' : ''}>
                                                    <div className="font-bold text-ink">{user.username}</div>
                                                    <div className="text-xs text-stone-400">{user.email}</div>
                                                </div>
                                            </div>
                                        </td>
                                        <td className="py-4 px-6">
                                            <div className="flex items-center gap-2 text-stone-500" title={`Signed up with ${user.provider || 'email'}`}>
                                                {user.provider === 'google' && <IconGoogle className="w-5 h-5" />}
                                                {user.provider === 'github' && <IconGithub className="w-5 h-5" />}
                                                {(!user.provider || user.provider === 'email') && <IconMail className="w-5 h-5 text-stone-300" />}
                                                <span className="text-xs capitalize hidden xl:inline-block text-stone-400">{user.provider || 'email'}</span>
                                            </div>
                                        </td>
                                        <td className="py-4 px-6">
                                            <span className={`inline-flex items-center gap-1.5 px-3 py-1 rounded-full text-xs font-bold uppercase tracking-wider ${user.role === 'admin'
                                                ? 'bg-indigo-50 text-indigo-600 border border-indigo-100'
                                                : 'bg-stone-100 text-stone-500 border border-stone-200'
                                                }`}>
                                                {user.role === 'admin' && <IconShield className="w-3 h-3" />}
                                                {user.role}
                                            </span>
                                        </td>
                                        <td className="py-4 px-6">
                                            <span className={`inline-flex items-center gap-1.5 px-3 py-1 rounded-full text-xs font-bold uppercase tracking-wider ${user.status === 'banned'
                                                ? 'bg-red-100 text-red-600 border border-red-200'
                                                : 'bg-emerald-50 text-emerald-600 border border-emerald-100'
                                                }`}>
                                                {user.status === 'banned' ? (
                                                    <>
                                                        <IconLock className="w-3 h-3" /> Banned
                                                    </>
                                                ) : (
                                                    <>
                                                        <IconCheck className="w-3 h-3" /> Active
                                                    </>
                                                )}
                                            </span>
                                        </td>
                                        <td className="py-4 px-6 text-right">
                                            {user.role !== 'admin' && (
                                                <button
                                                    onClick={() => handleToggleStatus(user)}
                                                    disabled={processingId === user.id}
                                                    className={`text-xs font-bold uppercase tracking-wider px-3 py-1.5 rounded-lg border transition-all cursor-pointer ${user.status === 'banned'
                                                        ? 'border-stone-200 text-stone-500 hover:border-stone-400 hover:text-stone-700 bg-white'
                                                        : 'border-red-200 text-red-500 hover:bg-red-50 hover:border-red-300 bg-white'
                                                        } ${processingId === user.id ? 'opacity-50 cursor-not-allowed' : ''}`}
                                                >
                                                    {user.status === 'banned' ? 'Unban' : 'Ban'}
                                                </button>
                                            )}
                                        </td>
                                    </tr>
                                ))
                            ) : (
                                <tr>
                                    <td colSpan={4} className="py-20 text-center text-stone-400 italic">
                                        No users found matching your search.
                                    </td>
                                </tr>
                            )}
                        </tbody>
                    </table>
                </div>
                <div className="p-4 border-t border-stone-100 bg-stone-50 flex justify-between items-center text-xs text-stone-400 uppercase tracking-widest font-bold">
                    <span>Total Users: {users.length}</span>
                </div>
            </div>
        </div>
    );
};

export default AdminUsers;
