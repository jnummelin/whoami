package main

import (
  "os"
  "fmt"
  "net/http"
  "log"
)

func main() {
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    fmt.Fprintf(os.Stdout, "Listening on :%s\n", port)
    hostname, _ := os.Hostname()
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(os.Stdout, "I'm %s\n", hostname)
 	      fmt.Fprintf(w, "I'm %s\n", hostname)
        fmt.Fprintf(w, "\n\n\n")
        fmt.Fprintf(w, "Remote address: %s\n\n", r.RemoteAddr)
        fmt.Fprintf(w, "\n\n\n")
        fmt.Fprintf(w, "Request headers:\n")
        fmt.Fprintf(w, "----------------\n")
        for k, v := range r.Header {
          fmt.Fprintf(w, "%s:%s\n", k, v)
        }
        fmt.Fprintf(w, "\n\n\n")
        fmt.Fprintf(w, "Environment:\n")
        fmt.Fprintf(w, "----------------\n")
        for _, e := range os.Environ() {
          fmt.Fprintln(w, "%s\n", e)
        }
    })


    log.Fatal(http.ListenAndServe(":" + port, nil))
}
