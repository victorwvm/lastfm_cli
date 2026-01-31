/*
Copyright © 2026 Victor Dantas <victordantas1524@gmail.com>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/victorwvm/lastfm_cli/service"
)

var topArtistsCmd = &cobra.Command{
	Use:   "topArtists",
	Short: "List the artists you've scroubbled the most.",
	Long:  ``,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey := os.Getenv("LASTFM_API_KEY")

		if username == "" {
			fmt.Print("Enter your Last.fm username: ")
			fmt.Scanln(&username)
		}

		if username == "" {
			fmt.Println("Invalid user!")
			return
		}

		response, err := service.GetTopArtist(apiKey, username)

		if err != nil {
			fmt.Printf("Error retrieving data: %v\n", err)
		}

		fmt.Printf("\n🎧 Showing the most listened-to artists of: %s\n", username)
		fmt.Println("--------------------------------------------------")

		for i, artist := range response.TopArtists.Artist {
			fmt.Printf("%d. %s - %s reproductions\n", i+1, artist.Name, artist.Playcount)
		}
		fmt.Println("--------------------------------------------------")
	},
}

func init() {
	rootCmd.AddCommand(topArtistsCmd)
}
