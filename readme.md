# where-to

I'm here, I'm there ;)

Jump from one place to another on your OS with the help of command completion and a simple config file.
![demo](./.rsrc/where-to.gif)

### Do note
Currently only works for zsh, but pull requests welcome!

# Dependencies

By default, where-to uses ls to show directories.

You can install [exa](https://github.com/ogham/exa) or [tree](https://gist.github.com/fscm/9eee2784f101f21515d66321180aef0f) for a better experience.

# Installation

```
git clone https://github.com/nanvenomous/where-to.git
cd where-to
sudo make install
sudo cp ./.completions/zsh/_to /usr/share/zsh/site-functions
```

You need to add the plugin to your shell

.zshrc:
```
eval "$(where-to init)"
```

Working on convenience functions, but for now just make a config file

~/.config/where-to.yaml
```
cho: "/home/natsu/projects/adiumads/cho"
where: "/home/natsu/projects/where-to"
ani: "/home/natsu/projects/ani-cli"
```

# Inspiration
This project is heavily inspired by [zoxide](https://github.com/ajeetdsouza/zoxide), but has a more declarative approach & focuses on autocompletion rather than fuzzy-finding.
