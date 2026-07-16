import type { WaybarModule } from '../../../types/waybar';

interface LogoModuleProps {
    module: WaybarModule;
    domId: string;
    pageTitle: string;
}

export default function LogoModule({ module, domId, pageTitle }: LogoModuleProps) {
    return (
        <div id={domId}>
            {module.format} {pageTitle} - Miha Šket
        </div>
    );
}
