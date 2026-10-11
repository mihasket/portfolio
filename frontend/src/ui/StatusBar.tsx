import { useState } from 'react';

function formatStatusDate(date: Date): string {
    const parts = new Intl.DateTimeFormat('en-US', {
        weekday: 'short',
        month: 'short',
        day: 'numeric',
        hour: '2-digit',
        minute: '2-digit',
        hourCycle: 'h23',
    }).formatToParts(date);

    const get = (type: Intl.DateTimeFormatPartTypes): string | undefined =>
        parts.find((p) => p.type === type)?.value;

    return `${get('weekday')} ${get('month')} ${get('day')} ${get('hour')}:${get('minute')}`;
}

export default function StatusBar() {
    const [date] = useState<string>(() => formatStatusDate(new Date()));

    return (
        <header className="bg-statusbar sticky top-0 z-10 flex h-10 w-full shrink-0 items-center justify-between md:static text-xl">
            <div className="px-4">
                <span className="pr-3">󰌽</span>
                <span>Home</span>
                <span className="hidden md:inline">- Miha Šket</span>
            </div>
            <div className="px-4">
                <span className="hidden pr-4 md:inline"></span>
                <span className="hidden pr-4 md:inline">󰖩</span>
                <span className="pr-3"></span>
                <span>{date}</span>
            </div>
        </header>
    );
}
