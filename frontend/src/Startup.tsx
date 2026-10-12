export default function Startup() {
    const time = new Date().getFullYear();

    return (
        <div id="intro-wrapper">
            <h1
                id="title"
                className="flex max-w-full flex-col items-center px-4 text-center text-[48px] lg:text-[72px]"
            >
                Miha Sket
            </h1>
            <p id="subtitle" className="text-white text-xl lg:text-2xl text-center">
                est. {time}
            </p>
        </div>
    );
}
