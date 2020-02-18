package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

var port = flag.Int("l", 9100, "listen port")

func main() {
	flag.Parse()

	cmdLine := strings.Join(os.Args[1:], " ")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		cmd := exec.Command("sh", "-c", cmdLine)
		resp, err := cmd.Output()

		if err != nil {
			fmt.Fprintln(os.Stderr, "ERROR:", err)
		}

		w.Write(resp)

	})

	err := http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", *port), nil)

	if err != nil {
		fmt.Fprintln(os.Stderr, "ERROR:", err)
		return
	}

}
