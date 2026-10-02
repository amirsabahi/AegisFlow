package main

import (
	"fmt"
	"strings"
)

var completionCommands = []string{
	"pending", "approve", "deny", "verify", "evidence", "simulate", "why",
	"diff-policy", "manifest", "supply-chain", "policy-pack", "test-action",
	"plugin", "status", "usage", "models", "providers", "policies", "tenants",
	"policy", "test", "version", "completion", "help",
}

func bashCompletionScript() string {
	return `_aegisctl_completion() {
	local cur commands
	cur="${COMP_WORDS[COMP_CWORD]}"
	commands="` + strings.Join(completionCommands, " ") + `"
	if [ "$COMP_CWORD" -eq 1 ]; then
		COMPREPLY=( $(compgen -W "${commands}" -- "${cur}") )
	fi
}
complete -F _aegisctl_completion aegisctl
`
}

func zshCompletionScript() string {
	return `#compdef aegisctl

_aegisctl() {
	local -a commands
	commands=(` + strings.Join(completionCommands, " ") + `)
	_describe 'command' commands
}
compdef _aegisctl aegisctl
`
}

func cmdCompletion(args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("usage: aegisctl completion <bash|zsh>")
	}
	switch args[0] {
	case "bash":
		fmt.Print(bashCompletionScript())
	case "zsh":
		fmt.Print(zshCompletionScript())
	default:
		return fmt.Errorf("unknown shell %q, expected \"bash\" or \"zsh\"", args[0])
	}
	return nil
}
