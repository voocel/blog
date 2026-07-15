import { useCallback, useEffect, useState } from 'react';
import type { HTMLAttributes } from 'react';
import { getCodeString } from 'rehype-rewrite';
import mermaid from 'mermaid';
import { useSettings } from '@/context/SettingsContext';

let seq = 0;
const nextId = () => `mermaid-${++seq}`;

interface CodeProps extends HTMLAttributes<HTMLElement> {
    node?: { children?: unknown };
}

/**
 * Custom `code` renderer for @uiw/react-md-editor.
 * Renders ```mermaid fences as diagrams; all other code blocks fall through unchanged.
 * Theme follows the app's effective light/dark setting.
 */
const MermaidCode = ({ className, children = [], node, ...props }: CodeProps) => {
    const { effectiveTheme } = useSettings();
    const [container, setContainer] = useState<HTMLElement | null>(null);

    const isMermaid = !!className && /^language-mermaid/.test(className.toLowerCase());
    const code = node?.children
        ? getCodeString(node.children as never)
        : Array.isArray(children)
            ? String(children[0] ?? '')
            : String(children ?? '');

    useEffect(() => {
        if (!container || !isMermaid || !code.trim()) return;
        let cancelled = false;
        // Fresh id per render avoids collisions under StrictMode / re-renders.
        const renderId = nextId();

        mermaid.initialize({
            startOnLoad: false,
            theme: effectiveTheme === 'dark' ? 'dark' : 'default',
        });

        mermaid
            .render(renderId, code)
            .then(({ svg, bindFunctions }) => {
                if (cancelled) return;
                container.innerHTML = svg;
                bindFunctions?.(container);
            })
            .catch((err: unknown) => {
                if (cancelled) return;
                const msg = err instanceof Error ? err.message : String(err);
                container.textContent = `Mermaid error: ${msg}`;
            });

        return () => {
            cancelled = true;
        };
    }, [container, isMermaid, code, effectiveTheme]);

    const ref = useCallback((el: HTMLElement | null) => {
        if (el) setContainer(el);
    }, []);

    if (isMermaid) {
        return <code ref={ref} data-name="mermaid" className={className} {...props} />;
    }

    return (
        <code className={className} {...props}>
            {children}
        </code>
    );
};

export default MermaidCode;
