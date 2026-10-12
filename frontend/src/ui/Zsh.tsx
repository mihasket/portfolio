interface ZshProps {
    command: string;
}

export default function Zsh({ command }: ZshProps) {
    return (
        <div>
            <div className="py-2 text-xl">
                <div className="flex p-0 m-0">
                    <span className="text-blue">~/.config/user/</span>
                </div>
                <div className="flex items-center">
                    <span className="text-maroon">&nbsp;</span>
                    <span className="text-green">{command}</span>
                    <span className="cursor"></span>
                </div>
            </div>
        </div>
    );
}
