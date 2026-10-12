import ProjectItem from './ProjectItem';

const projects = [
    {
        name: 'Car Repair Shop',
        description:
            'Complete redesign and rebuild of a legacy car repair shop website with a modern, responsive design and improved performance.',
        href: 'https://www.pomer.si/',
    },
    {
        name: 'HTTP 1.1 implementation',
        description:
            'Minimal implementation of an HTTP/1.1 server built directly on top of go TCP sockets.',
        href: 'https://github.com/mihasket/http-server/',
    },
    {
        name: 'Chip-8 Interpreter',
        description: 'A Rust program for a CHIP-8 emulator simulates the CHIP-8 virtual machine',
        href: 'https://github.com/mihasket/chip8/',
    },
    {
        name: 'Cozy',
        description: 'Desktop music player client built on MPD, built with Electron',
        href: 'https://github.com/mihasket/cozy/',
    },
    {
        name: 'mp3-tagger',
        description:
            'A C++ program for ID3v2.3 MP3 tagging modifies, or writes metadata in MP3 audio files according to the ID3v2.3 specification',
        href: 'https://github.com/mihasket/mp3-tagger/',
    },
];

export default function Projects() {
    return (
        <section>
            <h2 className="m-0 mb-3 flex items-center gap-2 text-xl font-semibold text-flamingo">
                Projects:
                <span className="h-px flex-1 bg-white/10" aria-hidden="true"></span>
            </h2>
            <ul className="m-0 list-none space-y-3 p-0">
                {projects.map((p) => (
                    <ProjectItem key={p.name} {...p} />
                ))}
            </ul>
        </section>
    );
}
