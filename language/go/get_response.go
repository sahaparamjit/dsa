package main

import (
	"fmt"
	"net/http"
)


func main() {
	
	fmt.Println(contentType("https://www.google.com"))

}

func contentType(url string) (string, error) {
	resp, err := http.Get(url)
	
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	return resp.Header.Get("Content-Type"), err
}