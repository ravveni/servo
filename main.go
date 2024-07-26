package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	currentDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	port := "8080"

	if len(os.Args) > 1 {
		port = os.Args[1]
	}

	fs := http.FileServer(http.Dir(currentDir))
	http.Handle("/", fs)

	fmt.Println("Serving", currentDir, "on port", port)

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		panic(err)
	}
}
