package main

import "fmt"
import "net/http"

func main(){
	url := "https://api.ipify.org/?format=json"

	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		fmt.Println("Something went wrong:", err)
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Something went wrong: ", err)
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusFound {
        distopiaHeader := resp.Header.Get("Distopia")
        fmt.Println("Distopia header value:", distopiaHeader)
    } else {
		fmt.Println("Status Code:", resp.StatusCode, "is not a redirect.")
	}
}
