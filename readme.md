# nav: a terminal navigation helper

A lightweight Go CLI tool that generates shell functions for enhanced terminal navigation with `up`, `dn`, `to`, and `t` commands.

## Installation & Setup

1. **Install the CLI tool:**
   ```bash
   go install github.com/nanvenomous/nav@latest
   # or build locally
   go build -o nav
   ```

2. **Add navigation functions to your shell:**
   ```bash
   eval "$(nav)"              # Current session only
   nav >> ~/.bashrc           # Add to bash config
   nav >> ~/.zshrc            # Add to zsh config
   ```

3. **Start using the navigation commands:**
   ```bash
   up 2        # Move up 2 directories
   dn mydir    # Move into mydir  
   to          # Interactive navigation with fzf
   ```

## Commands

### `up [levels]`
Move up directories (equivalent to `cd ..`)
- **Default:** `up` moves up 1 level
- **Multi-level:** `up 2` moves up 2 levels, `up 3` moves up 3 levels, etc.
- **Clears terminal** and lists directory contents after navigation

### `dn <directory>`  
Move down into a directory (equivalent to `cd <directory>`)
- **Directory-only:** Only accepts directories, not files
- **Tab completion:** Supports shell completion for directory names
- **Clears terminal** and lists directory contents after navigation

### `to`
Interactive directory navigation using `fzf`
- **File preview:** Shows current directory files at top, then interactive directory selection
- **Interactive selection:** Navigate directories with fzf
- **Parent navigation:** Includes `..` option to go up one level  
- **Continuous loop:** Keeps running until you exit (Escape or Ctrl+C)
- **Clears terminal** and lists contents after each selection
- **Requires fzf:** Must have `fzf` installed

## Features

### Smart List Command Detection
Automatically uses the best available directory listing tool:

1. **exa** (preferred): `exa --tree --level=0 --group-directories-first`
2. **tree** (fallback): `tree -C -L 1 --dirsfirst`
3. **ls variants:**
   - Linux: `ls --color=auto --group-directories-first -1`
   - macOS: `ls -G -1`
   - Other: `ls -1`

### Shell Integration
- **Cross-shell compatible:** Works with bash, zsh, and other POSIX shells
- **Tab completion:** Directory completion for `dn` command
- **Error handling:** Proper error messages and validation
- **Clean output:** Each command clears terminal before showing new directory

## Examples

```bash
# Navigate up multiple levels
up 3          # cd ../../../ + clear + list

# Move into subdirectory with completion
dn <TAB>      # Shows available directories
dn projects   # cd projects + clear + list

# Interactive navigation
to            # Opens fzf interface
              # Select directory or .. to go up
              # Continues until you exit
```

## Distribution

The CLI tool approach allows for easy distribution:

```bash
# Package managers
go install github.com/yourorg/nav@latest

# Direct download
curl -L https://github.com/yourorg/nav/releases/latest/download/nav > nav
chmod +x nav

# Then add to shell config
./nav >> ~/.bashrc
```

## Dependencies

### Required
- **Go 1.21+** (for building)
- **POSIX shell** (bash, zsh, etc.)

### Optional  
- **fzf** - Required for `to` command interactive navigation
- **exa** - Enhanced directory listings (recommended)
- **tree** - Tree-style directory listings (fallback)

## Technical Details

This lightweight tool:
- Generates clean, portable shell functions
- Provides proper error handling and validation
- Supports cross-platform directory operations
- Includes shell completion setup
- Zero external dependencies (single static binary)

The generated shell functions are self-contained and don't require the CLI binary at runtime - they only need standard shell utilities and optionally fzf for the `to` command.

## Why This Approach?

1. **Easy distribution:** Single binary that sets up everything
2. **Shell integration:** Functions run in your current shell, actually changing directories
3. **Portable:** Generated functions work across different shells and systems
4. **Self-contained:** No runtime dependency on the CLI tool once installed
5. **Package-friendly:** Can be distributed via package managers, GitHub releases, etc.

Perfect for dotfiles, team setups, or any scenario where you want enhanced terminal navigation!
