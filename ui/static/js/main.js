import { projectsPage } from "./projects.js";
import { homePage } from "./home.js";
import { terminalPage } from "./terminal.js";

document.addEventListener("DOMContentLoaded", function () {
  const path = window.location.pathname;

  if (path === "/") {
    homePage(path);
  } else if (path === "/projects/") {
    projectsPage();
  } else if (path === "/terminal/") {
    terminalPage(path);
  }
});
