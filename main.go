package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"os/exec"

	"github.com/mikerybka/util"
)

func main() {
	http.HandleFunc("/btc", func(w http.ResponseWriter, r *http.Request) {
		// RPC credentials
		username := util.RequireEnvVar("BTCRPC_USER")
		password := util.RequireEnvVar("BTCRPC_PASS")

		// JSON RPC payload
		data := `{
		"jsonrpc": "1.0",
		"id": "curltest",
		"method": "getblockchaininfo",
		"params": []
	}`

		req, err := http.NewRequest("POST", "http://127.0.0.1:8332", bytes.NewBuffer([]byte(data)))
		if err != nil {
			panic(err)
		}

		// Set headers
		req.Header.Set("Content-Type", "text/plain")

		// Set basic auth
		auth := username + ":" + password
		authHeader := "Basic " + base64.StdEncoding.EncodeToString([]byte(auth))
		req.Header.Set("Authorization", authHeader)

		// Do the request
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		// Read the response
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}

		// Print the result
		w.Write(body)
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
