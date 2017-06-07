package main

import "os"

func doSomething() error {
    err := os.MkdirAll("newdir", 0755)
    if err != nil {
        return err
    }
    defer os.RemoveAll("newdir")

    f, err := os.Create("newdir/newfile")
    if err != nil {
        return err
    }
    defer f.Close()
    return nil
}

func main() {
    doSomething()
}
