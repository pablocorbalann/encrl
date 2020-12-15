package main

import (
  "os"
  "fmt"
)

func checkIfFile(fileRoute string) bool{
  /*
  * This function is used to return true if a file
  * is found inside a directory. 
  *
  * Parameters:
  *   fileRoute -> The file to check
  *
  * Returns:
  *   true if found
  * */
  _, err := os.Stat(fileRoute)
  if err == nil {
    // The file actually exists
    return true
  } else if os.IsNotExist(err) {
    // THe file doesn't exists
    return false
  } else {
    // Schrodinger: file may or may not exist. See err for details.
    fmt.Println("[FATAL ERROR]:     Schrodinger location")
  }
}

func decodeCipherBytes(fileRoute string, cipherBytes []byte) {
  /*
  * This function is used to decode a given cipher in a byte(s)
  * format to a string, so it can be dumped to a file
  *
  * Parameters:
  *   fileRoute -> The file route to dump the information
  *   cipherBytes -> The cipher to be loaded to string
  * */
  if checkIfFile(fileRoute) {
    // File not empty, so we should remove the file to then write it again
    fmt.Printf("[RECREATING]:      %s", fileRoute)
    os.Remove(fileRoute)
  }
}
