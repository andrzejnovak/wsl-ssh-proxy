package main

import (
	"os"
	"os/exec"
	"syscall"
)

func main() {
	// Prepend "ssh" to the arguments so wsl.exe runs: ssh <original args>
	args := append([]string{"ssh"}, os.Args[1:]...)
	cmd := exec.Command(`C:\Windows\System32\wsl.exe`, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Prevent wsl.exe from creating a visible console window
	cmd.SysProcAttr = &syscall.SysProcAttr{
		CreationFlags: 0x08000000, // CREATE_NO_WINDOW
	}

	err := cmd.Run()
	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			os.Exit(exitError.ExitCode())
		}
		os.Exit(1)
	}
}
