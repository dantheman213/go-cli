package cli

import (
    "bytes"
    "errors"
    "os/exec"
    "runtime"
)

var linuxDarwinShellPathList []string = []string {
    "/bin/bash",
    "/usr/bin/bash",
    "/usr/local/bin/bash",
    "/bin/sh",
}

// Setup the command but don't run it, return exec.Cmd for granular manipulation.
func MakeCommand(command string) (*exec.Cmd, error) {
    return makeCommandWithPrefix(command)
}

// Setup the command and execute it right away. Return the stdout and stderr buffers.
func MakeAndRunCommand(command string) (cmd *exec.Cmd, stdout bytes.Buffer, stderr bytes.Buffer, err error) {
    cmd, err = makeCommandWithPrefix(command)
    if err != nil {
        return nil, bytes.Buffer{}, bytes.Buffer{}, err
    }

    cmd.Stdout = &stdout
    cmd.Stderr = &stderr

    if err := cmd.Start(); err != nil {
        return nil, bytes.Buffer{}, bytes.Buffer{}, err
    }

    return cmd, stdout, stderr, nil
}

// Setup the command and execute it right away. Return the stdout and stderr buffers together as a stream
func MakeAndRunCommandWithCombinedOutput(command string) (cmd *exec.Cmd, err error) {
    // TODO
    return nil, nil
}

func makeCommandWithPrefix(command string) (*exec.Cmd, error) {
    shell := ""
    prefix := "-c"

    if runtime.GOOS == "windows" {
        shell = "cmd.exe"
        prefix = "/C"
    } else {
        for _, path := range linuxDarwinShellPathList {
            if fileExists(path) {
                shell = path
                break
            }
        }

        if shell == "" {
            return nil, errors.New("could not find shell")
        }
    }

    return exec.Command(shell, prefix, command), nil
}
