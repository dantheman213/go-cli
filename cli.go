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
    "/bin/zsh",
}

// Setup the command but don't run it, return exec.Cmd for granular manipulation.
func MakeCommand(command string) (*exec.Cmd, error) {
    return makeCommandWithPrefix(command)
}

// Setup the command and execute it right away. Return the stdout and stderr buffers.
func MakeAndRunCommand(command string) (cmd *exec.Cmd, stdout *bytes.Buffer, stderr *bytes.Buffer, err error) {
    cmd, err = makeCommandWithPrefix(command)
    if err != nil {
        return nil, nil, nil, err
    }

    bufOut  := bytes.Buffer{}
    cmd.Stdout = &bufOut

    bufErr := bytes.Buffer{}
    cmd.Stderr = &bufErr

    if err := cmd.Start(); err != nil {
        return nil, nil, nil, err
    }

    return cmd, &bufOut, &bufErr, nil
}

func MakeAndRunCommandThenWait(command string) (cmd *exec.Cmd, stdout *bytes.Buffer, stderr *bytes.Buffer, err error) {
    cmd, stdout, stderr, err = MakeAndRunCommand(command)
    if err != nil {
        return nil, nil, nil, err
    }

    err = cmd.Wait()
    if err != nil {
        return nil, nil, nil, err
    }

    return cmd, stdout, stderr, nil
}

// Setup the command and execute it right away. Return the stdout and stderr buffers together as a stream
func MakeAndRunCommandWithCombinedOutput(command string) (cmd *exec.Cmd, out *bytes.Buffer, err error) {
    cmd, err = makeCommandWithPrefix(command)
    if err != nil {
        return nil, nil, err
    }

    buf := bytes.Buffer{}
    cmd.Stdout = &buf
    cmd.Stderr = &buf

    if err := cmd.Start(); err != nil {
        return nil, nil, err
    }

    return cmd, &buf, nil
}

// Setup the command execute it right away with combined stdout and stderr buffers then wait for command to finish
func MakeAndRunCommandWithCombinedOutputThenWait(command string) (cmd *exec.Cmd, stdout *bytes.Buffer, err error) {
    cmd, stdout, err = MakeAndRunCommandWithCombinedOutput(command)
    if err != nil {
        return nil, nil, err
    }

    err = cmd.Wait()
    if err != nil {
        return nil, nil, err
    }

    return cmd, stdout, nil
}

func makeCommandWithPrefix(command string) (*exec.Cmd, error) {
    shell := ""
    prefix := "-c"

    if runtime.GOOS == "windows" {
        shell = "cmd.exe"
        prefix = "/C"
    } else if runtime.GOOS == "linux" || runtime.GOOS == "darwin" {
        for _, path := range linuxDarwinShellPathList {
            if fileExists(path) {
                shell = path
                break
            }
        }

        if shell == "" {
            return nil, errors.New("could not find shell")
        }
    } else {
        return nil, errors.New("unsupported operating system detected")
    }

    return exec.Command(shell, prefix, command), nil
}
