package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	loopRequest()
}

func loopRequest() {
	for {
		resp, err := http.Get("http://85.192.35.76/api")
		if err != nil {
			log.Fatalln(err)
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}

		log.Println(resp.StatusCode, string(body))
	}
}