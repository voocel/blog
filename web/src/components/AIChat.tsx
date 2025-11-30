
import React, { useState, useRef, useEffect } from 'react';
import MDEditor from '@uiw/react-md-editor';
import { streamChatResponse } from '../services/geminiService';
import type { ChatMessage } from '../types';
import { IconX, IconSend, IconSparkles } from './Icons';
import { useBlog } from '../context/BlogContext';

const AIChat: React.FC = () => {
  const { posts } = useBlog();
  const [isOpen, setIsOpen] = useState(false);
  const [input, setInput] = useState('');
  const [messages, setMessages] = useState<ChatMessage[]>([
    { role: 'model', text: 'Greetings. I am Aether, the digital curator of the Voocel journal. How may I assist your exploration?' }
  ]);
  const [isTyping, setIsTyping] = useState(false);
  const messagesEndRef = useRef<HTMLDivElement>(null);
  const chatContainerRef = useRef<HTMLDivElement>(null);

  // Artistic Avatar URL - Beautiful/Ethereal style
  const AETHER_AVATAR_URL = "https://images.unsplash.com/photo-1524638431109-93d95c968f03?q=80&w=200&auto=format&fit=crop";

  const scrollToBottom = () => {
    messagesEndRef.current?.scrollIntoView({ behavior: 'smooth' });
  };

  useEffect(() => {
    scrollToBottom();
  }, [messages, isOpen]);

  // Click outside to close
  useEffect(() => {
    const handleClickOutside = (event: MouseEvent) => {
      if (
        isOpen &&
        chatContainerRef.current &&
        !chatContainerRef.current.contains(event.target as Node) &&
        !(event.target as Element).closest('button') // Prevent closing when clicking the toggle button itself
      ) {
        setIsOpen(false);
      }
    };

    document.addEventListener('mousedown', handleClickOutside);
    return () => {
      document.removeEventListener('mousedown', handleClickOutside);
    };
  }, [isOpen]);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    if (!input.trim() || isTyping) return;

    const userText = input.trim();
    setInput('');
    const newMessages = [...messages, { role: 'user', text: userText } as ChatMessage];
    setMessages(newMessages);
    setIsTyping(true);

    // Prepare context from context posts
    const blogContext = posts.filter(p => p.status === 'published').map(p => `Title: ${p.title}\nExcerpt: ${p.excerpt}`).join('\n\n');

    let currentResponse = "";
    setMessages(prev => [...prev, { role: 'model', text: '' }]);

    await streamChatResponse(
      newMessages.filter(m => m.text),
      userText,
      blogContext,
      (chunk) => {
        currentResponse += chunk;
        setMessages(prev => {
          const updated = [...prev];
          updated[updated.length - 1] = { role: 'model', text: currentResponse };
          return updated;
        });
      }
    );

    setIsTyping(false);
  };

  return (
    <div className="fixed bottom-8 right-8 z-50 flex flex-col items-end font-sans">
      {/* Chat Window */}
      {isOpen && (
        <div
          ref={chatContainerRef}
          className="mb-6 w-80 sm:w-96 bg-white/90 backdrop-blur-2xl rounded-3xl border border-white/60 shadow-[0_20px_50px_-12px_rgba(0,0,0,0.15)] overflow-hidden flex flex-col h-[600px] animate-slide-up origin-bottom-right ring-1 ring-white/50"
        >
          {/* Header */}
          <div className="bg-gradient-to-r from-stone-50/80 to-white/80 p-4 flex justify-between items-center border-b border-stone-100/50 backdrop-blur-md">
            <div className="flex items-center gap-3">
              <div className="relative">
                <img
                  src={AETHER_AVATAR_URL}
                  alt="Aether"
                  className="w-10 h-10 rounded-full object-cover border-2 border-white shadow-md"
                />
                <span className="absolute bottom-0 right-0 w-2.5 h-2.5 bg-teal-400 border-2 border-white rounded-full animate-pulse"></span>
              </div>
              <div className="flex flex-col">
                <span className="font-serif italic tracking-wide font-medium text-stone-800">Aether AI</span>
                <span className="text-[10px] uppercase tracking-widest text-teal-600/80 font-semibold">Online</span>
              </div>
            </div>
            <button
              onClick={() => setIsOpen(false)}
              className="w-8 h-8 flex items-center justify-center rounded-full hover:bg-stone-100 text-stone-400 hover:text-stone-600 transition-colors cursor-pointer"
            >
              <IconX className="w-4 h-4" />
            </button>
          </div>

          {/* Messages */}
          <div className="flex-1 overflow-y-auto p-5 space-y-6 scrollbar-thin scrollbar-thumb-stone-200 scrollbar-track-transparent">
            {messages.map((msg, idx) => (
              <div
                key={idx}
                className={`flex gap-3 items-end ${msg.role === 'user' ? 'flex-row-reverse' : 'flex-row'}`}
              >
                {/* Avatar for Model */}
                {msg.role === 'model' && (
                  <div className="flex-shrink-0 mb-1">
                    <img
                      src={AETHER_AVATAR_URL}
                      alt="AI"
                      className="w-8 h-8 rounded-full object-cover border border-white/50 shadow-sm"
                    />
                  </div>
                )}

                <div
                  className={`max-w-[80%] rounded-2xl px-4 py-2.5 text-sm leading-relaxed shadow-sm ${msg.role === 'user'
                    ? 'bg-stone-800 text-stone-50 rounded-br-sm shadow-stone-200'
                    : 'bg-white border border-white/80 shadow-[0_2px_8px_rgba(0,0,0,0.04)] text-stone-700 rounded-bl-sm'
                    }`}
                >
                  {/* Text Content */}
                  {msg.text && (
                    <div className={`markdown-content ${msg.role === 'user' ? 'text-stone-50' : 'text-stone-700'}`}>
                      <MDEditor.Markdown
                        source={msg.text}
                        style={{
                          backgroundColor: 'transparent',
                          color: 'inherit',
                          fontFamily: 'inherit',
                          fontSize: 'inherit'
                        }}
                      />
                    </div>
                  )}

                  {/* Loading Animation (Thinking Dots) */}
                  {msg.role === 'model' && idx === messages.length - 1 && isTyping && !msg.text && (
                    <div className="flex items-center gap-1 h-5 px-1">
                      <div className="w-1.5 h-1.5 rounded-full bg-teal-400 animate-sine" style={{ animationDelay: '0s' }}></div>
                      <div className="w-1.5 h-1.5 rounded-full bg-gold-400 animate-sine" style={{ animationDelay: '0.2s' }}></div>
                      <div className="w-1.5 h-1.5 rounded-full bg-rose-400 animate-sine" style={{ animationDelay: '0.4s' }}></div>
                    </div>
                  )}

                  {/* Streaming Cursor */}
                  {msg.role === 'model' && idx === messages.length - 1 && isTyping && msg.text && (
                    <span className="inline-block w-1.5 h-3.5 ml-1 bg-teal-400/80 animate-pulse align-middle rounded-full"></span>
                  )}
                </div>
              </div>
            ))}
            <div ref={messagesEndRef} />
          </div>

          {/* Input */}
          <form onSubmit={handleSubmit} className="p-4 bg-white/60 border-t border-stone-100 backdrop-blur-md">
            <div className="relative group">
              <input
                type="text"
                value={input}
                onChange={(e) => setInput(e.target.value)}
                placeholder="Ask Aether about the journal..."
                className="w-full pl-5 pr-12 py-3.5 bg-white/90 rounded-full border border-stone-200/80 focus:outline-none focus:border-teal-300 focus:ring-4 focus:ring-teal-50/50 text-sm text-ink placeholder-stone-400 transition-all shadow-sm group-hover:bg-white group-hover:shadow-md"
              />
              <button
                type="submit"
                disabled={!input.trim() || isTyping}
                className="absolute right-2 top-1/2 -translate-y-1/2 p-2 bg-stone-100 rounded-full text-stone-400 hover:text-white hover:bg-teal-500 disabled:opacity-30 disabled:hover:bg-stone-100 disabled:hover:text-stone-400 disabled:cursor-not-allowed transition-all duration-300 cursor-pointer shadow-sm hover:shadow-md hover:scale-105"
              >
                <IconSend className="w-4 h-4" />
              </button>
            </div>
          </form>
        </div>
      )}

      {/* Toggle Button with Enhanced Effects */}
      <button
        onClick={() => setIsOpen(!isOpen)}
        className={`group relative flex items-center justify-center w-12 h-12 rounded-full shadow-[0_8px_30px_rgba(0,0,0,0.12)] backdrop-blur-xl border border-white/60 transition-all duration-500 cursor-pointer ${isOpen
          ? 'bg-stone-800 rotate-90 scale-90'
          : 'bg-white/90 hover:scale-110 hover:-translate-y-1'
          }`}
      >
        {/* Outer Glow Ring - Spreads Outwards */}
        {!isOpen && (
          <span className="absolute -inset-10 rounded-full bg-gradient-to-tr from-teal-300/40 to-gold-300/40 opacity-0 group-hover:opacity-100 blur-3xl transition-all duration-700 scale-50 group-hover:scale-110"></span>
        )}

        {/* Rotating Border Ring */}
        {!isOpen && (
          <span className="absolute -inset-0.5 rounded-full border border-teal-500/30 border-t-transparent opacity-0 group-hover:opacity-100 animate-spin-slow transition-opacity duration-500"></span>
        )}

        {isOpen ? (
          <IconX className="w-5 h-5 text-white" />
        ) : (
          <div className="relative z-10">
            <IconSparkles className="w-6 h-6 text-teal-600 group-hover:text-teal-500 transition-colors duration-300" />

            {/* Inner Pulse */}
            <span className="absolute inset-0 rounded-full bg-teal-400/20 animate-ping duration-[2000ms] opacity-0 group-hover:opacity-100"></span>
          </div>
        )}
      </button>
    </div>
  );
};

export default AIChat;
