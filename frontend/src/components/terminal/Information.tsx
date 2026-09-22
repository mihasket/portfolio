import Zsh from './Zsh';

export default function Information() {
    return (
        <>
            <Zsh command="myself" />
            <div className="flex">
                <a href="https://github.com/mihasket/" target="_blank">
                    @github
                </a>
                <span className="text-flamingo ml-3"></span>
            </div>

            <div className="flex">
                <a href="mailto:sket.miha@proton.me" target="_blank">
                    @mail
                </a>
                <span className="text-flamingo ml-3">󰁥</span>
            </div>
            <div className="py-5">
                <samp>Software developer from Maribor, Slovenia.</samp>
                <br />
                <samp>Currently learning Go, Rust and C.</samp>
                <br />
                <samp>Love lifting weights and reading.</samp>
                <br />
            </div>

            <p className="websiteInfo">
                This website is made with Go and is designed to replicate the appearance of my Linux
                desktop environment, mainly{' '}
                <a href="https://github.com/Alexays/Waybar" target="_blank">
                    @waybar
                </a>
                . It is fully synchronized with my{' '}
                <a href="https://github.com/mihasket/dotfiles" target="_blank">
                    @dotfiles
                </a>
                , ensuring that any changes I push to my configuration are immediately reflected in
                the website’s look and feel.
            </p>

            <p className="p-0 m-0">To see my projects click on button 2 at the top of the bar!</p>

            <p className="p-0 m-0">
                Or simply visit <a href="/projects">@projects</a>
            </p>
        </>
    );
}
