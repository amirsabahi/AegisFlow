package main

import (
	"fmt"
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
	cur = "${COMP_WORDS[COMP_CWORD]}"
	commands "` + joinWithSpaces(completionCommands) + `"
	if ["$COMP_CWORD" -eq 1]; then 
		COMPREPLY = ( $(compgen -W "${commands} -- "${cur}"") )
	fi
}	
complete -F _aegisctl_completion aegisctl	
	`
}

func zshCompletionScript() string {
	return `
	_aegisctl() {
		local -a commands
		commands=(` + joinWithSpaces(completionCommands) + `)
		_describe 'command' commands		
}
_aegisctl		
`
}

func joinWithSpaces(items []string) string {
	result := ""
	for i, item := range items {
		if i > 0 {
			result += " "
		}
		result += item
	}
	return result
}

func cmdCompletation(args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("usage: aegisctl completion <bash|zsh>")
	}
	switch args[0] {
	case "bash":
		fmt.Print(bashCompletionScript())
	case "zsh":
		fmt.Print(zshCompletionScript())
	default:
		return fmt.Errorf("Unknown shell %q, expected \"bash\" or \"zsh\"", args[0])
	}
	return nil
}
