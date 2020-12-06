package main

import (
  "fmt"
  "os"
)

const (
  errorColor = "\033[1;31m%s\033[0m"
  debugColor = "\033[0;36m%s\033[0m"
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
    fmt.Printf(debugColor, "[READING FROM]: ") 
    fmt.Println(args[0])
    fmt.Printf(debugColor, "[WRITTING TO]: ")
    fmt.Println(args[1])
  } else {
    fmt.Printf(errorColor, "[FATAL ERROR]: Requiered 2 positional parameters - base file & output file\n")
  }
  return args
}

func main() {}
