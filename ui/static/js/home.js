export function homePage(path) {
  document.querySelectorAll(".path").forEach((el) => {
    el.innerText = " " + (path === "/" ? "~" : "~" + path);
  });
}
