package main

// Structure of each file inside RAM
type fileModel struct {
	data         []byte
	mime         string
	isCompressed bool
}
