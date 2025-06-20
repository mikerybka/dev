package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/mikerybka/util"
)

func main() {
	http.HandleFunc("POST /api/webhooks/github", func(w http.ResponseWriter, r *http.Request) {
		io.Copy(os.Stdout, r.Body)
		os.Stdout.Write([]byte("\n"))
	})
	port := util.EnvVar("PORT", "3000")
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Println(err)
	}
}
