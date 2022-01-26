package main

import (
    "fmt"
    "os"
    
    "github.com/kevinburke/ssh_config"
    "path/filepath"
)



func main() {
    f, _ := os.Open(filepath.Join(os.Getenv("HOME"), ".ssh", "config"))

    cfg, _ := ssh_config.Decode(f)

    for _, host := range cfg.Hosts {

        fmt.Printf("--------------\n")
        fmt.Println("patterns:", host.Patterns)
        // fmt.Println(host.String())
        for _, node := range host.Nodes {
            fmt.Println(node.String())
        }
    }

    // Print the config to stdout:
    // fmt.Println(cfg.String())

}
