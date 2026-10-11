import { useEffect, useState } from 'react';
import Startup from './Startup';
import Terminal from './ui/Terminal';
import Zsh from './ui/Zsh';
import StatusBar from './ui/StatusBar';
import Calendar from './ui/Calendar';
import CurrentSong from './ui/CurrentSong';

function App() {
    const [showIntro, setShowIntro] = useState(true);

    useEffect(() => {
        const timer = setTimeout(() => setShowIntro(false), 6300);
        return () => clearTimeout(timer);
    }, []);

    return (
        <main className="flex min-h-dvh w-full md:h-dvh">
            {showIntro ? (
                <div className="starfield flex justify-center items-center h-screen w-screen">
                    <Startup />
                </div>
            ) : (
                <div
                    id="content"
                    className="grain bg-wallpaper flex min-h-dvh w-full flex-col md:h-dvh"
                >
                    <StatusBar />

                    <div className="grid w-full flex-1 auto-rows-[minmax(12rem,auto)] grid-cols-1 md:min-h-0 md:grid-cols-2 md:grid-rows-[1fr_1fr] md:auto-rows-auto">
                        <div className="md:row-span-2 md:overflow-auto p-2 lg:pr-1 lg:p-1.5">
                            <Terminal>
                                <Zsh command="info" />

                                <div className="text-flamingo text-xl flex-center">
                                    <a href="https://github.com/mihasket/" target="_blank">
                                        @github
                                    </a>
                                    <span className="text-flamingo pl-3"></span>
                                </div>

                                <div className="text-flamingo text-xl flex-center">
                                    <a href="mailto:sket.miha@proton.me" target="_blank">
                                        @mail
                                    </a>
                                    <span className="pl-3">󰁥</span>
                                </div>

                                <div className="text-xl pt-4">
                                    <p>Software developer from Maribor, Slovenia.</p>
                                    <p>Currently learning Go, Rust and C.</p>
                                    <p>Love lifting weights and reading.</p>
                                </div>

                                <p className="text-xl pt-4">
                                    This website is made with Go and React
                                </p>

                                <Zsh command="github-contribution" />

                                <Calendar />

                                <Zsh command="" />
                            </Terminal>
                        </div>

                        <div className="md:overflow-auto p-2 lg:pl-1 lg:p-1.5">
                            <Terminal>
                                <Zsh command="currentsong" />

                                <CurrentSong />

                                <Zsh command="" />
                            </Terminal>
                        </div>

                        <div className="contents md:grid md:min-h-0 md:grid-cols-2 md:grid-rows-1">
                            <div className="md:overflow-auto p-2 lg:p-1.5 lg:pt-1 lg:pr-1">
                                <Terminal>
                                    <Zsh command="pfetch" />
                                </Terminal>
                            </div>
                            <div className="md:overflow-auto p-2 lg:p-1.5 lg:pt-1 lg:pl-1">
                                <Terminal>
                                    <Zsh command="slime-gif" />
                                </Terminal>
                            </div>
                        </div>
                    </div>
                </div>
            )}
        </main>
    );
}

export default App;
