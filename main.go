package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	var size int64 = 1024
	var unit string = "b"
	var fileName string = "output.tmp"

	// Declare flags
	flag.Int64Var(&size, "s", size, "File size value")
	flag.StringVar(&unit, "u", unit, "file size unit (b, kb, mb, gb, tb)")
	flag.StringVar(&fileName, "o", fileName, "File name")
	flag.Parse()

	// Calculate file size
	var bytes int64
	switch unit {
	case "b":
		bytes = size
	case "kb":
		bytes = size * 1024
	case "mb":
		bytes = size * 1024 * 1024
	case "gb":
		bytes = size * 1024 * 1024 * 1024
	case "tb":
		bytes = size * 1024 * 1024 * 1024 * 1024
	default:
		fmt.Println("Invalid unit: only b, kb, mb, gb and tb are allowed")
		return
	}

	f, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Parent directory does not exist:\n", err)
		return
	}
	defer f.Close()

	// Write the desired number of bytes to the file
	_, err = f.Write(make([]byte, bytes))
	if err != nil {
		fmt.Println(err)
		return
	}

	fileInfo, err := os.Stat(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	abs, err := filepath.Abs(fileName)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Created successfully\n- File name: %s\n- Full path: %s\n- File size: %d byte(s)\n", fileInfo.Name(), abs, fileInfo.Size())
}
