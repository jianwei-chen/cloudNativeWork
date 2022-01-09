package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"net/http"
	_ "net/http/pprof"

	"github.com/golang/glog"
)

func main() {
	flag.Set("v", "4")
	glog.V(2).Info("Starting http server...")
	http.HandleFunc("/healthz", healthz)
	http.HandleFunc("/", rootHandler)
	c, python, java := true, false, "no!"
	fmt.Println(c, python, java)
	err := http.ListenAndServe(":80", nil)
	if nil != err {
		log.Fatal(err)
	}
}

func healthz(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "200\n")
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("entering root handler")
	//io.WriteString(w, "--------Details of the http request header------------\n")
	for key, v := range r.Header {
		w.Header().Set(key, fmt.Sprintf("%s", v))
		//io.WriteString(w, fmt.Sprintf("%s=%s\n", key, v))
	}
	VERSION := os.Getenv("VERSION")
	fmt.Println(VERSION)
	w.Header().Set("VERSION", VERSION)
	w.WriteHeader(http.StatusOK)

	user := r.URL.Query().Get("user")
	if "" != user {
		io.WriteString(w, fmt.Sprintf("hello [%s]\n", user))
	} else {
		io.WriteString(w, "hello [stranger]\n")
	}

	fmt.Println(strings.Split(r.RemoteAddr, ":")[0], ":", 200)
}
