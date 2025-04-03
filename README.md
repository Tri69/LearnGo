# LearnGo
Repository Tutorial GoLang

### A. Basic GoLang <br>
file : main.go
```go
package main

import "fmt"

func main() {
  fmt.Println("Hello World")
}
```
Run Golang : ```go run main.go ```<br>
Build Go : ```go build main.go```

### B. Server API Golang<br>
file : server.go
```go
package main

import(
  "fmt"
  "net/http"
)

func HandlerIndex(w http.ResponseWriter, r *http.Request) {
  fmt.Fprint(w, "Test Api Handler")
}

func main() {
  http.HandlerFunc("/", HandlerIndex)
  http.HandlerFunc("/about", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "About Page Test")
})

  fmt.Println("Server is running in http://127.0.0.1:3001")
  http.ListenAndServe(":3001")
}
