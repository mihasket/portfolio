export interface WaybarModule {
    format: string;
    formatActive?: string; // Only used for workspaces
}

export interface WaybarData {
    waybar: {
        layer: string;
        height: number;
        modules: Record<string, WaybarModule>;
        layout: {
            left: string[];
            center: string[];
            right: string[];
        };
    };
}

export const MODULE_KEYS = {
    LOGO: 'custom/logo',
    WORKSPACE: 'hyprland/workspaces',
} as const;
