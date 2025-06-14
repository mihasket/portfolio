import { createZshPrompt } from "./helpers.js";

class Node {
  constructor(type = "terminal", splitDirection = null) {
    this.type = type; // 'terminal' or 'container'
    this.splitDirection = splitDirection; // 'horizontal' or 'vertical' for containers
    this.dom = null; // DOM element for terminals
    this.left = null;
    this.right = null;
    this.parent = null;
  }
}

export class WindowManager {
  constructor(containerId) {
    this.container = document.getElementById(containerId);
    this.root = null;
    this.activeNode = null;
    this.terminalCount = 0;
    this.MAX_TERMINALS = 4;

    this.setupKeyboardShortcuts();
  }

  findNodeByElement(node, element) {
    if (!node) return null;

    if (node.type === "terminal" && node.dom === element) {
      return node;
    }

    const leftResult = this.findNodeByElement(node.left, element);
    if (leftResult) return leftResult;

    return this.findNodeByElement(node.right, element);
  }

  setActiveTerminal(terminal) {
    document.querySelectorAll(".terminal").forEach((t) => {
      t.style.borderColor = "#a4a9b3";
    });

    terminal.style.borderColor = "#88c0d0";
    terminal.focus();
    this.activeNode = this.findNodeByElement(this.root, terminal);
  }

  createTerminal() {
    if (this.terminalCount >= this.MAX_TERMINALS) {
      console.log("Maximum terminals reached");
      return;
    }

    const terminal = document.createElement("div");
    terminal.classList.add("terminal", "w-100-p");

    terminal.appendChild(createZshPrompt());

    this.terminalCount++;
    return terminal;
  }

  addTerminal() {
    const terminal = this.createTerminal();
    if (!terminal) return;

    const node = new Node("terminal");
    node.dom = terminal;

    if (!this.root) {
      this.root = node;
      this.container.appendChild(terminal);
      this.activeNode = node;
    }
  }

  closeActiveTerminal() {
    if (!this.activeNode || this.activeNode.type !== "terminal") {
      return;
    }

    const nodeToRemove = this.activeNode;
    this.removeNode(nodeToRemove);

    if (nodeToRemove.dom) {
      nodeToRemove.dom.classList.add("closing");

      const removeTerminal = () => {
        if (nodeToRemove.dom) {
          nodeToRemove.dom.remove();
        }
      };

      const onAnimationEnd = () => {
        removeTerminal();
        clearTimeout(fallbackTimeout);
      };

      nodeToRemove.dom.addEventListener("animationend", onAnimationEnd, {
        once: true,
      });
      nodeToRemove.dom.addEventListener("animationcancel", onAnimationEnd, {
        once: true,
      });

      // Fallback: remove after animation duration + buffer
      const fallbackTimeout = setTimeout(() => {
        nodeToRemove.dom.removeEventListener("animationend", onAnimationEnd);
        nodeToRemove.dom.removeEventListener("animationcancel", onAnimationEnd);
        removeTerminal();
      }, 500);

      this.terminalCount--;
    }
  }

  removeNode(node) {
    if (!node.parent) {
      this.root = null;
      return;
    }
  }

  setupKeyboardShortcuts() {
    let keyBuffer = "";

    document.addEventListener("keyup", (event) => {
      // Only process shortcuts when body is focused
      if (event.target.localName !== "body") {
        return;
      }

      keyBuffer += event.key.toLowerCase();

      if (keyBuffer.length > 2) {
        keyBuffer = keyBuffer.slice(-2);
      }

      this.handleShortcut(keyBuffer);

      // Clear buffer after processing
      if (keyBuffer.length === 2) {
        keyBuffer = "";
      }
    });
  }

  handleShortcut(keys) {
    switch (keys) {
      case "ot": // Open terminal
        this.addTerminal();
        break;
      case "qt": // Quit terminal
        this.closeActiveTerminal();
        break;
    }
  }
}
