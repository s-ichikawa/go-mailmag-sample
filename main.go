package main

import (
    "log"
    "os"
    "net/http"
    "path"
    "path/filepath"
    "io"
    "io/ioutil"
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
            _, header, _ := r.FormFile("file")
            s, _ := header.Open()
            p := filepath.Join("files", header.Filename)
            buf, _ := ioutil.ReadAll(s)
            ioutil.WriteFile(p, buf, 0644)
            http.Redirect(w, r, "/"+path, 301)
        } else {
            http.Redirect(w, r, "/", 301)
        }

    })
    http.ListenAndServe(":8080", nil)
}
