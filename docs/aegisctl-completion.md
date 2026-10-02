# aegisctl Shell Completion

`aegisctl` can print command completion scripts for Bash and Zsh:

```bash
aegisctl completion bash
aegisctl completion zsh
```

## Bash Installation

With `aegisctl` on your `PATH`, load completions in the current shell with:

```bash
eval "$(aegisctl completion bash)"
```

Add that line to `~/.bashrc` to load completions in future Bash sessions.

## Zsh Script Generation

Add this to `~/.zshrc`, after initializing Zsh completions with `compinit`:

```zsh
autoload -Uz compinit
compinit
source <(aegisctl completion zsh)
```
