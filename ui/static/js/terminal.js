import { zshPath } from "./helpers.js";

export function terminalPage(path) {
  zshPath(path);
  checkKeyboardShortcut(path);
}

function createZshPrompt() {
  // Root container
  const container = document.createElement("div");
  container.classList.add("p-5");

  // Line 1: username@host path
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

  // Line 2: input prompt
  const inputLine = document.createElement("div");
  inputLine.classList.add("flex");

  const label = document.createElement("label");
  label.setAttribute("for", "fname");
  label.textContent = ">";

  const input = document.createElement("input");
  input.classList.add("pl-5");
  input.type = "text";
  input.id = "fname";
  input.name = "fname";

  const cursor = document.createElement("span");
  cursor.classList.add("blinking-cursor");

  inputLine.append(label, input, cursor);

  container.append(userLine, inputLine);

  return container;
}

function createTerminalDiv() {
  const terminal = document.createElement("div");
  terminal.classList.add("terminal");
  terminal.classList.add("w-100-p");

  terminal.appendChild(createZshPrompt());

  return terminal;
}

function checkKeyboardShortcut(path) {
  let keyBuffer = "";
  let activeTerminal = null;

  document.addEventListener("mouseover", function onEvent(event) {
    if (event.target.classList.contains("terminal")) {
      activeTerminal = event.target;
    }
  });

  document.addEventListener("keyup", function onEvent(event) {
    // Make sure you don't check shortcuts when the user is typing
    if (event.target.localName !== "body") {
      return;
    }

    keyBuffer += event.key.toLowerCase();

    if (keyBuffer.length > 2) {
      keyBuffer = keyBuffer.slice(-2);
    }

    if (keyBuffer === "ot") {
      keyBuffer = "";

      const main = document.getElementById("main");
      const terminal = createTerminalDiv();

      main.appendChild(terminal);

      zshPath(path);
    } else if (keyBuffer === "qt" && activeTerminal) {
      activeTerminal.classList.add("closing");

      activeTerminal.addEventListener(
        "animationend",
        () => {
          activeTerminal.remove();
          activeTerminal = null;
        },
        { once: true },
      );

      keyBuffer = "";
    }
  });
}
