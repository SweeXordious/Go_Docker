package main

import (
	"Go_Docker/commands"
	"time"
)

func main() {
	commands.ListContainer()
	id, _ := commands.RunImageByName("nginx")
	commands.ListContainer()
	time.Sleep(30 * time.Second)
	commands.StopContainer(id)
}
