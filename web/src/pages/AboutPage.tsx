import React, { useEffect, useRef, useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { IconSparkles, IconBrain, IconMessage, IconSend } from '@/components/Icons';
import AnimatedNavWidget from '@/components/AnimatedNavWidget';


// --- Reusable Reveal Component (Same as HomePage) ---
interface RevealProps {
    children: React.ReactNode;
    className?: string;
    delay?: number;
    threshold?: number;
}

const Reveal: React.FC<RevealProps> = ({ children, className = "", delay = 0, threshold = 0.1 }) => {
    const ref = useRef<HTMLDivElement>(null);
    const [isVisible, setIsVisible] = useState(false);

    useEffect(() => {
        const observer = new IntersectionObserver(
            ([entry]) => {
                if (entry.isIntersecting) {
                    setIsVisible(true);
                    observer.disconnect();
                }
            },
            { threshold }
        );

        if (ref.current) {
            observer.observe(ref.current);
        }

        return () => observer.disconnect();
    }, [threshold]);

    return (
        <div
            ref={ref}
            className={`${isVisible ? 'animate-blur-in' : 'opacity-0'} ${className}`}
            style={{ animationDelay: `${delay}ms`, animationFillMode: 'both' }}
        >
            {children}
        </div>
    );
};

const AboutPage: React.FC = () => {
    const navigate = useNavigate();
    useEffect(() => {
        window.scrollTo(0, 0);
    }, []);

    return (
        <div className="min-h-screen pt-20 pb-32 bg-[var(--color-base)] text-ink">
            <div className="fixed top-8 left-8 z-50">
                <AnimatedNavWidget
                    isCompact={true}
                    disableFixed={true}
                    showBackButton={true}
                    onBackClick={() => navigate('/')}
                />
            </div>
            {/* Hero Section */}
            <section className="relative py-20 md:py-32 px-6 overflow-hidden">
                <div className="max-w-4xl mx-auto text-center relative z-10">
                    <Reveal>
                        <p className="text-xs uppercase tracking-[0.3em] text-gold-600 mb-6 font-bold">The Philosophy</p>
                    </Reveal>
                    <Reveal delay={200}>
                        <h1 className="text-5xl md:text-7xl font-serif font-bold mb-8 leading-tight">
                            Curating the <span className="italic text-gold-600">Digital</span> Soul.
                        </h1>
                    </Reveal>
                    <Reveal delay={400}>
                        <p className="text-xl md:text-2xl text-[var(--color-text-secondary)] font-serif italic max-w-2xl mx-auto leading-relaxed">
                            "We believe that code is poetry, and design is the silent language of connection. Voocel is a sanctuary for those who seek beauty in the binary."
                        </p>
                    </Reveal>
                </div>

                {/* Background Elements */}
                <div className="absolute top-0 left-1/2 -translate-x-1/2 w-[800px] h-[800px] bg-gold-100/30 rounded-full blur-3xl -z-10 opacity-50 pointer-events-none" />
            </section>

            {/* The Author Section */}
            <section className="max-w-6xl mx-auto px-6 mb-32">
                <div className="grid grid-cols-1 md:grid-cols-2 gap-12 md:gap-24 items-center">
                    <Reveal delay={200} className="relative group">
                        <div className="aspect-[4/5] rounded-2xl overflow-hidden relative z-10">
                            <img
                                src="https://images.unsplash.com/photo-1544005313-94ddf0286df2?fit=crop&w=800&q=80"
                                alt="Alex Chen"
                                className="w-full h-full object-cover transition-transform duration-700 group-hover:scale-105"
                            />
                            <div className="absolute inset-0 bg-gradient-to-t from-stone-900/40 to-transparent opacity-60" />
                        </div>
                        {/* Decorative Frame */}
                        <div className="absolute -inset-4 border border-gold-200 rounded-2xl -z-10 translate-x-4 translate-y-4 transition-transform duration-500 group-hover:translate-x-2 group-hover:translate-y-2" />
                    </Reveal>

                    <div className="space-y-8">
                        <Reveal delay={300}>
                            <h2 className="text-3xl md:text-4xl font-serif font-bold mb-2">Alex Chen</h2>
                            <p className="text-xs uppercase tracking-widest text-gold-600">Founder & Lead Editor</p>
                        </Reveal>

                        <Reveal delay={400}>
                            <p className="text-[var(--color-text-secondary)] leading-relaxed mb-6">
                                With over a decade of experience in both software engineering and visual arts, Alex founded Voocel to bridge the gap between technical precision and aesthetic emotion.
                            </p>
                            <p className="text-[var(--color-text-secondary)] leading-relaxed">
                                "I wanted to create a space where technology doesn't feel cold. Every line of code should serve a human purpose, and every pixel should tell a story. This journal is my exploration of that intersection."
                            </p>
                        </Reveal>

                        <Reveal delay={500}>
                            <div className="flex gap-8 pt-4">
                                <div>
                                    <div className="text-3xl font-serif font-bold text-ink">12+</div>
                                    <div className="text-xs uppercase tracking-wider text-[var(--color-text-muted)] mt-1">Years Exp.</div>
                                </div>
                                <div>
                                    <div className="text-3xl font-serif font-bold text-ink">150+</div>
                                    <div className="text-xs uppercase tracking-wider text-[var(--color-text-muted)] mt-1">Articles</div>
                                </div>
                                <div>
                                    <div className="text-3xl font-serif font-bold text-ink">10k</div>
                                    <div className="text-xs uppercase tracking-wider text-[var(--color-text-muted)] mt-1">Readers</div>
                                </div>
                            </div>
                        </Reveal>
                    </div>
                </div>
            </section>

            {/* Values Grid */}
            <section className="bg-[var(--color-surface)] py-24 border-y border-[var(--color-border-subtle)]">
                <div className="max-w-6xl mx-auto px-6">
                    <Reveal className="text-center mb-16">
                        <h2 className="text-3xl font-serif font-bold">Our Core Values</h2>
                    </Reveal>

                    <div className="grid grid-cols-1 md:grid-cols-3 gap-12">
                        <Reveal delay={100} className="text-center group cursor-default">
                            <div className="w-16 h-16 mx-auto bg-[var(--color-surface-alt)] rounded-full flex items-center justify-center mb-6 group-hover:bg-gold-50 transition-colors duration-500">
                                <IconSparkles className="w-8 h-8 text-[var(--color-text-muted)] group-hover:text-gold-600 transition-colors" />
                            </div>
                            <h3 className="text-lg font-bold mb-3 font-serif">Aesthetic First</h3>
                            <p className="text-[var(--color-text-secondary)] text-sm leading-relaxed">
                                Beauty is not an afterthought. It is the foundation of functionality and the key to user engagement.
                            </p>
                        </Reveal>

                        <Reveal delay={200} className="text-center group cursor-default">
                            <div className="w-16 h-16 mx-auto bg-[var(--color-surface-alt)] rounded-full flex items-center justify-center mb-6 group-hover:bg-gold-50 transition-colors duration-500">
                                <IconBrain className="w-8 h-8 text-[var(--color-text-muted)] group-hover:text-gold-600 transition-colors" />
                            </div>
                            <h3 className="text-lg font-bold mb-3 font-serif">Intelligent Design</h3>
                            <p className="text-[var(--color-text-secondary)] text-sm leading-relaxed">
                                Leveraging AI and modern algorithms to enhance human creativity, not replace it.
                            </p>
                        </Reveal>

                        <Reveal delay={300} className="text-center group cursor-default">
                            <div className="w-16 h-16 mx-auto bg-[var(--color-surface-alt)] rounded-full flex items-center justify-center mb-6 group-hover:bg-gold-50 transition-colors duration-500">
                                <IconMessage className="w-8 h-8 text-[var(--color-text-muted)] group-hover:text-gold-600 transition-colors" />
                            </div>
                            <h3 className="text-lg font-bold mb-3 font-serif">Open Dialogue</h3>
                            <p className="text-[var(--color-text-secondary)] text-sm leading-relaxed">
                                Fostering a community where ideas flow freely and diverse perspectives are celebrated.
                            </p>
                        </Reveal>
                    </div>
                </div>
            </section>

            {/* Contact Section */}
            <section className="max-w-4xl mx-auto px-6 py-32 text-center">
                <Reveal>
                    <h2 className="text-4xl font-serif font-bold mb-6">Let's Create Together</h2>
                    <p className="text-[var(--color-text-secondary)] mb-10 max-w-lg mx-auto">
                        Whether you have a project in mind or just want to discuss the future of digital design, we'd love to hear from you.
                    </p>

                    <div className="flex flex-col sm:flex-row gap-4 justify-center items-center">
                        <button className="px-8 py-4 bg-ink text-white rounded-full font-medium tracking-wide hover:bg-gold-600 transition-colors shadow-lg hover:shadow-xl flex items-center gap-3 group cursor-pointer">
                            <span>Start a Project</span>
                            <IconSend className="w-4 h-4 group-hover:translate-x-1 transition-transform" />
                        </button>
                        <button className="px-8 py-4 bg-[var(--color-surface)] border border-[var(--color-border)] text-ink rounded-full font-medium tracking-wide hover:border-gold-400 transition-colors flex items-center gap-3 cursor-pointer">
                            <span>hello@voocel.com</span>
                        </button>
                    </div>
                </Reveal>
            </section>
        </div>
    );
};

export default AboutPage;
