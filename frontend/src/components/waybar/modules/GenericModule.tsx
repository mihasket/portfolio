import type { WaybarModule } from '../../../types/waybar';

interface GenericModuleProps {
    module: WaybarModule;
    domId: string;
}

export default function GenericModule({ module, domId }: GenericModuleProps) {
    const paddingClass =
        domId === 'custom-dateicon'
            ? 'py-3 pl-3 pr-1'
            : domId === 'clock'
              ? 'py-3 pr-3 pl-1'
              : 'p-3';

    return (
        <div className={paddingClass}>
            <p>{module.format}</p>
        </div>
    );
}
