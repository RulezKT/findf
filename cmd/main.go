package main

import "findf"

func main() {

	// findfiles.FindDir("files")

	folder := findf.Dir("cmd")

	findf.File(folder, "main.go")
}
