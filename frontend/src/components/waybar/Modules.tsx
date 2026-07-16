import type { WaybarModule } from '../../types/waybar';
import { MODULE_KEYS } from '../../types/waybar';
import LogoModule from './modules/LogoModule';
import WorkspaceModule from './modules/WorkspaceModule';
import GenericModule from './modules/GenericModule';

interface ModulesProps {
    layoutKeys: string[];
    modules: Record<string, WaybarModule>;
    pageTitle: string;
}

function toDomId(key: string): string {
    return key.replace(/\//g, '-');
}

export default function Modules({ layoutKeys, modules, pageTitle }: ModulesProps) {
    const currentPath = typeof window !== 'undefined' ? window.location.pathname : '/';

    return (
        <>
            {layoutKeys.map((key) => {
                const module = modules[key];
                if (!module) return null;

                const domId = toDomId(key);

                switch (key) {
                    case MODULE_KEYS.LOGO:
                        return (
                            <LogoModule
                                key={key}
                                module={module}
                                domId={domId}
                                pageTitle={pageTitle}
                            />
                        );
                    case MODULE_KEYS.WORKSPACE:
                        return (
                            <WorkspaceModule
                                key={key}
                                module={module}
                                domId={domId}
                                currentPath={currentPath}
                            />
                        );
                    default:
                        return <GenericModule key={key} module={module} domId={domId} />;
                }
            })}
        </>
    );
}
