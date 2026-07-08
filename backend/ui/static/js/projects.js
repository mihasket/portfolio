export function projectsPage() {
  updateTmuxDate();
  setupProjectButtons();
}

function updateTmuxDate() {
  const tmuxDateElement = document.getElementById("tmuxDate");
  if (!tmuxDateElement) {
    return;
  }

  const now = new Date();
  const date = now.toISOString().split("T")[0];
  const time = now.toTimeString().slice(0, 5);

  tmuxDateElement.innerText = `${date} | ${time}`;
}

function setupProjectButtons() {
  const projectButtons = document.querySelectorAll(".projectButton");

  projectButtons.forEach((button) => {
    button.addEventListener("click", handleProjectButtonClick);
  });
}

function handleProjectButtonClick(event) {
  const button = event.currentTarget;

  document.querySelectorAll(".project-content").forEach((div) => {
    div.classList.remove("active");
  });

  const projectId = button.id.replace(".md", "");
  const selectedContent = document.getElementById(projectId);

  if (selectedContent) {
    selectedContent.classList.add("active");
  }
}
