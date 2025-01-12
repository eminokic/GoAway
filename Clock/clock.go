package main

import "time"
import "fmt"

func main() {
  now := time.Now()      // current local time
  sec := now.Unix()      // number of seconds since January 1, 1970 UTC
  nsec := now.UnixNano() // number of nanoseconds since January 1, 1970 UTC

  fmt.Println(now)  // time.Time
  fmt.Println(sec)  // int64
  fmt.Println(nsec) // int64
}
