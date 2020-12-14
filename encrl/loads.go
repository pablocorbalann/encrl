/*
* In this file we store all the configuration for the codifications, as well
* as the functions that load the codifications table from a json file
*
* All the codifications tables are inside the codifications/ file. For example
* codifications/caesar.json
*/
package main 

import (
  "encoding/json"
  "fmt"
  "io/ioutil"
  "os"
  "flag"
  "path/filepath"
)

func loadArguments() (string, string, string) {
  /*
  * This function is used to parse all the arguments from the "os" module. It
  * also validates the arguments before using them.
  * Then he returns a list with the arguments
  */
  // Use the flag package to read all the flags and ten parse them using
  // the Parse method
  readingFile := flag.String("r", "", "The file to read from")
  writingFile := flag.String("w", "", "The file to write at")
  codificationTable := flag.String("c", "caesar", "The codification the program should use.")
  flag.Parse()
  fmt.Printf("\n[PARSING]:         r %s | w  %s\n", *readingFile, *writingFile)
  fmt.Printf("[CODIFICATION]:    %s\n", *codificationTable)
  // Create basic variables to then return them and check if there•
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
    
 
func loadCipher (cipher string) map[string]string {
  /*
  * This function loads a cipher from the codifications/ folder to then return it.
  * The default cipher that is loaded is the caesar cipher.
  */
  cipherToLoad, pathErr := filepath.Abs(fmt.Sprintf("../codifications/%s.json", cipher))
  if pathErr != nil {
    fmt.Printf(errorColor, "[FATAL ERROR]:      Can not get the absolute path of codification\n")
    os.Exit(1)
  }
  jsonCipher, err := ioutil.ReadFile(cipherToLoad)
  if err != nil {
    fmt.Printf(errorColor, "[FATAL ERROR]:      Can not load the json cipher file\n")
    os.Exit(1)
  }
  // Inform the user about the loading cipher
  fmt.Printf("[CIPHER]:          %s\n", cipherToLoad)
  // Parse the json file to a cipher
  var jsonCodificationData map[string]string
  err = json.Unmarshal(jsonCipher, &jsonCodificationData)
  if err != nil {
    fmt.Println(errorColor, "[FATAL ERROR]:      Can not Unmarshal the json cipher file\n")
    os.Exit(1)
  }
  fmt.Printf(noticeColor, "[CIPHER]:          ")
  fmt.Println("Loaded correctly")
  return jsonCodificationData
}
