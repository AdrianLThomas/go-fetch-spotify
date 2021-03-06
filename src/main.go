package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/AdrianLThomas/go-fetch-spotify/client"
)

func main() {
	artistName := flag.String("artist", "", "Artist name to lookup")
	spotifyToken := flag.String("token", "", "Auth token to use for query")
	flag.Parse()

	if *artistName == "" {
		fmt.Println("--artist not set")
		os.Exit(1)
	}
	if *spotifyToken == "" {
		fmt.Println("--token not set")
		os.Exit(1)
	}

	spotifyClient := client.NewSpotifyClient(*spotifyToken)
	artist, err := spotifyClient.Fetch(*artistName)

	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}

	fmt.Printf("%s\n", artist)
}
