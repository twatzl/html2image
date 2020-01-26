package main

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"flag"
	"fmt"
	"github.com/twatzl/html2image/html2image"
	"github.com/unrolled/render" // or "gopkg.in/unrolled/render.v1"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)


func main() {
	binPath := flag.String("path", "/usr/local/bin/wkhtmltoimage", "wkhtmltoimage bin path")
	imgRootDir := flag.String("img.dir", "./tmp/", "generated image local dir")
	port := flag.String("web.port", "8080", "web server port")
	flag.Parse()
	imgRender := html2image.ImageRender{}
	imgRender.BinaryPath = binPath
	r := render.New()
	mux := http.NewServeMux()
	staticHandler := http.FileServer(http.Dir(*imgRootDir))

	mux.HandleFunc("/to/img.png", func(w http.ResponseWriter, req *http.Request) {
		imgRender.RenderBytes(w, req, "png")
	})
	mux.HandleFunc("/to/img.jpg", func(w http.ResponseWriter, req *http.Request) {
		imgRender.RenderBytes(w, req, "jpg")
	})
	mux.HandleFunc("/show/img/", func(w http.ResponseWriter, req *http.Request) {
		req.URL.Path = req.URL.Path[9:]
		staticHandler.ServeHTTP(w, req)
	})
	mux.HandleFunc("/api/v1/to/img.json", func(w http.ResponseWriter, req *http.Request) {
		imgRender.RenderJson(r, w, req, imgRootDir)
	})
	if len(*port) == 0 {
		*port = "8080"
	}
	http.ListenAndServe(":"+*port, mux)
}
