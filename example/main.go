package main

import (
    "fmt"
    "github.com/dantheman213/go-cli"
    "runtime"
)

func main() {
    if runtime.GOOS != "windows" {
        command := `cd /tmp && echo "Hello World!" > hello_world.txt && cat /tmp/hello_world.txt`
        cmd, stdout, stderr, err := cli.RunCommand(command)
        if err != nil {
            panic(err)
        }

        _ = cmd.Wait()
        fmt.Println(stdout.String())
        fmt.Println(stderr.String())
    } else {
        fmt.Println("TODO")
    }
}
