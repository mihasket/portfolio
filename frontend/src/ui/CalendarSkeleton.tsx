export function CalendarSkeleton() {
    return (
        <div className="contribution-calendar" aria-busy="true" aria-label="Loading contributions">
            <div className="contribution-stats">
                <p className="p-0 m-0 text-xl w-fit rounded bg-white/10 text-transparent select-none animate-pulse">
                    000 contributions in the last year.
                </p>
                <p className="w-fit rounded bg-white/10 text-transparent select-none animate-pulse mt-1">
                    hint: hover to see the number!
                </p>
            </div>
            <div className="calendar-container">
                <div className="calendar-grid animate-pulse">
                    {Array.from({ length: 53 * 7 }, (_, i) => (
                        <div key={i} className="day intensity-0 cursor-default" />
                    ))}
                </div>
            </div>
        </div>
    );
}
