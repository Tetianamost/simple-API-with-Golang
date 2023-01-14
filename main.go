import (
 "fmt"
 "net/http"
)

func main() {
 http.HandleFunc("/", Handler)
 http.ListenAndServe(":8080", nil)
}

func handleGet(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Handling GET request")
}

func Handler(w http.ResponseWriter, r *http.Request) {
    if r.Method == "GET" {
        handleGet(w, r)
    } else {
        fmt.Fprintf(w, "Hello World!")
    }
}