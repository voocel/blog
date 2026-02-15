import React, { useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import { motion } from 'framer-motion';
import AnimatedNavWidget from '@/components/AnimatedNavWidget';
import SEO from '@/components/SEO';
import { IconSparkles, IconBrain, IconMessage, IconGithub, IconMail } from '@/components/Icons';
import { useTranslation } from '@/hooks/useTranslation';

interface Pillar {
  title: string;
  desc: string;
  icon: React.FC<{ className?: string }>;
  tag: string;
}

interface AIProject {
  name: string;
  path: string;
  desc: string;
  topic: string;
}

const AboutPage: React.FC = () => {
  const navigate = useNavigate();
  const { locale } = useTranslation();

  useEffect(() => {
    window.scrollTo(0, 0);
  }, []);

  const isZh = locale === 'zh';

  const copy = isZh
    ? {
        seoTitle: '关于 - Voocel',
        badge: 'VOOCEL / ABOUT',
        title: '把 AI 工程实践写成可读、可复用的长期资产',
        subtitle: 'AI + Engineering Journal',
        intro:
          '这不是个人简历页，而是一个持续迭代的 AI 工程现场。围绕大模型应用、Agent 系统和基础设施实现，记录真实问题与真实解法。',
        highlights: [
          { value: 'LLM 应用', label: '主题方向' },
          { value: 'MCP / Agent', label: '系统方向' },
          { value: 'GitHub 52+ Repos', label: '开源积累' },
        ],
        pillarTitle: '写作与开发原则',
        aiTitle: 'GitHub AI 方向',
        aiIntro: '基于 GitHub（voocel）当前公开项目，核心方向集中在 LLM 中间层、MCP 生态和 Agent 工程化。',
        stackTitle: '当前技术栈',
        notesTitle: '构建方式',
        notes: [
          'AI 落地优先：先解决真实场景，再追求抽象优雅。',
          '结构稳定：模型调用、业务逻辑、基础设施三层解耦。',
          '持续迭代：每次改动都可验证、可回滚、可解释。',
        ],
        ctaTitle: '保持联系',
        ctaDesc: '欢迎交流 LLM 应用、Agent 工程化、MCP 生态相关问题，或对文章提出改进建议。',
      }
    : {
        seoTitle: 'About - Voocel',
        badge: 'VOOCEL / ABOUT',
        title: 'Turning AI engineering practice into reusable long-term knowledge',
        subtitle: 'AI + Engineering Journal',
        intro:
          'This is not a resume page. It is an evolving AI engineering workspace documenting LLM applications, agent systems, and real implementation trade-offs.',
        highlights: [
          { value: 'LLM Apps', label: 'Main Focus' },
          { value: 'MCP / Agent', label: 'System Track' },
          { value: 'GitHub 52+ Repos', label: 'Open Source Base' },
        ],
        pillarTitle: 'Writing & Engineering Principles',
        aiTitle: 'GitHub AI Focus',
        aiIntro: 'Based on public projects in voocel, the current focus is LLM middleware, MCP ecosystem tooling, and production-oriented agent systems.',
        stackTitle: 'Current Stack',
        notesTitle: 'How It Is Built',
        notes: [
          'AI delivery first: solve concrete production cases before abstraction.',
          'Stable structure: model access, business logic, and infra are decoupled.',
          'Continuous iteration: each change should be testable and explainable.',
        ],
        ctaTitle: 'Let’s Connect',
        ctaDesc: 'Open to discussions on LLM apps, agent engineering, MCP tooling, and collaboration.',
      };

  const pillars: Pillar[] = isZh
    ? [
        {
          title: '表达要清楚',
          desc: '每篇内容都尽量给出背景、问题和可执行结论，降低读者理解成本。',
          icon: IconMessage,
          tag: 'Clarity',
        },
        {
          title: '技术要落地',
          desc: '优先分享真实场景里的方案，不追热点堆概念，强调可复现。',
          icon: IconBrain,
          tag: 'Practical',
        },
        {
          title: '体验要克制',
          desc: '界面和交互保持简洁，视觉服务内容，不让样式喧宾夺主。',
          icon: IconSparkles,
          tag: 'Intentional',
        },
      ]
    : [
        {
          title: 'Clarity First',
          desc: 'Every article aims to provide context, concrete problems, and actionable takeaways.',
          icon: IconMessage,
          tag: 'Clarity',
        },
        {
          title: 'Practical Depth',
          desc: 'Real-world solutions over hype, with reproducible patterns and tradeoffs.',
          icon: IconBrain,
          tag: 'Practical',
        },
        {
          title: 'Intentional UX',
          desc: 'A restrained visual system where design supports reading, not distraction.',
          icon: IconSparkles,
          tag: 'Intentional',
        },
      ];

  const techStack = [
    'LLM Orchestration',
    'MCP',
    'Agent Workflow',
    'Go 1.25',
    'Gin',
    'PostgreSQL',
    'GORM',
    'React 19',
    'Vite',
    'TypeScript',
    'Tailwind CSS',
    'Docker Compose',
    'Nginx',
  ];

  const aiProjects: AIProject[] = isZh
    ? [
        {
          name: 'openclaw-mini',
          path: 'https://github.com/voocel/openclaw-mini',
          topic: 'Agent 应用',
          desc: '轻量化的 AI Agent 实践项目，聚焦最小可用闭环和快速验证工作流。',
        },
        {
          name: 'litellm',
          path: 'https://github.com/voocel/litellm',
          topic: 'LLM 网关',
          desc: '统一多模型调用与路由策略，降低不同模型接口差异带来的接入复杂度。',
        },
        {
          name: 'mcp-sdk-go',
          path: 'https://github.com/voocel/mcp-sdk-go',
          topic: 'MCP 工具链',
          desc: '围绕 Model Context Protocol 的 Go SDK 能力，服务模型与工具之间的标准化交互。',
        },
        {
          name: 'agentcore',
          path: 'https://github.com/voocel/agentcore',
          topic: 'Agent 框架',
          desc: '面向生产场景的 Agent 组织方式，聚焦可观测、可维护和流程化执行。',
        },
      ]
    : [
        {
          name: 'openclaw-mini',
          path: 'https://github.com/voocel/openclaw-mini',
          topic: 'Agent App',
          desc: 'A lightweight AI agent project focused on minimal viable loops and fast workflow validation.',
        },
        {
          name: 'litellm',
          path: 'https://github.com/voocel/litellm',
          topic: 'LLM Gateway',
          desc: 'A unified layer for multi-model routing and access, reducing provider integration complexity.',
        },
        {
          name: 'mcp-sdk-go',
          path: 'https://github.com/voocel/mcp-sdk-go',
          topic: 'MCP Tooling',
          desc: 'Go SDK work around Model Context Protocol for standardized model-tool interaction.',
        },
        {
          name: 'agentcore',
          path: 'https://github.com/voocel/agentcore',
          topic: 'Agent Framework',
          desc: 'Production-oriented agent architecture with focus on observability and maintainable workflow execution.',
        },
      ];

  return (
    <div className="relative min-h-screen bg-[var(--color-base)] text-ink pt-20 pb-24 overflow-hidden">
      <SEO title={copy.seoTitle} />

      <div className="pointer-events-none absolute -top-24 -left-20 w-[28rem] h-[28rem] rounded-full bg-amber-200/30 blur-3xl" />
      <div className="pointer-events-none absolute top-64 -right-24 w-[24rem] h-[24rem] rounded-full bg-orange-200/30 blur-3xl" />
      <div className="pointer-events-none absolute inset-0 opacity-[0.035] [background-image:linear-gradient(to_right,currentColor_1px,transparent_1px),linear-gradient(to_bottom,currentColor_1px,transparent_1px)] [background-size:42px_42px]" />

      <div className="fixed top-8 left-8 z-50">
        <AnimatedNavWidget
          isCompact={true}
          disableFixed={true}
          showBackButton={true}
          onBackClick={() => navigate('/')}
        />
      </div>

      <main className="relative z-10 max-w-6xl mx-auto px-6">
        <motion.section
          className="grid grid-cols-1 lg:grid-cols-[1.25fr_0.75fr] gap-8 items-end mb-16"
          initial={{ opacity: 0, y: 18 }}
          animate={{ opacity: 1, y: 0 }}
          transition={{ duration: 0.45, ease: 'easeOut' }}
        >
          <div>
            <p className="text-xs uppercase tracking-[0.26em] text-gold-600 font-bold mb-5">{copy.badge}</p>
            <h1 className="text-4xl md:text-6xl font-serif font-bold leading-[1.06] mb-5 max-w-4xl">{copy.title}</h1>
            <p className="text-sm uppercase tracking-[0.2em] text-[var(--color-text-muted)] mb-6">{copy.subtitle}</p>
            <p className="text-base md:text-xl text-[var(--color-text-secondary)] leading-relaxed max-w-3xl">{copy.intro}</p>
          </div>

          <div className="rounded-3xl border border-[var(--color-border)] bg-[var(--color-elevated)]/80 backdrop-blur-xl p-7 shadow-sm">
            <div className="flex items-center gap-3 mb-4">
              <img src="/logo.svg" alt="Voocel Logo" className="w-9 h-9 rounded-lg" />
              <div>
                <p className="font-bold text-base">Voocel</p>
                <p className="text-xs uppercase tracking-[0.16em] text-[var(--color-text-muted)]">Journal System</p>
              </div>
            </div>
            <p className="text-sm text-[var(--color-text-secondary)] leading-relaxed">
              {isZh
                ? '站点持续围绕工程深度与表达质量更新。所有设计选择都服务于阅读体验与知识沉淀。'
                : 'The site evolves around engineering depth and editorial quality. Every design choice serves reading and knowledge retention.'}
            </p>
          </div>
        </motion.section>

        <motion.section
          className="grid grid-cols-1 sm:grid-cols-3 gap-4 mb-20"
          initial={{ opacity: 0, y: 18 }}
          animate={{ opacity: 1, y: 0 }}
          transition={{ delay: 0.08, duration: 0.45, ease: 'easeOut' }}
        >
          {copy.highlights.map((item) => (
            <div
              key={item.label}
              className="rounded-2xl border border-[var(--color-border)] bg-[var(--color-surface)]/90 backdrop-blur px-5 py-6"
            >
              <p className="text-[10px] uppercase tracking-[0.2em] text-[var(--color-text-muted)] mb-2">{item.label}</p>
              <p className="text-2xl font-serif font-bold text-ink">{item.value}</p>
            </div>
          ))}
        </motion.section>

        <motion.section
          className="mb-20"
          initial={{ opacity: 0, y: 18 }}
          animate={{ opacity: 1, y: 0 }}
          transition={{ delay: 0.16, duration: 0.45, ease: 'easeOut' }}
        >
          <h2 className="text-2xl md:text-3xl font-serif font-bold mb-8">{copy.pillarTitle}</h2>
          <div className="grid grid-cols-1 lg:grid-cols-3 gap-5">
            {pillars.map((item) => {
              const Icon = item.icon;
              return (
                <article
                  key={item.title}
                  className="group rounded-2xl border border-[var(--color-border)] bg-[var(--color-surface)] px-6 py-7 hover:-translate-y-0.5 hover:shadow-md transition-all duration-300"
                >
                  <div className="flex items-center justify-between mb-4">
                    <div className="w-11 h-11 rounded-xl bg-[var(--color-surface-alt)] flex items-center justify-center">
                      <Icon className="w-5 h-5 text-gold-600" />
                    </div>
                    <span className="text-[10px] uppercase tracking-[0.18em] text-[var(--color-text-muted)]">{item.tag}</span>
                  </div>
                  <h3 className="text-lg font-bold mb-2">{item.title}</h3>
                  <p className="text-sm leading-relaxed text-[var(--color-text-secondary)]">{item.desc}</p>
                </article>
              );
            })}
          </div>
        </motion.section>

        <motion.section
          className="mb-20"
          initial={{ opacity: 0, y: 18 }}
          animate={{ opacity: 1, y: 0 }}
          transition={{ delay: 0.2, duration: 0.45, ease: 'easeOut' }}
        >
          <div className="flex items-end justify-between mb-7 gap-4">
            <div>
              <h2 className="text-2xl md:text-3xl font-serif font-bold mb-2">{copy.aiTitle}</h2>
              <p className="text-sm text-[var(--color-text-secondary)] max-w-3xl">{copy.aiIntro}</p>
            </div>
            <a
              href="https://github.com/voocel"
              target="_blank"
              rel="noopener noreferrer"
              className="shrink-0 text-xs uppercase tracking-[0.16em] text-[var(--color-text-muted)] hover:text-ink transition-colors"
            >
              github.com/voocel
            </a>
          </div>

          <div className="grid grid-cols-1 lg:grid-cols-3 gap-5">
            {aiProjects.map((repo) => (
              <a
                key={repo.name}
                href={repo.path}
                target="_blank"
                rel="noopener noreferrer"
                className="group rounded-2xl border border-[var(--color-border)] bg-[var(--color-surface)] px-6 py-6 hover:-translate-y-0.5 hover:shadow-md transition-all duration-300"
              >
                <p className="text-[10px] uppercase tracking-[0.2em] text-gold-600 font-bold mb-2">{repo.topic}</p>
                <h3 className="text-lg font-bold mb-3 group-hover:text-orange-600 transition-colors">{repo.name}</h3>
                <p className="text-sm leading-relaxed text-[var(--color-text-secondary)]">{repo.desc}</p>
              </a>
            ))}
          </div>
        </motion.section>

        <motion.section
          className="grid grid-cols-1 lg:grid-cols-[1fr_1fr] gap-10 mb-20"
          initial={{ opacity: 0, y: 18 }}
          animate={{ opacity: 1, y: 0 }}
          transition={{ delay: 0.24, duration: 0.45, ease: 'easeOut' }}
        >
          <div>
            <h2 className="text-2xl md:text-3xl font-serif font-bold mb-6">{copy.stackTitle}</h2>
            <div className="flex flex-wrap gap-2.5">
              {techStack.map((item) => (
                <span
                  key={item}
                  className="px-3 py-1.5 rounded-lg bg-[var(--color-surface)] border border-[var(--color-border)] text-sm text-[var(--color-text-secondary)]"
                >
                  {item}
                </span>
              ))}
            </div>
          </div>

          <div>
            <h2 className="text-2xl md:text-3xl font-serif font-bold mb-6">{copy.notesTitle}</h2>
            <div className="relative pl-6 space-y-5">
              <div className="absolute left-1 top-1 bottom-1 w-px bg-[var(--color-border)]" />
              {copy.notes.map((note, idx) => (
                <div key={note} className="relative">
                  <span className="absolute -left-[22px] top-[5px] w-2.5 h-2.5 rounded-full bg-gold-500" />
                  <p className="text-sm leading-relaxed text-[var(--color-text-secondary)]">{note}</p>
                  {idx < copy.notes.length - 1 && <div className="mt-5 border-b border-[var(--color-border-subtle)]" />}
                </div>
              ))}
            </div>
          </div>
        </motion.section>

        <motion.section
          className="rounded-3xl border border-[var(--color-border)] bg-gradient-to-br from-[var(--color-surface)] to-[var(--color-surface-alt)] p-8 md:p-10"
          initial={{ opacity: 0, y: 18 }}
          animate={{ opacity: 1, y: 0 }}
          transition={{ delay: 0.3, duration: 0.45, ease: 'easeOut' }}
        >
          <h2 className="text-2xl md:text-3xl font-serif font-bold mb-4">{copy.ctaTitle}</h2>
          <p className="text-[var(--color-text-secondary)] mb-8 max-w-2xl">{copy.ctaDesc}</p>
          <div className="flex flex-wrap gap-3">
            <a
              href="https://github.com/voocel"
              target="_blank"
              rel="noopener noreferrer"
              className="inline-flex items-center gap-2 px-5 py-3 rounded-xl bg-ink text-[var(--color-base)] font-medium hover:opacity-90 transition-opacity"
            >
              <IconGithub className="w-4 h-4" />
              GitHub
            </a>
            <a
              href="mailto:hello@voocel.com"
              className="inline-flex items-center gap-2 px-5 py-3 rounded-xl border border-[var(--color-border)] bg-[var(--color-base)]/70 text-ink font-medium hover:border-gold-500 transition-colors"
            >
              <IconMail className="w-4 h-4" />
              hello@voocel.com
            </a>
          </div>
        </motion.section>
      </main>
    </div>
  );
};

export default AboutPage;
