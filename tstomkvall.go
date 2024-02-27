package main

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func Contains[T comparable](A []T, e T) bool {
	for _, v := range A {
		if v == e {
			return true
		}
	}
	return false
}

func main() {
	// get CWD
	path, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}

	// read files on CWD
	files, err := os.ReadDir(path)
	if err != nil {
		log.Fatalln(err)
	}

	allowedFiletypes := []string{".ts"}

	// append .ts files in CWD to arguments list and process them one-by-one
	for _, file := range files {
		fileExtension := filepath.Ext(file.Name())
		fileBaseName := strings.TrimSuffix(file.Name(), fileExtension)
		args := []string{}

		if Contains(allowedFiletypes, strings.ToLower(fileExtension)) {
			var sb strings.Builder

			filePathTS := filepath.Join(path, file.Name())
			sb.WriteString(fileBaseName)
			sb.WriteString(".mkv")
			filePathMKV := filepath.Join(path, sb.String())

			// .ts file is simply remuxed into .mkv container by copying all
			// streams from input file and leaving them untouched in output
			args = append(args, "-y")
			args = append(args, "-i")
			args = append(args, filePathTS)
			args = append(args, "-c")
			args = append(args, "copy")
			args = append(args, filePathMKV)
		}

		// don't do anything if file is not a .ts file
		if len(args) == 0 {
			continue
		}

		// run FFmpeg on the .ts file
		if err := exec.Command("ffmpeg", args...).Run(); err != nil {
			log.Fatal(err)
		}
		log.Printf("Remuxed \"%s\" to .mkv\n", file.Name())
	}
}
