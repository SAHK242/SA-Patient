package main

import (
	"os"

	"patient/cmd"
)

func main() {
	if err := os.Setenv("TZ", "Asia/Saigon"); err != nil {
		panic(err)
	}

	cmd.Start()
}
