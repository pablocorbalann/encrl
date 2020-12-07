package main

import (
  "fmt"
  "os"
  "flag"
  "path/filepath"
)

const (
  infoColor    = "\033[1;34m%s\033[0m"
  noticeColor  = "\033[1;36m%s\033[0m"
  alertColor = "\033[1;33m%s\033[0m"
  errorColor   = "\033[1;31m%s\033[0m"
  debugColor   = "\033[0;36m%s\033[0m"
)

func parseArguments() (string, string, string) {
  /*
  * This function is used to parse all the arguments from the "os" module. It
  * also validates the arguments before using them.
  * Then he returns a list with the arguments
  */
  // Use the flag package to read all the flags and ten parse them using
  // the Parse method
  readingFile := flag.String("r", "", "The file to read from.")
  writingFile := flag.String("w", "", "The file to write at")
  codificationTable := flag.String("c", "caesar", "The codification the program should use.")
  flag.Parse()
  fmt.Printf("\n[PARSING]:         r %s | w  %s\n", *readingFile, *writingFile)
  fmt.Printf("[CODIFICATION]:    %s\n", *codificationTable)
  // Create basic variables to then return them and check if there 
  // is a reading file and a writing file
  var readingFileRoute, writingFileRoute string
  if *readingFile == "" || *writingFile == "" {
    // One of the flags is missing so we can't parse them
    fmt.Printf(errorColor, "[FATAL ERROR]:     Parse error, one of the flags is missing\n")
    os.Exit(1)
  }
  // Convert from the relative path to the absolute path using the filepath module
  readingFileRoute, readingErr := filepath.Abs(*readingFile)
  writingFileRoute, writingErr := filepath.Abs(*writingFile)
  if readingErr == nil && writingErr == nil {
    // There actually were two paths, so we should inform the user
    fmt.Printf(noticeColor, "\n[READING FROM]:    ")
    fmt.Println(readingFileRoute)
    fmt.Printf(noticeColor, "[WRITTING TO]:     ")
    fmt.Println(writingFileRoute)
    // Return the complete routes
    return readingFileRoute, writingFileRoute, *codificationTable
  } else {
    fmt.Printf(errorColor, "[FATAL ERROR]:     One of the routes is not valid\n")
    os.Exit(1)
  }
  return "", "", ""
}

func main() {
  encrlIcon := `
d88888b      d8b   db       .o88b.      d8888b.      db      
88           888o  88      d8P  Y8      88   8D      88      
88ooooo      88V8o 88      8P           88oobY'      88      
88~~~~~      88 V8o88      8b           88 8b        88      
88.          88  V888      Y8b  d8      88  88.      88booo. 
Y88888P      VP   V8P       Y88P'       88   YD      Y88888P 
`
  fmt.Printf(noticeColor, encrlIcon)
  readingFile, writingFile, codification := parseArguments()
  fmt.Println(readingFile, writingFile, codification)
}
