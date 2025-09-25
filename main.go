package main

import (
	_ "embed"
	"fmt"
	"os"
)

//go:embed scripts/nav-functions.sh
var shellFunctions string

const version = "1.0.2"

func main() {
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "init":
			fmt.Print(shellFunctions)
		case "version", "--version", "-v":
			fmt.Printf("nav %s\n", version)
		case "help", "--help", "-h":
			showHelp()
		default:
			fmt.Fprintf(os.Stderr, "Unknown command: %s\n", os.Args[1])
			os.Exit(1)
		}
	} else {
		// Default behavior when run without arguments
		fmt.Print(shellFunctions)
	}
}

func showHelp() {
	fmt.Printf("Terminal Navigation Helper v%s\n\n", version)
	fmt.Printf("Usage:\n")
	fmt.Printf("  nav              Output shell functions for evaluation\n")
	fmt.Printf("  nav init         Output shell functions for evaluation (alias)\n")
	fmt.Printf("  nav version      Show version information\n")
	fmt.Printf("  nav help         Show this help message\n\n")
	fmt.Printf("To set up navigation functions in your shell:\n")
	fmt.Printf("  eval \"$(nav)\"             # for current session\n")
	fmt.Printf("  nav >> ~/.bashrc          # for bash\n")
	fmt.Printf("  nav >> ~/.zshrc           # for zsh\n")
	fmt.Printf("  source ~/.bashrc          # reload your shell\n\n")
	fmt.Printf("This will add these commands to your shell:\n")
	fmt.Printf("  up [levels] - Move up directories (default: 1)\n")
	fmt.Printf("  dn <dir>    - Move down into directory\n")
	fmt.Printf("  to          - Interactive navigation with fzf\n")
	fmt.Printf("  t           - Clear and list current directory\n\n")
	fmt.Printf("Each command clears the terminal and lists directory contents.\n")
}
