package main

import (
	"io/ioutil"
	"net/http"
	"time"
)

func main() {

	c := http.Client{Timeout: time.Millisecond}

	res, err := c.Get("http://google.com")
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	println(string(body))

}
