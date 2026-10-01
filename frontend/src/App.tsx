import { useEffect, useState } from 'react';

function App() {
    const time = new Date().getFullYear();
    const [showIntro, setShowIntro] = useState(true);

    useEffect(() => {
        const timer = setTimeout(() => setShowIntro(false), 6300);
        return () => clearTimeout(timer);
    }, []);

    return (
        <div className="flex justify-center items-center h-screen">
            {showIntro ? (
                <div id="intro-wrapper">
                    <h1 id="title">Hello World</h1>
                    <p id="subtitle" className="text-white text-2xl text-center">
                        est. {time}
                    </p>
                </div>
            ) : (
                <div>
                    <p className="text-white text-2xl text-center">testing</p>
                </div>
            )}
        </div>
    );
}

export default App;
