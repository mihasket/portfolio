import Zsh from '../terminal/Zsh';

export default function Home() {
    return (
        <div className="flex-1 flex w-full h-full overflow-y-auto">
            <div className="flex w-full">
                <div className="terminal w-full">
                    {/*<Information title="desktop" />*/}

                    <Zsh command="github-contribution" />

                    {/*<GithubCalendar data={githubCalendar} />*/}

                    <Zsh command="" />
                </div>
            </div>

            <div className="flex flex-col w-full">
                <div className="terminal h-full">
                    <Zsh command="sfetch" />

                    {/*<LatestTrack data={lastFM} />*/}

                    <Zsh command="" />
                </div>

                <div className="flex h-full">
                    <div className="terminal w-full">
                        <Zsh command="pfetch" />
                        <pre className="text-white">{`   __
-=(o '.       miha@mp3
   '.-.\\      os      Golang
   /|  \\\\     host    miha
   '|  ||     ascii   snd
    _\\_):,_   bar     waybar
`}</pre>
                        <Zsh command="" />
                    </div>

                    <div className="terminal w-full">
                        <Zsh command="slime-gif" />
                        <img className="max-w-full" src="/img/slime.gif" alt="Moving slime" />
                    </div>
                </div>
            </div>
        </div>
    );
}
