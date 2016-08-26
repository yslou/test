package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/garyburd/redigo/redis"
	"github.com/yslou/test/lib"
)

var (
	ErrNotSupportedExtension = fmt.Errorf("Not supported extenstion")
	db                       redis.Conn
	tokens                   map[string]model.User
)

func loadPage(path string) (body []byte, err error) {
	if strings.HasSuffix(path, "/") {
		path = filepath.Join(path, "index.html")
	}
	if !strings.HasSuffix(path, ".html") && !strings.HasSuffix(path, ".js") {
		return body, ErrNotSupportedExtension
	}
	f := filepath.Join("www", path)
	return ioutil.ReadFile(f)
}

func genericHandler(w http.ResponseWriter, r *http.Request) {
	b, err := loadPage(r.URL.Path)
	if err == nil {
		w.Write(b)
	} else {
		http.Error(w, err.Error(), 404)
	}

	//  w.Header().Set("Content-Type", "application/json; charset=utf-8")
}

// login process
func login(w http.ResponseWriter, r *http.Request) {
	// read credential from body
	cert, err := model.ReadCert(r.Body)
	if err != nil {
		fmt.Println("login: syntax error")
		http.Error(w, "Permission Denied", 403)
		return
	}
	// verify identity
	res, err := redis.Bool(db.Do("EXISTS", "_user_"+cert.Login))
	if err != nil {
		fmt.Println("login: db error ", cert.Login)
		http.Error(w, "Internal Error", 500)
		return
	} else if !res {
		fmt.Println("login: unknow user ", cert.Login)
		http.Error(w, "Permission Denied", 403)
		return
	}
	rec, err := redis.String(db.Do("GET", "_user_"+cert.Login))
	user, err := model.JSONUser(rec)
	if err != nil {
		fmt.Println("login: db error ", cert.Login)
		http.Error(w, "Internal Error", 500)
		return
	}
	if cert.Password != user.Password {
		fmt.Println("login: invalid pwd for ", cert.Login)
		http.Error(w, "Permission Denied", 403)
		return
	}
	// generate token
	token := model.NewToken(cert)
	// store token
	tokens[token] = user
	// send token
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	var t model.Ticket
	t.Login = cert.Login
	t.Token = token
	json.NewEncoder(w).Encode(t)
}

// config personal profile
func config(w http.ResponseWriter, r *http.Request) {
	// verify token
	token := r.Header.Get("myAuth")
	user, ok := tokens[token]
	if !ok {
		fmt.Println("config: Unauthorized")
		http.Error(w, "Unauthorized", 401)
		return
	}
	// store config
	cfg, err := model.ReadUser(r.Body)
	if err != nil {
		fmt.Println("config: syntax error")
		http.Error(w, "Invalid operation", 403)
		return
	}
	cfg.Password = user.Password
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
	var err error
	// init db conn
	redisAddr := flag.String("redis-addr", ":6379", "Redis server address")
	db, err = redis.Dial("tcp", *redisAddr)
	if err != nil {
		fmt.Println("failed connect to redis server")
		return
	}

	http.HandleFunc("/login:", login)
	http.HandleFunc("/config:", config)
	http.HandleFunc("/admin:", admin)
	http.HandleFunc("/post:", post)
	http.HandleFunc("/get:", get)
	http.HandleFunc("/", genericHandler)
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("what? ", err)
	}
}
