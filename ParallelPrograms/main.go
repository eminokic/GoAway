// Creating a parallel program using GoLang.
package main

import "fmt"
import "runtime"
import "sync"
// Declare a wait group to run the two programs in parallel.
var waitGroup sync.WaitGroup

// Declare the main function to run hello and goodbye in parallel. 
func main() { // Open up main block. 
  fmt.Println("OS: ", runtime.GOOS)
  fmt.Println("Arch: ", runtime.GOARCH)
  fmt.Println("CPUs: ", runtime.NumCPU())
  fmt.Println("GoRoutines: ", runtime.NumGoroutine())

  // So the first programs stats are setup.
  // Now it is time to call the functions to do so.

  fmt.Println("\n Preparing Wait Group ...")
  waitGroup.Add(1)

  // Call the corresponding programs to run in parallel.
  go Hello() 
  Goodbye() 

  // Similar to above, print the stats of the computer/thread running the second program in parallel. 

  fmt.Println("\n\n OS: ", runtime.GOOS)
  fmt.Println("Arch: ", runtime.GOARCH)
  fmt.Println("CPUs : ", runtime.NumCPU())
  fmt.Println("GoRoutines: ", runtime.NumGoroutine())
  
  waitGroup.Wait() 
} // Close main block.

// The hello function is a loop that prints hello.
func Hello() {
    for i := 0; i < 10; i++ {
      // Print off hello as many iterations as wanted.
      fmt.Println("Hello ... ", i)

    } // Close for loop.

    // Use the wait group to see when hello is complete.
    waitGroup.Done()
} // Close Hello Helper Function.

// Declaring the Goodbye function to be run in parallel.
func Goodbye() {
    for j := 0; j < 10; j++ {
        // Print off goodbye as many iterations as wanted.
        fmt.Println("Goodbye ... ", j)

    } // Close for loop.
} // Close Goodbye Helper Function.

