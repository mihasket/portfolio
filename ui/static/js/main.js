import { projectsPage } from "./projects.js";
import { homePage } from "./home.js";

document.addEventListener("DOMContentLoaded", async function () {
  const path = window.location.pathname;

  if (path === "/") {
    homePage(path);
  } else if (path === "/projects") {
    projectsPage();
  }
});
