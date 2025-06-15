export function zshPath(path) {
  document.querySelectorAll(".path").forEach((el) => {
    el.innerText = " " + (path === "/" ? "~" : path);
  });
}

export function createZshPrompt() {
  const container = document.createElement("div");
  container.classList.add("py-2");

  const userLine = document.createElement("div");
  userLine.classList.add("flex", "p-0", "m-0");

  const userSpan = document.createElement("span");
  userSpan.classList.add("yellow");
  userSpan.textContent = "miha";

  const atSpan = document.createElement("span");
  atSpan.textContent = "@";

  const hostSpan = document.createElement("span");
  hostSpan.classList.add("orange");
  hostSpan.textContent = "mp3";

  const pathSpan = document.createElement("span");
  pathSpan.classList.add("path", "blue", "pl-5");

  userLine.append(userSpan, atSpan, hostSpan, pathSpan);

  const inputLine = document.createElement("div");
  inputLine.classList.add("flex", "items-center");

  const label = document.createElement("span");
  label.textContent = ">";

  const cursor = document.createElement("span");
  cursor.classList.add("blinking-cursor");

  inputLine.append(label, cursor);

  container.append(userLine, inputLine);

  return container;
}
