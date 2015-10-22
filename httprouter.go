package main

import (
    "fmt"
    "github.com/julienschmidt/httprouter"
    "net/http"
    "encoding/json"
)

type requestStruct struct {
	Name string `json:"name"`
}

type responseStruct struct {
	Greeting string `json:"greeting"`
}

func hello(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {
    fmt.Fprintf(rw, "Hello, %s!\n", p.ByName("name"))
}

func helloPost(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {
    decoder := json.NewDecoder(req.Body)
    var rqst requestStruct
    var resp responseStruct   
    decoder.Decode(&rqst)
    resp.Greeting = "Hello,"+rqst.Name+"!"
    json.NewEncoder(rw).Encode(resp)
}

func main() {
    mux := httprouter.New()
    mux.GET("/hello/:name", hello)
    mux.POST("/hello", helloPost)
    server := http.Server{
            Addr:        "0.0.0.0:8080",
            Handler: mux,
    }
    server.ListenAndServe()
}
