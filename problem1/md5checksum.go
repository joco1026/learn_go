package main

import (
  "io/ioutil"
  "crypto/md5"
  "fmt"
  "os"
)

func main() {

  contents, err := ioutil.ReadFile("checkme")
  input, err := ioutil.ReadAll(os.Stdin)

  if err != nil {
    panic(err)
  }

  inputhasher := md5.New()
  inputhasher.Write(input)

  filehasher := md5.New()
  filehasher.Write(contents)

  //print contents of file and its hash
  fmt.Printf("%s\n", contents)
  fmt.Printf("%x\n", filehasher.Sum(nil))

  //print standard input and its hash
  fmt.Printf("%s\n", input)
  fmt.Printf("%x\n", inputhasher.Sum(nil))

  if err != nil {
    panic(err)
  }
}


