# urfave/cli v3 Demo Application

A comprehensive demonstration of [urfave/cli v3](https://github.com/urfave/cli) features including commands, subcommands, flags, validation, and more.

## Features Demonstrated

### 🏗️ Basic Application Structure
- Root command with metadata (name, version, authors, copyright)
- Help system and version information
- Global flags accessible across all commands

### 🔧 Commands & Subcommands
- **greet**: Simple greeting command with multiple options
- **math**: Mathematical operations with subcommands
  - `add`: Add two floating-point numbers
  - `factorial`: Calculate factorial with validation
- **file**: File operations
  - `create`: Create files with content and overwrite protection
  - `info`: Display file information and statistics
- **config**: Configuration management with built-in validation

### 🚩 Flag Types & Features
- **String flags**: Text input with default values and aliases
- **Integer flags**: Numeric input with validation
- **Float flags**: Decimal number support
- **Boolean flags**: Simple on/off switches
- **Count flags**: Repeat counting (factorial example)
- **Required flags**: Mandatory parameters with validation
- **Flag aliases**: Short and long form options (`-n`, `--name`)
- **Flag validation**: Custom validation functions with error messages

### ✨ Advanced Features
- **Before hooks**: Execute code before command processing
- **Command aliases**: Multiple ways to invoke commands (`greet`, `g`, `hello`)
- **Argument validation**: Custom validation with helpful error messages
- **Error handling**: Graceful error reporting and exit codes
- **Help generation**: Automatic help text generation
- **Shell completion**: Built-in shell completion support

## Usage Examples

### Basic Greeting
```bash
./demo greet --name "World" --lang en
# Output: Hello, World!
```

### Advanced Greeting Options
```bash
./demo greet --name "CLI User" --lang es --uppercase --count 3
# Output: HOLA, CLI USER! (repeated 3 times)
```

### Mathematical Operations
```bash
./demo math add --first 15.5 --second 24.3
# Output: 15.50 + 24.30 = 39.80

./demo math factorial --number 6
# Output: 6! = 720
```

### File Operations
```bash
./demo file create myfile.txt --content "Hello urfave/cli!"
./demo file info myfile.txt
```

### Configuration with Validation
```bash
./demo config --host localhost --port 8080 --ssl
# Output: Configuration with HTTPS URL

./demo config --host "ab" --port 70000
# Output: Validation errors for invalid values
```

### Global Flags
```bash
./demo --verbose greet --name "User"
# Output: Verbose mode enabled, then greeting
```

## Building & Running

```bash
# Initialize and install dependencies
go mod init urfavecli-demo
go get github.com/urfave/cli/v3@latest

# Build the application
go build -o demo

# Run with help
./demo --help

# Test specific commands
./demo greet --help
./demo math --help
```

## Key urfave/cli v3 Features Highlighted

1. **Type-safe flag handling** - Strongly typed flags with built-in parsing
2. **Hierarchical commands** - Nested subcommands with individual help
3. **Flexible validation** - Custom validators with clear error messages  
4. **Rich metadata** - Authors, copyright, version information
5. **Hook system** - Before/After hooks for custom logic
6. **Auto-generated help** - Beautiful help text with minimal configuration
7. **Shell completion** - Built-in support for command completion
8. **Error handling** - Graceful error reporting and exit codes

This demo showcases the power and simplicity of urfave/cli v3 for building robust command-line applications in Go.