package main

import (
	"chat/system"

	"github.com/argus-labs/go-ecs/pkg/cardinal"
)

func main() {
	world := cardinal.NewWorld()

	cardinal.RegisterSystem(world, system.UserChatSystem)

	world.StartGame()
}
