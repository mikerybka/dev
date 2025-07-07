package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os/exec"

	"github.com/btcsuite/btcd/rpcclient"
	"github.com/mikerybka/util"
)

func main() {
	http.HandleFunc("/btc", func(w http.ResponseWriter, r *http.Request) {
		connCfg := &rpcclient.ConnConfig{
			Host:         "localhost:8332",
			User:         util.RequireEnvVar("BTCRPC_USER"),
			Pass:         util.RequireEnvVar("BTCRPC_PASS"),
			HTTPPostMode: true, // Bitcoin core only supports HTTP POST
			DisableTLS:   true, // If TLS is not configured
		}

		client, err := rpcclient.New(connCfg, nil)
		if err != nil {
			log.Fatalf("Error creating new client: %v", err)
		}
		defer client.Shutdown()

		res, err := client.GetBlockChainInfo()
		if err != nil {
			log.Fatalf("Error getting blockchain info: %v", err)
		}

		fmt.Fprintln(w, res.VerificationProgress)
	})
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
