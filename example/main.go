package main

import (
    "fmt"
    "github.com/dantheman213/go-cli"
    "runtime"
)

func main() {
    if runtime.GOOS == "windows" {
        fmt.Println("TODO")
    } else {
        runLinuxDarwinCommandSampleOne()
        runLinuxDarwinCommandSampleTwo()
    }
}

func runLinuxDarwinCommandSampleOne() {
    command := `cd /tmp && echo "Hello World!" > hello_world.txt && cat /tmp/hello_world.txt`
    _, stdout, stderr, err := cli.MakeAndRunCommandThenWait(command)
    if err != nil {
        panic(err)
    }

    fmt.Println(stdout.String())
    fmt.Println(stderr.String())
}

func runLinuxDarwinCommandSampleTwo() {
    command := `echo "Hello, Path! $PATH"`
    cmd, stdout, stderr, err := cli.MakeAndRunCommand(command)
    if err != nil {
        panic(err)
    }

    err = cmd.Wait()
    if err != nil {
        panic(err)
    }

    fmt.Println("OUTPUT: " + stdout.String())
    fmt.Println("ERROR: " + stderr.String())
}
