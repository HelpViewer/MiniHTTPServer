package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	port := flag.String("port", "8080", "Port (8080)")
	dir := flag.String("dir", "www/", "Path (www/)")
	addrF := flag.String("addr", "127.0.0.1:", "IPAddress (127.0.0.1:)")
	flag.Parse()

	absDir, err := filepath.Abs(*dir)
	if err != nil {
		log.Fatalf("File system error: %s: %v", *dir, err)
	}

	if _, err := os.Stat(absDir); os.IsNotExist(err) {
		log.Fatalf("Dir %s missing!", absDir)
	}

	http.Handle("/", http.FileServer(http.Dir(absDir)))

	addr := *addrF + *port
	log.Printf("Running with path %s on http://%s", absDir, addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
