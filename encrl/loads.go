/*
* In this file we store all the configuration for the codifications, as well
* as the functions that load the codifications table from a json file
*
* All the codifications tables are inside the codifications/ file. For example
* codifications/caesar.json
*/
package main 

import (
  "os"
  "fmt"
)
func loadCipher (cipher string) {
  /*
  * This function loads a cipher from the codifications/ folder to then return it.
  * The default cipher that is loaded is the caesar cipher.
  */
  cipherToLoad := fmt.Sprintf("../codifications/%s.json", cipher)
  jsonCipher, err := os.Open(cipherToLoad)
  if err != nil {
    fmt.Printf(errorColor, "[FATAL ERROR]:      Can not load the json cipher file\n")
    os.Exit(1)
  }
  // Inform the user about the loading cipher
  fmt.Printf(noticeColor, "[CIPHER]:          ")
  fmt.Println(cipherToLoad)
  fmt.Sprintf("%s", jsonCipher)
}
