package cli

import (
    "bytes"
    "os/exec"
    "runtime"
)

// Setup the command and execute it right away. Return the stdout and stderr.
func RunCommand(command string) (cmd *exec.Cmd, stdout bytes.Buffer, stderr bytes.Buffer, err error) {
    cmd = makeCommandWithPrefix(command)

    cmd.Stdout = &stdout
    cmd.Stderr = &stderr

    if err := cmd.Start(); err != nil {
        return nil, bytes.Buffer{}, bytes.Buffer{}, err
    }

    return cmd, stdout, stderr, nil
}

// Setup the command and execute it right away. Return the stdout and stderr together as a stream
func RunCommandCombinedOutput(command string) (cmd *exec.Cmd, err error) {
    // TODO
    return nil, nil
}

// Setup the command but don't run it, return exec.Cmd for granular manipulation.
func SetupCommand(command string) *exec.Cmd {
    cmd := makeCommandWithPrefix(command)
    return cmd
}

func makeCommandWithPrefix(command string) *exec.Cmd {
    shell := "bash"
    prefix := "-c"
    if runtime.GOOS == "windows" {
        shell = "cmd.exe"
        prefix = "/C"
    }

    return exec.Command(shell, prefix, command)
}
