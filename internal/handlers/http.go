package handlers

import (
  "os"
  "log"
  "time"
  "fmt"
  "net/http"
  "strconv"
  "io/ioutil"
  "github/diploma/internal/calculations"
)

func GraphHandler(wr http.ResponseWriter, r *http.Request) {
    // ABOUT SECTION HTML CODE
    defer func(since time.Time) {
  		log.Print("execution-time", time.Since(since).String())
  	}(time.Now())

    query := r.URL.Query()
    m := clean(query["m"])
    L := clean(query["l"])
    w := clean(query["w"])
    a := clean(query["a"])
    calculations.Calculations(m, L, w, a)
    wr.Header().Set("Content-Type", "text/html; charset=utf-8")
    f, err := os.Open("html/exclude.html")
    if err != nil {
        fmt.Println(err)
    }

    b, err := ioutil.ReadAll(f)
    wr.Write(b)
}

func clean(npt []string) float64{
    s, err := strconv.ParseFloat(npt[0], 64)
    if err != nil {
        fmt.Println(err)
    }
    return s
}

func InputHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/html; charset=utf-8")

    f, err := os.Open("html/input.html")
    if err != nil {
        fmt.Println(err)
    }

    b, err := ioutil.ReadAll(f)
    w.Write(b)
}
