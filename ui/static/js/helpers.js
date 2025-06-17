export function zshPath(path) {
  document.querySelectorAll(".path").forEach((el) => {
    el.innerText = " " + (path === "/" ? "~" : "~" + path);
  });
}

export async function loadZshPrompt() {
  try {
    const response = await fetch("/templates/zsh");
    const html = await response.text();

    return html;
  } catch (error) {
    console.error("Failed to load zsh template:", error);

    return null;
  }
}
