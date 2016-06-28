package main
// Import used functions
import (
  "os"
  "fmt"
  "log"
  "net/http"
)

// Define the main function
func main() {
  // Read the arguments passed on execution time
  // We only care for the first (0 is the program name)
  url := os.Args[1]

  // Try to get the website with http.Get
  response, err := http.Get(url)

  // If the http.Get failed, err will hold the error message
  if err != nil {
    // Print a message followed by the actual error
    fmt.Println("Looks like", url, "is down or not available!\n")
    fmt.Println("http.Get Error:")
    log.Fatal(err)
  }
  // If it didn't fail and the status code is not 503
  // assume the site is up.
  if response.Status != "503" {
    fmt.Println("Looks like", url, "is up!\nStatus Code:", response.Status)
  }
}
