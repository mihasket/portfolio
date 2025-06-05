// TODO: Organize this better
const path = window.location.pathname;

document.querySelectorAll('.path').forEach(el => {
    el.innerText = " " + (path === "/" ? "~" : path);
});

if (path === '/projects') {
    const now = new Date();
    const date = now.toISOString().split('T')[0];

    const hours = String(now.getHours()).padStart(2, '0');
    const minutes = String(now.getMinutes()).padStart(2, '0');
    const time = `${hours}:${minutes}`;

    document.getElementById("tmuxDate").innerText = `${date} | ${time}`;
}

function showProject(projectName) {
    const contents = document.querySelectorAll('.project-content');
    contents.forEach(div => div.classList.remove('active'));

    const selected = document.getElementById(projectName.replace(".md", ""));
    if (selected) {
        selected.classList.add('active');
    }
}

document.addEventListener('DOMContentLoaded', function() {
    document.querySelectorAll('.projectButton').forEach(el => {
        document.getElementById(el.id).addEventListener('click', () => showProject(el.id));
    });
});
