package main

import (
	"world-engine/system"

	"github.com/argus-labs/monorepo/pkg/cardinal"
)

func main() {
	world := cardinal.NewWorld()

	cardinal.RegisterSystem(world, system.PlayerSetUpdater, cardinal.WithHook(cardinal.PreUpdate))
	cardinal.RegisterSystem(world, system.PlayerSpawnSystem)
	cardinal.RegisterSystem(world, system.MovePlayerSystem)
	cardinal.RegisterSystem(world, system.PlayerLeaveSystem)
	cardinal.RegisterSystem(world, system.OnlineStatusUpdater)

	world.StartGame()
}
