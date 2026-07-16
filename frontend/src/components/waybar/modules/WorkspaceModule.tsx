import type { WaybarModule } from '../../../types/waybar';

const PAGES = [
    { name: '1', path: '/' },
    { name: '2', path: '/projects' },
];

interface WorkspacesModuleProps {
    module: WaybarModule;
    domId: string;
    currentPath: string;
}

export default function WorkspacesModule({ module, domId, currentPath }: WorkspacesModuleProps) {
    return (
        <nav id={domId} className="flex">
            {PAGES.map((page) => {
                const isActive = currentPath === page.path;
                const icon = isActive && module.formatActive ? module.formatActive : module.format;

                return (
                    <a key={page.path} className="px-2" href={page.path}>
                        {icon}
                    </a>
                );
            })}
        </nav>
    );
}
