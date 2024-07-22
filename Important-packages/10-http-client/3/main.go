package main

import (
	"bytes"
	"io"
	"net/http"
	"os"
)

func main() {

	c := http.Client{}
	jsonBody := bytes.NewBuffer([]byte(`{"name":"Augusto"}`))
	res, err := c.Post("http://google.com", "application/json", jsonBody)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	io.CopyBuffer(os.Stdout, res.Body, nil)
}
