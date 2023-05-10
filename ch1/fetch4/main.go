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
		if !strings.HasPrefix(url, "https://") || !strings.HasPrefix(url, "http://") {
			url = "https://" + url
		}

		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Have error : %v\n", err)
			os.Exit(0)
		}

		_, err = io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Fetch url: %q have error %v\n", url, err)
			os.Exit(0)
		}
		fmt.Printf("Fetch url: %q http status is %d\n", url, resp.StatusCode)

	}
}
