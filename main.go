package main

import (
    "log"
    "os"
    "net/http"
    "path"
    "path/filepath"
    "io"
)

func main() {
    cwd, err := os.Getwd()
    if err != nil {
        log.Fatal(err)
    }

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        if ok, err := path.Match("/data/*.html", r.URL.Path); err != nil || !ok {
            http.NotFound(w, r)
            return
        }

        name := filepath.Join(cwd, "data", filepath.Base(r.URL.Path))
        f, err := os.Open(name)
        if err != nil {
            http.NotFound(w, r)
            return
        }
        defer f.Close()
        io.Copy(w, f)
    })

    http.HandleFunc("/upload", func(w http.ResponseWriter, r *http.Request) {
        if r.Method == "POST" {
            stream, header, err := r.FormFile("file")
            if err != nil {
                http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
                return
            }
            p := filepath.Join("files", header.Filename)
            println(p)
            f, err := os.Create(p)
            if err != nil {
                http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
                return
            }
            defer f.Close()
            io.Copy(f, stream)
            http.Redirect(w, r, path.Join("/files", p), 301)
        } else {
            http.Redirect(w, r, "/", 301)
        }

    })
    http.ListenAndServe(":8080", nil)
}