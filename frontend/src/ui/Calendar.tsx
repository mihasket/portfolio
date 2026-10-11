import { useEffect, useState } from 'react';

interface DayData {
    date: string;
    intensity: string;
    count: number;
}

interface CalendarData {
    github: {
        total: number;
        contributions: DayData[][];
    };
}

export default function Calendar() {
    const [data, setData] = useState<CalendarData | null>(null);
    const [error, setError] = useState<string | null>(null);

    useEffect(() => {
        const controller = new AbortController();

        async function fetchCalendar() {
            try {
                const res = await fetch('/api/v1/github', { signal: controller.signal });
                if (!res.ok) {
                    throw new Error(`Request failed with status ${res.status}`);
                }
                const json: CalendarData = await res.json();
                setData(json);
            } catch (err) {
                if (err instanceof DOMException && err.name === 'AbortError') return;
                setError(err instanceof Error ? err.message : 'Unknown error');
            }
        }

        fetchCalendar();
        return () => controller.abort();
    }, []);

    if (error) return <p>Failed to load contributions: {error}</p>;
    if (!data) return null;

    return (
        <div className="contribution-calendar">
            <div className="contribution-stats">
                <p className="p-0 m-0 text-xl">
                    {data.github.total} contributions in the last year.
                </p>
                <p>hint: hover to see the number!</p>
            </div>
            <div className="calendar-container">
                <div className="calendar-grid">
                    {data.github.contributions.flat().map((day) => (
                        <div
                            key={day.date}
                            className={`day intensity-${day.intensity}`}
                            title={`${day.count} contributions on ${day.date}`}
                            data-date={day.date}
                            data-count={day.count}
                        ></div>
                    ))}
                </div>
            </div>
        </div>
    );
}
