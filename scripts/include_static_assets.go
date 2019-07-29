package main

import (
    "fmt"
    "io"
    "io/ioutil"
    "os"
    "strings"
)

func main() {
    fmt.Printf("\n\ninclude_static_assets script\n")
    fmt.Printf("==============================\n")
    fmt.Printf("Removing older staticassets.go\n")
    os.Remove(os.Args[1] + "staticassets.go")
    fs, _ := ioutil.ReadDir(os.Args[1] + "static_assets")
    out, _ := os.Create(os.Args[1] + "staticassets.go")
    out.Write([]byte("package " + os.Args[2] + " \n\nconst (\n"))
    for _, f := range fs {
        varName := strings.Replace(f.Name(), ".", "_", -1)
        varName = strings.Replace(varName, " ", "-", -1)
        out.Write([]byte(varName + " = `"))
        
        f, err := os.Open("./static_assets/" + f.Name())
        if err != nil {
            fmt.Printf("%s: %s\n", f.Name(), err)
            out.Write([]byte("Error Reading!"))
        } else {
            fmt.Printf("Writing %s...", f.Name())
            io.Copy(out, f)
            fmt.Printf("Complete!\n")
        }
        out.Write([]byte("`\n"))

        f.Close()
    }
    out.Write([]byte(")\n"))
    fmt.Printf("All Done!\n\n")
}