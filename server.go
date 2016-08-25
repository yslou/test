package main


import (
    "fmt"
    "io/ioutil"
    "net/http"
    "path/filepath"
    "strings"
)

var NotSupportedExtension = fmt.Errorf("Not supported extenstion")

func loadPage(path string) (body []byte, err error) {
    if strings.HasSuffix(path, "/") {
        path = filepath.Join(path, "index.html")
    }
    if !strings.HasSuffix(path, ".html") && !strings.HasSuffix(path, ".js") {
        return body, NotSupportedExtension
    }
    f := filepath.Join("www", path)
    return ioutil.ReadFile(f)
}

func handler(w http.ResponseWriter, r *http.Request) {
    b, err := loadPage(r.URL.Path);
    if err == nil {
        w.Write(b);
    } else {
        http.Error(w, err.Error(), 404);
    }

//  w.Header().Set("Content-Type", "application/json; charset=utf-8")
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}




