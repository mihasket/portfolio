import { projectsPage } from "./projects.js";
import { homePage } from "./home.js";
import { WindowManager } from "./tree.js";
import { Terminal } from "./terminal.js";

document.addEventListener("DOMContentLoaded", async function () {
  const path = window.location.pathname;

  if (path === "/") {
    homePage(path);
  } else if (path === "/projects") {
    projectsPage();
  } else if (path === "/terminal") {
    const windowManager = new WindowManager("main");
    windowManager.setupKeyboardShortcuts(path);
    await windowManager.addTerminal(path);

    const terminal = new Terminal();
    terminal.init();
  }
});
