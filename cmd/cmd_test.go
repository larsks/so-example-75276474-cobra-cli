package cmd

import (
	"testing"
)

func TestRootCmdWithArgs(t *testing.T) {
	rootCmd.SetArgs([]string{"-t", "12"})
	if err := rootCmd.Execute(); err != nil {
		t.Errorf("failed to execute rootCmd")
	}

	if threads != 12 {
		t.Errorf("expected 12, got %d", threads)
	}
}

func TestRootCmdInvalidArgs(t *testing.T) {
	rootCmd.SetArgs([]string{"--arg-that-does-not-exist"})
	if err := rootCmd.Execute(); err == nil {
		t.Errorf("command succeeded when it should have failed")
	}
}

func TestFooCmdWithArgs(t *testing.T) {
	rootCmd.SetArgs([]string{"foo", "-c", "2"})
	if err := rootCmd.Execute(); err != nil {
		t.Errorf("failed to execute rootCmd")
	}

	if count != 2 {
		t.Errorf("execpted 2, got %d", count)
	}
}
