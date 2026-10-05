import { useEffect, useState } from 'react';
import Startup from './Startup';

function App() {
    const [showIntro, setShowIntro] = useState(true);

    useEffect(() => {
        const timer = setTimeout(() => setShowIntro(false), 6300);
        return () => clearTimeout(timer);
    }, []);

    return (
        <main className="starfield flex justify-center items-center h-screen">
            {showIntro ? (
                <Startup />
            ) : (
                <div className="content flex">
                    <Startup isIntro={false} />
                    <p className=" content text-2xl">menu</p>
                </div>
            )}
        </main>
    );
}

export default App;
