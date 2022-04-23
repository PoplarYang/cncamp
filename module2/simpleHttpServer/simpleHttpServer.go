package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/http/pprof"
	"os"
	"strings"
	"time"

	"github.com/golang/glog"
)

func Get200(w http.ResponseWriter, r *http.Request) {

	// write request header to response
	if len(r.Header) > 0 {
		for k, v := range r.Header {
			// write to header
			w.Header().Add(k, strings.Join(v, ","))
			// replace existing header
			// w.Header().Set("content-type", "application/json")

			// write to reponse
			io.WriteString(w, fmt.Sprintf("%s: %s\n", k, strings.Join(v, ",")))
		}
	}
	// write status code
	w.WriteHeader(http.StatusOK)

	// get env
	EnvVersion := os.Getenv("VERSION")
	if EnvVersion == "" {
		EnvVersion = "unknown"
	}

	// write content
	content := fmt.Sprintf("status ok, version: %s\n", EnvVersion)
	w.Write([]byte(content))
	glog.Infof("%s %s %s %d UserAgent: %s\n", r.RemoteAddr, r.Method, r.RequestURI, 200, r.UserAgent())
}

func Get403(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusForbidden)
	w.Write([]byte("status forbidden\n"))
	glog.Infof("%s %s %s %d UserAgent: %s\n", r.RemoteAddr, r.Method, r.RequestURI, 403, r.UserAgent())
}

func Get404(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("status not found\n"))
	glog.Infof("%s %s %s %d UserAgent: %s\n", r.RemoteAddr, r.Method, r.RequestURI, 404, r.UserAgent())
}

func Get500(w http.ResponseWriter, r *http.Request) {
	glog.Warning("internal error")
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("status internal server error\n"))
	glog.Warningf("%s %s %s %d UserAgent: %s\n", r.RemoteAddr, r.Method, r.RequestURI, 500, r.UserAgent())
}

func Healthz(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("status ok"))
	glog.Warningf("%s %s %s %d UserAgent: %s\n", r.RemoteAddr, r.Method, r.RequestURI, 200, r.UserAgent())
}

func SlowRequest(w http.ResponseWriter, r *http.Request) {
	time.Sleep(time.Second * 3)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("status ok"))
	glog.Warningf("%s %s %s %d UserAgent: %s\n", r.RemoteAddr, r.Method, r.RequestURI, 200, r.UserAgent())
}

type ServerConf struct {
	ListenAddr string
	// Version    string
}

func main() {
	flag.Parse()
	defer glog.Flush()
	serverConf := ServerConf{ListenAddr: ":8080"}

	http.HandleFunc("/", Get200)
	http.HandleFunc("/healthz", Healthz)
	http.HandleFunc("/get/200", Get200)
	http.HandleFunc("/get/403", Get403)
	http.HandleFunc("/get/404", Get404)
	http.HandleFunc("/get/500", Get500)
	http.HandleFunc("/get/slow", SlowRequest)
	mux := http.NewServeMux()
	mux.HandleFunc("/debug/pprof/", pprof.Index)
	mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	mux.HandleFunc("/debug/pprof/trace", pprof.Trace)

	s := &http.Server{
		Addr:           serverConf.ListenAddr,
		Handler:        nil,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20, // 1M bytes
	}

	glog.Info("SimpleHttpServer Start ...")
	err := s.ListenAndServe()
	if err != nil {
		glog.Fatal(err)
	}
}
