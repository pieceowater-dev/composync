package main

import "composync/cmd"

func main() {
	err := cmd.Execute()
	if err != nil {
		return
	}
}
