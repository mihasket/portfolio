export class Terminal {
  constructor(inputId, bodyId, cursorId) {
    this.input = document.getElementById(inputId);
    this.body = document.getElementById(bodyId);
    this.cursor = document.getElementById(cursorId);
    this.history = [];
    this.historyIndex = -1;

    this.commands = {
      slimegif: "Shows a gif of a slime!",
      github: "Shows the github contribution calendar",
      sfetch: "Shows the latest song i have listened to!",
      nvim: "Opens neovim to show my projects",
      pfetch: "Displays website information",
      help: "Show available commands",
      clear: "Clear the terminal screen",
      whoami: "Display current user",
      date: "Display current date and time",
      echo: "Display text",
      exit: "Exit the terminal",
    };
  }

  init() {
    this.input.addEventListener("keydown", (e) => {
      this.handleKeydown(e);
    });
    this.input.addEventListener("input", () => {
      this.updateCursor();
    });
    this.input.focus();

    this.updateCursor();
  }

  handleKeydown(e) {
    if (e.key === "Enter") {
      e.preventDefault();
      this.executeCommand();
    } else if (e.key === "ArrowUp") {
      e.preventDefault();
      this.navigateHistory(-1);
    } else if (e.key === "ArrowDown") {
      e.preventDefault();
      this.navigateHistory(1);
    } else if (e.key === "Tab") {
      e.preventDefault();
      this.autoComplete();
    }
  }

  executeCommand() {
    const command = this.input.value.trim();

    if (command) {
      this.history.push(command);
      this.historyIndex = this.history.length;

      this.processCommand(command);
    }

    this.input.readOnly = true;
    this.createNewInputLine();
    this.updateCursor();
  }

  processCommand(command) {
    const [cmd, ...args] = command.split(" ");

    switch (cmd.toLowerCase()) {
      case "help":
        this.showHelp();
        break;
      case "clear":
        this.clearTerminal();
        break;
      default:
        this.addToOutput(
          `Command not found: ${cmd}. Type 'help' for available commands.`,
          "error-output",
        );
        break;
    }
  }

  showHelp() {
    this.addToOutput("", "command-output");

    Object.entries(this.commands).forEach(([cmd, description]) => {
      const helpLine = document.createElement("div");
      helpLine.className = "command-output";
      helpLine.innerHTML = `<span class="help-command">${cmd.padEnd(12)}</span><span class="help-description">${description}</span>`;
      this.body.insertBefore(helpLine, this.body.lastElementChild);
    });

    this.addToOutput("", "command-output");
  }

  clearTerminal() {
    // Remove all output except the input line
    const inputLine = this.body.lastElementChild;
    this.body.innerHTML = "";
    this.body.appendChild(inputLine);
  }

  addToOutput(command, className = "command-output") {
    const outputDiv = document.createElement("div");
    outputDiv.className = className;
    outputDiv.innerHTML = command;
    this.body.appendChild(outputDiv);
    this.scrollToBottom();
  }

  createNewInputLine() {
    // Remove old input line
    const oldInputLine = this.body.lastElementChild;
    oldInputLine.remove();

    const zshPath = document.createElement("div");
    zshPath.classList.add("flex", "p-0", "m-0");

    const path = document.createElement("span");
    path.classList.add("blue");
    path.textContent = "~/mihasket/portfolio";

    const zshInput = document.createElement("div");
    zshInput.classList.add("flex", "items-center");

    const leftArrow = document.createElement("span");
    leftArrow.classList.add("maroon");
    leftArrow.textContent = "\u00A0";

    const input = document.createElement("input");
    input.id = "terminal-input";
    input.classList.add("terminal-input", "green");
    input.type = "text";
    input.setAttribute("autocomplete", "off");
    input.setAttribute("spellcheck", "false");

    const cursor = document.createElement("span");
    cursor.id = "cursor";
    cursor.classList.add("cursor");

    zshPath.appendChild(path);
    zshInput.appendChild(leftArrow);
    zshInput.appendChild(input);
    zshInput.appendChild(cursor);

    this.body.appendChild(zshPath);
    this.body.appendChild(zshInput);

    this.input = input;
    this.cursor = cursor;

    // Reattach event listeners
    input.addEventListener("keydown", (e) => this.handleKeydown(e));
    input.addEventListener("input", () => this.updateCursor());
    input.focus();
  }

  navigateHistory(direction) {
    if (this.history.length === 0) return;

    this.historyIndex += direction;

    if (this.historyIndex < 0) {
      this.historyIndex = 0;
    } else if (this.historyIndex >= this.history.length) {
      this.historyIndex = this.history.length;
      this.input.value = "";
    } else {
      this.input.value = this.history[this.historyIndex];
    }

    this.updateCursor();
  }

  autoComplete() {
    const input = this.input.value;
    const matches = Object.keys(this.commands).filter((cmd) =>
      cmd.startsWith(input.toLowerCase()),
    );

    if (matches.length === 1) {
      this.input.value = matches[0];
      this.updateCursor();
    } else if (matches.length > 1) {
      this.addToOutput(matches.join("  "), "command-output");
    }
  }

  updateCursor() {
    // Position cursor after the text
    const textWidth = this.input.value.length;
    this.cursor.style.marginLeft = `${textWidth * 0.6}em`;
  }

  scrollToBottom() {
    this.body.scrollTop = this.body.scrollHeight;
  }
}
