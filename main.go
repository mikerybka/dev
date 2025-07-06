package main

import (
	"fmt"
	"io"
	"net/http"
	"os/exec"

	"github.com/mikerybka/util"
)

func main() {
	http.HandleFunc("POST /api/eval", func(w http.ResponseWriter, r *http.Request) {
		b, err := io.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}
		cmd := exec.Command("/bin/bash", "-c", string(b))
		cmd.Stdout = w
		cmd.Stderr = w
		err = cmd.Run()
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
	})
	port := util.RequireEnvVar("PORT")
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Println(err)
	}
}
