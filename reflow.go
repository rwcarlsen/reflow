
package main

import (
  "fmt"
  "strings"
  "os"
  "io/ioutil"
  "log"
)

var max = 80

func main() {
  // get text from stdin
  data, err := ioutil.ReadAll(os.Stdin)
  if err != nil {
    log.Fatal(err)
  }
  text := string(data)

  // split text into words
  text = strings.Replace(text, "\n", " ", -1)
  words := []string{}
  for {
    i := strings.IndexAny(text, " \t")
    if i == -1 {
      words = append(words, text)
      break
    }
    words = append(words, text[:i+1])
    text = text[i+1:]
  }

  // build char limited lines
  lines := []string{}
  chars, start := 0, 0
  for i, w := range words {
    chars += len(w)
    if chars > max {
      line := strings.Join(words[start:i], "")
      lines = append(lines, line)
      chars, start = 0, i
    }
  }
  line := strings.Join(words[start:], "")
  lines = append(lines, line)

  fmt.Println(strings.Join(lines, "\n"))
}

