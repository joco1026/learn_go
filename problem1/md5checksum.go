package main

import (
  "io"
)

func main() {
  file := "checkme"
  contents := io.ReadFile(file)

}



/*



import (
    "crypto/md5"
    "fmt"
    "io"
)

func main() {
    h := md5.New()
    io.WriteString(h, "The fog is getting thicker!")
    fmt.Printf("%x", h.Sum(nil))
}

--------

import (
  "bufio"
  "crypto/md5"
  "fmt"
  "os"
)

func main() {
  data := bufio.NewReader(os.Stdin)
  stats, _ := os.Stdin.Stat()

  if stats.Size() > 0 {
    str, _, _ := data.ReadLine()
    fmt.Printf("%x\n", md5.Sum([]byte(str)))
  } else {
    fmt.Printf("%x\n", md5.Sum([]byte("")))
  }
}

--------
*/

