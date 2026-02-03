import React, { createContext, useContext, useState, type ReactNode } from 'react';
import type { MediaFile, VisitLog, DashboardOverview, User, Comment } from '../types';
import { metaService } from '../services/metaService';
import { authService } from '../services/authService';

interface AdminContextType {
    files: MediaFile[];
    visitLogs: VisitLog[];
    dashboardStats: DashboardOverview | null;
    adminUsers: User[];
    allComments: Comment[];

    // Refresh Functions
    refreshAdminData: () => Promise<void>;
    refreshFiles: () => Promise<void>;
    refreshVisitLogs: () => Promise<void>;
    refreshDashboardOverview: () => Promise<void>;
    refreshAdminUsers: () => Promise<void>;
    refreshAllComments: () => Promise<void>;

    // File CRUD
    addFile: (file: MediaFile) => Promise<void>;
    deleteFile: (id: number) => Promise<void>;
}

const AdminContext = createContext<AdminContextType | undefined>(undefined);

export const AdminProvider: React.FC<{ children: ReactNode }> = ({ children }) => {
    const [files, setFiles] = useState<MediaFile[]>([]);
    const [visitLogs, setVisitLogs] = useState<VisitLog[]>([]);
    const [dashboardStats, setDashboardStats] = useState<DashboardOverview | null>(null);
    const [adminUsers, setAdminUsers] = useState<User[]>([]);
    const [allComments, setAllComments] = useState<Comment[]>([]);

    const refreshAdminData = async () => {
        try {
            const [fetchedFiles, fetchedLogs] = await Promise.all([
                metaService.getFiles(),
                metaService.getVisitLogs()
            ]);
            setFiles(fetchedFiles);
            setVisitLogs(fetchedLogs);
        } catch (err) {
            console.error("Failed to load admin data", err);
        }
    };

    const refreshFiles = async () => {
        try {
            const fetchedFiles = await metaService.getFiles();
            setFiles(fetchedFiles);
        } catch (err) {
            console.error("Failed to refresh files", err);
        }
    };

    const refreshVisitLogs = async () => {
        try {
            const fetchedLogs = await metaService.getVisitLogs();
            setVisitLogs(fetchedLogs);
        } catch (err) {
            console.error("Failed to refresh visit logs", err);
        }
    };

    const refreshDashboardOverview = async () => {
        try {
            const stats = await metaService.getDashboardOverview();
            setDashboardStats(stats);
        } catch (err) {
            console.error("Failed to refresh dashboard overview", err);
        }
    };

    const refreshAdminUsers = async () => {
        try {
            const users = await authService.getUsers();
            setAdminUsers(users);
        } catch (err) {
            console.error("Failed to refresh users", err);
        }
    };

    const refreshAllComments = async () => {
        try {
            const comments = await import('../services/commentService').then(m => m.commentService.getAllComments());
            setAllComments(comments);
        } catch (err) {
            console.error("Failed to refresh all comments", err);
        }
    };

    const addFile = async (file: MediaFile) => {
        try {
            const newFile = await metaService.addFile(file);
            setFiles(prev => [newFile, ...prev]);
        } catch (err) {
            console.error("Failed to add file", err);
            throw err;
        }
    };

    const deleteFile = async (id: number) => {
        try {
            await metaService.deleteFile(id);
            setFiles(prev => prev.filter(f => f.id !== id));
        } catch (err) {
            console.error("Failed to delete file", err);
            throw err;
        }
    };

    return (
        <AdminContext.Provider value={{
            files,
            visitLogs,
            dashboardStats,
            adminUsers,
            allComments,
            refreshAdminData,
            refreshFiles,
            refreshVisitLogs,
            refreshDashboardOverview,
            refreshAdminUsers,
            refreshAllComments,
            addFile,
            deleteFile
        }}>
            {children}
        </AdminContext.Provider>
    );
};

export const useAdmin = () => {
    const context = useContext(AdminContext);
    if (!context) {
        throw new Error('useAdmin must be used within an AdminProvider');
    }
    return context;
};
