package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"bytes"
)

func main() {
	body, _ := json.Marshal(map[string]string{
		"id": "29392423",
	})
	nbody := bytes.NewBuffer(body)

	req, err := http.NewRequest("POST", "https://google.com", nbody)
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if r, err := ioutil.ReadAll(resp.Body); err == nil {
		fmt.Println(string(r))
	} else {
		panic(err)
	}
}
