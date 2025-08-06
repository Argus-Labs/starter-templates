package main

import (
	"world-engine/system"

	"github.com/argus-labs/monorepo/pkg/cardinal"
)

func main() {
	world := cardinal.NewWorld()

	cardinal.RegisterSystem(world, system.UserChatSystem)

	world.StartGame()
}
