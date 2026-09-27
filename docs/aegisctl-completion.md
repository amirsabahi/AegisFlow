# aegisctl Shell Completion

`aegisctl` can print command completion scripts for Bash and Zsh:

```bash
aegisctl completion bash
aegisctl completion zsh
```

## Bash Installation

With `aegisctl` on your `PATH`, load completions in the current shell with:

```bash
source <(aegisctl completion bash)
```

Add that `source` line to `~/.bashrc` to load it in future Bash sessions.

## Zsh Script Generation

`aegisctl completion zsh` prints its shell-specific completion script for use with your Zsh completion setup.
