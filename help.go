package main

import (
	"fmt"
	"os"
)

// Description for each arguments
var helpArgs map[string]string = map[string]string{
	"domain":   "Required to redirect from http to https",
	"port":     "TCP Port to be listened",
	"origins":  "Which domains to be allowed by CORS policy",
	"port-tls": "Use this to listen for HTTPS requests",
	"cert":     "SSL certificate file, required if \"--port-tls\" specified",
	"key":      "SSL secret key file, required if \"--port-tls\" specified",
}

// Displays usage help
func help() {
	fmt.Println("")
	fmt.Println("USAGE:")
	fmt.Printf("    %s [arguments] <public_root_directory>\n", os.Args[0])
	fmt.Println("")
	fmt.Println("ARGUMENTS:")
	for argName, argDesc := range helpArgs {
		fmt.Printf("    --%s=...  : %s\n", argName, argDesc)
	}
	fmt.Println("")
}
