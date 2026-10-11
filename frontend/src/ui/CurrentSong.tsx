import { useEffect, useState } from 'react';

interface Song {
    name: string;
    artist: string;
    album: string;
    imageUrlLarge: string;
    imageUrlMedium: string;
    timestamp: number;
}

interface lastFM {
    lastFM: Song;
}

function formatTimestamp(timestamp: number): string {
    return new Date(timestamp * 1000).toLocaleString('en-GB', {
        day: 'numeric',
        month: 'short',
        year: 'numeric',
        hour: '2-digit',
        minute: '2-digit',
        hourCycle: 'h23',
    });
}

export default function CurrentSong() {
    const [data, setData] = useState<Song | null>(null);
    const [error, setError] = useState<string | null>(null);

    useEffect(() => {
        const controller = new AbortController();

        async function fetchCurrentSong() {
            try {
                const res = await fetch('/api/v1/lastfm', { signal: controller.signal });
                if (!res.ok) {
                    throw new Error(`Request failed with status ${res.status}`);
                }
                const json: lastFM = await res.json();
                setData(json.lastFM);
            } catch (err) {
                if (err instanceof DOMException && err.name === 'AbortError') return;
                setError(err instanceof Error ? err.message : 'Unknown error');
            }
        }

        fetchCurrentSong();
        return () => controller.abort();
    }, []);

    if (error) return <p>Failed to load current song: {error}</p>;
    if (!data) return null;

    return (
        <div className="flex items-center gap-3 p-3 text-sm md:gap-0 md:p-0 md:text-xl">
            {/* Image */}
            <div className="shrink-0 md:p-2">
                {data?.imageUrlLarge ? (
                    <picture>
                        <source
                            media="(min-width: 1024px) and (min-height: 800px)"
                            srcSet={data.imageUrlLarge}
                        />
                        <img
                            src={data.imageUrlMedium}
                            alt="lastFM is having problems"
                            className="h-16 w-16 rounded object-cover desk:md:h-[300px] desk:md:w-[300px] desk:md:rounded-none"
                        />
                    </picture>
                ) : (
                    <div className="flex h-16 w-16 items-center justify-center rounded bg-black/30 text-2xl desk:md:h-[300px] desk:md:w-[300px] desk:md:rounded-none desk:md:text-6xl">
                        
                    </div>
                )}
            </div>

            {/* Info */}
            <div className="grid min-w-0 flex-1 grid-cols-[minmax(0,1fr)] gap-x-4 md:flex-none md:grid-cols-[auto_1fr] md:p-2">
                <span className="hidden text-blue md:block">title</span>
                <span className="truncate font-semibold md:overflow-visible md:whitespace-normal md:font-normal">
                    {data?.name}
                </span>

                <span className="hidden text-blue md:block">artist</span>
                <span className="truncate md:overflow-visible md:whitespace-normal">
                    {data?.artist}
                </span>

                <span className="hidden text-blue md:block">album</span>
                <span className="truncate opacity-70 md:overflow-visible md:whitespace-normal md:opacity-100">
                    {data?.album}
                </span>

                <span className="hidden text-blue md:block">date</span>
                {data && data.timestamp > 0 ? (
                    <span className="truncate text-xs opacity-70 md:overflow-visible md:text-xl md:opacity-100">
                        {formatTimestamp(data.timestamp)}
                    </span>
                ) : (
                    <div className="mt-1 flex items-center text-xs md:mt-0 md:text-xl">
                        <div className="equalizer mr-2">
                            <span></span>
                            <span></span>
                            <span></span>
                        </div>
                        <span className="truncate">Currently listening! ^_^</span>
                    </div>
                )}
            </div>
        </div>
    );
}
