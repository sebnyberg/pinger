package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	"os"
	"strings"
)

func main() {
	url := os.Getenv("PING_URL")
	if url == "" {
		panic("PING_URL is required")
	}
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		s, err := httputil.DumpRequest(r, false)
		if err != nil {
			fmt.Println("failed to dump request", err)
			return
		}
		fmt.Println("REQUEST:")
		fmt.Println(string(s))
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println("failed to GET underlying URL:", err)
			return
		}
		s, err = httputil.DumpResponse(resp, false)
		if err != nil {
			fmt.Println("failed to dump response", err)
			return
		}
		fmt.Println("RESPONSE:")
		fmt.Println(string(s))
		for k, vs := range resp.Header {
			w.Header().Set(k, strings.Join(vs, ""))
		}
		_, err = io.Copy(w, resp.Body)
		if err != nil && err != io.EOF {
			fmt.Println("failed to copy response from underlying ping to response:", err)
			return
		}
		if err := resp.Body.Close(); err != nil {
			fmt.Println("failed to copy response from underlying ping to response:", err)
			return
		}
	})
	fmt.Println("using Ping URL:", url)
	if err := http.ListenAndServe(":8000", nil); err != nil {
		fmt.Println(err)
	}
	fmt.Println("exiting...")
}
