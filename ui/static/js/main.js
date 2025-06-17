import { projectsPage } from "./projects.js";
import { homePage } from "./home.js";
import { WindowManager } from "./tree.js";

document.addEventListener("DOMContentLoaded", function () {
  const path = window.location.pathname;

  if (path === "/") {
    homePage(path);
  } else if (path === "/projects") {
    projectsPage();
  } else if (path === "/terminal") {
    const windowManager = new WindowManager("main");
    windowManager.setupKeyboardShortcuts();
  }
});
