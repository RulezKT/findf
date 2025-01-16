package findf

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func Dir(filesdir string) string {

	dir, err := filepath.Abs("./")
	if err != nil {
		log.Fatal(err)
	}

	for {
		dirToFind := filepath.Join(dir, filesdir)

		_, err := os.Stat(dirToFind)
		if err == nil {
			// fmt.Println("Folder exist")
			return dirToFind
		}

		dir = filepath.Dir(dir)

		if dir == filepath.Dir(dir) {
			fmt.Println("Didn't find neccesary folder")
			log.Fatal(err)
		}
	}
}

func File(dir string, file string) string {

	fileToFind := filepath.Join(dir, file)

	_, err := os.Stat(fileToFind)
	if err == nil {
		// fmt.Println("File exist")
		return fileToFind
	}

	fmt.Println("Didn't find neccesary file")
	log.Fatal(err)
	return ""
}
