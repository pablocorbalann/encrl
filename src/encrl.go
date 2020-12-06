package main

import (
  "fmt"
  "os"
  "path/filepath"
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
    // Convert from the relative path to the absolute path using the filepath module
    readingFileRoute, readingErr := filepath.Abs(args[0])
    writingFileRoute, writingErr := filepath.Abs(args[1])
    if readingErr == nil && writingErr == nil {
      // There actually were two paths, so we should inform the user
      fmt.Printf(noticeColor, "\n[READING FROM]:    ")
      fmt.Println(readingFileRoute)
      fmt.Printf(noticeColor, "[WRITTING TO]:     ")
      fmt.Println(writingFileRoute)
    } else {
      fmt.Printf(errorColor, "[FATAL ERROR]:     One of the routes is not valid\n")
    }
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
