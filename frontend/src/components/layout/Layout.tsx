import type { ReactNode } from 'react';
import Waybar from './Waybar';

interface LayoutProps {
    title: string;
    children: ReactNode;
}

export default function Layout({ title, children }: LayoutProps) {
    return (
        <>
            <title>{`${title} - Miha Šket`}</title>

            <header>
                <Waybar pageTitle={title} />
            </header>

            <main id="main" className="flex items-stretch justify-center w-full">
                {children}
            </main>

            <div className="mobile"></div>
        </>
    );
}
