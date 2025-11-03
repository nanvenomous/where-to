package main

import (
	_ "embed"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

//go:embed scripts/nav-functions.sh
var shellFunctions string

//go:embed version
var version string

type Config struct {
	GitStatusCommand *string `yaml:"gitStatusCommand"`
}

func ptr[T any](thg T) *T {
	return &thg
}

func getDefaultConfig() Config {
	return Config{
		GitStatusCommand: ptr("git status -sb"),
	}
}

func getConfigPath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(homeDir, ".config", "where-to", "config.yaml"), nil
}

func loadConfig() Config {
	configPath, err := getConfigPath()
	if err != nil {
		return Config{} // Return empty config if path error
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		// If config doesn't exist, return empty config (disabled by default)
		return Config{}
	}

	// Unmarshal into empty config to detect what was actually set
	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		fmt.Fprintf(os.Stderr, "Warning: Failed to parse config file: %v\n", err)
		return Config{}
	}

	// Return config as-is (nil = disabled, value = use that value)
	return config
}

func createDefaultConfig() error {
	configPath, err := getConfigPath()
	if err != nil {
		return err
	}

	// Check if config already exists
	if _, err := os.Stat(configPath); err == nil {
		return fmt.Errorf("config file already exists at %s", configPath)
	}

	// Create config directory if it doesn't exist
	configDir := filepath.Dir(configPath)
	if err := os.MkdirAll(configDir, 0755); err != nil {
		return err
	}

	// Create default config
	config := getDefaultConfig()
	data, err := yaml.Marshal(config)
	if err != nil {
		return err
	}

	if err := os.WriteFile(configPath, data, 0644); err != nil {
		return err
	}

	fmt.Printf("Created default config at %s\n", configPath)
	return nil
}

func generateShellFunctions(config Config) string {
	result := shellFunctions

	// Replace the git status command placeholder
	if config.GitStatusCommand == nil || *config.GitStatusCommand == "" {
		// If nil or empty, disable git status checks
		result = strings.ReplaceAll(result, "{{GIT_STATUS_COMMAND}}", "")
	} else {
		result = strings.ReplaceAll(result, "{{GIT_STATUS_COMMAND}}", *config.GitStatusCommand)
	}

	return result
}

func main() {
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "init":
			config := loadConfig()
			fmt.Print(generateShellFunctions(config))
		case "config":
			if err := createDefaultConfig(); err != nil {
				fmt.Fprintf(os.Stderr, "Error: %v\n", err)
				os.Exit(1)
			}
		case "version", "--version", "-v":
			fmt.Printf("where-to %s\n", strings.TrimSpace(version))
		case "help", "--help", "-h":
			showHelp()
		default:
			fmt.Fprintf(os.Stderr, "Unknown command: %s\n", os.Args[1])
			os.Exit(1)
		}
	} else {
		// Default behavior when run without arguments
		config := loadConfig()
		fmt.Print(generateShellFunctions(config))
	}
}

func showHelp() {
	configPath, _ := getConfigPath()
	fmt.Printf("Terminal Navigation Helper v%s\n\n", strings.TrimSpace(version))
	fmt.Printf("Usage:\n")
	fmt.Printf("  where-to              Output shell functions for evaluation\n")
	fmt.Printf("  where-to init         Output shell functions for evaluation (alias)\n")
	fmt.Printf("  where-to config       Create default config file\n")
	fmt.Printf("  where-to version      Show version information\n")
	fmt.Printf("  where-to help         Show this help message\n\n")
	fmt.Printf("To set up navigation functions in your shell:\n")
	fmt.Printf("  eval \"$(where-to)\"             # for current session\n")
	fmt.Printf("  where-to >> ~/.bashrc          # for bash\n")
	fmt.Printf("  where-to >> ~/.zshrc           # for zsh\n")
	fmt.Printf("  source ~/.bashrc          # reload your shell\n\n")
	fmt.Printf("This will add these commands to your shell:\n")
	fmt.Printf("  up [levels] - Move up directories (default: 1)\n")
	fmt.Printf("  dn <dir>    - Move down into directory\n")
	fmt.Printf("  to          - Interactive navigation with fzf\n")
	fmt.Printf("  t [dir]     - Clear and list current/specified directory\n\n")
	fmt.Printf("Configuration:\n")
	fmt.Printf("  Config file: %s\n", configPath)
	fmt.Printf("  Run 'where-to config' to create an example config file.\n")
	fmt.Printf("  Add gitStatusCommand to enable git status display in repositories.\n")
	fmt.Printf("  Example: gitStatusCommand: 'git status -sb'\n\n")
	fmt.Printf("Each command clears the terminal and lists directory contents.\n")
}
