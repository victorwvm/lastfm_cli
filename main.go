/*
Copyright © 2026 Victor Dantas <victordantas1524@gmail.com>
*/
package main

import (
	"github.com/joho/godotenv"
	"github.com/victorwvm/lastfm_cli/cmd"
)

func main() {
	_ = godotenv.Load()

	cmd.Execute()
}
