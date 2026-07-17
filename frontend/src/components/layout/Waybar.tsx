import { useFetch } from '../../hooks/useFetch';
import type { WaybarData } from '../../types/waybar';
import Modules from '../waybar/Modules';

interface WaybarProps {
    pageTitle: string;
}

export default function Waybar({ pageTitle }: WaybarProps) {
    const { data, loading, error } = useFetch<WaybarData>('/api/v1/waybar');

    if (loading || error || !data) return null;

    return (
        <div id="waybar" className="relative flex w-screen items-center justify-between">
            <div className="modules-left flex z-10">
                <Modules
                    layoutKeys={data.waybar.layout.left}
                    modules={data.waybar.modules}
                    pageTitle={pageTitle}
                />
            </div>

            <div className="modules-center flex-abs-center">
                <Modules
                    layoutKeys={data.waybar.layout.center}
                    modules={data.waybar.modules}
                    pageTitle={pageTitle}
                />
            </div>

            <div className="modules-right flex z-10 py-0">
                <Modules
                    layoutKeys={data.waybar.layout.right}
                    modules={data.waybar.modules}
                    pageTitle={pageTitle}
                />
            </div>
        </div>
    );
}
