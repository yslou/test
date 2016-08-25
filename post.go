package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
)

var message = []byte(`
{
    "github_url": "https://gist.github.com/YOUR_ACCOUNT/GIST_ID",
    "contact_email": ""
}`)

var url = "https://192.168.1.1"

func main() {
    req, err := http.NewRequest("POST", url, bytes.NewBuffer(message))
    req.Header.Set("X-Custom-Header", "myvalue")
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    fmt.Println("response Status:", resp.Status)
    fmt.Println("response Headers:", resp.Header)
    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println("response Body:", string(body))
}
