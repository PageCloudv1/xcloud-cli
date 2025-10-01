package main

import (
	"testing"
)

func TestVersion(t *testing.T) {
	if version == "" {
		t.Error("Version should not be empty")
	}
}

func TestRootCommand(t *testing.T) {
	if rootCmd.Use != "xcloud" {
		t.Errorf("Expected command name 'xcloud', got '%s'", rootCmd.Use)
	}
}

func TestCommands(t *testing.T) {
	commands := []string{"version", "deploy", "status", "logs", "init", "create"}

	for _, cmdName := range commands {
		found := false
		for _, cmd := range rootCmd.Commands() {
			if cmd.Use == cmdName || cmd.Use == cmdName+" [nome-do-projeto]" || cmd.Use == cmdName+" [tipo] [nome]" {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Command '%s' not found", cmdName)
		}
	}
}

// Benchmark para testar performance de inicialização
func BenchmarkRootCommand(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rootCmd.Use = "xcloud"
	}
}
