export default function Startup() {
    const time = new Date().getFullYear();

    return (
        <div id="intro-wrapper">
            <h1 id="title">Hello World</h1>
            <p id="subtitle" className="text-white text-2xl text-center">
                est. {time}
            </p>
        </div>
    );
}
