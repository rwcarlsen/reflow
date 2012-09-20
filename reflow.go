
package main

import (
  "fmt"
  "strings"
  "os"
  "io/ioutil"
  "log"
  "flag"
  "strconv"
)

var max = 70

func main() {
  flag.Parse()
  if flag.NArg() > 0 {
    if n, err := strconv.Atoi(flag.Arg(0)); err == nil {
      max = n
    }
  }
        
  // get text from stdin
  data, err := ioutil.ReadAll(os.Stdin)
  if err != nil {
    log.Fatal(err)
  }
  text := string(data)

  // stop at first line break
  end := strings.Index(text, "\n\n")
  if end >=0 {
    text = text[:end]
  }

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
    tailLen := len(w)-len(strings.Trim(w, " \t"))
    if chars - tailLen > max {
      line := strings.Join(words[start:i], "")
      lines = append(lines, strings.Trim(line, " \t"))
      chars, start = 0, i
    }
  }
  line := strings.Join(words[start:], "")
  lines = append(lines, strings.Trim(line, " \t"))
  
  fmt.Print(strings.Join(lines, "\n"))
}

