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

var username string
var topTracksCmd = &cobra.Command{
	Use:   "topTracks",
	Short: "List of most listened to songs",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		apiKey := os.Getenv("LASTFM_API_KEY")

		fmt.Println("Enter your Last.fm username:")
		fmt.Scanln(&username)

		if username == "" {
			fmt.Println("Invalid user!")
			return
		}

		response, err := service.GetTopTracks(apiKey, username)

		if err != nil {
			fmt.Printf("Error retrieving data: %v\n", err)
		}

		fmt.Printf("\n🎧 Showing the most listened to songs by:%s\n", username)
		fmt.Println("--------------------------------------------------")

		for i, t := range response.TopTracks.Track {
			fmt.Printf("%d. %s - %s reproductions\n", i+1, t.Name, t.Playcount)
		}

		fmt.Println("--------------------------------------------------")
	},
}

func init() {
	rootCmd.AddCommand(topTracksCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// topTracksCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// topTracksCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
