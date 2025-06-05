function homePage(path) {
    document.querySelectorAll('.path').forEach(el => {
        el.innerText = " " + (path === "/" ? "~" : path);
    });
}

function projectsPage() {
    const now = new Date();
    const date = now.toISOString().split('T')[0];

    const hours = String(now.getHours()).padStart(2, '0');
    const minutes = String(now.getMinutes()).padStart(2, '0');
    const time = `${hours}:${minutes}`;

    document.getElementById("tmuxDate").innerText = `${date} | ${time}`;

    const projectButtons = document.querySelectorAll('.projectButton');

    projectButtons.forEach(button => {
        document.getElementById(button.id).addEventListener('click', () => {
            const contents = document.querySelectorAll('.project-content');
            contents.forEach(div => div.classList.remove('active'));

            const selected = document.getElementById(button.id.replace(".md", ""));
            if (selected) {
                selected.classList.add('active');
            }
        });
    });
}

// Main function
document.addEventListener('DOMContentLoaded', function() {
    const path = window.location.pathname;

    if (path === '/') {
        homePage(path);
    }
    else if (path === '/projects/') {
        projectsPage();
    }
});
