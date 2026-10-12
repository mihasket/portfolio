import type { ReactNode } from 'react';

interface TerminalProps {
    children: ReactNode;
}

export default function Terminal({ children }: TerminalProps) {
    return <div className="terminal h-full w-full">{children}</div>;
}
