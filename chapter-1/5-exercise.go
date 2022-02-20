package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		prefix := "http://"
		if !strings.HasPrefix(url, prefix) && !strings.HasPrefix(url, "https://") {
			url = prefix + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprint(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		fmt.Println(resp.Status)
		if _, err := io.Copy(os.Stdout, resp.Body); err != nil {
			fmt.Fprint(os.Stderr, "fetch: reading %s: %v\n", url, err)
		}
	}
}
