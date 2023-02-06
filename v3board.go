package main

import (
	"log"
	"v3board/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatalf("boot server failed, err: %v\n", err)
	}
}
