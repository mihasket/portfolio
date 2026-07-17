import { useEffect, useState } from 'react';
import { tryCatch, fetchJson } from '../types/result';

const API_URL = import.meta.env.DEV ? '' : import.meta.env.VITE_API_URL;

export function useFetch<T>(endpoint: string) {
    const [data, setData] = useState<T | null>(null);
    const [error, setError] = useState<string | null>(null);
    const [loading, setLoading] = useState(true);

    useEffect(() => {
        const controller = new AbortController();

        async function fetchData() {
            setLoading(true);

            const result = await tryCatch(
                fetchJson<T>(`${API_URL}${endpoint}`, {
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

        fetchData();

        return () => {
            controller.abort();
        };
    }, [endpoint]);

    return { data, error, loading };
}
