
import React, { useState } from 'react';
import { IconGithub, IconMail } from './Icons';

interface SettingsModalProps {
    isOpen: boolean;
    onClose: () => void;
}

const SettingsModal: React.FC<SettingsModalProps> = ({ isOpen, onClose }) => {
    if (!isOpen) return null;

    const [activeTab, setActiveTab] = useState('site');

    return (
        <div className="fixed inset-0 z-50 flex items-center justify-center p-4 bg-black/20 backdrop-blur-sm">
            <div className="bg-[#fdfaf6] w-full max-w-4xl h-[80vh] rounded-3xl shadow-2xl overflow-hidden flex flex-col animate-fade-in-up">

                {/* Header */}
                <div className="px-8 py-6 border-b border-stone-100 flex items-center justify-between bg-white/50 backdrop-blur-md">
                    <div className="flex gap-8 text-sm font-bold text-stone-400">
                        <button
                            onClick={() => setActiveTab('site')}
                            className={`${activeTab === 'site' ? 'text-red-500 relative after:absolute after:-bottom-6 after:left-0 after:w-full after:h-0.5 after:bg-red-500' : 'hover:text-stone-600'} transition-colors`}
                        >
                            网站设置
                        </button>
                        <button className="hover:text-stone-600 cursor-not-allowed opacity-50">色彩配置</button>
                        <button className="hover:text-stone-600 cursor-not-allowed opacity-50">首页布局</button>
                        <button className="hover:text-stone-600 cursor-not-allowed opacity-50">预览</button>
                    </div>

                    <div className="flex items-center gap-4">
                        <button className="text-stone-400 hover:text-stone-600 text-sm font-medium" onClick={onClose}>取消</button>
                        <button className="bg-red-500 text-white px-4 py-1.5 rounded-lg text-sm font-bold shadow-red-200 shadow-md hover:bg-red-600 transition-colors">
                            导入密钥
                        </button>
                    </div>
                </div>

                {/* Content */}
                <div className="flex-1 overflow-y-auto p-8 bg-stone-50/30">
                    <div className="grid grid-cols-2 gap-12 max-w-4xl mx-auto">
                        {/* Left Column */}
                        <div className="space-y-8">
                            <div>
                                <h3 className="text-xs font-bold text-stone-400 mb-4">Favicon</h3>
                                <div className="w-20 h-20 bg-stone-100 rounded-2xl flex items-center justify-center border-2 border-dashed border-stone-200 cursor-pointer hover:border-red-300 hover:bg-red-50 transition-colors">
                                    <span className="text-2xl">🍬</span>
                                </div>
                            </div>

                            <div>
                                <h3 className="text-xs font-bold text-stone-400 mb-3">站点标题</h3>
                                <input type="text" defaultValue="YYsuni" className="w-full bg-stone-200/50 border-none rounded-lg px-4 py-3 text-stone-700 font-medium focus:ring-2 focus:ring-red-200 outline-none" />
                            </div>

                            <div>
                                <h3 className="text-xs font-bold text-stone-400 mb-3">站点描述</h3>
                                <textarea
                                    defaultValue="YYsuni 的个人博客，记录前端开发、探索、笔记。"
                                    className="w-full bg-stone-200/50 border-none rounded-lg px-4 py-3 text-stone-700 font-medium h-32 resize-none focus:ring-2 focus:ring-red-200 outline-none"
                                />
                            </div>

                            <div>
                                <h3 className="text-xs font-bold text-stone-400 mb-3">备案信息</h3>
                                <div className="flex gap-4">
                                    <input type="text" placeholder="备案号" className="flex-1 bg-stone-200/50 border-none rounded-lg px-4 py-3 text-sm focus:ring-2 focus:ring-red-200 outline-none" />
                                    <input type="text" placeholder="备案链接" className="flex-1 bg-stone-200/50 border-none rounded-lg px-4 py-3 text-sm focus:ring-2 focus:ring-red-200 outline-none" />
                                </div>
                            </div>
                        </div>

                        {/* Right Column */}
                        <div className="space-y-8">
                            <div>
                                <h3 className="text-xs font-bold text-stone-400 mb-4">Avatar</h3>
                                <div className="w-20 h-20 rounded-full overflow-hidden border-4 border-white shadow-sm">
                                    <img src="https://images.unsplash.com/photo-1534528741775-53994a69daeb?auto=format&fit=crop&w=200&q=80" className="w-full h-full object-cover" />
                                </div>
                            </div>

                            <div>
                                <h3 className="text-xs font-bold text-stone-400 mb-3">用户名</h3>
                                <input type="text" defaultValue="Suni" className="w-full bg-stone-200/50 border-none rounded-lg px-4 py-3 text-stone-700 font-medium focus:ring-2 focus:ring-red-200 outline-none" />
                            </div>

                            <div>
                                <h3 className="text-xs font-bold text-stone-400 mb-3">社交按钮</h3>
                                <div className="space-y-3">
                                    {/* Mail */}
                                    <div className="flex items-center gap-2 bg-stone-200/50 p-2 rounded-lg">
                                        <div className="w-8 h-8 bg-white rounded flex items-center justify-center text-stone-400"><IconMail className="w-4 h-4" /></div>
                                        <input type="text" defaultValue="yysuni1001@gmail.com" className="flex-1 bg-transparent border-none text-sm text-stone-600 focus:outline-none" />
                                        <button className="text-xs text-red-400 hover:text-red-500 px-2 font-medium">删除</button>
                                    </div>

                                    {/* Juejin/Twitter */}
                                    <div className="flex items-center gap-2 bg-stone-200/50 p-2 rounded-lg">
                                        <div className="w-8 h-8 bg-white rounded flex items-center justify-center text-stone-400"><span className="text-xs font-bold">X</span></div>
                                        <input type="text" defaultValue="https://twitter.com/voocel" className="flex-1 bg-transparent border-none text-sm text-stone-600 focus:outline-none" />
                                        <button className="text-xs text-red-400 hover:text-red-500 px-2 font-medium">删除</button>
                                    </div>

                                    {/* Github */}
                                    <div className="flex items-center gap-2 bg-stone-200/50 p-2 rounded-lg">
                                        <div className="w-8 h-8 bg-white rounded flex items-center justify-center text-stone-400"><IconGithub className="w-4 h-4" /></div>
                                        <input type="text" defaultValue="https://github.com/voocel" className="flex-1 bg-transparent border-none text-sm text-stone-600 focus:outline-none" />
                                        <button className="text-xs text-red-400 hover:text-red-500 px-2 font-medium">删除</button>
                                    </div>

                                    <button className="w-full py-2 border border-dashed border-stone-300 rounded-lg text-stone-400 text-sm hover:border-stone-400 hover:text-stone-500 transition-colors">
                                        + 添加按钮
                                    </button>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    );
};

export default SettingsModal;
