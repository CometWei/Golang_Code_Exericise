package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	cmd := exec.Command("gedit", "../Get_http_File/robots.txt")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("cmd.Start() failed with %s\n", err)
	}
	fmt.Printf("combined out:\n%s\n", string(out))
}
