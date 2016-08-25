package main


import (
    "fmt"
    "io/ioutil"
    "net/http"
    "path/filepath"
    "strings"

    "github.com/garyburd/redigo/redis"
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

func genericHandler(w http.ResponseWriter, r *http.Request) {
    b, err := loadPage(r.URL.Path);
    if err == nil {
        w.Write(b);
    } else {
        http.Error(w, err.Error(), 404);
    }

//  w.Header().Set("Content-Type", "application/json; charset=utf-8")
}

// login process
func login(w http.ResponseWriter, r *http.Request) {
    // readbody
    // verify identity
    // generate token
    // send token
    // store token
}

// config personal profile
func config(w http.ResponseWriter, r *http.Request) {
    // verify token
    // store config
}

// admin list users & config
func admin(w http.ResponseWriter, r *http.Request) {
    // verify token
    // list admin profile
}

func post(w http.ResponseWriter, r *http.Request) {
    // verify token
    // mapping identify
    // store latlng
}

func get(w http.ResponseWriter, r *http.Request) {
    // verify token
    // mapping identify
    // read friends latlng
}

func main() {
    // init db conn
    http.HandleFunc("/login:", genericHandler)
    http.HandleFunc("/config:", genericHandler)    
    http.HandleFunc("/admin:", genericHandler)
    http.HandleFunc("/post:", genericHandler)
    http.HandleFunc("/get:", genericHandler)
    http.HandleFunc("/", genericHandler)
    http.ListenAndServe(":8080", nil)
}




