package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/valyala/fasthttp"
)

// Global variables
var pathMap *map[string]fileModel = nil
var origins string = ""

// Main entry
func main() {
	// Initialize required parameters
	domain := "localhost"
	port := "8080"
	portTLS := ""
	certFile := ""
	keyFile := ""

	// Get parameters from flags
	flag.StringVar(&domain, "domain", domain, helpArgs["domain"])
	flag.StringVar(&port, "port", port, helpArgs["port"])
	flag.StringVar(&origins, "origins", origins, helpArgs["origins"])
	flag.StringVar(&portTLS, "port-tls", portTLS, helpArgs["port-tls"])
	flag.StringVar(&certFile, "cert", certFile, helpArgs["cert"])
	flag.StringVar(&keyFile, "key", keyFile, helpArgs["key"])
	flag.Parse()

	// If user doesn't put required arguments, then show help info
	if flag.NArg() == 0 {
		help()
		return
	}

	// Public root directory
	dir := flag.Args()[0]

	// If using TLS, check for certificate and secret key availability
	if len(portTLS) > 0 && len(keyFile) == 0 {
		fmt.Println("You forgot \"--key=...\" flag!")
		return
	}
	if len(portTLS) > 0 && len(certFile) == 0 {
		fmt.Println("You forgot \"--cert=...\" flag!")
		return
	}

	// Load all files at public directory to RAM
	load(dir)

	// Check if currently working on local or production environtment
	domainAsIP, _ := fasthttp.ParseIPv4(net.IP{}, []byte(domain))
	isLocal := (domain == "localhost") || (domainAsIP != nil)

	// Pre-process origins
	if isLocal {
		if len(portTLS) > 0 {
			origins = domain + ":" + portTLS + "," + origins
		} else {
			origins = domain + ":" + port + "," + origins
		}
	} else {
		origins = domain + "," + origins
	}

	fmt.Printf("Kuda is listening at %s port...\n", port)

	// Remember, emptied port-tls means disabled TLS
	var err error = nil
	if len(portTLS) > 0 {
		fmt.Printf("Kuda also securely listening at %s port...\n", portTLS)

		// Listen for incoming HTTP request and redirect them to HTTPS
		go fasthttp.ListenAndServe(":"+port, func(ctx *fasthttp.RequestCtx) {
			path := string(ctx.URI().Path())
			if isLocal {
				ctx.Redirect("https://"+domain+":"+portTLS+"/"+path, 302)
			} else {
				ctx.Redirect("https://"+domain+"/"+path, 302)
			}
		})

		// Listen for incoming HTTPS request:
		err = fasthttp.ListenAndServeTLS(":"+portTLS, certFile, keyFile, handler)
	} else {
		// Listen for incoming HTTP request:
		err = fasthttp.ListenAndServe(":"+port, handler)
	}

	log.Fatalln(err)
}
