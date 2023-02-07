// Fetch print the content found at a URL
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
		if strings.HasPrefix(url, "http://") {
			resp, err := http.Get(url)
			if err != nil {
				fmt.Fprint(os.Stderr, "fetch: %v\n", err)
				os.Exit(1)
			}
			//fmt.Printf("%s", resp.Body)
			io.Copy(os.Stdout, resp.Body)
			//b, err := io.ReadAll(resp.Body)
			resp.Body.Close()
			if _, err := io.Copy(os.Stdout, resp.Body); err != nil {
				//log.Fatal(err)
				//os.Exit(1)
				fmt.Println("Http Status:", resp.Status)
				//continue
			} else {
				fmt.Fprint(os.Stderr, "fetch: reading %s: %v\n", url, err)
				os.Exit(1)
			}
		} else {
			fmt.Print("Please add prefix 'http://'!!!")
		}
		//fmt.Printf("%d", c)
	}

}
