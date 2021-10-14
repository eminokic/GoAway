// Declare the main package.
package main
// Import the formatting package provided by Go.
import "fmt"

// Declaring the main function for the go program to run.
func main() {
  // Ask the user a question using the print line funciton.
  fmt.Println("How has your day been? (good/bad)")
  // Declare the variable take user's input.
  var response string

  // Scan the users input.
  fmt.Scanln(&response)

  if(response == "good") {
      fmt.Println("That is awesome, have a good day!")
  } else {
    fmt.Println("Aw, I hope it gets better ...")
  }
}

