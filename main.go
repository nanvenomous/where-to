package main

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/urfave/cli/v3"
)

func main() {
	app := &cli.Command{
		Name:    "demo",
		Usage:   "A demo application showcasing urfave/cli v3 features",
		Version: "1.0.0",
		Authors: []any{
			"Demo Author <demo@example.com>",
		},
		Copyright: "Copyright (c) 2024 Demo Inc.",
		Commands: []*cli.Command{
			{
				Name:    "greet",
				Aliases: []string{"g", "hello"},
				Usage:   "Greet someone with optional customization",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "name",
						Aliases:  []string{"n"},
						Value:    "World",
						Usage:    "Name to greet",
						Required: false,
					},
					&cli.StringFlag{
						Name:    "lang",
						Aliases: []string{"l"},
						Value:   "en",
						Usage:   "Language for greeting (en, es, fr)",
					},
					&cli.BoolFlag{
						Name:    "uppercase",
						Aliases: []string{"u"},
						Usage:   "Convert greeting to uppercase",
					},
					&cli.IntFlag{
						Name:    "count",
						Aliases: []string{"c"},
						Value:   1,
						Usage:   "Number of times to repeat the greeting",
					},
				},
				Action: greetAction,
			},
			{
				Name:  "math",
				Usage: "Mathematical operations",
				Commands: []*cli.Command{
					{
						Name:  "add",
						Usage: "Add two numbers",
						Flags: []cli.Flag{
							&cli.Float64Flag{
								Name:     "first",
								Aliases:  []string{"a"},
								Usage:    "First number",
								Required: true,
							},
							&cli.Float64Flag{
								Name:     "second",
								Aliases:  []string{"b"},
								Usage:    "Second number",
								Required: true,
							},
						},
						Action: addAction,
					},
					{
						Name:  "factorial",
						Usage: "Calculate factorial of a number",
						Flags: []cli.Flag{
							&cli.IntFlag{
								Name:     "number",
								Aliases:  []string{"n"},
								Usage:    "Number to calculate factorial for",
								Required: true,
							},
						},
						Action: factorialAction,
					},
				},
			},
			{
				Name:  "file",
				Usage: "File operations",
				Commands: []*cli.Command{
					{
						Name:      "create",
						Usage:     "Create a new file",
						ArgsUsage: "[filename]",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:    "content",
								Aliases: []string{"c"},
								Usage:   "Content to write to file",
								Value:   "Hello, World!",
							},
							&cli.BoolFlag{
								Name:  "force",
								Usage: "Overwrite existing file",
							},
						},
						Action: createFileAction,
					},
					{
						Name:      "info",
						Usage:     "Show file information",
						ArgsUsage: "[filename]",
						Action:    fileInfoAction,
					},
				},
			},
			{
				Name:  "config",
				Usage: "Configuration management with validation",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "host",
						Usage:    "Server host",
						Required: true,
						Validator: func(host string) error {
							if len(host) < 3 {
								return fmt.Errorf("host must be at least 3 characters long")
							}
							return nil
						},
					},
					&cli.IntFlag{
						Name:  "port",
						Usage: "Server port",
						Value: 8080,
						Validator: func(port int) error {
							if port < 1 || port > 65535 {
								return fmt.Errorf("port must be between 1 and 65535")
							}
							return nil
						},
					},
					&cli.BoolFlag{
						Name:  "ssl",
						Usage: "Enable SSL",
					},
				},
				Action: configAction,
			},
		},
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "verbose",
				Usage: "Enable verbose output",
			},
			&cli.StringFlag{
				Name:  "output",
				Usage: "Output format (json, yaml, text)",
				Value: "text",
			},
		},
		Before: func(ctx context.Context, cmd *cli.Command) (context.Context, error) {
			if cmd.Bool("verbose") {
				fmt.Println("Verbose mode enabled")
			}
			return ctx, nil
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			fmt.Printf("Welcome to the urfave/cli v3 demo app!\n")
			fmt.Printf("Version: %s\n", cmd.Root().Version)
			fmt.Printf("Use '%s help' to see available commands.\n", cmd.Root().Name)
			return nil
		},
	}

	if err := app.Run(context.Background(), os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func greetAction(ctx context.Context, cmd *cli.Command) error {
	name := cmd.String("name")
	lang := cmd.String("lang")
	uppercase := cmd.Bool("uppercase")
	count := cmd.Int("count")

	greetings := map[string]string{
		"en": "Hello",
		"es": "Hola",
		"fr": "Bonjour",
	}

	greeting, exists := greetings[lang]
	if !exists {
		return fmt.Errorf("unsupported language: %s", lang)
	}

	message := fmt.Sprintf("%s, %s!", greeting, name)
	if uppercase {
		message = strings.ToUpper(message)
	}

	for range count {
		fmt.Println(message)
	}

	return nil
}

func addAction(ctx context.Context, cmd *cli.Command) error {
	a := cmd.Float64("first")
	b := cmd.Float64("second")
	result := a + b

	fmt.Printf("%.2f + %.2f = %.2f\n", a, b, result)
	return nil
}

func factorialAction(ctx context.Context, cmd *cli.Command) error {
	n := cmd.Int("number")

	if n < 0 {
		return fmt.Errorf("factorial is not defined for negative numbers")
	}

	if n > 20 {
		return fmt.Errorf("factorial too large for display (max: 20)")
	}

	result := factorial(n)
	fmt.Printf("%d! = %d\n", n, result)
	return nil
}

func factorial(n int) int64 {
	if n <= 1 {
		return 1
	}
	return int64(n) * factorial(n-1)
}

func createFileAction(ctx context.Context, cmd *cli.Command) error {
	if cmd.Args().Len() == 0 {
		return fmt.Errorf("filename is required")
	}

	filename := cmd.Args().First()
	content := cmd.String("content")
	force := cmd.Bool("force")

	if _, err := os.Stat(filename); err == nil && !force {
		return fmt.Errorf("file '%s' already exists (use --force to overwrite)", filename)
	}

	if err := os.WriteFile(filename, []byte(content), 0644); err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}

	fmt.Printf("File '%s' created successfully\n", filename)
	return nil
}

func fileInfoAction(ctx context.Context, cmd *cli.Command) error {
	if cmd.Args().Len() == 0 {
		return fmt.Errorf("filename is required")
	}

	filename := cmd.Args().First()
	info, err := os.Stat(filename)
	if err != nil {
		return fmt.Errorf("failed to get file info: %v", err)
	}

	fmt.Printf("File: %s\n", filename)
	fmt.Printf("Size: %d bytes\n", info.Size())
	fmt.Printf("Mode: %s\n", info.Mode())
	fmt.Printf("Modified: %s\n", info.ModTime().Format(time.RFC3339))
	fmt.Printf("IsDir: %t\n", info.IsDir())

	return nil
}

func configAction(ctx context.Context, cmd *cli.Command) error {
	host := cmd.String("host")
	port := cmd.Int("port")
	ssl := cmd.Bool("ssl")

	protocol := "http"
	if ssl {
		protocol = "https"
	}

	fmt.Printf("Configuration:\n")
	fmt.Printf("  Host: %s\n", host)
	fmt.Printf("  Port: %d\n", port)
	fmt.Printf("  SSL: %t\n", ssl)
	fmt.Printf("  URL: %s://%s:%d\n", protocol, host, port)

	return nil
}
