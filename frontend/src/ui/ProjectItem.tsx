interface ProjectItemProps {
    name: string;
    description: string;
    href: string;
}

export default function ProjectItem({ name, description, href }: ProjectItemProps) {
    return (
        <li className="group flex gap-3 rounded px-2 py-1 transition-colors hover:bg-white/5">
            <span className="text-flamingo transition-transform group-hover:translate-x-1">›</span>
            <div className="min-w-0">
                <a
                    href={href}
                    target="_blank"
                    rel="noreferrer"
                    className="text-flamingo underline-offset-4 hover:underline"
                >
                    @{name}
                </a>
                <p className="m-0 text-sm opacity-70">{description}</p>
            </div>
        </li>
    );
}
