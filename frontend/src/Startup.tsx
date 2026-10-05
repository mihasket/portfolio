interface StartupProps {
    isIntro?: boolean;
}

export default function Startup({ isIntro = true }: StartupProps) {
    const time = new Date().getFullYear();

    return (
        <div
            id={isIntro ? 'intro-wrapper' : undefined}
            className={isIntro ? undefined : 'startup-static'}
        >
            <h1
                id={isIntro ? 'title' : undefined}
                className={isIntro ? undefined : 'startup-title-static'}
            >
                Hello World
            </h1>

            <p
                id={isIntro ? 'subtitle' : undefined}
                className={`text-white text-2xl text-center ${
                    isIntro ? '' : 'startup-subtitle-static'
                }`}
            >
                est. {time}
            </p>
        </div>
    );
}
