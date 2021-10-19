// Import the package, in this case, main.
package main

// Import GoLang's formatting library.
import "fmt"

// Importing GoLang's buffer and input/output library.
import "bufio"

// Importing GoLang's OS library.
import "os"

// Defining the main function to be run to collect a user's favorite songs.
func main() {
  // Defining a reader to read the user's input.
  reader := bufio.NewReader(os.Stdin)
  // Printing a prompt for the user to read in.
  fmt.Println("Enter your spotify playlist...")
  // Now defining a variable to take in the user's input following the prompt.
  playlist, _ := reader.ReadString('\n')
  // Asking the user for the second prompt and taking in the second input.
  fmt.Println("Enter your favorite song currently on spotify and or apple music...")
  favoriteSong, _ := reader.ReadString('\n')

  fmt.Println("Playlist: " + playlist + ", Favorite Song: " + favoriteSong)
}

