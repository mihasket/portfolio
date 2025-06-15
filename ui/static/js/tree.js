import { createZshPrompt } from "./helpers.js";

class Node {
  constructor(type = "terminal", splitDirection = null) {
    this.type = type;
    this.splitDirection = splitDirection;
    this.dom = null;
    this.left = null;
    this.right = null;
    this.parent = null;
  }
}

export class WindowManager {
  constructor(mainId) {
    this.main = document.getElementById(mainId);
    this.root = null;
    this.activeNode = null;
    this.terminalCount = 0;
    this.MAX_TERMINALS = 4;
  }

  addTerminal() {
    const terminal = this.createTerminal();
    if (!terminal) return;

    const newTerminalNode = new Node("terminal");
    newTerminalNode.dom = terminal;

    // First terminal
    if (!this.root) {
      this.root = newTerminalNode;
      this.main.appendChild(terminal);
      this.activeNode = newTerminalNode;

      this.setActiveTerminal(terminal);

      return;
    }

    const nodeToSplit = this.findRightmostTerminal(this.root);
    const splitDirection =
      this.terminalCount % 2 === 0 ? "vertical" : "horizontal";

    this.splitNode(nodeToSplit, newTerminalNode, splitDirection);
    this.activeNode = newTerminalNode;
    this.setActiveTerminal(newTerminalNode.dom);
  }

  splitNode(existingNode, newNode, splitDirection) {
    const containerNode = new Node("container", splitDirection);
    const container = this.createContainer(splitDirection);

    containerNode.parent = existingNode.parent;
    existingNode.parent = containerNode;
    newNode.parent = containerNode;

    containerNode.left = existingNode;
    containerNode.right = newNode;

    // Update the parent's reference to point to the new container
    if (containerNode.parent) {
      if (containerNode.parent.left === existingNode) {
        containerNode.parent.left = containerNode;
      } else {
        containerNode.parent.right = containerNode;
      }
    } else {
      this.root = containerNode;
    }

    const existingDOM = existingNode.dom;
    const parentDOM = existingDOM.parentNode;

    parentDOM.removeChild(existingDOM);

    container.appendChild(existingDOM);
    container.appendChild(newNode.dom);

    parentDOM.appendChild(container);

    containerNode.dom = container;
  }

  setActiveTerminal(terminal) {
    document.querySelectorAll(".terminal").forEach((t) => {
      t.style.borderColor = "#a4a9b3";
    });

    terminal.style.borderColor = "#88c0d0";
    terminal.focus();
    this.activeNode = this.findNodeByElement(this.root, terminal);
  }

  findNodeByElement(node, element) {
    if (!node) return null;

    if (node.type === "terminal" && node.dom === element) {
      return node;
    }

    const leftResult = this.findNodeByElement(node.left, element);
    if (leftResult) {
      return leftResult;
    }

    return this.findNodeByElement(node.right, element);
  }

  findRightmostTerminal(node) {
    if (!node) {
      return null;
    }

    if (node.type === "terminal") {
      return node;
    }

    return (
      this.findRightmostTerminal(node.right) ||
      this.findRightmostTerminal(node.left)
    );
  }

  createTerminal() {
    if (this.terminalCount >= this.MAX_TERMINALS) {
      console.log("Maximum terminals reached");
      return;
    }

    const terminal = document.createElement("div");
    terminal.classList.add("terminal", "flex", "flex-1");

    terminal.appendChild(createZshPrompt());

    this.terminalCount++;
    return terminal;
  }

  createContainer(splitDirection) {
    const container = document.createElement("div");
    container.classList.add("flex", "flex-1", "min-w-0-p", "min-h-0-p");

    if (splitDirection === "vertical") {
      container.classList.add("flex-row");
    } else {
      container.classList.add("flex-col");
    }

    return container;
  }

  closeActiveTerminal() {
    if (!this.activeNode || this.activeNode.type !== "terminal") {
      console.log("No active terminal to close");
      return;
    }

    this.removeNode(this.activeNode);
    this.terminalCount--;

    // Find next active terminal
    const nextTerminal = this.findRightmostTerminal(this.root);
    if (nextTerminal) {
      this.activeNode = nextTerminal;
      this.setActiveTerminal(nextTerminal.dom);
    } else {
      this.activeNode = null;
    }
  }

  removeNode(nodeToRemove) {
    if (!nodeToRemove.parent) {
      nodeToRemove.dom.remove();
      this.root = null;
      return;
    }

    const parent = nodeToRemove.parent;
    const sibling = parent.left === nodeToRemove ? parent.right : parent.left;

    nodeToRemove.dom.remove();

    if (!parent.parent) {
      this.root = sibling;
      sibling.parent = null;

      parent.dom.remove();
      this.main.appendChild(sibling.dom);
    } else {
      const grandparent = parent.parent;
      sibling.parent = grandparent;

      if (grandparent.left === parent) {
        grandparent.left = sibling;
      } else {
        grandparent.right = sibling;
      }

      const grandparentDOM = parent.dom.parentNode;
      grandparentDOM.removeChild(parent.dom);
      grandparentDOM.appendChild(sibling.dom);
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
