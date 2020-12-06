package main

import (
  "fmt"
  "os"
)

const (
  infoColor    = "\033[1;34m%s\033[0m"
  noticeColor  = "\033[1;36m%s\033[0m"
  alertColor = "\033[1;33m%s\033[0m"
  errorColor   = "\033[1;31m%s\033[0m"
  debugColor   = "\033[0;36m%s\033[0m"
)

func parseArguments()[]string {
  /*
  * This function is used to parse all the arguments from the "os" module. It
  * also validates the arguments before using them.
  * Then he returns a list with the arguments
  */
  args := os.Args[1:]
  if len(args) >= 2 {
    // There actually were two paths, so we should inform the user
    fmt.Printf(noticeColor, "\n[READING FROM]:    ") 
    fmt.Println(args[0])
    fmt.Printf(noticeColor, "[WRITTING TO]:     ")
    fmt.Println(args[1])
  } else {
    fmt.Printf(errorColor, "\n[FATAL ERROR]:     Requiered 2 positional parameters - base file & output file\n")
  }
  return args
}

func main() {
  encrlIcon := `
 ___  _  _  __  ___ _
| __|| \| |/ _|| o \ |
| _| | \\ ( (_ |   / |_
|___||_|\_|\__||_|\\___|
`
  fmt.Printf(noticeColor, encrlIcon)
  args := parseArguments()
  fmt.Println(args)
}
