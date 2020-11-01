package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// Load all files inside public root directory into RAM, only executed once.
func load(publicDir string) {
	// Add temporary map of file
	loaded := map[string]fileModel{}

	// Load files one by one
	fmt.Printf("Loading from \"%s\":\n", publicDir)
	filepath.Walk(publicDir, func(path string, info os.FileInfo, err2 error) error {
		// Check for premature error
		if err2 != nil {
			log.Fatalln(err2.Error())
			return err2
		}

		// We just need files, so skip to its inside
		if info.IsDir() {
			return nil
		}

		// Convert absolute path to URI without leading slash
		uriPath := strings.TrimPrefix(path, publicDir+"/")
		fileData, err3 := ioutil.ReadFile(path)
		if err3 != nil {
			log.Fatalln(err3.Error())
			return err3
		}

		isCompressed := false

		// Get MIME of correspond file
		var fileMime string
		if strings.HasSuffix(uriPath, ".js") {
			fileMime = "application/javascript"
			isCompressed = true
		} else if strings.HasSuffix(uriPath, ".css") {
			fileMime = "text/css"
			isCompressed = true
		} else if strings.HasSuffix(uriPath, ".svg") {
			fileMime = "image/svg+xml"
			isCompressed = true
		} else {
			fileMime = http.DetectContentType(fileData)
			isCompressed = (fileMime == "image/x-icon") || strings.HasPrefix(fileMime, "text/") || strings.HasPrefix(fileMime, "font/")
		}
		fileMime = strings.TrimSuffix(fileMime, "; charset=utf-8")
		fmt.Printf("\t- %s (%s)\n", uriPath, fileMime)

		// If supposed to not compressed, then just add into map and skip compression process
		if !isCompressed {
			loaded[uriPath] = fileModel{
				data:         fileData,
				mime:         fileMime,
				isCompressed: false,
			}
			return nil
		}

		// Do GZIP compression
		var compressedData bytes.Buffer
		compressor, err4 := gzip.NewWriterLevel(&compressedData, gzip.BestCompression)
		if err4 != nil {
			log.Fatalln(err4.Error())
			return err4
		}
		_, err5 := compressor.Write(fileData)
		if err5 != nil {
			log.Fatalln(err5.Error())
			return err5
		}
		compressor.Close()

		// Add compressed file to map
		loaded[uriPath] = fileModel{
			data:         compressedData.Bytes(),
			mime:         fileMime,
			isCompressed: true,
		}

		return nil
	})

	// Assign the created file map to a global pointer, then report via CLI
	pathMap = &loaded
	fmt.Println("")
	fmt.Println("All Files Loaded!")
	fmt.Println("")
}
