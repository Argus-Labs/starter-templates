package main

import (
	"github.com/argus-labs/monorepo/pkg/cardinal"
)

func main() {
	world := cardinal.NewWorld()

	// Register systems
	// cardinal.RegisterSystem(world, system.ExampleSystem)

	world.StartGame()
}
