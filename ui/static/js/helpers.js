export function zshPath(path) {
  document.querySelectorAll(".path").forEach((el) => {
    el.innerText = " " + (path === "/" ? "~" : path);
  });
}
