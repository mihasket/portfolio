interface ZshProps {
    command: string;
    path?: string;
}

export default function Zsh({ command, path }: ZshProps) {
    const displayPath = !path || path === '/' ? '~' : '~' + path;

    return (
        <>
            <div className="py-2">
                <div className="flex p-0 m-0">
                    <span className="text-3xl text-blue">{' ' + displayPath}</span>
                </div>
                <div className="flex items-center">
                    <span className="text-maroon">&nbsp;</span>
                    <span className="terminal-input text-green">{command}</span>
                    <span className="cursor"></span>
                </div>
            </div>
        </>
    );
}
