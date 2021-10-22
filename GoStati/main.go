package main

import "fmt"

type status struct {
  does_playlist_exist bool
  do_songs_exist bool
}

func main() {
  stat := status{does_playlist_exist: false, do_songs_exist: false}
  fmt.Println(stat)
  if(!stat.does_playlist_exist){
    fmt.Println("creating playlist")
  } else if (!stat.do_songs_exist) {
    fmt.Println("Adding Songs to playlist...")
  }
}
