package main

import (
    "net/http"
    "io/ioutil"
    "fmt"
    "crypto/md5"
    "encoding/hex"
)

func handler(w http.ResponseWriter, r *http.Request) {
  //fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
  body, err := ioutil.ReadAll(r.Body)
  if err != nil {
    panic(err)
  }

  bs := []byte(body)
  enc := md5.Sum(bs)
  st := hex.EncodeToString(enc[:])


  fmt.Fprintf(w, st, r.URL.Path[1:])
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}


// curl -i -X POST -H 'Content-Type: application/json' -d '{"name": "New item", "year": "2009"}' http://rest-api.io/items

