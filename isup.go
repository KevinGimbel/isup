package main
// Import used packages
import (
  "os"
  "fmt"
  "log"
  "net/http"
  "strings"
)

// Check if a valid HTTP/S protocol is present.
// http.Get throws an error otherwise.
func isValidUrlProtocol(url string) bool {
  isHttps := strings.Contains(url, "https")
  isHttp := strings.Contains(url, "http")
  isValid := (isHttps != false) || (isHttp != false)
  return isValid
}

// Print invalid URL protocol message
func printInvalidFormatError() {
  fmt.Println("Could not find protocol. Please specify a protocol.\n")
  fmt.Println("Example calls:")
  fmt.Println("$", os.Args[0], "'http://kevingimbel.com'")
  fmt.Println("$", os.Args[0], "'https://github.com' \n")
  os.Exit(1)
}

// Print down message if the page cannot be reached
func printDownMessage(url string, err error) {
  // Print a message followed by the actual error
  fmt.Println("Looks like", url, "is down or not available!\n")
  fmt.Println("http.Get Error:")
  log.Fatal(err)
  os.Exit(1)
}

// Print the up message.
// If the status is 503 (internal server error), a special message is printed.
func printUpMessage(url string, response *http.Response) {
  // Print a message based on the status.
  if response.Status != "503" {
    fmt.Println("Looks like", url, "is up!\nStatus Code:", response.Status)
  } else {
    fmt.Println("Looks like", url, "is up but has a internal error. Status Code: 503")
  }
  os.Exit(0)
}

// Define the main function
// This function runs when the program is executed.
func main() {
  // Read the arguments passed on execution time
  // We only care for the first (0 is the program name)
  url := os.Args[1]

  // Check if the URL has a valid protocol
  if (isValidUrlProtocol(url) == false ) {
    printInvalidFormatError()
  }

  // Try to get the website with http.Get
  response, err := http.Get(url)

  // If the http.Get failed, err will hold the error message
  if err != nil {
    // Print a message followed by the actual error
    printDownMessage(url, err)
  }

  // if we reach this point we were able to make the HTTP Request and got a response.
  // Print the Up Message with the URL and response
  printUpMessage(url, response)
}
