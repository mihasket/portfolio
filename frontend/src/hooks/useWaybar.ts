import { useEffect, useState } from 'react';
import type { WaybarData } from '../types/waybar';
import { tryCatch, fetchJson } from '../types/result';

const API_URL = import.meta.env.DEV ? '' : import.meta.env.VITE_API_URL;

export function useWaybar() {
    const [data, setData] = useState<WaybarData | null>(null);
    const [error, setError] = useState<string | null>(null);
    const [loading, setLoading] = useState(true);

    useEffect(() => {
        const controller = new AbortController();

        async function fetchWaybar() {
            setLoading(true);

            const result = await tryCatch(
                fetchJson<WaybarData>(`${API_URL}/api/v1/waybar`, {
                    signal: controller.signal,
                })
            );

            if (result.error) {
                if (result.error instanceof DOMException && result.error.name === 'AbortError') {
                    return;
                }

                setError(result.error.message);
            } else {
                setData(result.data);
                setError(null);
            }

            setLoading(false);
        }

        fetchWaybar();

        return () => {
            controller.abort();
        };
    }, []);

    return { data, error, loading };
}
