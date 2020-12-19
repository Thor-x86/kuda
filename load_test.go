package main

import (
	"bytes"
	"compress/gzip"
	"io/ioutil"
	"path/filepath"
	"testing"
)

func TestLoad(t *testing.T) {
	publicPath := filepath.Join("./", "test")
	load(publicPath)

	expectations := map[string]fileModel{
		"index.html": {
			data:         loadFileFromDisk(t, "index.html", true),
			mime:         "text/html",
			isCompressed: true,
		},
		"index.css": {
			data:         loadFileFromDisk(t, "index.css", true),
			mime:         "text/css",
			isCompressed: true,
		},
		"index.js": {
			data:         loadFileFromDisk(t, "index.js", true),
			mime:         "application/javascript",
			isCompressed: true,
		},
		"roboto.ttf": {
			data:         loadFileFromDisk(t, "roboto.ttf", true),
			mime:         "font/ttf",
			isCompressed: true,
		},
		"test.jpg": {
			data:         loadFileFromDisk(t, "test.jpg", false),
			mime:         "image/jpeg",
			isCompressed: false,
		},
		"test.png": {
			data:         loadFileFromDisk(t, "test.png", false),
			mime:         "image/png",
			isCompressed: false,
		},
		"test.svg": {
			data:         loadFileFromDisk(t, "test.svg", true),
			mime:         "image/svg+xml",
			isCompressed: true,
		},
		"with_index/index.html": {
			data:         loadFileFromDisk(t, "with_index/index.html", true),
			mime:         "text/html",
			isCompressed: true,
		},
		"with_index/wkwkwk@&$!-_#=+.jpg": {
			data:         loadFileFromDisk(t, "with_index/wkwkwk@&$!-_#=+.jpg", false),
			mime:         "image/jpeg",
			isCompressed: false,
		},
		"without-index/never gonna give you up.jpg": {
			data:         loadFileFromDisk(t, "without-index/never gonna give you up.jpg", false),
			mime:         "image/jpeg",
			isCompressed: false,
		},
	}

	for eachKey, eachFile := range *pathMap {
		eachExpect, ok := expectations[eachKey]
		if !ok {
			t.Errorf("\"%s\" is exist in pathMap but not exist on disk", eachKey)
			continue
		}
		if eachFile.mime != eachExpect.mime {
			t.Errorf("\"%s\" has \"%s\" MIME, but expected \"%s\"", eachKey, eachFile.mime, eachExpect.mime)
			continue
		}
		if eachFile.isCompressed != eachExpect.isCompressed {
			t.Errorf("\"%s\" has \"%t\" compression, but expected \"%t\"", eachKey, eachFile.isCompressed, eachExpect.isCompressed)
			continue
		}
		if string(eachFile.data) != string(eachExpect.data) {
			t.Errorf("\"%s\" has different data from original file", eachKey)
		}
	}
}

// Test helper, it helps load manually from disk as comparation
func loadFileFromDisk(t *testing.T, relativePath string, isCompress bool) []byte {
	// Get public path
	publicPath := filepath.Join("./", "test")

	// Absolute path of file
	path := filepath.Join(publicPath, relativePath)

	// Read file from disk, directly return the result if isCompress=false
	result, err2 := ioutil.ReadFile(path)
	if !isCompress {
		return result
	}
	if err2 != nil {
		t.Fatal(err2)
		return nil
	}

	// Compress if isCompress=true
	var compressedResult bytes.Buffer
	compressor, err3 := gzip.NewWriterLevel(&compressedResult, gzip.BestCompression)
	if err3 != nil {
		t.Fatal(err3)
		return nil
	}
	_, err4 := compressor.Write(result)
	if err4 != nil {
		t.Fatal(err4)
		return nil
	}
	compressor.Close()

	return compressedResult.Bytes()
}
